package resolve

import (
	"os"
	"path/filepath"
	"regexp"
)

var core = []string{"hello", "world"}
var extensions = []string{"", ".js", ".node"}

type packageJSON struct {
	Main string `json:"main"`
}

type Dependency struct {
	File     *os.File
	Pathname string
}

// require(X) from module at path Y.
// Implements Node's [`require.resolve` algorithm](http://nodejs.org/docs/v0.4.8/api/all.html#all_Together...).
func Resolve(x string, y string) (*Dependency, error) {
	y = filepath.Clean(y)

	// require(X) from module at path Y
	// 1. If X is a core module,
	// a. return the core module
	// b. STOP
	// 2. If X begins with './' or '/' or '../'
	// a. LOAD_AS_FILE(Y + X)
	// b. LOAD_AS_DIRECTORY(Y + X)
	// 3. LOAD_NODE_MODULES(X, dirname(Y))
	// 4. THROW "not found"

	isFileOrDir, err := regexp.MatchString("^\\.?\\.?\\/", x)

	if err != nil {
		return nil, err
	}

	if isFileOrDir {
		dependency, err := loadAsFile(y + "/" + x)
		if err != nil {
			return nil, err
		}
		if dependency != nil {
			dependency.Pathname = filepath.Clean(dependency.Pathname)
			return dependency, nil
		}

		dependency, err = loadAsDir(y + "/" + x)
		if err != nil {
			return nil, err
		}
		if dependency != nil {
			dependency.Pathname = filepath.Clean(dependency.Pathname)
			return dependency, nil
		}
	} else {
		dependency, err := loadNodeModules(x, y)

		if err != nil {
			return nil, err
		}

		if dependency != nil {
			dependency.Pathname = filepath.Clean(dependency.Pathname)
			return dependency, nil
		}
	}

	return nil, nil
}
