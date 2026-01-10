package main

import (
	"odn/internal/config"
	"os/exec"
	"path/filepath"
)

func openInObsidian(file string) error {
	v := "obsidian://open?vault=My%20Vault&file=Daily%20Writing/" + file
	cmd := exec.Command("open", v)
	return cmd.Run()
}

func openInNvim(file string) error {
	fileToOpen := filepath.Join(config.VaultPath, file)
	cmd := exec.Command("nvim", fileToOpen)
	return cmd.Run()
}
