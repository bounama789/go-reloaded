package main

import "testing"

type Test struct {
	input    string
	expected string
}

var tests = []Test{
	{
		input:    "1E (hex) files were added",
		expected: "30 files were added",
	},
	{
		input:    "It has been 10 (bin) years",
		expected: "It has been 2 years",
	},
	{
		input:    "Ready, set, go (up) !",
		expected: "Ready, set, GO!",
	},
	{
		input:    "I should stop SHOUTING (low)",
		expected: "I should stop shouting",
	},
	{
		input:    "Welcome to the Brooklyn bridge (cap)",
		expected: "Welcome to the Brooklyn Bridge",
	},
	{
		input:    "This is so exciting (up, 2)",
		expected: "This is SO EXCITING",
	},
	{
		input:    "I was sitting over there ,and then BAMM !!",
		expected: "I was sitting over there, and then BAMM!!",
	},
	{
		input:    "I was thinking ... You were right",
		expected: "I was thinking... You were right",
	},
	{
		input:    "I am exactly how they describe me: ' awesome '",
		expected: "I am exactly how they describe me: 'awesome'",
	},
	{
		input:    "As Elton John said: ' I am the most well-known homosexual in the world '",
		expected: "As Elton John said: 'I am the most well-known homosexual in the world'",
	},
	{
		input:    "There it was. A amazing rock!",
		expected: "There it was. An amazing rock!",
	},
	{
		input:    "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
		expected: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
	},
	{
		input:    "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
		expected: "Simply add 66 and 2 and you will see the result is 68.",
	}, {
		input:    "There is no greater agony than bearing a untold story inside you.",
		expected: "There is no greater agony than bearing an untold story inside you.",
	},
	{
		input:    "Punctuation tests are ... kinda boring ,don't you think !?",
		expected: "Punctuation tests are... kinda boring, don't you think!?",
	},
	{
		input:    "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
		expected: "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
	}, {
		input:    "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
		expected: "I have to pack 5 outfits. Packed 26 just to be sure",
	},
	{
		input:    "Don not be sad ,because sad backwards is das . And das not good",
		expected: "Don not be sad, because sad backwards is das. And das not good",
	},
	{
		input:    "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
		expected: "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
	},
	{
		input:    "I am exactly how they describe me: ' awesome '",
		expected: "I am exactly how they describe me: 'awesome'",
	},
	{
		input:    "(cap),(up)(low)abcd(up)",
		expected: ", ABCD",
	},
	{
		input:    "'I am exactly ' 'h'ow they describe me': ' awesome '",
		expected: "'I am exactly' 'h'ow they describe me': 'awesome'",
	},
	{
		input:    "Don not be sad ,because sad (cap) (up, 2) backwards is das . And das not good",
		expected: "Don not be sad, BECAUSE SAD backwards is das. And das not good",
	},
	{
		input:    "(max(max(cap,2)))",
		expected: "(Max (Max))",
	},
	{
		input:    ",,,",
		expected: ",,,",
	},
}

func TestAll(t *testing.T) {
	for _, test := range tests {
		if output := process(test.input); output != test.expected {
			t.Errorf("Output %q not equal to \n expected %q", output, test.expected)
		}
	}
}
