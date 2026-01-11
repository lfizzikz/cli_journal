package main

import (
	"errors"
	"odn/internal/config"
	"os"
	"os/exec"
	"path/filepath"
)

func openInObsidian(file string) error {
	v := "obsidian://open?vault=My%20Vault&file=Daily%20Writing/" + file
	cmd := exec.Command("open", v)
	return cmd.Run()
}

func openInNvim(file string) error {
	f := filepath.Join(config.VaultPath, file)
	cmd := exec.Command("nvim", f)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func OpenInDefaultEditor(file string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	if cfg.Editor == "nvim" {
		openInNvim(file)
		return nil
	}
	if cfg.Editor == "obsidian" {
		openInObsidian(file)
		return nil
	}
	return errors.New("couldnt open in editor")
}
