package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	configs := make(map[string]Config)

	configs["default"] = Config{
		Server: ":6379",
	}

	Connect(configs)
	client := Client()
	assert.NotEqual(t, client, nil, "Redis connect error")
}