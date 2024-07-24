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

func TestSaveSavesDataPersistently(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/kvtest.store"
	s, err := kv.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	s.Set("A", "1")
	s.Set("B", "2")
	s.Set("C", "3")
	err = s.Save()
	if err != nil {
		t.Fatal(err)
	}
	s2, err := kv.OpenStore(path)
	if err != nil {
		t.Fatal(err)
	}
	if v, _ := s2.Get("A"); v != "1" {
		t.Errorf("want A=1, got %s", v)
	}
	if v, _ := s2.Get("B"); v != "2" {
		t.Errorf("want B=2, got %s", v)
	}
	if v, _ := s2.Get("C"); v != "3" {
		t.Errorf("want C=3, got %s", v)
	}
}
