// Build application for goweb
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	ProjectName      string
	ProjectDirectory string
	OutputDirectory  string
	Configuration    string
	Architecture     string
	OperatingSystem  string
	Publish          bool
)

func init() {
	flag.StringVar(&ProjectName, "name", "main", "Name of the project")
	flag.StringVar(&ProjectDirectory, "D", ".", "Directory with main function for build")
	flag.StringVar(&OutputDirectory, "O", "./bin", "Directory to output  build files")
	flag.StringVar(&Configuration, "C", "Debug", "Configuration to build the project under (Debug/Release)")
	flag.StringVar(&Architecture, "A", "amd64", "Architecture to build application under")
	flag.StringVar(&OperatingSystem, "OS", "linux", "Operating systme to build application under")
	flag.BoolVar(&Publish, "P", false, "Enable publish mode in build")
	flag.Parse()
}
func getProjectName() {
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: could read directory %s", err)
		//		return ""
	}

	name := filepath.Base(dir)
	log.Printf("Project Name: %s", name)
	ProjectName = name
}
func configureBuildDir() string {
	outputDir := filepath.Join(OutputDirectory, Configuration, OperatingSystem, Architecture)
	_, dirErr := os.ReadDir(outputDir)
	if dirErr != nil {
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			log.Fatalf("ERROR: unable to make build directory %s: %s", outputDir, err)
		}
	}
	return outputDir
}
func setBuildEnv() {
	cmd := exec.Command("go", "env", "-w", fmt.Sprintf("GOOS=%s", OperatingSystem), fmt.Sprintf("GOARCH=%s", Architecture), "CGO_ENABLED=0")
	cmd.Run()
}

func main() {
	if ProjectName != "" {
		getProjectName()
	}
	outputDir := configureBuildDir()

	log.Println("Setting up build environment...")
	setBuildEnv()

	log.Println("Building project...")
	cmd := exec.Command("go", "build", "-o", fmt.Sprintf("%s/%s", outputDir, ProjectName))
	cmd.Dir = ProjectDirectory
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		log.Fatalf("ERROR: could not build project:  %s", err)
	}

	log.Printf("Project built to %s", OutputDirectory)
}
