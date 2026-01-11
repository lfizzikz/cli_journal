package parseflags

import (
	"flag"
	"fmt"
	"odn/internal/config"
)

func ParseConfigFlags(args []string) (config.Config, error) {
	fs := flag.NewFlagSet("config", flag.ContinueOnError)

	var editor string

	fs.StringVar(&editor, "editor", "", "default editor when opening files")

	if err := fs.Parse(args); err != nil {
		return config.Config{}, err
	}

	if extra := fs.Args(); len(extra) > 0 {
		return config.Config{}, fmt.Errorf("unknown config argument: %q (use \"default\" or --editor)", extra[0])
	}

	cfg := config.Config{Editor: config.Editor(editor)}
	if err := cfg.Validate(); err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}
