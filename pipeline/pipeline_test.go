package pipeline_test

import (
	"bytes"
	"testing"

	"github.com/ezebunandu/pipeline"
	"github.com/google/go-cmp/cmp"
)

func TestStdoutPrintsMessageToOutput(t *testing.T) {
	t.Parallel()
	want := "Hello, world\n"
	p := pipeline.FromString(want)
	buf := new(bytes.Buffer)
	p.Output = buf
	p.Stdout()
	if p.Error != nil {
		t.Fatal(p.Error)
	}
	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}
