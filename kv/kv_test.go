package kv_test

import (
	"testing"

	"github.com/ezebunandu/kv"
)

func TestGetReturnsNotOKIfKeyDoesNotExist(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
	_, ok := s.Get("key")
	if ok {
		t.Error("unexpected okay")
	}
}

func TestGetReturnsValueAndOKIfKeyExists(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
	s.Set("key", "value")
	v, ok := s.Get("key")
	if !ok {
		t.Error("not okay")
	}
	if v != "value" {
		t.Errorf("expected 'value' got %q", v)
	}
}

func TestSetUpdatesExistingKeyToNewValue(t *testing.T) {
	t.Parallel()
	s, err := kv.OpenStore("dummy path")
	if err != nil {
		t.Fatal(err)
	}
	s.Set("key", "value")
	s.Set("key", "new value")
	v, ok := s.Get("key")
	if !ok {
		t.Error("not okay")
	}
	if v != "new value" {
		t.Errorf("expected 'new value' got %v", v)
	}
}
