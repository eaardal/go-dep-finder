package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go-dep-finder <root_directory_path> <dependency>")
		return
	}

	rootDir := os.Args[1]
	dependency := os.Args[2]

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name() == "go.mod" {
			projectDir := filepath.Dir(path)
			projectName := filepath.Base(projectDir)

			if checkDependency(path, dependency) {
				fmt.Printf("Project: %s, Dependency Version: %s\n", projectName, getDependencyVersion(path, dependency))
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path: %v\n", err)
		return
	}
}

func checkDependency(filePath, dependency string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, dependency) {
			return true
		}
	}

	return false
}

func getDependencyVersion(filePath, dependency string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "Unknown"
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.Contains(line, dependency) {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				return parts[1]
			}
			return "Version not found"
		}
	}

	return "Unknown"
}
