package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

type rsyslogType int

const (
	rsyslogUnknown rsyslogType = iota
	rsyslogAction
	rsyslogInput
	rsyslogQueue
	rsyslogResource
)

type rsyslogExporter struct {
	debug   bool
	started bool
	logfile *os.File
	scanner *bufio.Scanner
	pointStore
}

func newRsyslogExporter(logPath string) (*rsyslogExporter, error) {
	debug := false
	if len(logPath) > 0 {
		debug = true
	}

	e := &rsyslogExporter{
		debug:   debug,
		scanner: bufio.NewScanner(os.Stdin),
		pointStore: pointStore{
			pointMap: make(map[string]*point),
			lock:     &sync.RWMutex{},
		},
	}

	if e.debug {
		var err error
		e.logfile, err = os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return e, fmt.Errorf("could not open log file")
		}
		log.SetOutput(e.logfile)
		log.Println("starting")
	}

	return e, nil
}

func (re *rsyslogExporter) handleStatLine(line string) {
	pstatType := getStatType(re.scanner.Text())

	switch pstatType {
	case rsyslogAction:
		a := newActionFromJSON(re.scanner.Bytes())
		for _, p := range a.toPoints() {
			re.add(p)
		}

	case rsyslogInput:
		i := newInputFromJSON(re.scanner.Bytes())
		for _, p := range i.toPoints() {
			re.add(p)
		}

	case rsyslogQueue:
		q := newQueueFromJSON(re.scanner.Bytes())
		for _, p := range q.toPoints() {
			re.add(p)
		}

	case rsyslogResource:
		r := newResourceFromJSON(re.scanner.Bytes())
		for _, p := range r.toPoints() {
			re.add(p)
		}

	default:
	}
}

// Describe sends the description of currently known metrics collected
// by this Collector to the provided channel. Note that this implementation
// does not necessarily send the "super-set of all possible descriptors" as
// defined by the Collector interface spec, depending on the timing of when
// it is called. The rsyslog exporter does not know all possible metrics
// it will export until the first full batch of rsyslog impstats messages
// are received via stdin. This is ok for now.
func (re *rsyslogExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- prometheus.NewDesc(
		prometheus.BuildFQName("", "rsyslog", "scrapes"),
		"times exporter has been scraped",
		nil, nil,
	)

	keys := re.keys()

	for _, k := range keys {
		p, err := re.get(k)
		if err != nil {
			ch <- p.promDescription()
		}
	}
}

// Collect is called by Prometheus when collecting metrics.
func (re *rsyslogExporter) Collect(ch chan<- prometheus.Metric) {
	if re.debug {
		log.Print("Collect waiting for lock")
	}

	keys := re.keys()

	for _, k := range keys {
		p, err := re.get(k)
		if err != nil {
			continue
		}

		metric := prometheus.MustNewConstMetric(
			p.promDescription(),
			p.promType(),
			p.promValue(),
		)

		ch <- metric
	}
}

func (re *rsyslogExporter) run() {
	for re.scanner.Scan() {
		if strings.Contains(re.scanner.Text(), "EOF") {
			os.Exit(0)
		}
		re.handleStatLine(re.scanner.Text())
	}
}
