package resolve

import (
	"os"
	"strings"
)

func nodeModulesPaths(start string) []string {
	// NODE_MODULES_PATHS(START)
	// 1. let PARTS = path split(START)
	// 2. let ROOT = index of first instance of "node_modules" in PARTS, or 0
	// 3. let I = count of PARTS - 1
	// 4. let DIRS = []
	// 5. while I > ROOT,
	//    a. if PARTS[I] = "node_modules" CONTINUE
	//    c. DIR = path join(PARTS[0 .. I] + "node_modules")
	//    b. DIRS = DIRS + DIR
	//    c. let I = I - 1
	// 6. return DIRS

	parts := strings.Split(start, string(os.PathSeparator))

	root := 0

	for index, part := range parts {
		if part == "node_modules" {
			root = index
			break
		}
	}

	i := len(parts) - 1

	dirs := []string{}

	for i > root {
		if parts[i] == "node_modules" {
			continue
		}

		dir := strings.Join(append(parts[0:i+1], "node_modules"), string(os.PathSeparator))

		dirs = append(dirs, dir)
		i = i - 1
	}

	return dirs
}
