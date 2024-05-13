package count_test

import (
	"bytes"
	"testing"

	"github.com/ezebunandu/count"
)

func TestCountCountsLines(t *testing.T) {
	t.Parallel()
	buf := bytes.NewBufferString("1\n2\n3\n")
	c := count.NewCounter()
	c.Input = buf
	got := c.Lines()
	want := 3
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
