package soac

import (
	"reflect"
	"testing"
)

func TestNewParam(t *testing.T) {
	actual := NewParam()
	expected := Param{39, 0, 49}
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

func TestNewChanger(t *testing.T) {
	actual := NewChanger()
	p := NewParam()
	expected := Changer{"", &p}

	if actual.val != expected.val {
		t.Errorf("got %v\nwant %v", actual.val, expected.val)
	}
	if !reflect.DeepEqual(actual.p, expected.p) {
		t.Errorf("got %v\nwant %v", actual.p, expected.p)
	}
}
