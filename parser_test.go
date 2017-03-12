package erowidparser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListExperiences(t *testing.T) {
	files, err := listExperiences()
	assert.NoError(t, err)
	assert.True(t, len(files) > 0)
}

func TestParseExperiences(t *testing.T) {
	files, err := listExperiences()
	assert.NoError(t, err)
	assert.True(t, len(files) > 0)

	file := files[0]
	_, err = parseExperience(file.Name())
	assert.NoError(t, err)
}

func TestRandExperiences(t *testing.T) {
	exps, err := randExperiences()
	assert.NoError(t, err)
	for _, exp := range exps {
		fmt.Println(exp.Name())
	}
}
