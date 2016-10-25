package main

import (
	"bytes"
	"encoding/json"
)

type input struct {
	Name      string `json:"name"`
	Submitted int64  `json:"submitted"`
}

func newInputFromJSON(b []byte) *input {
	dec := json.NewDecoder(bytes.NewReader(b))
	var pstat input
	dec.Decode(&pstat)
	return &pstat
}

func (i *input) toPoints() []*point {
	points := make([]*point, 1)

	points[0] = &point{
		Name:        "input_submitted",
		Type:        counter,
		Value:       i.Submitted,
		Description: "messages submitted",
		LabelValue:  i.Name,
	}

	return points
}
