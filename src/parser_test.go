// Package src provides main function on src files
package src

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// func TestParser(t *testing.T) {
// 	tests := []struct {
// 		name       string
// 		file       *os.File
// 		scanner    *bufio.Scanner
// 		wantWidth  int
// 		wantHeight int
// 		wantTurns  int
// 		wantRet    int
// 	}{
// 		{
// 			name:       "CreatePackagePalletandTruck",
// 			file:       createTestFile("test1.txt", "5 5 10\nPackage 1 1 Green\nPallet 1 1\nTruck 3 4 4000 5\n"),
// 			wantWidth:  5,
// 			wantHeight: 5,
// 			wantTurns:  10,
// 			wantRet:    0,
// 		},
// 		{
// 			name:       "InvalidFirstLine",
// 			file:       createTestFile("test2.txt", "invalid input\n"),
// 			wantWidth:  0,
// 			wantHeight: 0,
// 			wantTurns:  0,
// 			wantRet:    1,
// 		},
// 		{
// 			name:       "InvalidNumberOfParts",
// 			file:       createTestFile("test3.txt", "5 5 10\nP 1 1 E\nT 1 1 S 2 2\nP 2 2 S\nT 3 3 S 3 3\n"),
// 			wantWidth:  5,
// 			wantHeight: 5,
// 			wantTurns:  10,
// 			wantRet:    1,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.scanner = bufio.NewScanner(tt.file)
// 			tt.scanner.Scan()
// 			gotWidth, gotHeight, gotTurns, gotRet := Parser(tt.file, tt.scanner)
// 			if gotWidth != tt.wantWidth || gotHeight != tt.wantHeight || gotTurns != tt.wantTurns || gotRet != tt.wantRet {
// 				t.Errorf("Parser() = (%d, %d, %d, %d), want (%d, %d, %d, %d)", gotWidth, gotHeight, gotTurns, gotRet, tt.wantWidth, tt.wantHeight, tt.wantTurns, tt.wantRet)
// 			}
// 		})
// 	}
// }

func TestParserError(t *testing.T) {
	file, _ := os.CreateTemp("", "test.txt")
	scanner := bufio.NewScanner(errorReader{})
	scanner.Scan()
	_, ret := Parser(file, scanner)
	if ret != 1 {
		t.Errorf("Expected return value to be 1, got %d", ret)
	}
}

type errorReader struct{}

func (errorReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("test error")
}

// func createTestFile(name, contents string) *os.File {
// 	tmpfile, err := ioutil.TempFile("", name)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if _, err := tmpfile.Write([]byte(contents)); err != nil {
// 		log.Fatal(err)
// 	}

// 	if _, err := tmpfile.Seek(0, 0); err != nil {
// 		log.Fatal(err)
// 	}
// 	return tmpfile
// }
