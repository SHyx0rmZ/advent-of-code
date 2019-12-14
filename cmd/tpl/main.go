package main

import (
	"os"
	"path/filepath"
	"text/template"
)

const tplProblem = `package {{.Package}}

import (
	"fmt"
	"io"
)

type problem struct {}

func Problem() *problem {
	return &problem{}
}
{{range .Methods}}
func (p problem) {{.}}WithReader(r io.Reader) (string, error) {
	_, err := p.parse(r)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", 0), nil
}
{{end}}
func (problem) parse(r io.Reader) ([]int, error) {
	var es []int
	return es, nil
}
`

const tplProblemTest = `package {{.Package}}_test

import (
	"strings"
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/{{.Package}}"
)
{{range .Methods}}
func TestProblem_{{.}}(t *testing.T) {
	r, err := {{$.Package}}.Problem().{{.}}WithReader(strings.NewReader(""))
	if r != "" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "", nil)
	}
}
{{end}}`

type vars struct {
	Package string
	Methods []string
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	path := os.Args[1]

	err := os.MkdirAll(path, 0755)
	if err != nil {
		panic(err)
	}

	vars := vars{
		filepath.Base(path),
		[]string{"PartOne", "PartTwo"},
	}

	err = create(filepath.Join(path, "problem.go"), tplProblem, vars)
	if err != nil {
		panic(err)
	}

	err = create(filepath.Join(path, "problem_test.go"), tplProblemTest, vars)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Join(path, "data"), 0755)
	if err != nil {
		panic(err)
	}

	err = create(filepath.Join(path, "data", "input.txt"), "", vars)
	if err != nil {
		panic(err)
	}
}

func create(path string, content string, vars vars) error {
	t, err := template.New(filepath.Base(path)).Parse(content)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, vars)
	if err != nil {
		return err
	}

	return nil
}
