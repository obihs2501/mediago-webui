package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// executeMediago executes the mediago binary with given parameters
func executeMediago(url, format, cookies, proxy string) (string, error) {
	mediagoPath := findMediagoBinary()
	if mediagoPath == "" {
		return "", fmt.Errorf("mediago binary not found. Please ensure mediago.exe is in the same directory or in PATH")
	}

	args := []string{}

	if format != "" && format != "best" {
		args = append(args, "-f", format)
	}

	args = append(args, "-o", "downloads/%(title)s.%(ext)s")

	if cookies != "" {
		args = append(args, "--cookies", cookies)
	}

	if proxy != "" {
		args = append(args, "--proxy", proxy)
	}

	args = append(args, url)

	// Create downloads directory
	os.MkdirAll("downloads", 0755)

	// Execute command
	cmd := exec.Command(mediagoPath, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return string(output), fmt.Errorf("download failed: %s", err.Error())
	}

	return string(output), nil
}

// getExtractors gets the list of supported extractors
func getExtractors() (string, error) {
	mediagoPath := findMediagoBinary()
	if mediagoPath == "" {
		return "", fmt.Errorf("mediago binary not found")
	}

	cmd := exec.Command(mediagoPath, "--list-extractors")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// findMediagoBinary finds the mediago executable
func findMediagoBinary() string {
	// Get executable directory
	exePath, err := os.Executable()
	if err == nil {
		exeDir := filepath.Dir(exePath)
		// Check in same directory as executable
		paths := []string{
			filepath.Join(exeDir, "mediago.exe"),
			filepath.Join(exeDir, "mediago"),
		}
		for _, path := range paths {
			if _, err := os.Stat(path); err == nil {
				return path
			}
		}
	}

	// Check current working directory
	cwd, _ := os.Getwd()
	paths := []string{
		filepath.Join(cwd, "mediago.exe"),
		filepath.Join(cwd, "mediago"),
		filepath.Join(cwd, "..", "mediago.exe"),
		filepath.Join(cwd, "..", "mediago"),
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			abs, _ := filepath.Abs(path)
			return abs
		}
	}

	// Check PATH
	if path, err := exec.LookPath("mediago"); err == nil {
		return path
	}

	return ""
}

// Helper to sanitize file paths
func sanitizePath(path string) string {
	path = strings.ReplaceAll(path, "..", "")
	path = strings.ReplaceAll(path, "~", "")
	return path
}
