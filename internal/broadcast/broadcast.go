// Package broadcast implements broadcast type including
// the reading of the broadcast.yaml, calling transformations,
// and parsing the template.

package broadcast

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"path"

	"github.com/mahgoh/bm-ws/broadcasts/internal/transformer"
	"gopkg.in/yaml.v3"
)

// Broadcast is the representation of a broadcast that is
// being processed to a HTML output.
type Broadcast struct {
	Version  string
	Theme    string // has to exist in dir ./themes
	Title    string
	Subtitle string
	Header   struct {
		Headline string // no-escape
	}
	Footer struct {
		Signature string // no-escape
	}
	Topics []*Topic
}

// Topic is the representation of a text entry inside a
// broadcast, similar to an article.
type Topic struct {
	Heading string
	Content string // no-escape, transformed by transformer
}

// NewBroadcast reads the broadcast.yaml in the specified
// source directory and creates a new Broadcast element.
func NewBroadcast(src string) *Broadcast {
	filePath := path.Join(src, "broadcast.yaml")
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var broadcast *Broadcast

	if err = yaml.Unmarshal(file, &broadcast); err != nil {
		log.Fatal(err)
	}

	return broadcast
}

// Transform executes transformations on the headline,
// signature, and content of each topic.
func (b *Broadcast) Transform() {

	b.Header.Headline = transformer.Transform(b.Header.Headline)
	b.Footer.Signature = transformer.Transform(b.Footer.Signature)

	for _, t := range b.Topics {
		t.Content = transformer.Transform(t.Content)
	}
}

// Parse loads and parses the specified template
// returning a bytes buffer of the HTML output
func (b *Broadcast) Parse() *bytes.Buffer {
	funcMap := template.FuncMap{
		"increment": func(i int) int {
			return i + 1
		},
		"noEscape": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	tmplPath := path.Join("themes", fmt.Sprintf("%s.tmpl.html", b.Theme))
	tmpl := template.Must(template.New(path.Base(tmplPath)).Funcs(funcMap).ParseFiles(tmplPath))

	buf := &bytes.Buffer{}

	if err := tmpl.Execute(buf, b); err != nil {
		log.Fatal(err)
	}

	return buf
}
