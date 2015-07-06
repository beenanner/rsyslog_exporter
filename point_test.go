package main

import "testing"

func TestaddCounter(t *testing.T) {
	s1 := &point{
		Name:  "my counter",
		Type:  counter,
		Value: int64(10),
	}

	s2 := &point{
		Name:  "my counter",
		Type:  counter,
		Value: int64(5),
	}

	err := s1.add(s2)
	if err != nil {
		t.Error(err)
	}

	if expect := int64(15); s1.Value != expect {
		t.Errorf("expected '%d', got '%d'", expect, s1.Value)
	}

	s3 := &point{
		Name:  "my gauge",
		Type:  gauge,
		Value: int64(10),
	}

	err = s1.add(s3)
	if err == nil {
		t.Errorf("incompatible point types should raise error")
	}
}

func TestaddGauge(t *testing.T) {
	s1 := &point{
		Name:  "my gauge",
		Type:  gauge,
		Value: int64(10),
	}

	s2 := &point{
		Name:  "my gauge",
		Type:  gauge,
		Value: int64(5),
	}

	err := s1.add(s2)
	if err != nil {
		t.Error(err)
	}

	if want, got := int64(5), s1.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	s3 := &point{
		Name:  "my counter",
		Type:  counter,
		Value: int64(10),
	}

	err = s1.add(s3)
	if err == nil {
		t.Errorf("incompatible point types should raise error")
	}
}
