package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

func mustWriteFile(name string, content []byte) {
	if err := os.WriteFile(name, content, 0644); err != nil {
		log.Fatalf("create file %s: %v", name, err)
	}
}

func generate(day int) {
	folderName := fmt.Sprintf("day%02d", day)
	_, err := os.Stat(folderName)
	if err == nil {
		log.Fatalf("folder %s already exists", folderName)
	}
	if err := os.Mkdir(folderName, 0755); err != nil {
		log.Fatalf("create folder %s: %v", folderName, err)
	}

	tpl, err := template.ParseFiles("day_template.txt")
	if err != nil {
		log.Fatalf("parse template: %v", err)
	}
	buffer := new(bytes.Buffer)
	if err := tpl.Execute(buffer, map[string]any{"Day": day}); err != nil {
		log.Fatalf("execute template: %v", err)
	}
	mustWriteFile(path.Join(folderName, "day.go"), buffer.Bytes())
	mustWriteFile(path.Join(folderName, "input.txt"), []byte{})
	mustWriteFile(path.Join(folderName, "example.txt"), []byte{})

	if err := updateImports("main.go"); err != nil {
		log.Fatalf("update imports: %v", err)
	}
}

func updateImports(s string) error {
	mainFile, err := os.Open(s)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}

	var days []string
	entries, err := os.ReadDir(".")
	if err != nil {
		return fmt.Errorf("read dir: %w", err)
	}
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "day") {
			day := entry.Name()[3:]
			days = append(days, day)
		}
	}

	buffer := new(bytes.Buffer)
	mainContent, err := os.ReadFile(mainFile.Name())
	if err != nil {
		return fmt.Errorf("read main file: %w", err)
	}
	var inAutoImport bool
	for _, line := range strings.Split(string(mainContent), "\n") {
		if strings.TrimSpace(line) == "// <AUTOMATIC-IMPORT>" {
			inAutoImport = true
			_, _ = fmt.Fprintf(buffer, line+"\n")
			// Write all the automatic imports
			for _, d := range days {
				_, _ = fmt.Fprintf(buffer, "\t_ \"github.com/denysvitali/aoc-2024/day%s\"\n", d)
			}
		}
		if strings.TrimSpace(line) == "// </AUTOMATIC-IMPORT>" {
			inAutoImport = false
		}
		if inAutoImport {
			continue
		}
		_, _ = fmt.Fprintln(buffer, line)
	}
	return os.WriteFile(mainFile.Name(), buffer.Bytes(), 0644)
}
