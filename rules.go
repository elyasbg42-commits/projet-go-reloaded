package main

import (
	"regexp"
	"strings"
)

//Transformer le texte en une liste de mots, avec strings.Fields.
//Process applique toutes les règles, ligne par ligne, en conservant les retours à la ligne du texte d'origine.
//Traiter une balise comme un élèment à part de la liste qui agit sur ce qui la précède puis dispartaît.

func Process(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		//line = convertTags(line)    // etape 3 et 4
		//line = fixPunctuation(line) // etape 6
		//line = fixQuotes(line)      // etape 6
		//line = fixArticles(line)    // etape 5
		lines[i] = line
	}
	return strings.Join(lines, "\n")
}

// tokenRe reconnait soit une balise complète, soit un mot (suite de non-espaces).
// utiliser strconv.Parseint puis strconv.Atoi.
// remplancer le mot par sa valeur et supprime la balise.
var tokenRe = regexp.MustCompile(`\((?:hex|bin)\)|\((?:up|low|cap)(?::\s*,\s*\d+)?\)|[^\s]+`)
var tagRe = regexp.MustCompile(`^\((hex|bin|up|low|cap)(?::\s*,\s*(\d+))?\)$`)
