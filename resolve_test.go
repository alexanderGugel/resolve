package resolve

import (
	"os"
	"testing"
)

type testCase struct {
	require  string
	from     string
	resolved string
}

func TestResolveExisting(t *testing.T) {
	pwd, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	testCases := []testCase{
		testCase{"./test/hello.js", pwd, pwd + "/test/hello.js"},
		testCase{"./test/hello", pwd, pwd + "/test/hello.js"},
		testCase{"./test/other_file.js", pwd, pwd + "/test/other_file.js"},
		testCase{"./test/other_file", pwd, pwd + "/test/other_file.js"},
		testCase{"./test/just_dir/hello_1", pwd, pwd + "/test/just_dir/hello_1.js"},
		testCase{"./test/just_dir/hello_2", pwd, pwd + "/test/just_dir/hello_2.js"},
		testCase{"./test/just_dir/index", pwd, pwd + "/test/just_dir/index.js"},
		testCase{"./test/just_dir", pwd, pwd + "/test/just_dir/index.js"},
		testCase{"./test/module_with_main", pwd, pwd + "/test/module_with_main/main.js"},
		testCase{"./test/module_with_main/package.json", pwd, pwd + "/test/module_with_main/package.json"},
		testCase{"./test/module_without_main", pwd, pwd + "/test/module_without_main/index.js"},
		testCase{"./test/not_here/", pwd, ""},
		testCase{"./test/somewhere_else", pwd, ""},
		testCase{"./test/not_found/module.js", pwd, ""},
		testCase{"./test/module_without_main", pwd, pwd + "/test/module_without_main/index.js"},
		// FIXME
		// testCase{"module_1", pwd + "/test/module_0", pwd + "/test/module_0/node_modules/module_1/index.js"},
		testCase{"module_2", pwd + "/test/module_0", ""},
		testCase{"module_1", pwd + "/test/module_0/node_modules/module_1/node_modules/module_2", pwd + "/test/module_0/node_modules/module_1/index.js"},
	}

	for _, testCase := range testCases {
		dependency, err := Resolve(testCase.require, testCase.from)

		if testCase.resolved == "" {

			if err == nil {
				t.Errorf("got no error when resolving %q; expected error", testCase.require)
			}

			if dependency != nil {
				t.Errorf("got dependency %q when resolving %q; expected nil", dependency.Pathname, testCase.require)
			}

		} else {

			if err != nil {
				t.Errorf("got error %q when resolving %q, expected nil", err, testCase.require)
			}

			if dependency == nil {
				t.Errorf("got no dependency when resolving %q; expected %q", testCase.require, testCase.resolved)
				continue
			}

			if dependency.Pathname != testCase.resolved {
				t.Errorf("got dependency %q when resolving %q; expected %q", dependency.Pathname, testCase.require, testCase.resolved)
			}

		}
	}
}
