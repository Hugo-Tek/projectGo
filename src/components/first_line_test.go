package components

import (
	"bufio"
	"strings"
	"testing"
)

func TestParseFirstLine(t *testing.T) {
	// Test a valid input
	input := "10 20 30\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	width, height, turns, ret := ParseFirstLine(scanner)
	if ret != 0 {
		t.Errorf("Expected return value to be 0, got %d", ret)
	}
	if width != 10 {
		t.Errorf("Expected width to be 10, got %d", width)
	}
	if height != 20 {
		t.Errorf("Expected height to be 20, got %d", height)
	}
	if turns != 30 {
		t.Errorf("Expected turns to be 30, got %d", turns)
	}
}

func TestParseFirstLineInvalidLess3Parts(t *testing.T) {
	// Test invalid input with less than 3 elements
	input := "10 20\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	_, _, _, ret := ParseFirstLine(scanner)
	if ret == 0 {
		t.Error("Expected return value to be non-zero, got 0")
	}
}

func TestParseFirstLineInvalidMore3Parts(t *testing.T) {
	// Test invalid input with more than 3 elements
	input := "10 20 30 40\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	_, _, _, ret := ParseFirstLine(scanner)
	if ret == 0 {
		t.Error("Expected return value to be non-zero, got 0")
	}
}

func TestParseFirstLineWidthNonInteger(t *testing.T) {
	// Test invalid input where width is not an integer
	input := "a 20 30\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	_, _, _, ret := ParseFirstLine(scanner)
	if ret == 0 {
		t.Error("Expected return value to be non-zero, got 0")
	}
}

func TestParseFirstLineHeightNonInteger(t *testing.T) {
	// Test invalid input where height is not an integer
	input := "10 b 30\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	_, _, _, ret := ParseFirstLine(scanner)
	if ret == 0 {
		t.Error("Expected return value to be non-zero, got 0")
	}
}

func TestParseFirstLineTurnsNonInteger(t *testing.T) {
	// Test invalid input where turns is not an integer
	input := "10 20 c\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	_, _, _, ret := ParseFirstLine(scanner)
	if ret == 0 {
		t.Error("Expected return value to be non-zero, got 0")
	}
}
