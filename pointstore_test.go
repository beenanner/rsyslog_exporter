package main

import "testing"

func TestPointStore(t *testing.T) {
	ps := newPointStore()

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

	err := ps.add(s1)
	if err != nil {
		t.Error(err)
	}

	got, err := ps.get(s1.Name)
	if err != nil {
		t.Error(err)
	}

	if want, got := int64(10), got.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	err = ps.add(s2)
	if err != nil {
		t.Error(err)
	}

	got, err = ps.get(s2.Name)
	if err != nil {
		t.Error(err)
	}

	if want, got := int64(15), got.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	s3 := &point{
		Name:  "my gauge",
		Type:  gauge,
		Value: int64(20),
	}

	err = ps.add(s3)
	if err != nil {
		t.Error(err)
	}

	got, err = ps.get(s3.Name)
	if err != nil {
		t.Error(err)
	}

	if want, got := int64(20), got.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}

	s4 := &point{
		Name:  "my gauge",
		Type:  gauge,
		Value: int64(15),
	}

	err = ps.add(s4)
	if err != nil {
		t.Error(err)
	}

	got, err = ps.get(s4.Name)
	if err != nil {
		t.Error(err)
	}

	if want, got := int64(15), got.Value; want != got {
		t.Errorf("want '%d', got '%d'", want, got)
	}
}
