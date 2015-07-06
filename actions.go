package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type action struct {
	Name              string `json:"name"`
	Processed         int64  `json:"processed"`
	Failed            int64  `json:"failed"`
	Suspended         int64  `json:"suspended"`
	SuspendedDuration int64  `json:"suspended.duration"`
	Resumed           int64  `json:"resumed"`
}

func newActionFromJSON(b []byte) *action {
	dec := json.NewDecoder(bytes.NewReader(b))
	var pstat action
	dec.Decode(&pstat)
	pstat.Name = strings.ToLower(pstat.Name)
	pstat.Name = strings.Replace(pstat.Name, " ", "_", -1)
	return &pstat
}

func (a *action) toPoints() []*point {
	points := make([]*point, 5)

	points[0] = &point{
		Name:        fmt.Sprintf("%s_processed", a.Name),
		Type:        counter,
		Value:       a.Processed,
		Description: "messages processed",
	}

	points[1] = &point{
		Name:        fmt.Sprintf("%s_failed", a.Name),
		Type:        counter,
		Value:       a.Failed,
		Description: "messages failed",
	}

	points[2] = &point{
		Name:        fmt.Sprintf("%s_suspended", a.Name),
		Type:        counter,
		Value:       a.Suspended,
		Description: "times suspended",
	}

	points[3] = &point{
		Name:        fmt.Sprintf("%s_suspended_duration", a.Name),
		Type:        counter,
		Value:       a.SuspendedDuration,
		Description: "time spent suspended",
	}

	points[4] = &point{
		Name:        fmt.Sprintf("%s_resumed", a.Name),
		Type:        counter,
		Value:       a.Resumed,
		Description: "times resumed",
	}

	return points
}
