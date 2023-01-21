package common_test

import (
	"testing"

	"github.com/mhdiiilham/homies/common"
	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	assertion := assert.New(t)

	t.Run("success read config file", func(t *testing.T) {
		configEnv := "../config.test.yaml"

		expected := &common.Configuration{
			Env:  "test",
			Port: 8099,
			Database: common.DatabaseCredential{
				Dialect:  "postgre",
				Host:     "localhost",
				Port:     5432,
				Username: "postgre",
				Password: "root",
				Name:     "postgre",
			},
			ExternalClient: common.ExternalClient{
				Veryfi: common.Veryfi{
					Url:      "https://api.veryfi.com/api/v8/partner/documents",
					ClientId: "fake-client-id",
					Username: "fake-client-username",
					ApiKey:   "fake-client-api-key",
				},
			},
		}

		cfg, actualErr := common.ReadConfig(configEnv)
		assertion.Nil(actualErr)
		assertion.Equal(expected, cfg)
	})

	t.Run("failed read config file", func(t *testing.T) {
		configEnv := "../config.test-failed.yaml"

		cfg, actualErr := common.ReadConfig(configEnv)
		assertion.Nil(cfg)
		assertion.Equal(common.ErrConfigNotFound, actualErr)
	})
}

func TestConfigurationGetPort(t *testing.T) {
	cfg := common.Configuration{Port: 8099}
	expected := ":8099"

	assert.Equal(t, expected, cfg.GetPort())
}
