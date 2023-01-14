// Package errors provides main function on errors files
package errors

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	// Test error input
	err := fmt.Errorf("test error")
	ret := Check(err)
	if ret != 1 {
		t.Errorf("Expected return value to be 1, got %d", ret)
	}

	// Test non-error input
	ret = Check(nil)
	if ret != 0 {
		t.Errorf("Expected return value to be 0, got %d", ret)
	}
}
