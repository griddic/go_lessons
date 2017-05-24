package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func Test_Byte(t *testing.T) {
	assert.Equal(t, "5B", humanize(5))
}