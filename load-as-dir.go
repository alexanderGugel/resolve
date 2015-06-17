package resolve

import (
	"encoding/json"
	"os"
)

func loadAsDir(x string) (*Dependency, error) {
	// LOAD_AS_DIRECTORY(X)
	// 1. If X/package.json is a file,
	// a. Parse X/package.json, and look for "main" field.
	// b. let M = X + (json main field)
	// c. LOAD_AS_FILE(M)
	// 2. LOAD_AS_FILE(X/index)

	file, err := os.Open(x + "/package.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	manifest := packageJSON{}

	parser := json.NewDecoder(file)
	if err = parser.Decode(&manifest); err != nil {
		return nil, err
	}

	if manifest.Main != "" {
		return loadAsFile(x + "/" + manifest.Main)
	}

	return loadAsFile(x + "/" + "index")
}
