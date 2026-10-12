package main

import "testing"

type testcase struct {
	name string
	in   string
	want string
}

// RunCases exécute une série de cas sur une fonction texte.
func RunCases(t *testing.T, f func(string) string, cases []testcase) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := f(c.in); got != c.want {
				t.Errorf("\nentrée : %q\nattendu : %q\nobtenu : %q", c.in, c.want, got)
			}
		})
	}
}

func TestConvertTags(t *testing.T) {
	RunCases(t, convertTags, []testcase{
		{"hex", "1E (hex) files were added", "30 files were added"},
		{"bin", "10 (bin) years", "2 years"},
		{"hex invalide", "zz (hex) test", "zz test"},
		{"up", "Ready, set, go (up) !", "Ready, set, GO !"},
		{"low", "i should stop SHOUTING (low)", "i should stop shouting"},
		{"cap", "Welcom to the Brooklyn bridge (cap)", "Welcom to the Brooklyn Bridge"},
		{"up avec nombre", "This is so exciting (up, 2)", "This is SO EXCITING"},
		{"nombre trop grand", "so exciting (up, 10)", "SO EXCITING"},
		{"cap avec nombre", "one two three (cap, 3)", "One Two Three"},
		{"balise en début", "(up) hello", "hello"},
		{"pas d'espace double", "a b (up)", "a B"},
		{"balise collée au mot", "42(hex) and 10(bin)and", "66 and 2 and"},
		{"balise invalide conservée", "THE(law, 3) winter", "THE(law, 3) winter"},
	})
}

func TestFixPunctuation(t *testing.T) {
	RunCases(t, FixPunctuation, []testcase{
		{"simple", "I was sitting over there, and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"points de suspension", "I was thinking...you were right", "I was thinking... you were right"},
		{"mélange", "Punctuation test are...kinda boring, what do you think ?", "Punctuation test are... kinda boring, what do you think?"},
		{"groupe !?", "Really!?Yes", "Really!? Yes"},
		{"fin de ligne", "Hello world.", "Hello world."},
	})
}

func TestFixQuotes(t *testing.T) {
	RunCases(t, FixQuotes, []testcase{
		{"un mot", "I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"plusieurs mots", "As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
		{"deux paires", "He said ' hi ' and she said ' bye '", "He said 'hi' and she said 'bye'"},
		{"déjà collées", "I am 'awesome'", "I am 'awesome'"},
		{"contraction", "I don't know", "I don't know"},
	})
}

func TestFixArticles(t *testing.T) {
	RunCases(t, FixArticles, []testcase{
		{"majuscule", "There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"minuscule", "bearing a untold story inside you.", "bearing an untold story inside you."},
		{"h", "a hour ago", "an hour ago"},
		{"consonne", "a good day", "a good day"},
		{"mot contenant a", "banana apple", "banana apple"},
		{"dernier mot", "there is a", "there is a"},
	})
}

func TestProcess(t *testing.T) {
	RunCases(t, Process, []testcase{
		{
			name: "exemple Dickens",
			in:   "it (cap) was the best of times, it was the worst of times (up), it was the age of wisdom, it was the age of foolishness (cap,6),it was the epoch of belief,it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, It Was THE(law, 3) winter of despair.",
			want: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, It Was THE(law, 3) winter of despair.",
		},
		{
			name: "hex et bin",
			in:   "Simply add 42(hex) and 10(bin)and you will see the result is 68",
			want: "Simply add 66 and 2 and you will see the result is 68",
		},
		{
			name: "a/an",
			in:   "There is no greater agony than bearing a untold story inside you.",
			want: "There is no greater agony than bearing an untold story inside you.",
		},
		{
			name: "ponctuation",
			in:   "Punctuation test are...kinda boring, what do you think ?",
			want: "Punctuation test are... kinda boring, what do you think?",
		},
		{
			name: "apostrophes",
			in:   "As Elton John said: ' I am the most well-known homosexual in the world '",
			want: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
		{
			name: "plusieurs lignes",
			in:   "hello (up)\nworld, ok\n",
			want: "HELLO\nworld, ok\n",
		},
		{
			name: "texte vide",
			in:   "",
			want: "",
		},
	})
}
