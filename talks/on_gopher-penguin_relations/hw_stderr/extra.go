package main

/*
// START1 OMIT
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
// END1 OMIT
*/

/*
// START2 OMIT
type Writer interface {
        Write(p []byte) (n int, err error)
}
// END2 OMIT
*/

/*
// START3 OMIT
var (
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
// END3 OMIT
*/

/*
// START4 OMIT
type File
	// ...
    func (f *File) Write(b []byte) (n int, err error)
	// ...
// END4 OMIT
*/
