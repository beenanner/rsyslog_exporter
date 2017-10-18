package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/prometheus/common/log"
)

type input struct {
	Name      string `json:"name"`
	Submitted int64  `json:"submitted"`
}

func newInputFromJSON(b []byte) *input {
	dec := json.NewDecoder(bytes.NewReader(b))
	var pstat input
	err := dec.Decode(&pstat)
	if err != nil {
		log.Errorf("Could not unmarshall json input '%s': %s", b, err)
	}
	return &pstat
}

func (i *input) toPoints() []*point {
	points := make([]*point, 1)

	points[0] = &point{
		Name:        fmt.Sprintf("%s_submitted", i.Name),
		Type:        counter,
		Value:       i.Submitted,
		Description: "messages submitted",
	}

	return points
}
