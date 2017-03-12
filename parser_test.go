package erowidparser

import "testing"
import "github.com/stretchr/testify/assert"

func TestListExperiences(t *testing.T) {
	err := listExperiences()
	assert.NoError(t, err)
}
