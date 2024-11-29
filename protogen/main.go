package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

type BufConfig struct {
	Deps []string `yaml:"deps"`
}

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

	var wg sync.WaitGroup
	errorsChan := make(chan error, len(protoFiles))

	// Process each proto file concurrently
	for _, file := range protoFiles {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()

			content, err := os.ReadFile(file)
			if err != nil {
				errorsChan <- fmt.Errorf("failed to read file %s: %v", file, err)
				return
			}

			matched, err := regexp.Match(`(?i)option\s+go_package\s*=\s*".*github\.com/zenrocklabs/zenbtc.*"`, content)
			if err != nil {
				errorsChan <- fmt.Errorf("failed to match regex in file %s: %v", file, err)
				return
			}

			if matched {
				fmt.Printf("Processing file %s\n", file)

				// Process buf commands concurrently
				var cmdWg sync.WaitGroup
				cmdWg.Add(2)

				go func() {
					defer cmdWg.Done()
					cmd := exec.Command("buf", "generate", "--template", "proto/buf.gen.gogo.yaml", file)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err := cmd.Run(); err != nil {
						errorsChan <- fmt.Errorf("failed to run buf generate (gogo) for file %s: %v", file, err)
					}
				}()

				go func() {
					defer cmdWg.Done()
					cmd := exec.Command("buf", "generate", "--template", "proto/buf.gen.python.yaml", file)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err := cmd.Run(); err != nil {
						errorsChan <- fmt.Errorf("failed to run buf generate (python) for file %s: %v", file, err)
					}
				}()

				cmdWg.Wait()
			}
		}(file)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(errorsChan)
	}()

	// Collect and handle errors
	var errors []error
	for err := range errorsChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		fmt.Println("\nEncountered errors during processing:")
		for _, err := range errors {
			fmt.Printf("- %v\n", err)
		}
		if len(errors) == len(protoFiles) {
			log.Fatal("All file processing failed")
		}
	}

	fmt.Println("Proto files generated.")

	// Generate Pulsar files
	if err := generatePulsarFiles(); err != nil {
		log.Fatalf("Failed to generate Pulsar files: %v", err)
	}

	// Move proto files to the right places
	srcDir := filepath.Join("../github.com", "zenrocklabs", "zenbtc", "x", "zenbtc", "types")
	dstDir := filepath.Join("../", "x", "zenbtc", "types")

	if err := copyDir(srcDir, dstDir); err != nil {
		log.Fatalf("Failed to copy files: %v", err)
	}

	// Clean up only if the directory exists
	if err := os.RemoveAll("../github.com"); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Failed to remove github.com directory: %v", err)
	}
}

func copyDir(src string, dst string) error {
	var wg sync.WaitGroup
	errorsChan := make(chan error, 100) // Buffer for potential errors

	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
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
			return os.MkdirAll(dstPath, info.Mode())
		}

		// Handle file copy concurrently
		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := copyFile(path, dstPath); err != nil {
				errorsChan <- fmt.Errorf("failed to copy %s to %s: %v", path, dstPath, err)
			}
		}()

		return nil
	})

	// Wait for all copies to complete
	go func() {
		wg.Wait()
		close(errorsChan)
	}()

	// Collect any errors
	var errors []error
	for err := range errorsChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return fmt.Errorf("multiple errors during copy: %v", errors)
	}

	return err
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

func generatePulsarFiles() error {
	// Check if protoc-gen-go-pulsar is installed
	if _, err := exec.LookPath("protoc-gen-go-pulsar"); err != nil {
		return fmt.Errorf("protoc-gen-go-pulsar is not installed. Please run: go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar")
	}

	// Get project root directory using git
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	projectRoot, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get project root: %v", err)
	}
	protoRoot := strings.TrimSpace(string(projectRoot))

	// Change to proto directory
	protoDir := filepath.Join(protoRoot, "proto")
	if err := os.Chdir(protoDir); err != nil {
		return fmt.Errorf("failed to change to proto directory: %v", err)
	}

	// Update buf modules
	cmd = exec.Command("buf", "dep", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to update buf modules: %v", err)
	}

	// Build buf
	cmd = exec.Command("buf", "build", "-v")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to build buf: %v", err)
	}

	// Generate proto pulsar code
	cmd = exec.Command("buf", "generate", "-v", "--template", "buf.gen.pulsar.yaml")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate pulsar files: %v", err)
	}

	fmt.Println("Pulsar files generated.")
	return nil
}

func generatePythonDependencies() {
	bufYamlPath := "proto/buf.yaml"
	bufYamlContent, err := os.ReadFile(bufYamlPath)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", bufYamlPath, err)
	}

	var bufConfig BufConfig
	if err = yaml.Unmarshal(bufYamlContent, &bufConfig); err != nil {
		log.Fatalf("Failed to parse %s: %v", bufYamlPath, err)
	}

	var wg sync.WaitGroup
	errorsChan := make(chan error, len(bufConfig.Deps))

	for _, dep := range bufConfig.Deps {
		wg.Add(1)
		go func(dep string) {
			defer wg.Done()

			fmt.Printf("Generating python dependencies for %s\n", dep)
			cmd := exec.Command("buf", "generate", "--template", "proto/buf.gen.python.yaml", dep)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				errorsChan <- fmt.Errorf("failed to generate dependencies for %s: %v", dep, err)
			}
		}(dep)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(errorsChan)
	}()

	// Collect and handle errors
	var errors []error
	for err := range errorsChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		fmt.Println("\nEncountered errors during Python dependency generation:")
		for _, err := range errors {
			fmt.Printf("- %v\n", err)
		}
		if len(errors) == len(bufConfig.Deps) {
			log.Fatal("All Python dependency generation failed")
		}
	}

	fmt.Println("Python dependencies generated.")
}
