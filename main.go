package main

import (
	"flag"
	"log"
	"log/syslog"
	"net/http"
	"os"
	"os/signal"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	promlog "github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
)

var (
	listenAddress = flag.String("web.listen-address", ":9104", "Address to listen on for web interface and telemetry.")
	metricPath    = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
)

// landingPage contains the HTML served at '/'.
// TODO: Make this nicer and more informative.
var landingPage = []byte(`<html>
<head><title>Rsyslog exporter</title></head>
<body>
<h1>Rsyslog exporter</h1>
<p><a href='` + *metricPath + `'>Metrics</a></p>
</body>
</html>
`)

func init() {
	prometheus.MustRegister(version.NewCollector("rsyslog_exporter"))
}

func main() {
	logwriter, e := syslog.New(syslog.LOG_NOTICE|syslog.LOG_SYSLOG, "rsyslog_exporter")
	if e == nil {
		log.SetOutput(logwriter)
	}

	flag.Parse()
	exporter := newRsyslogExporter()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		log.Print("interrupt received, exiting")
		os.Exit(0)
	}()

	go func() {
		exporter.run()
	}()

	promlog.Infoln("Starting rsyslog_exporter", version.Info())
	promlog.Infoln("Build context", version.BuildContext())

	prometheus.MustRegister(exporter)
	http.Handle(*metricPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(landingPage)
	})

	log.Printf("Listening on %s", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
