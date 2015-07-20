package main

import (
	"bufio"
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
	started bool
	logfile *os.File
	scanner *bufio.Scanner
	pointStore
}

func newRsyslogExporter() *rsyslogExporter {
	e := &rsyslogExporter{
		scanner: bufio.NewScanner(os.Stdin),
		pointStore: pointStore{
			pointMap: make(map[string]*point),
			lock:     &sync.RWMutex{},
		},
	}
	return e
}

func (re *rsyslogExporter) handleStatLine(buf []byte) {
	pstatType := getStatType(buf)

	switch pstatType {
	case rsyslogAction:
		a := newActionFromJSON(buf)
		for _, p := range a.toPoints() {
			re.add(p)
		}

	case rsyslogInput:
		i := newInputFromJSON(buf)
		for _, p := range i.toPoints() {
			re.add(p)
		}

	case rsyslogQueue:
		q := newQueueFromJSON(buf)
		for _, p := range q.toPoints() {
			re.add(p)
		}

	case rsyslogResource:
		r := newResourceFromJSON(buf)
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
		re.handleStatLine(re.scanner.Bytes())
	}
}
