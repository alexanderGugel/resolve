package resolve

func loadNodeModules(x string, start string) (*Dependency, error) {
	// LOAD_NODE_MODULES(X, START)
	// 1. let DIRS=NODE_MODULES_PATHS(START)
	// 2. for each DIR in DIRS:
	//    a. LOAD_AS_FILE(DIR/X)
	//    b. LOAD_AS_DIRECTORY(DIR/X)

	dirs := nodeModulesPaths(start)

	for _, dir := range dirs {
		dependency, err := loadAsFile(dir + "/" + x)
		if err != nil {
			return nil, err
		}
		if dependency != nil {
			return dependency, nil
		}

		dependency, err = loadAsDir(dir + "/" + x)
		if err != nil {
			return nil, err
		}
		if dependency != nil {
			return dependency, nil
		}
	}

	return nil, nil
}
