package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gosimple/slug"
)

type Node struct {
	Parent     *Node
	Children   []*Node
	Name       string
	Title      string
	Identifier string
	Path       string

	indent int
}

func (n *Node) String() string {
	var out bytes.Buffer

	out.WriteString(strings.Repeat("  ", n.indent) + n.Name)
	out.WriteString("\n")

	for _, v := range n.Children {
		out.WriteString(v.String())
	}

	return out.String()
}

func BuildContext(path string, root *Node) []*Node {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	if len(dirs) == 0 {
		log.Fatalf("Folder '%s' has no contents to process.", path)
	}

	children := make([]*Node, 0)

	for _, v := range dirs {
		ext := filepath.Ext(filepath.Join(path, v.Name()))
		if !v.IsDir() && ext != ".md" {
			continue
		}

		node := &Node{}
		node.Parent = root
		node.Children = nil
		node.Name = strings.Title(v.Name())
		node.Title = getTitleFromFileName(v)
		node.Identifier = slug.Make(v.Name())
		node.Path = getRelativePath(path, v)
		node.indent = root.indent + 1

		if v.IsDir() {
			node.Children = BuildContext(filepath.Join(path, v.Name()), node)
		}

		children = append(children, node)
	}

	return children
}

func main() {
	root := Node{}
	root.Parent = nil
	root.Name = "markdown"
	root.Identifier = "root"

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	markdownFolder := filepath.Join(cwd, "../", "markdown")
	root.Children = BuildContext(markdownFolder, &root)

	template, err := template.ParseFiles("README.template")
	if err != nil {
		panic(err)
	}
	file, err := os.Create("../README.md")
	if err != nil {
		panic(err)
	}
	err = template.Execute(file, root)
	if err != nil {
		panic(err)
	}
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

func getRelativePath(parentDirectory string, file os.FileInfo) string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	rootDir := filepath.Join(pwd, "../")
	fullPath := filepath.Join(parentDirectory, file.Name())
	return strings.ReplaceAll(fullPath, rootDir+"/", "")
}
