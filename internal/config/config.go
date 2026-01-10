package config

import "fmt"

const VaultPath = "/Users/trevornance/Documents/My Vault/Daily Writing/"

var EditorToUse = "nvim"

type Editor string

const (
	EditorNvim     Editor = "nvim"
	EditorObsidian Editor = "obsidian"
)

type Config struct {
	Editor Editor `json:"editor"`
}

func Default() Config {
	return Config{
		Editor: EditorNvim,
	}
}

func (c Config) Validate() error {
	switch c.Editor {
	case EditorNvim, EditorObsidian:
		return nil
	default:
		return fmt.Errorf("unsupported editor: %q (allowed: %q, %q)", c.Editor, EditorNvim, EditorObsidian)
	}
}
