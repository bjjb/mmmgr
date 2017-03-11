package cfg

import (
	"github.com/spf13/viper"
	"testing"
)

func TestUnmarshalKey(t *testing.T) {
	type myType struct {
		Foo string `json:"foo"`
	}

	Cfg = &cfg{viper.New()}
	Cfg.SetDefault("k", map[string]string{"foo": "bar"})

	b := new(myType)
	err := UnmarshalKey("k", b)
	if err != nil {
		t.Error(err)
	}
	want := "bar"
	got := b.Foo
	if got != want {
		t.Errorf("UnmarshalKey(%q) => expected %q, got %q", "k", want, got)
	}
}
