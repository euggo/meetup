package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// BGN2 OMIT
type fileInfo struct {
	path string
	data string
	err  error
}

func newFileInfo(path string) *fileInfo { // HLargs
	// open file at path
	// decompress data while reading file
	// return a pointer to an instance of fileInfo
	// END2 OMIT

	f, err := os.Open(path)
	if err != nil {
		return &fileInfo{path: path, err: err}
	}
	defer f.Close() //nolint

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return &fileInfo{path: path, err: err}
	}
	defer gzr.Close() //nolint

	data, err := ioutil.ReadAll(gzr)
	if err != nil {
		return &fileInfo{path: path, err: err}
	}

	return &fileInfo{path: path, data: string(data)}
}

// BGN1 OMIT
func gzipFilePaths(dir string) ([]string, error) { // HLargs
	// read directory listing
	// iterate over listed files gathering paths of gzipped files
	// return slice of paths and a potential error
	// END1 OMIT

	wrapErr := func(err error) error {
		return fmt.Errorf("cannot get gzip paths: %s", err)
	}

	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, wrapErr(err)
	}

	for k := len(fis) - 1; k >= 0; k-- {
		if !isGzipFile(fis[k]) {
			fis = append(fis[:k], fis[k+1:]...) // remove from slice
		}
	}

	paths := make([]string, 0, len(fis))
	for _, fi := range fis {
		paths = append(paths, pathFromInfo(dir, fi))
	}

	return paths, nil
}

func isGzipFile(fi os.FileInfo) bool {
	return !fi.IsDir() && strings.HasSuffix(fi.Name(), ".gz")
}

func pathFromInfo(dir string, fi os.FileInfo) string {
	return filepath.Join(dir, fi.Name())
}
