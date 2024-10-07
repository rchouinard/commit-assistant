package git

import (
	"fmt"
	"os/exec"
	"sort"
	"strings"
)

func IsGitInstalled() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

func IsGitRepo() (bool, error) {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	if err := cmd.Run(); err != nil {
		return false, err
	}
	return true, nil
}

func HasUncommitedChanges() (bool, error) {
	cmd := exec.Command("git", "diff-index", "--quiet", "HEAD", "--")
	return cmd.Run() != nil, nil
}

func GetRoot() (string, error) {
	root, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", fmt.Errorf("could not find git root: %w", err)
	}
	return strings.TrimSpace(string(root)), nil
}

func GetChangedFiles() ([]string, error) {
	cmd := exec.Command("git", "ls-files", "--modified", "--others", "--exclude-standard")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("could not get changed files: %w", err)
	}

	files := strings.FieldsFunc(string(out), func(r rune) bool {
		return r == '\n'
	})
	sort.Strings(files)
	return files, nil
}

func GetStagedFiles() ([]string, error) {
	root, err := GetRoot()
	if err != nil {
		return nil, fmt.Errorf("could not get staged files: %w", err)
	}

	cmd := exec.Command("git", "diff", "--cached", "--name-only", "--relative", root)
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("could not get staged files: %w", err)
	}

	return strings.FieldsFunc(string(out), func(r rune) bool {
		return r == '\n'
	}), nil
}

func StageFiles(files []string) error {
	args := append([]string{"add"}, files...)
	cmd := exec.Command("git", args...)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not stage files: %w", err)
	}
	return nil
}

func DiffFiles(files []string) (string, error) {
	args := append([]string{"diff", "--staged", "--"}, files...)
	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("could not generate diff: %w", err)
	}
	return string(out), nil
}
