package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type testCase struct {
	cmd  string
	args []string
}

var includeDirs = map[string][]testCase{
	"js": {{cmd: "npm", args: []string{"run", "test:prod"}}},
	"go": {
		{cmd: "go", args: []string{"test", "-v", "./..."}},
		{cmd: "go", args: []string{"vet", "-v", "./..."}},
	},
}

func includeDir(dir string) bool {
	return false
}

func getCommitFiles() (_ []string, err error) {
	files, err := exec.Command("git", "diff", "--cached", "--name-only").Output()
	if err != nil {
		return
	}
	return strings.Split(strings.TrimSpace(string(files)), "\n"), nil
}

func getCommitTests() (res map[string][]testCase, err error) {
	files, err := getCommitFiles()
	if err != nil {
		return
	}
	res = make(map[string][]testCase)
	for _, filename := range files {
		ext := filepath.Ext(filename)
		if ext == ".md" || ext == ".gitignore" || ext == ".json" {
			continue
		}
		fnames := strings.Split(filename, "/")
		if len(fnames) < 2 {
			continue
		}
		tests, ok := includeDirs[fnames[0]]
		if !ok {
			continue
		}
		if fnames[0] == "go" {
			res[fnames[0]] = tests
		} else {
			dirname := filepath.Join(fnames[:2]...)
			res[dirname] = tests
		}
	}
	return
}

func runTestCase(dir string, test testCase) (err error) {
	cmd := exec.Command(test.cmd, test.args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("")

	cwd, _ := os.Getwd()
	dirs, err := getCommitTests()
	if err != nil {
		log.Fatal(err)
	}

	for dir := range dirs {
		log.Printf("Changed: %s", dir)
	}

	for dir, tests := range dirs {
		dirPath := filepath.Join(cwd, dir)
		for _, test := range tests {
			log.Printf("Running: %s %s", test.cmd, strings.Join(test.args, " "))
			err := runTestCase(dirPath, test)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Done: %s %s", test.cmd, strings.Join(test.args, " "))
		}
	}
}
