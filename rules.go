package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Process applique toutes les règles, ligne par ligne,
// en conservant les retours à la ligne du texte d'origine.
func Process(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		line = convertTags(line)
		line = FixPunctuation(line)
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

// convertTags applique les balises d'une ligne puis les supprime.
// Le texte est transformé en liste de mots ; chaque balise agit sur les
// n derniers mots déjà présents dans la liste, puis disparaît.
func convertTags(line string) string {
	var words []string
	last := 0

	for _, loc := range tagRe.FindAllStringSubmatchIndex(line, -1) {
		// loc[0:2] = balise entière, loc[2:4] = nom, loc[4:6] = nombre (ou -1)
		words = append(words, strings.Fields(line[last:loc[0]])...)

		tag := line[loc[2]:loc[3]]
		n := 1
		if loc[4] != -1 {
			n, _ = strconv.Atoi(line[loc[4]:loc[5]])
		}
		applyTag(words, tag, n)
		last = loc[1]
	}
	words = append(words, strings.Fields(line[last:])...)

	return strings.Join(words, " ")
}

// applyTag modifie en place les n derniers mots de la liste.
// Si n dépasse le nombre de mots, tous les mots sont modifiés.
func applyTag(words []string, tag string, n int) {
	start := len(words) - n
	if start < 0 {
		start = 0
	}
	for i := start; i < len(words); i++ {
		switch tag {
		case "hex":
			words[i] = convertBase(words[i], 16)
		case "bin":
			words[i] = convertBase(words[i], 2)
		case "up":
			words[i] = strings.ToUpper(words[i])
		case "low":
			words[i] = strings.ToLower(words[i])
		case "cap":
			words[i] = capitalize(words[i])
		}
	}
}

// convertBase convertit un nombre écrit dans la base donnée en décimal.
// Si le mot n'est pas un nombre valide, il est laissé tel quel.
func convertBase(word string, base int) string {
	n, err := strconv.ParseInt(word, base, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(n, 10)
}

// capitalize met la première lettre en majuscule et le reste en minuscules.
func capitalize(word string) string {
	runes := []rune(strings.ToLower(word))
	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			break
		}
	}
	return string(runes)
}

// punctRe capture un groupe de ponctuation avec les espaces qui l'entourent.
var punctRe = regexp.MustCompile(`\s*([.,!?:;]+)\s*`)

// FixPunctuation colle la ponctuation au mot précédent
// et place un espace après le groupe.
func FixPunctuation(line string) string {
	line = punctRe.ReplaceAllString(line, "$1 ")
	return strings.TrimSpace(line)
}
