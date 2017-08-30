package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func newQueueFromJSON(b []byte) (*queue, error) {
	dec := json.NewDecoder(bytes.NewReader(b))
	var pstat queue
	err := dec.Decode(&pstat)
	if err != nil {
		return nil, fmt.Errorf("failed to decode queue stat `%v`: %v", string(b), err)
	}
	return &pstat, nil
}

func (q *queue) toPoints() []*point {
	points := make([]*point, 6)

	points[0] = &point{
		Name:        "queue_size",
		Type:        gauge,
		Value:       q.Size,
		Description: "messages currently in queue",
		LabelValue:  q.Name,
	}

	points[1] = &point{
		Name:        "queue_enqueued",
		Type:        counter,
		Value:       q.Enqueued,
		Description: "total messages enqueued",
		LabelValue:  q.Name,
	}

	points[2] = &point{
		Name:        "queue_full",
		Type:        counter,
		Value:       q.Full,
		Description: "times queue was full",
		LabelValue:  q.Name,
	}

	points[3] = &point{
		Name:        "queue_discarded_full",
		Type:        counter,
		Value:       q.DiscardedFull,
		Description: "messages discarded due to queue being full",
		LabelValue:  q.Name,
	}

	points[4] = &point{
		Name:        "queue_discarded_not_full",
		Type:        counter,
		Value:       q.DiscardedNf,
		Description: "messages discarded when queue not full",
		LabelValue:  q.Name,
	}

	points[5] = &point{
		Name:        "queue_max_size",
		Type:        gauge,
		Value:       q.MaxQsize,
		Description: "maximum size queue has reached",
		LabelValue:  q.Name,
	}

	return points
}
