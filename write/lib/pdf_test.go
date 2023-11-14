package lib

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func TestPDF(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	testDir := path.Dir(filename) + "/test/"
	testFile := testDir + "test.pdf"
	testInput := testDir + "test.md"
	t.Logf("Current test filename: %s", testFile)
	testContent, err := os.ReadFile(testInput)
	if err != nil {
		t.Fatal("Error reading test file", err)
	}

	err = Pdf(testFile, string(testContent))
	if err != nil {
		t.Fatal("Error creating PDF", err)
	}
}

func TestReadmePDF(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	testDir := path.Dir(filename) + "/../../"
	testFile := testDir + "README.pdf"
	testInput := testDir + "README.md"
	t.Logf("Current test filename: %s", testFile)
	testContent, err := os.ReadFile(testInput)
	if err != nil {
		t.Fatal("Error reading test file", err)
	}

	err = Pdf(testFile, string(testContent))
	if err != nil {
		t.Fatal("Error creating PDF", err)
	}
}
