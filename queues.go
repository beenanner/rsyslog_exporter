package main

import (
	"bytes"
	"encoding/json"
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
		Name:        "size",
		Type:        gauge,
		Value:       q.Size,
		Description: "messages currently in queue",
		LabelValue:  q.Name,
	}

	points[1] = &point{
		Name:        "enqueued",
		Type:        counter,
		Value:       q.Enqueued,
		Description: "total messages enqueued",
		LabelValue:  q.Name,
	}

	points[2] = &point{
		Name:        "full",
		Type:        counter,
		Value:       q.Full,
		Description: "times queue was full",
		LabelValue:  q.Name,
	}

	points[3] = &point{
		Name:        "discarded_full",
		Type:        counter,
		Value:       q.DiscardedFull,
		Description: "messages discarded due to queue being full",
		LabelValue:  q.Name,
	}

	points[4] = &point{
		Name:        "discarded_not_full",
		Type:        counter,
		Value:       q.DiscardedNf,
		Description: "messages discarded when queue not full",
		LabelValue:  q.Name,
	}

	points[5] = &point{
		Name:        "max_queue_size",
		Type:        gauge,
		Value:       q.MaxQsize,
		Description: "maximum size queue has reached",
		LabelValue:  q.Name,
	}

	return points
}
