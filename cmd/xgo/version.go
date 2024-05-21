package main

import "fmt"

const VERSION = "1.0.36"
const REVISION = "0ed0c7229779392de059107132377800ddf51502+1"
const NUMBER = 227

func getRevision() string {
	revSuffix := ""
	if isDevelopment {
		revSuffix = "_DEV"
	}
	return fmt.Sprintf("%s %s%s BUILD_%d", VERSION, REVISION, revSuffix, NUMBER)
}
