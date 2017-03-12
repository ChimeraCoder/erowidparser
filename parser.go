package erowidparser

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var rnd = rand.New(rand.NewSource(time.Now().Unix()))

var rel_path = filepath.Join("e", "http80", "www.erowid.org", "experiences")

var root string

var experienceRegexp = regexp.MustCompile(`exp\.phpquery=ID=\d*.html`)

func listExperiences() ([]os.FileInfo, error) {
	var experiences []os.FileInfo
	files, err := ioutil.ReadDir(filepath.Join(root, rel_path))
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if experienceRegexp.MatchString(f.Name()) {
			experiences = append(experiences, f)
		}
	}
	return experiences, err
}

func parseExperience(filename string) (string, error) {
	const startPrefix = `<!--   Start Body   -->`
	const endSuffix = `Exp Year`

	f, err := os.Open(filepath.Join(root, rel_path, filename))
	if err != nil {
		return "", err
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		return "", err
	}

	selection := doc.Find(".report-text-surround")
	html, err := selection.Html()
	if err != nil {
		return "", err
	}

	i := strings.Index(html, startPrefix)
	if i > 0 {
		html = html[i+len(startPrefix):]
	}

	i = strings.Index(html, endSuffix)
	if i > 0 {
		html = html[:i]
	}

	doc, err = goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}

	return doc.Text(), nil
}

func randExperiences() ([]os.FileInfo, error) {
	experiences, err := listExperiences()
	if err != nil {
		return nil, err
	}

	result := make([]os.FileInfo, len(experiences))
	list := rnd.Perm(len(experiences))
	for i, experience := range experiences {
		result[list[i]] = experience
	}
	return result, nil
}
