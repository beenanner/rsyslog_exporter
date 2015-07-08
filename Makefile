VERSION := 0.0.1
TARGET := rsyslog_exporter
GOFLAGS := -ldflags "-X main.Version $(VERSION)"

include Makefile.COMMON
