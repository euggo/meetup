package main

import (
	"compress/gzip"
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

func newFileInfo(path string) *fileInfo {
	// open file at path, decompress while reading
	// return *fileInfo
	// END2 OMIT
	fo := &fileInfo{path: path}

	f, err := os.Open(path)
	if err != nil {
		fo.err = err
		return fo
	}
	defer func() {
		_ = f.Close()
	}()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		fo.err = err
		return fo
	}
	defer func() {
		_ = gzr.Close()
	}()

	data, err := ioutil.ReadAll(gzr)
	if err != nil {
		fo.err = err
		return fo
	}

	fo.data = string(data)

	return fo
}

// BGN1 OMIT
func gzipFilePaths(dir string) ([]string, error) {
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
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

// END1 OMIT

func isGzipFile(fi os.FileInfo) bool {
	return !fi.IsDir() && strings.HasSuffix(fi.Name(), ".gz")
}

func pathFromInfo(dir string, fi os.FileInfo) string {
	return filepath.Join(dir, fi.Name())
}
