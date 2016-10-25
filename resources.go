package main

import (
	"bytes"
	"encoding/json"
)

type resource struct {
	Name     string `json:"name"`
	Utime    int64  `json:"utime"`
	Stime    int64  `json:"stime"`
	Maxrss   int64  `json:"maxrss"`
	Minflt   int64  `json:"minflt"`
	Majflt   int64  `json:"majflt"`
	Inblock  int64  `json:"inblock"`
	Outblock int64  `json:"oublock"`
	Nvcsw    int64  `json:"nvcsw"`
	Nivcsw   int64  `json:"nivcsw"`
}

func newResourceFromJSON(b []byte) *resource {
	dec := json.NewDecoder(bytes.NewReader(b))
	var pstat resource
	dec.Decode(&pstat)
	return &pstat
}

func (r *resource) toPoints() []*point {
	points := make([]*point, 9)

	points[0] = &point{
		Name:        "utime",
		Type:        counter,
		Value:       r.Utime,
		Description: "user time used in microseconds",
		LabelValue:  r.Name,
	}

	points[1] = &point{
		Name:        "stime",
		Type:        counter,
		Value:       r.Stime,
		Description: "system time used in microsends",
		LabelValue:  r.Name,
	}

	points[2] = &point{
		Name:        "maxrss",
		Type:        gauge,
		Value:       r.Maxrss,
		Description: "maximum resident set size",
		LabelValue:  r.Name,
	}

	points[3] = &point{
		Name:        "minflt",
		Type:        counter,
		Value:       r.Minflt,
		Description: "total minor faults",
		LabelValue:  r.Name,
	}

	points[4] = &point{
		Name:        "majflt",
		Type:        counter,
		Value:       r.Majflt,
		Description: "total major faults",
		LabelValue:  r.Name,
	}

	points[5] = &point{
		Name:        "inblock",
		Type:        counter,
		Value:       r.Inblock,
		Description: "filesystem input operations",
		LabelValue:  r.Name,
	}

	points[6] = &point{
		Name:        "oublock",
		Type:        counter,
		Value:       r.Outblock,
		Description: "filesystem output operations",
		LabelValue:  r.Name,
	}

	points[7] = &point{
		Name:        "nvcsw",
		Type:        counter,
		Value:       r.Nvcsw,
		Description: "voluntary context switches",
		LabelValue:  r.Name,
	}

	points[8] = &point{
		Name:        "nivcsw",
		Type:        counter,
		Value:       r.Nivcsw,
		Description: "involuntary context switches",
		LabelValue:  r.Name,
	}

	return points
}
