package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	logPath       = flag.String("logpath", "", "Log file to write to for debugging purposes")
	listenAddress = flag.String("web.listen-address", ":9104", "Address to listen on for web interface and telemetry.")
	metricPath    = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
)

func main() {
	flag.Parse()
	exporter, err := newRsyslogExporter(*logPath)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		exporter.run()
	}()

	prometheus.MustRegister(exporter)
	http.Handle(*metricPath, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
<head><title>Rsyslog exporter</title></head>
<body>
<h1>Rsyslog exporter</h1>
<p><a href='` + *metricPath + `'>Metrics</a></p>
</body>
</html>
`))
	})

	err = http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		panic(err)
	}
}
