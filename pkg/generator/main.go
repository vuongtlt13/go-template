package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"yourapp/pkg/generator/backend"
	"yourapp/pkg/generator/core"
)

func main() {
	cmd := ""
	flag.StringVar(&cmd, "cmd", "", "Command to run (generate)")
	modelName := flag.String("model", "", "Model name in PascalCase (e.g. User)")
	component := flag.String("component", "repository", "Component to generate (repository, schema, ...)")
	printOnly := flag.Bool("print", false, "Print output instead of writing to file")
	flag.Parse()

	if cmd == "generate" {
		if *modelName == "" {
			fmt.Println("Error: --model is required (PascalCase, e.g. User)")
			os.Exit(1)
		}
		modelSnake := core.SnakeCase(*modelName)
		modelFile := filepath.Join("internal/model", modelSnake+".go")
		if _, err := os.Stat(modelFile); os.IsNotExist(err) {
			fmt.Printf("Error: model file %s does not exist\n", modelFile)
			os.Exit(1)
		}
		entity, err := core.ParseGoStructFile(modelFile)
		if err != nil {
			fmt.Printf("Error parsing model: %v\n", err)
			os.Exit(1)
		}

		// Hiện tại chỉ hỗ trợ repository và schema
		if *component == "repository" {
			gen := backend.NewRepositoryGenerator(entity, "pkg/generator/templates", true)
			content, err := gen.Generate()
			if err != nil {
				fmt.Printf("Error generating repository: %v\n", err)
				os.Exit(1)
			}
			if *printOnly {
				fmt.Println(content)
			} else {
				outputDir := filepath.Join(gen.ComponentFolder())
				if err := core.EnsureDir(outputDir); err != nil {
					fmt.Printf("Error creating output dir: %v\n", err)
					os.Exit(1)
				}
				outputFile := filepath.Join(outputDir, gen.OutputFilename())
				if err := os.WriteFile(outputFile, []byte(content), 0644); err != nil {
					fmt.Printf("Error writing file: %v\n", err)
					os.Exit(1)
				}
				_ = exec.Command("goimports", "-w", outputFile).Run()
				fmt.Printf("Generated: %s\n", outputFile)
			}
			return
		}
		if *component == "schema" {
			gen := backend.NewSchemaGenerator(entity, "pkg/generator/templates", true)
			content, err := gen.Generate()
			if err != nil {
				fmt.Printf("Error generating schema: %v\n", err)
				os.Exit(1)
			}
			if *printOnly {
				fmt.Println(content)
			} else {
				outputDir := filepath.Join(gen.ComponentFolder())
				if err := core.EnsureDir(outputDir); err != nil {
					fmt.Printf("Error creating output dir: %v\n", err)
					os.Exit(1)
				}
				outputFile := filepath.Join(outputDir, gen.OutputFilename())
				if err := os.WriteFile(outputFile, []byte(content), 0644); err != nil {
					fmt.Printf("Error writing file: %v\n", err)
					os.Exit(1)
				}
				_ = exec.Command("goimports", "-w", outputFile).Run()
				fmt.Printf("Generated: %s\n", outputFile)
			}
			return
		}
		if *component == "datatable" {
			gen := backend.NewDataTableGenerator(entity, "pkg/generator/templates", true)
			content, err := gen.Generate()
			if err != nil {
				fmt.Printf("Error generating datatable: %v\n", err)
				os.Exit(1)
			}
			if *printOnly {
				fmt.Println(content)
			} else {
				outputDir := filepath.Join(gen.ComponentFolder())
				if err := core.EnsureDir(outputDir); err != nil {
					fmt.Printf("Error creating output dir: %v\n", err)
					os.Exit(1)
				}
				outputFile := filepath.Join(outputDir, gen.OutputFilename())
				if err := os.WriteFile(outputFile, []byte(content), 0644); err != nil {
					fmt.Printf("Error writing file: %v\n", err)
					os.Exit(1)
				}
				_ = exec.Command("goimports", "-w", outputFile).Run()
				fmt.Printf("Generated: %s\n", outputFile)
			}
			return
		}
		if *component == "service" {
			gen := backend.NewServiceGenerator(entity, "pkg/generator/templates", true)
			content, err := gen.Generate()
			if err != nil {
				fmt.Printf("Error generating service: %v\n", err)
				os.Exit(1)
			}
			if *printOnly {
				fmt.Println(content)
			} else {
				outputDir := filepath.Join(gen.ComponentFolder())
				if err := core.EnsureDir(outputDir); err != nil {
					fmt.Printf("Error creating output dir: %v\n", err)
					os.Exit(1)
				}
				outputFile := filepath.Join(outputDir, gen.OutputFilename())
				if err := os.WriteFile(outputFile, []byte(content), 0644); err != nil {
					fmt.Printf("Error writing file: %v\n", err)
					os.Exit(1)
				}
				_ = exec.Command("goimports", "-w", outputFile).Run()
				fmt.Printf("Generated: %s\n", outputFile)
			}
			return
		}
		if *component == "handler" {
			gen := backend.NewHandlerGenerator(entity, "pkg/generator/templates", true)
			content, err := gen.Generate()
			if err != nil {
				fmt.Printf("Error generating handler: %v\n", err)
				os.Exit(1)
			}
			if *printOnly {
				fmt.Println(content)
			} else {
				outputDir := filepath.Join(gen.ComponentFolder())
				if err := core.EnsureDir(outputDir); err != nil {
					fmt.Printf("Error creating output dir: %v\n", err)
					os.Exit(1)
				}
				outputFile := filepath.Join(outputDir, gen.OutputFilename())
				if err := os.WriteFile(outputFile, []byte(content), 0644); err != nil {
					fmt.Printf("Error writing file: %v\n", err)
					os.Exit(1)
				}
				_ = exec.Command("goimports", "-w", outputFile).Run()
				fmt.Printf("Generated: %s\n", outputFile)
			}
			return
		}
		fmt.Println("Unsupported component. Only 'repository', 'schema', 'datatable', 'service', and 'handler' are supported.")
		return
	}

	fmt.Println("Usage: go run main.go --cmd=generate --model=User [--component=repository] [--print]")
}
