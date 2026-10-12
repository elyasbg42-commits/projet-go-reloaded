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
