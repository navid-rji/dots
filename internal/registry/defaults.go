package registry

import "github.com/navid-rji/dots/internal/config"

func defaults() map[string]config.App {
	return map[string]config.App{
		"nvim": {Paths: []string{"~/.config/nvim/init.lua"}},
		"tmux": {Paths: []string{"~/.tmux.conf"}},
		"zsh":  {Paths: []string{"~/.zshrc"}},
		"bash": {Paths: []string{"~/.bashrc"}},
		"git":  {Paths: []string{"~/.gitconfig"}},
	}
}
