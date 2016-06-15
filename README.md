**Reading CSV file independent of the platform.**

#Usage

```go
func Example() {
	// testFile is a CSV file with CR line endings.
	testFile := bytes.NewBufferString("a,b,c\r1,2,3\r").Bytes()

	// First try reading the csv file the normal way.
	// The CSV reader doesn't recognize the '\r' line ending.
	r1 := csv.NewReader(bytes.NewReader(testFile))
	lines1, err := r1.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Without macreader: %#v\n", lines1)
	// Output: Without macreader: [][]string{[]string{"a", "b", "c\r1", "2", "3"}}

	// Now try reading the csv file using macreader.
	// It should work as expected.
	r2 := csv.NewReader(New(bytes.NewReader(testFile)))
	lines2, err := r2.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("With macreader: %#v\n", lines2)
	// With macreader: [][]string{[]string{"a", "b", "c"}, []string{"1", "2", "3"}}
}
```