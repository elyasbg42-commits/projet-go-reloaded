package main

import "testing"

type testcase struct {
	name string
	in string
	want string
}

// RunCase execute une série de cas sur une fonction texte
func RunCases(t *testing.T, f func(string)string, cases []testcase) {
	t.Helper()
	for _, c:= range cases {
		t.Run(c.name, func(t *testing.T)) {
			if got:= f(c.in); got != c.want{
				t.Errorf("\nentrée : %q\nattendu : %q", c.in, got, c.want)
			}
		}
	}
}

func TestProcess(t *testing.T) {
	RunCases(t, Process, []testCase) {
		{
		"exemple Dickens",
		"it (cap) was the best of times, it was the worst of times (up), it was the age of wisdom, it was the age of foolishness (cap,6),it was the epoch of belief,it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope  IT WAS THE(law, 3) winter of despair.",
		"it was the best of times, it was the worst of Times, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
	},
	{"hex et bin","Simply add 42(hex) and 10(bin)and you will see the result is 68", "Simply add 66 and 2 and you will see the result is 68"},
	{"a/an","There is no greater agony than bearing a untold story inside you.", "Ther is no greater agony than bearing an untol story inside you."}
	{"ponctuation","Punctuation test are...kinda boring, what do you think ?", "Ponctuatuation tets are...kinda boring, what do you thing ?"},
	{"plusieurs lignes", "hello (up)\nworld, ok\n","HELLO\nworld, ok\n"},
	{"texte vide", "", ""},
}
}
