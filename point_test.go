package main

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestAddCounter(t *testing.T) {
	p1 := &point{
		Name:  "my counter",
		Type:  counter,
		Value: int64(10),
	}

	p2 := &point{
		Name:  "my counter",
		Type:  counter,
		Value: int64(5),
	}

	err := p1.add(p2)
	if err != nil {
		t.Error(err)
	}

	if expect := int64(15); p1.Value != expect {
		t.Errorf("expected '%d', got '%d'", expect, p1.Value)
	}

	p3 := &point{
		Name:  "my gauge",
		Type:  gauge,
		Value: int64(10),
	}

	err = p1.add(p3)

	if err != ErrIncompatiblePointType {
		t.Errorf("incompatible point types should raise error")
	}

	if want, got := float64(15), p1.promValue(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := prometheus.ValueType(1), p1.promType(); want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}

	wanted := `Desc{fqName: "rsyslog_my counter", help: "", constLabels: {}, variableLabels: []}`
	if want, got := wanted, p1.promDescription().String(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}
}

func TestAddGauge(t *testing.T) {
	p1 := &point{
		Name:  "my gauge",
		Type:  gauge,
		Value: int64(10),
	}

	p2 := &point{
		Name:  "my gauge",
		Type:  gauge,
		Value: int64(5),
	}

	err := p1.add(p2)
	if err != nil {
		t.Error(err)
	}

	if want, got := int64(5), p1.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	p3 := &point{
		Name:  "my counter",
		Type:  counter,
		Value: int64(10),
	}

	err = p1.add(p3)
	if err != ErrIncompatiblePointType {
		t.Errorf("incompatible point types should raise error")
	}

	if want, got := float64(5), p1.promValue(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	if want, got := prometheus.ValueType(2), p1.promType(); want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}

	wanted := `Desc{fqName: "rsyslog_my gauge", help: "", constLabels: {}, variableLabels: []}`
	if want, got := wanted, p1.promDescription().String(); want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

}

func TestAddNotHandled(t *testing.T) {
	p1 := &point{
		Name:  "my gauge",
		Type:  gauge,
		Value: int64(10),
	}

	p2 := &point{
		Name:  "bad",
		Type:  99,
		Value: int64(5),
	}

	err := p1.add(p2)
	if err != ErrUnknownPointType {
		t.Errorf("incompatible point types should raise error")
	}
}
