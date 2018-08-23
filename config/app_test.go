package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	assert.Equal(t, App.Common.AppId, "2016032101228633", "Common config unmarshal error")
	assert.NotEmpty(t, App.DB["default"], "Database config unmarshal error")
	assert.NotEmpty(t, App.Redis["default"], "Redis config unmarshal error")
	assert.NotEmpty(t, App.Nsq["default"], "Nsq config unmarshal error")
	assert.NotEmpty(t, App.Log, "Nsq config unmarshal error")
}
