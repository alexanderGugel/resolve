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
	defer file.Close()

	if err == nil {
		manifest := packageJSON{}

		parser := json.NewDecoder(file)
		if err = parser.Decode(&manifest); err != nil {
			return nil, err
		}

		if manifest.Main != "" {
			m := x + string(os.PathSeparator) + manifest.Main
			dependency, err := loadAsFile(m)

			if err == nil && dependency != nil {
				return dependency, err
			}
		}
	}

	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	return loadAsFile(x + string(os.PathSeparator) + "index")
}
