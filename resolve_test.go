package resolve

import (
	"os"
	"testing"
)

func TestResolve(t *testing.T) {
	pwd, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	testcases := map[string]string{
		"./test/hello.js": "test/hello.js",
		"./test/hello": "test/hello.js",
		"./test/other-file.js": "test/other-file.js",
		"./test/other-file": "test/other-file.js",
		"./test/just-dir/hello-1": "test/just-dir/hello-1.js",
		"./test/just-dir/hello-2": "test/just-dir/hello-2.js",
		"./test/just-dir/index": "test/just-dir/index.js",

		// FIXME
		"./test/just-dir": "test/just-dir/index.js",

		"./test/module-with-main": "test/module-with-main/main.js",
		"./test/module-with-main/package.json": "test/module-with-main/package.json",
		"./test/module-without-main": "test/module-without-main/index.js",
	}

	for required, resolved := range testcases {
		dependency, err := Resolve(required, pwd)
		if err != nil {
			t.Error("Should not have thrown an error: ", err)
		}
		expected := pwd + string(os.PathSeparator) + resolved

		if dependency == nil || dependency.Pathname != expected {
			t.Error("Expected ", expected, " got ", dependency)
		}
	}

}
