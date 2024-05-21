package main

import "fmt"

const VERSION = "1.0.36"
const REVISION = "5aed1aa769e2337af71072f7e2e92f8cf8318863+1"
const NUMBER = 228

func getRevision() string {
	revSuffix := ""
	if isDevelopment {
		revSuffix = "_DEV"
	}
	return fmt.Sprintf("%s %s%s BUILD_%d", VERSION, REVISION, revSuffix, NUMBER)
}
