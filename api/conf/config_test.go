package conf_test

import (
	"blog/conf"
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {
	err := conf.LoadConfigFromToml("../etc/config.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.Conf())
}
