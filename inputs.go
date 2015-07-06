package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		Name:        fmt.Sprintf("%s_submitted", i.Name),
		Type:        counter,
		Value:       i.Submitted,
		Description: "messages submitted",
	}

	return points
}
