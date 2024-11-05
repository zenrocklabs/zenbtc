package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func main() {
	processProtoFiles()
}

func processProtoFiles() {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	fmt.Printf("Program executed from: %s\n", cwd)

	fmt.Println("Generating proto code")

	// Collect all .proto files
	var protoFiles []string
	if err = filepath.WalkDir("./proto", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".proto" {
			protoFiles = append(protoFiles, path)
		}
		return nil
	}); err != nil {
		log.Fatalf("Error walking directory: %v", err)
	}

	fmt.Println(protoFiles)

	// Process each proto file
	for _, file := range protoFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Printf("Failed to read file %s: %v", file, err)
			continue
		}

		matched, err := regexp.Match(`(?i)option\s+go_package\s*=\s*".*github\.com/zenrocklabs/zenbtc.*"`, content)
		if err != nil {
			log.Printf("Failed to match regex in file %s: %v", file, err)
			continue
		}

		if matched {
			fmt.Printf("Processing file %s\n", file)

			cmd1 := exec.Command("buf", "generate", "--template", "proto/buf.gen.gogo.yaml", file)
			cmd1.Stdout = os.Stdout
			cmd1.Stderr = os.Stderr
			if err := cmd1.Run(); err != nil {
				log.Printf("Failed to run buf generate for file %s: %v", file, err)
				continue
			}

			cmd2 := exec.Command("buf", "generate", "--template", "proto/buf.gen.python.yaml", file)
			cmd2.Stdout = os.Stdout
			cmd2.Stderr = os.Stderr
			if err := cmd2.Run(); err != nil {
				log.Printf("Failed to run buf generate for file %s: %v", file, err)
				continue
			}
		}
	}

	fmt.Println("Proto files generated.")

	// Move proto files to the right places
	srcDir := filepath.Join("github.com", "zenrocklabs", "zenbtc")
	if err := copyDir(srcDir, "./"); err != nil {
		log.Fatalf("Failed to copy files: %v", err)
	}
	if err := os.RemoveAll("./github.com"); err != nil {
		log.Fatalf("Failed to remove github.com directory: %v", err)
	}
}

func copyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get the relative path
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			// Create the directory in the destination
			return os.MkdirAll(dstPath, info.Mode())
		} else {
			// Copy the file
			return copyFile(path, dstPath)
		}
	})
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.ReadFrom(in)
	if err != nil {
		return err
	}
	return out.Close()
}
