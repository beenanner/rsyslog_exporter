VERSION := 0.0.3
TARGET := rsyslog_exporter
GOFLAGS := -ldflags "-X main.Version $(VERSION)"
ROOTPKG := github.com/digitalocean/$(TARGET)

include Makefile.COMMON
