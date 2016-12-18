package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	processedFiles      = map[string]bool{}
	inconsistencies     []string
	totalNumberOfChecks int
)

func main() {
	dirs := os.Args[1:]
	greet(dirs)

	for _, baseDir := range dirs {
		println()
		println("--------------------Base directory is set to " + baseDir + "---------------------")
		runValidations(baseDir, dirs)
		printResults()
		resetResults()
	}

	printTotalNumberOfCompares()
}

func runValidations(baseDir string, dirs []string) {
	basefiles, err := ioutil.ReadDir(baseDir)
	handleError(err)

	for _, file := range basefiles {
		processFile(file, baseDir, dirs)
		processedFiles[file.Name()] = true
	}
}

func processFile(file os.FileInfo, baseDir string, dirs []string) {
	for _, remoteDir := range dirs {
		if isProcessed(file.Name()) || sameDirectory(baseDir, remoteDir) {
			continue
		}

		if !isConsistent(file, remoteDir) {
			logInconsistency(file, baseDir, remoteDir)
		}
	}
}

func isProcessed(fileName string) bool {
	return processedFiles[fileName]
}

func sameDirectory(dir1 string, dir2 string) bool {
	return strings.Compare(dir1, dir2) == 0
}

func logInconsistency(file os.FileInfo, baseDir string, remoteDir string) {
	msg := "File " + filepath.Join(baseDir, file.Name()) + " inconsistent with " + remoteDir
	inconsistencies = append(inconsistencies, msg)
}

func isConsistent(baseFile os.FileInfo, remoteDir string) bool {
	totalNumberOfChecks++
	if _, err := os.Stat(filepath.Join(remoteDir, baseFile.Name())); err != nil {
		return false
	}
	return true
}

func greet(dirs []string) {
	fmt.Println("Running cosistency check for:")

	for i, dir := range dirs {
		fmt.Printf("Dir %d - %s\n", i, dir)
	}
	println()
}

func printResults() {
	for _, dif := range inconsistencies {
		fmt.Println(dif)
	}
}

func printTotalNumberOfCompares() {
	println()
	fmt.Printf("Total number of checks performed: %d\n\n", totalNumberOfChecks)
}

func resetResults() {
	inconsistencies = nil
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
