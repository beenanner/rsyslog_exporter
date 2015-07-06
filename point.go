package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type pointType int

const (
	counter pointType = iota
	gauge
)

type point struct {
	Name        string
	Description string
	Type        pointType
	Value       int64
}

func (s *point) add(newPoint *point) error {
	switch s.Type {
	case gauge:
		if newPoint.Type != gauge {
			return fmt.Errorf("incompatible point type")
		}
		s.Value = newPoint.Value
	case counter:
		if newPoint.Type != counter {
			return fmt.Errorf("incompatible point type")
		}
		s.Value = s.Value + newPoint.Value
	}
	return nil
}

func (p *point) promDescription() *prometheus.Desc {
	return prometheus.NewDesc(
		prometheus.BuildFQName("", "rsyslog", p.Name),
		p.Description,
		nil, nil,
	)
}

func (p *point) promType() prometheus.ValueType {
	if p.Type == counter {
		return prometheus.CounterValue
	}
	return prometheus.GaugeValue
}

func (p *point) promValue() float64 {
	return float64(p.Value)
}
