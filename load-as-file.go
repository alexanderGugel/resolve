package resolve

import (
	"os"
)

func loadAsFile(x string) (*Dependency, error) {
	// LOAD_AS_FILE(X)
	// 1. If X is a file, load X as JavaScript text.  STOP
	// 2. If X.js is a file, load X.js as JavaScript text.  STOP
	// 3. If X.node is a file, load X.node as binary addon.  STOP

	for _, extension := range extensions {
		filename := x + extension
		file, err := os.Open(filename)
		if err != nil {
			file.Close()
			if os.IsNotExist(err) {
				continue
			} else {
				return nil, err
			}
		}

		fi, err := file.Stat()
		if err != nil {
			file.Close()
			return nil, err
		}

		if fi.Mode().IsDir() {
			file.Close()
			continue
		}

		return &Dependency{file, filename}, nil
	}

	return nil, nil
}
