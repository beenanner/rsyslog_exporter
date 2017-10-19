package main

import "testing"

var (
	resourceLog = []byte(`{"name":"resource-usage","origin":"impstats","utime":10,"stime":20,"maxrss":30,"minflt":40,"majflt":50,"inblock":60,"oublock":70,"nvcsw":80,"nivcsw":90,"openfiles":100}`)
)

func TestNewResourceFromJSON(t *testing.T) {
	logType := getStatType(resourceLog)
	if logType != rsyslogResource {
		t.Errorf("detected pstat type should be %d but is %d", rsyslogResource, logType)
	}

	pstat := newResourceFromJSON([]byte(resourceLog))

	if want, got := "resource-usage", pstat.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(10), pstat.Utime; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(20), pstat.Stime; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(30), pstat.Maxrss; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(40), pstat.Minflt; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(50), pstat.Majflt; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(60), pstat.Inblock; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(70), pstat.Outblock; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(80), pstat.Nvcsw; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(90), pstat.Nivcsw; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := int64(100), pstat.Openfiles; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}
}

func TestResourceToPoints(t *testing.T) {
	pstat := newResourceFromJSON([]byte(resourceLog))
	points := pstat.toPoints()

	point := points[0]
	if want, got := "resource-usage_utime", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(10), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := counter, point.Type; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	point = points[1]
	if want, got := "resource-usage_stime", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(20), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := counter, point.Type; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	point = points[2]
	if want, got := "resource-usage_maxrss", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(30), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := gauge, point.Type; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	point = points[3]
	if want, got := "resource-usage_minflt", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(40), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := counter, point.Type; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	point = points[4]
	if want, got := "resource-usage_majflt", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(50), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := counter, point.Type; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	point = points[5]
	if want, got := "resource-usage_inblock", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(60), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := counter, point.Type; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	point = points[6]
	if want, got := "resource-usage_oublock", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(70), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := counter, point.Type; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	point = points[7]
	if want, got := "resource-usage_nvcsw", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(80), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := counter, point.Type; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	point = points[8]
	if want, got := "resource-usage_nivcsw", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(90), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := counter, point.Type; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	point = points[9]
	if want, got := "resource-usage_openfiles", point.Name; want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := int64(100), point.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := counter, point.Type; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}
}
