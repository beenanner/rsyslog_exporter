package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/prometheus/common/log"
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
	err := dec.Decode(&pstat)
	if err != nil {
		log.Errorf("Could not unmarshall json input '%s': %s", b, err)
	}
	return &pstat
}

func (r *resource) toPoints() []*point {
	points := make([]*point, 9)

	points[0] = &point{
		Name:        fmt.Sprintf("%s_utime", r.Name),
		Type:        counter,
		Value:       r.Utime,
		Description: "user time used in microseconds",
	}

	points[1] = &point{
		Name:        fmt.Sprintf("%s_stime", r.Name),
		Type:        counter,
		Value:       r.Stime,
		Description: "system time used in microsends",
	}

	points[2] = &point{
		Name:        fmt.Sprintf("%s_maxrss", r.Name),
		Type:        gauge,
		Value:       r.Maxrss,
		Description: "maximum resident set size",
	}

	points[3] = &point{
		Name:        fmt.Sprintf("%s_minflt", r.Name),
		Type:        counter,
		Value:       r.Minflt,
		Description: "total minor faults",
	}

	points[4] = &point{
		Name:        fmt.Sprintf("%s_majflt", r.Name),
		Type:        counter,
		Value:       r.Majflt,
		Description: "total major faults",
	}

	points[5] = &point{
		Name:        fmt.Sprintf("%s_inblock", r.Name),
		Type:        counter,
		Value:       r.Inblock,
		Description: "filesystem input operations",
	}

	points[6] = &point{
		Name:        fmt.Sprintf("%s_oublock", r.Name),
		Type:        counter,
		Value:       r.Outblock,
		Description: "filesystem output operations",
	}

	points[7] = &point{
		Name:        fmt.Sprintf("%s_nvcsw", r.Name),
		Type:        counter,
		Value:       r.Nvcsw,
		Description: "voluntary context switches",
	}

	points[8] = &point{
		Name:        fmt.Sprintf("%s_nivcsw", r.Name),
		Type:        counter,
		Value:       r.Nivcsw,
		Description: "involuntary context switches",
	}

	return points
}
