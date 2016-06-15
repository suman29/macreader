package macreader

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReaderForMac(t *testing.T) {
	testFile := bytes.NewBufferString("a,b,c\r1,2,3\r").Bytes()

	r2 := csv.NewReader(New(bytes.NewReader(testFile)))
	lines2, err := r2.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	expectedOutput := [][]string{[]string{"a", "b", "c"}, []string{"1", "2", "3"}}
	assert.Equal(t, expectedOutput, lines2)
}

func TestReaderForUnix(t *testing.T) {
	testFile := bytes.NewBufferString("a,b,c\n1,2,3\n").Bytes()

	r2 := csv.NewReader(New(bytes.NewReader(testFile)))
	lines2, err := r2.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	expectedOutput := [][]string{[]string{"a", "b", "c"}, []string{"1", "2", "3"}}
	assert.Equal(t, expectedOutput, lines2)
}

func TestReaderForWindows(t *testing.T) {
	testFile := bytes.NewBufferString("a,b,c\r\n1,2,3\r\n").Bytes()

	r2 := csv.NewReader(New(bytes.NewReader(testFile)))
	lines2, err := r2.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	expectedOutput := [][]string{[]string{"a", "b", "c"}, []string{"1", "2", "3"}}
	assert.Equal(t, expectedOutput, lines2)
}

