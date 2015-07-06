package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type queue struct {
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	Enqueued      int64  `json:"enqueued"`
	Full          int64  `json:"full"`
	DiscardedFull int64  `json:"discarded.full"`
	DiscardedNf   int64  `json:"discarded.nf"`
	MaxQsize      int64  `json:"maxqsize"`
}

func newQueueFromJSON(b []byte) *queue {
	dec := json.NewDecoder(bytes.NewReader(b))
	var pstat queue
	dec.Decode(&pstat)
	pstat.Name = strings.ToLower(pstat.Name)
	pstat.Name = strings.Replace(pstat.Name, " ", "_", -1)
	return &pstat
}

func (q *queue) toPoints() []*point {
	points := make([]*point, 6)

	points[0] = &point{
		Name:        fmt.Sprintf("%s_size", q.Name),
		Type:        gauge,
		Value:       q.Size,
		Description: "messages currently in queue",
	}

	points[1] = &point{
		Name:        fmt.Sprintf("%s_enqueued", q.Name),
		Type:        counter,
		Value:       q.Enqueued,
		Description: "total messages enqueued",
	}

	points[2] = &point{
		Name:        fmt.Sprintf("%s_full", q.Name),
		Type:        counter,
		Value:       q.Full,
		Description: "times queue was full",
	}

	points[3] = &point{
		Name:        fmt.Sprintf("%s_discarded_full", q.Name),
		Type:        counter,
		Value:       q.DiscardedFull,
		Description: "messages discarded due to queue being full",
	}

	points[4] = &point{
		Name:        fmt.Sprintf("%s_discarded_not_full", q.Name),
		Type:        counter,
		Value:       q.DiscardedNf,
		Description: "messages discarded when queue not full",
	}

	points[5] = &point{
		Name:        fmt.Sprintf("%s_max_queue_size", q.Name),
		Type:        gauge,
		Value:       q.MaxQsize,
		Description: "maximum size queue has reached",
	}

	return points
}
