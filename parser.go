package erowidparser

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

var rel_path = filepath.Join("e", "http80", "www.erowid.org", "experiences")

var root string

var experienceRegexp = regexp.MustCompile(`exp\.phpquery=ID=\d*.html`)

func listExperiences() error {
	files, err := ioutil.ReadDir(filepath.Join(root, rel_path))
	if err != nil {
		return err
	}
	for _, f := range files {
		if experienceRegexp.MatchString(f.Name()) {
			fmt.Println(f.Name())
		}
	}
	return err
}
