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
		"./test/hello.js":                      "test/hello.js",
		"./test/hello":                         "test/hello.js",
		"./test/other_file.js":                 "test/other_file.js",
		"./test/other_file":                    "test/other_file.js",
		"./test/just_dir/hello_1":              "test/just_dir/hello_1.js",
		"./test/just_dir/hello_2":              "test/just_dir/hello_2.js",
		"./test/just_dir/index":                "test/just_dir/index.js",
		"./test/just_dir":                      "test/just_dir/index.js",
		"./test/module_with_main":              "test/module_with_main/main.js",
		"./test/module_with_main/package.json": "test/module_with_main/package.json",
		"./test/module_without_main":           "test/module_without_main/index.js",
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

	testcases := []string{"./test/not_here/", "./test/somewhere_else", "./test/not_found/module.js"}

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
