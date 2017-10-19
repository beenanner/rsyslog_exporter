package main

import "testing"

var (
	inputLog = []byte(`{"name":"test_input", "origin":"imuxsock", "submitted":1000}`)
)

func TestGetInput(t *testing.T) {
	logType := getStatType(inputLog)
	if logType != rsyslogInput {
		t.Errorf("detected pstat type should be %d but is %d", rsyslogInput, logType)
	}

	pstat := newInputFromJSON([]byte(inputLog))

	if want, got := "test_input", pstat.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(1000), pstat.Submitted; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}
}

func TestInputtoPoints(t *testing.T) {
	pstat := newInputFromJSON([]byte(inputLog))
	points := pstat.toPoints()

	point := points[0]
	if want, got := "test_input_submitted", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(1000), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}
}
