package errors

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestScannerNoError(t *testing.T) {
	// Test scanner with no error
	scanner := bufio.NewScanner(strings.NewReader("test input"))
	ret := ScannerError(scanner)
	if ret != 0 {
		t.Errorf("Expected return value to be 0, got %d", ret)
	}
}

func TestScannerError(t *testing.T) {
	// Test scanner with error
	scanner := bufio.NewScanner(errorReader{})
	scanner.Scan()
	ret := ScannerError(scanner)
	if ret != 1 {
		t.Errorf("Expected return value to be 1, got %d", ret)
	}
}

type errorReader struct{}

func (errorReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("test error")
}
