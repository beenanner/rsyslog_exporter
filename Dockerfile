FROM quay.io/prometheus/golang-builder as builder

ADD .   /go/src/github.com/hrak/rsyslog_exporter
WORKDIR /go/src/github.com/hrak/rsyslog_exporter

RUN make

FROM        quay.io/prometheus/busybox:latest
MAINTAINER  The Prometheus Authors <prometheus-developers@googlegroups.com>

COPY --from=builder /go/src/github.com/hrak/rsyslog_exporter/rsyslog_exporter  /bin/rsyslog_exporter

EXPOSE      9104
ENTRYPOINT  [ "/bin/rsyslog_exporter" ]
