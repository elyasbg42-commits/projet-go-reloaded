package main

import (
	"regexp"
	"strings"
)

// Process applique toutes les règles, ligne par ligne,
// en conservant les retours à la ligne du texte d'origine.
func Process(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		//line = convertTags(line)
		//line = FixPunctuation(line)
		//line = FixQuotes(line)
		//line = FixArticles(line)
		lines[i] = line
	}
	return strings.Join(lines, "\n")
}

// tagRe reconnaît une balise valide : (hex) (bin) (up) (low) (cap)
// avec, éventuellement, un nombre : (up, 2) ou (cap,6).
// Elle peut être collée à un mot : 42(hex) ou (cap,6),it
var tagRe = regexp.MustCompile(`\(\s*(hex|bin|up|low|cap)\s*(?:,\s*(\d+)\s*)?\)`)
