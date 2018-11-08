package novel

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"unicode"
)

var (
	base = [][]byte{
		[]byte("SoMeThInG"),
		[]byte("DATA"),
	}
)

// DataStore manages data.
type DataStore struct {
	ls [][]byte
}

// NewDataStore creates a new DataStore with lines as it's initial contents.
func NewDataStore(lines ...string) (*DataStore, error) {
	ls := make([][]byte, len(lines)+len(base))

	for i, v := range base {
		ls[i] = v
	}

	var err error

	for i, line := range lines {
		for _, r := range line {
			if !unicode.IsLetter(r) {
				err = &BadValueError{
					string(r),
					"must only contain letters",
				}
				break
			}
		}

		if err != nil {
			break
		}

		ls[i+len(base)] = []byte(line)
	}

	ds := DataStore{
		ls: ls,
	}

	if err != nil {
		err = fmt.Errorf("cannot create DataStore properly: %s", err)
	}

	return &ds, err
}

// Lowercased returns the current DataStore data lowercased.
func (ds *DataStore) Lowercased() []byte {
	var bs []byte

	for _, line := range ds.ls {
		bs = append(bs, bytes.ToLower(line)...)
		bs = append(bs, '\n')
	}

	return bs
}

// Report prints the current state of DataStore.
func (ds *DataStore) Report() {
	f, err := os.Create("datastore.dat")
	if err != nil {
		panic(err)
	}
	defer f.Close() //nolint

	for _, line := range ds.ls {
		if _, err = f.Write(line); err != nil {
			panic(err)
		}

		if _, err = f.WriteString("\n"); err != nil {
			panic(err)
		}
	}
}

// RollBack will rollback the managed data to the base value.
func (ds *DataStore) RollBack() {
	ds.ls = ds.ls[:len(base)]
}

// OutFile wraps a file.
type OutFile struct {
	*os.File
}

// NewOutFile creates a new OutFile with file created as it's initial contents.
func NewOutFile(file string) (*OutFile, error) {
	f, err := os.Create(file)
	if err != nil {
		return nil, fmt.Errorf("cannot create OutFile: %s", err)
	}

	of := OutFile{
		File: f,
	}

	return &of, nil
}

func (of *OutFile) Write(p []byte) (int, error) {
	r := bytes.NewReader(p)
	sc := bufio.NewScanner(r)
	var n int

	for sc.Scan() {
		l, err := of.File.Write([]byte("output: "))
		n += l
		if err != nil {
			return n, err
		}

		l, err = of.File.Write(sc.Bytes())
		n += l
		if err != nil {
			return n, err
		}

		l, err = of.File.Write([]byte("\n"))
		n += l
		if err != nil {
			return n, err
		}
	}

	return n, nil
}
