package config_test

import (
	"github.com/atiwoto1/thims/internal/config"
	"testing"
)

func TestConfig_GetEnv(t *testing.T) {
	c := config.NewConfig(getEnv)
	testCases := []struct {
		name, key          string
		fallback, expected string
	}{
		{name: "should default value if not set", key: "env", fallback: "dev", expected: "dev"},
		{name: "when key is set, don't use fallback", key: "port", fallback: "5000", expected: "8080"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			v := c.GetEnvOrDefault(tc.key, tc.fallback)
			if v != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, v)
			}
		})
	}
}

func getEnv(key string) string {
	if key == "PORT" {
		return "8080"
	}
	return ""
}
