package erowidparser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	_, err = ParseExperience(file.Name())
	assert.NoError(t, err)
}

func TestRandExperiences(t *testing.T) {
	exps, err := RandExperiences()
	assert.NoError(t, err)
	for _, exp := range exps {
		fmt.Println(exp.Name())
	}
}

func TestConcatenateAll(t *testing.T) {
	f, err := os.OpenFile(filepath.Join(Root, "all_experiences.txt"), os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	exps, err := listExperiences()
	if err != nil {
		panic(err)
	}

	for _, exp := range exps {
		log.Printf("Parsing %s", exp.Name())
		text, err := ParseExperience(exp.Name())
		assert.NoError(t, err)
		if err != nil {
			panic(err)
		}
		_, err = f.Write([]byte(text))
		assert.NoError(t, err)
		if err != nil {
			panic(err)
		}
	}
}
