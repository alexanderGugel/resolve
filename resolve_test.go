package resolve

import (
	"os"
	"testing"
)

func TestResolveExisting(t *testing.T) {
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
		"./test/just-dir": "test/just-dir/index.js",
		"./test/module-with-main": "test/module-with-main/main.js",
		"./test/module-with-main/package.json": "test/module-with-main/package.json",
		"./test/module-without-main": "test/module-without-main/index.js",
	}

	for required, resolved := range testcases {
		dependency, err := Resolve(required, pwd)

		if err != nil {
			t.Errorf("got error %q when resolving %q, expected nil", err, required)
		}

		expected := pwd + string(os.PathSeparator) + resolved

		if dependency == nil {
			t.Errorf("got no dependency when resolving %q; expected %q", required, expected)
			continue
		}

		if dependency.Pathname != expected {
			t.Errorf("got dependency %q when resolving %q; expected %q", dependency.Pathname, required, expected)
		}
	}
}

func TestResolveMissing(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	testcases := []string{"./test/not-here/", "./test/somewhere-else", "./test/not-found/module.js"}

	for _, required := range testcases {
		dependency, err := Resolve(required, pwd)

		if err == nil {
			t.Errorf("got no error when resolving %q; expected error", required)
		}
		if dependency != nil {
			t.Errorf("got dependency %q when resolving %q; expected nil", dependency.Pathname, required)
		}
	}
}
