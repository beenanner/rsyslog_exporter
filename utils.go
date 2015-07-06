package main

import "strings"

func getStatType(line string) rsyslogType {
	if strings.Contains(line, "processed") {
		return rsyslogAction
	} else if strings.Contains(line, "submitted") {
		return rsyslogInput
	} else if strings.Contains(line, "enqueued") {
		return rsyslogQueue
	} else if strings.Contains(line, "utime") {
		return rsyslogResource
	}
	return rsyslogUnknown
}
