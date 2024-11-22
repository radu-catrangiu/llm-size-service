package evaluator

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path"
	"strings"
)

type EvaluatorResponse struct {
	Size        string `json:"size"`
	RoundedSize string `json:"roundedSize"`
}

type Evaluator struct {
	model   string
	hfToken string
	repoDir string
	repoUri string
}

func New(model string, hfToken string) *Evaluator {
	e := &Evaluator{
		model:   model,
		hfToken: hfToken,
	}

	e.computeRepositoryDir()
	e.computeRepositoryUri()

	return e
}

func (e *Evaluator) computeRepositoryUri() {
	authComp := ""

	if e.hfToken != "" {
		authComp = fmt.Sprintf(":%s@", e.hfToken)
	}

	e.repoUri = fmt.Sprintf("https://%shuggingface.co/%s", authComp, e.model)
	slog.Debug("Repository URI", "uri", e.repoUri)
}

func (e *Evaluator) computeRepositoryDir() {
	m := strings.ToLower(e.model)
	m = strings.ReplaceAll(m, "/", "--")
	e.repoDir = path.Join("/tmp", m)
	slog.Debug("Repository directory", "dir", e.repoDir)
}

func (e *Evaluator) cloneRepo() error {
	cmd := exec.Command("git", "clone", "--no-checkout", e.repoUri, e.repoDir)

	err := cmd.Run()
	if err != nil {
		slog.Error("Error cloning repository", "model", e.model, "error", err)
		return fmt.Errorf("error cloning repository: %w", err)
	}

	return nil
}

func (e *Evaluator) getFiles() (res *gitLfsResponse, err error) {
	cmd := exec.Command("git-lfs", "ls-files", "-s", "--json")

	cmd.Dir = e.repoDir

	resBytes, err := cmd.Output()
	if err != nil {
		slog.Error("Error listing repo files", "model", e.model, "error", err)
		return nil, fmt.Errorf("error listing repo files: %w", err)
	}

	res = &gitLfsResponse{}
	err = json.Unmarshal(resBytes, res)
	if err != nil {
		slog.Error("Error unmarshalling git-lfs response", "model", e.model, "error", err)
		return nil, fmt.Errorf("error unmarshalling git-lfs response: %w", err)
	}

	return res, nil
}

func (e *Evaluator) deleteRepo() error {
	err := os.RemoveAll(e.repoDir)
	if err != nil {
		slog.Error("Error removing repository", "dir", e.repoDir, "error", err)
		return fmt.Errorf("error removing repository: %w", err)
	}
	return nil
}

func (e *Evaluator) GetSize() (*EvaluatorResponse, error) {
	err := e.cloneRepo()
	if err != nil {
		return nil, err
	}

	files, err := e.getFiles()
	if err != nil {
		return nil, err
	}

	err = e.deleteRepo()
	if err != nil {
		return nil, err
	}

	totalBytes := files.computeSizeBytes()

	size := convertBytesToK8sSize(totalBytes, false)
	roundedSize := convertBytesToK8sSize(totalBytes, true)

	res := &EvaluatorResponse{
		Size:        size,
		RoundedSize: roundedSize,
	}

	return res, nil
}
