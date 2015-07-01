package soac

import (
	"testing"
)

func TestNewChanger(t *testing.T) {
	actual := NewChanger()
	expected := &Changer{39, 0, 49}

	if actual.fg != expected.fg {
		t.Errorf("got %v\nwant %v", actual.fg, expected.fg)
	}
	if actual.attr != expected.attr {
		t.Errorf("got %v\nwant %v", actual.attr, expected.attr)
	}
	if actual.bg != expected.bg {
		t.Errorf("got %v\nwant %v", actual.bg, expected.bg)
	}
}
