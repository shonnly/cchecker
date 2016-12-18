package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var baseDir string
var remoteDirs []string
var inconsistencies []string

func main() {
	dirs := os.Args[1:]
	baseDir = dirs[0]
	remoteDirs = dirs[1:]

	greet()
	runValidations()
	printResults()
}

func runValidations() {
	basefiles, err := ioutil.ReadDir(baseDir)
	handleError(err)

	for _, file := range basefiles {
		processFile(file)
	}
}

func processFile(file os.FileInfo) {
	for _, remoteDir := range remoteDirs {
		if !isConsistent(file, remoteDir) {
			logInconsistency(file, remoteDir)
		}
	}
}

func logInconsistency(file os.FileInfo, remoteDir string) {
	msg := "File " + filepath.Join(baseDir, file.Name()) + " inconsistent with " + remoteDir
	inconsistencies = append(inconsistencies, msg)
}

func isConsistent(baseFile os.FileInfo, remoteDir string) bool {
	if _, err := os.Stat(filepath.Join(remoteDir, baseFile.Name())); err != nil {
		return false
	}
	return true
}

func greet() {
	fmt.Println("Running Cosistency Check:")
	println("Base dir - ", baseDir)

	for i, dir := range remoteDirs {
		fmt.Printf("Remote %d - %s\n", i, dir)
	}
}

func printResults() {
	for _, dif := range inconsistencies {
		fmt.Println(dif)
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
