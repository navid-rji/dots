# dots

> Open any app's config file in your editor — no more hunting through `~/.config`.

`dots` is a small command-line dotfile manager. Instead of remembering where each
program keeps its config, just run `dots <app>` and it opens the right file in your
editor.

```console
$ dots nvim      # opens ~/.config/nvim/init.lua
$ dots zsh       # opens ~/.zshrc
$ dots git       # opens ~/.gitconfig
```

> [!NOTE]
> **Early days.** This is a work in progress — and my first Go project. Expect rough
> edges and breaking changes. Feedback, ideas, and issues are very welcome.

## Why

Config files live in a dozen different places: `~/.config/nvim/`, `~/.zshrc`,
`~/.tmux.conf`, `~/.config/hypr/`, and on and on. `dots` keeps a small registry that
maps an app name to its config path, so editing a dotfile is one short command away
instead of a `find` expedition.

## Install

Requires Go 1.26+.

```console
go install github.com/navid-rji/dots@latest
```

This installs a `dots` binary into `$GOPATH/bin` (usually `~/go/bin`) — make sure that
directory is on your `PATH`.

## Usage

The first time you run `dots`, it asks which command should open your files (for
example `nvim`, `code`, or `emacsclient -c {} -n`). You can change this later in the
config file.

| Command | Description |
|---|---|
| `dots <app>` | Open an app's config in your editor |
| `dots list` (`ls`) | List known apps and their config paths |
| `dots add <app> <path>` (`a`) | Register a new app → path mapping |
| `dots update <app> <path>` (`u`) | Change an existing app's path |
| `dots dots` | Open dots' own config file |

`dots` ships with sensible defaults for `nvim`, `tmux`, `zsh`, `bash`, and `git`.
Anything you add or override lives in your own config and takes precedence.

### Example

```console
$ dots add hyprland ~/.config/hypr/hyprland.conf
added hyprland -> ~/.config/hypr/hyprland.conf

$ dots list
bash         ~/.bashrc
git          ~/.gitconfig
hyprland     ~/.config/hypr/hyprland.conf
nvim         ~/.config/nvim/init.lua
tmux         ~/.tmux.conf
zsh          ~/.zshrc

$ dots hyprland   # opens it in your editor
```

## Configuration

`dots` stores its config as TOML at `~/.config/dots/config.toml`. It respects
`$XDG_CONFIG_HOME`, or you can point it somewhere else entirely with `$DOTS_CONFIG_DIR`.

```toml
editor = "nvim"

[apps.hyprland]
paths = ["~/.config/hypr/hyprland.conf"]
```

**Editor command.** Set `editor` to whatever opens your files. Use `{}` as a
placeholder for the file path; if you leave it out, the path is appended at the end:

```toml
editor = "code"                    # runs: code <path>
editor = "emacsclient -c {} -n"    # runs: emacsclient -c <path> -n
```

If `editor` is empty, `dots` falls back to `$VISUAL`, then `$EDITOR`, then `vi`.

Paths support `~` and environment variables (for example `$XDG_CONFIG_HOME/foo`).

## Roadmap

### Phase 0 — MVP (core commands)

- [x] App registry mapping app name → config path(s), with built-in defaults
- [x] `dots <app>` — resolve an app and open its config in your editor
- [x] Configurable editor command (with `{}` path placeholder)
- [x] dots' own config file at `~/.config/dots/config.toml`
- [x] `dots list` / `ls` — show known apps and their paths
- [x] `dots add <app> <path>` — register a custom mapping
- [x] `dots update <app> <path>` — change an existing mapping
- [x] `dots dots` — open dots' own config
- [x] `dots clear <app>` — drop a mapping (restoring the built-in default if one exists)
- [ ] `dots list` also shows whether each config file exists on disk
- [ ] Protect the dots config file from accidental deletion or overwrite

### Phase 1 — Discovery & custom search paths

- [ ] Configurable search paths
- [ ] `dots scan` / `discover` — walk the search paths and known locations, surface the
      configs found on disk, and let you register them into the registry

### Phase 2 — Multiple configs & interactive selection

- [ ] When an app has more than one config, show a picker to choose which to open
      (today `dots <app>` opens the first registered path)
- [ ] `dots <app> --list` — list all known configs for an app and pick one
- [ ] `dots` with no args (or `-i` / `dots pick`) — fuzzy-find across all registered
      apps and open on <kbd>Enter</kbd> (arrow keys / vim bindings for navigation)

### Phase 3 — Fetch & install configs

- [ ] `dots get <url> [dest]` — download a config file from a URL
- [ ] `--as <app>` — download and register it as an app, backing up any existing file first

### Phase 4 — TUI

- [ ] `dots tui` — full-screen interactive interface

## Built with

- [Cobra](https://github.com/spf13/cobra) — command-line framework
- [go-toml](https://github.com/pelletier/go-toml) — TOML parsing

## Contributing

This is an early personal project, but issues, ideas, and pull requests are welcome.
If you hit a bug or have a suggestion, open an issue.

## License

[MIT](LICENSE) © Navid Rajaei
