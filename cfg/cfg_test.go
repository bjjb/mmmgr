package cfg

import (
	"github.com/spf13/viper"
	"strings"
	"testing"
)

func TestUnmarshalKey(t *testing.T) {
	type myType struct {
		Foo string `json:"foo"`
	}

	Cfg = &cfg{viper.New()}
	Cfg.SetConfigType("json")
	json := `{"foo":{"quxes":[null,{"flobs":{"b":true}}]}}`
	if err := Cfg.Viper.ReadConfig(strings.NewReader(json)); err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", Cfg.AllSettings())

	foo := &struct {
		Quxes []struct {
			Flobs map[string]bool `json:"flobs"`
		} `json:"quxes"`
	}{}

	if err := UnmarshalKey("foo", foo); err != nil {
		t.Error(err)
	}
	if !foo.Quxes[1].Flobs["b"] {
		t.Errorf("expected true")
	}
}
