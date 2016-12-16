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
	validate()
	printResults()
}

func validate() {
	basefiles, err := ioutil.ReadDir(baseDir)
	handleError(err)

	for _, file := range basefiles {
		for _, remoteDir := range remoteDirs {
			isConsistent := isConsistent(file, remoteDir)
			if !isConsistent {
				difMsg := "File " + filepath.Join(baseDir, file.Name()) + " inconsistent with " + remoteDir
				inconsistencies = append(inconsistencies, difMsg)
			}
		}
	}
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
