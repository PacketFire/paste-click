package service

import (
	"os"
	"reflect"
	"testing"
)

var (
	wantedConfig = Config{
		Addr:          ":8080",
		Logging:       true,
		StorageDriver: "fs",
		SiteName:      "paste.click",
	}
)

func TestConfigInit(t *testing.T) {
	t.Run("ServerConfig parses environment variables", func(t *testing.T) {
		os.Setenv("ADDR", ":8080")
		defer os.Unsetenv("ADDR")

		conf := New()
		if !reflect.DeepEqual(conf, wantedConfig) {
			t.Errorf("Config is not initializing correctly. got %v want %v", conf, wantedConfig)
		}
	})
}
