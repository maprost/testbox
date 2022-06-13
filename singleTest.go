package testbox

import (
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/maprost/testbox/must"
)

func SingleTest(t testing.TB) {
	SingleTestWithSkip(t, 2)
}

func SingleTestWithSkip(t testing.TB, skip int) {
	fn := functionName(t, skip)
	//fmt.Printf("args: %s contains '%s'\n", strings.Join(os.Args, ", "), fn)

	skipTest := true
	for _, arg := range os.Args {
		if strings.Contains(arg, fn) {
			skipTest = false
		}
	}

	if skipTest {
		t.Skipf("skip test '%s', can only run as single test (args: %s)", fn, strings.Join(os.Args, ", "))
	}
}

func functionName(t testing.TB, skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	must.BeTrue(t, ok)
	fn := runtime.FuncForPC(pc)

	// prepare function name
	// remove '/' in front
	functionName := fn.Name()
	//fmt.Printf("origin function name: %s\n", functionName)
	lastSlash := strings.LastIndex(functionName, "/")
	if lastSlash != -1 {
		functionName = functionName[lastSlash+1:]
	}
	//fmt.Printf("cut function name: %s\n", functionName)

	// cut the file name
	split := strings.Split(functionName, ".")
	if len(split) == 1 {
		functionName = split[0]
	} else {
		functionName = split[1]
	}

	//fmt.Printf("split function name: %s\n", functionName)
	return functionName
}
