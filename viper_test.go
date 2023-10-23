package belajar_golang_viper

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)
}
