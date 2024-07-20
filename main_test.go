package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	expected := "8000"

	actual := getEnv("PORT", "8000")

	assert.Equal(t, expected, actual)
}
