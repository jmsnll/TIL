package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gosimple/slug"
)

type Topic struct {
	Name       string
	Identifier string
	Entries    []Entry
}

type Entry struct {
	Title        string
	RelativePath string
}

func main() {
	template, err := template.ParseFiles("README.template")
	if err != nil {
		panic(err)
	}
	file, err := os.Create("../README.md")
	if err != nil {
		panic(err)
	}
	err = template.Execute(file, EnumerateTopics())
	if err != nil {
		panic(err)
	}
}

func EnumerateTopics() []Topic {
	topics := make([]Topic, 0)

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	markdownFolder := filepath.Join(cwd, "../", "markdown")

	dirs, err := ioutil.ReadDir(markdownFolder)
	if err != nil {
		panic(err)
	}

	if len(dirs) == 0 {
		panic("No entries found to use for generation.")
	}

	for _, v := range dirs {
		if ignoreProjectFolders(v) || !v.IsDir() {
			continue
		}

		topic := Topic{
			Name:       strings.Title(v.Name()),
			Identifier: slug.Make(v.Name()),
		}

		topicDir, err := ioutil.ReadDir(filepath.Join(markdownFolder, v.Name()))
		if err != nil {
			panic(err)
		}

		for _, t := range topicDir {
			extension := filepath.Ext(filepath.Join(markdownFolder, v.Name(), t.Name()))
			if extension == ".md" {
				entry := Entry{
					Title:        getTitleFromFileName(t),
					RelativePath: filepath.Join("markdown", v.Name(), t.Name()),
				}
				topic.Entries = append(topic.Entries, entry)
			}
		}
		topics = append(topics, topic)
	}
	return topics
}

func ignoreProjectFolders(file os.FileInfo) bool {
	prefixes := []string{".", "_"}

	for _, prefix := range prefixes {
		if strings.HasPrefix(file.Name(), prefix) {
			return true
		}
	}

	return false
}

func getTitleFromFileName(file os.FileInfo) string {
	name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
	name = strings.Replace(name, "-", " ", -1)
	return strings.Title(name)
}
