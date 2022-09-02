// Package transformer implements transformation rules that
// allow text formatting.

package transformer

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	Rules []*Rule
)

// Rule is the representation of a transformation rule that
// consists of a regular expression to match occurences and
// a replace function that is executed if a match is found.
type Rule struct {
	Pattern *regexp.Regexp        // regular expression to match
	Replace func([]string) string // replacer function that is called for each match (match is provided as parameter)
}

// Register rules - patterns and their replacer function
// that are executed when a match is found.
//
// Rules:
// - bold:
//     pattern: **example**
//		 transformed: <b>example</b>
// - italic:
//		 pattern: _example_
//     transformed: <i>example</i>
// - link:
//     pattern: [example](https://example.com)
//     transformed: <a href="https://example.com" target="_blank">example</a>
// - code:
// 		 pattern: `example`
//		 transformed: <code>example</code>
//
func init() {
	Rules = []*Rule{
		{
			Pattern: regexp.MustCompile(`\*{2}([^\*]+)\*{2}`),
			Replace: func(m []string) string {
				return fmt.Sprintf("<b>%s</b>", m[1])
			},
		},
		{
			Pattern: regexp.MustCompile(`_{1}([^_]+)_{1}`),
			Replace: func(m []string) string {
				return fmt.Sprintf("<i>%s</i>", m[1])
			},
		},
		{
			Pattern: regexp.MustCompile(`\[([^\]]+)\]\(([^\)]+)\)`),
			Replace: func(m []string) string {
				return fmt.Sprintf("<a href=\"%s\" target=\"_blank\">%s</a>", m[2], m[1])
			},
		},
		{
			Pattern: regexp.MustCompile("`([^`]+)`"),
			Replace: func(m []string) string {
				return fmt.Sprintf("<code>%s</code>", m[1])
			},
		},
	}
}

// Transform runs the pattern of each rule on given string
// and executes the replacer function if a match is found.
func Transform(s string) string {
	for _, p := range Rules {
		matches := p.Pattern.FindAllStringSubmatch(s, -1)

		if matches != nil {
			for _, m := range matches {
				s = strings.ReplaceAll(s, m[0], p.Replace(m))
			}
		}
	}

	return s
}
