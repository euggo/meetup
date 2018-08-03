package runtest

import (
	"reflect"
	"runtime"
	"testing"
)

// Run ...
func Run(fn func(*testing.T)) {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	tests := []testing.InternalTest{{Name: name, F: fn}}
	matchAll := func(t string, pat string) (bool, error) { return true, nil }
	testing.Main(matchAll, tests, nil, nil)
}
