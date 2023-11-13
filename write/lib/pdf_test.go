package lib

import (
	"path"
	"runtime"
	"testing"
)

func TestPDF(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	testFile := path.Dir(filename) + "/test.pdf"
	t.Logf("Current test filename: %s", testFile)

	err := Pdf(testFile, `
# Hello world

Sup bro! I'll see you right, this sweet Undie 500 is as flat stick as a beaut seabed. Mean while, in the pub, Jonah Lomu and Fred Dagg were up to no good with a bunch of snarky utes. The chronic force of his packing a sad was on par with James and the Giant Peach's pretty suss foreshore and seabed issue.

Put the jug on will you bro, all these carked it pavlovas can wait till later. The first prize for reffing the game goes to... Jim Hickey and his good as Jafa, what a hottie. Bro, gumboots are really pearler good with hard yakka jelly tip icecreams, aye. You have no idea how nuclear-free our buzzy pohutukawa trees were aye. Every time I see those sweet as hangis it's like the op shop all over again aye, cook your own eggs Jake. Anyway, James Cook is just Cardigan Bay in disguise, to find the true meaning of life, one must start wobbling with the Tui, mate.

After the Monopoly, the New Zealand version with Queen Street and stuff is skived off, you add all the cracker marmite shortages to the vivid you've got yourself a meal. Technology has allowed stuffed cuzzies to participate in the global conversation of outrageously awesome pieces of pounamu. The next Generation of epic ankle biters have already flogged over at Lake Taupo. What's the hurry a Taniwha? There's plenty of kiwiburgers in South Pacific. The fish n' chip shop holds the most rip-off community in the country..
	`)
	if err != nil {
		t.Fatal("Error creating PDF", err)
	}
}
