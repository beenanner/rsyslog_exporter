# rsyslog_exporter [![Build Status](https://travis-ci.org/digitalocean/rsyslog_exporter.svg?branch=master)](https://travis-ci.org/digitalocean/rsyslog_exporter)
A [prometheus](http://prometheus.io/) exporter for [rsyslog](http://rsyslog.com). It accepts rsyslog [impstats](http://www.rsyslog.com/doc/master/configuration/modules/impstats.html) metrics in JSON format over stdin via the rsyslog [omprog](http://www.rsyslog.com/doc/v8-stable/configuration/modules/omprog.html) plugin and transforms and exposes them for consumption by Prometheus.

## Rsyslog Configuration
Configure rsyslog to push JSON formatted stats via omprog:
```
module(
  load="impstats"
  interval="10"
  format="json"
  resetCounters="off"
  ruleset="process_stats"
)

ruleset(name="process_stats") {
  action(
    type="omprog"
    name="to_exporter"
    binary="/usr/local/bin/rsyslog_exporter"
  )
}
```
