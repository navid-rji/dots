# dots

[![CI](https://github.com/navid-rji/dots/actions/workflows/ci.yml/badge.svg)](https://github.com/navid-rji/dots/actions/workflows/ci.yml)

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

### Homebrew

```console
brew install navid-rji/tap/dots
```

Or add the tap permanently, then install `dots` like any other formula:

```console
brew tap navid-rji/tap
brew install dots
```

### Go

Requires Go 1.26+.

```console
go install github.com/navid-rji/dots@latest
```

This installs a `dots` binary into `$GOPATH/bin` (usually `~/go/bin`) — make sure that
directory is on your `PATH`.

`dots` runs on **macOS and Linux**.

## Usage

The first time you run `dots`, it asks two things: which command should open your
files (for example `nvim`, `code`, or `emacsclient -c {} -n`), and whether to include
the built-in defaults. You can change both later in the config file.

| Command | Description |
| --- | --- |
| `dots <app>` | Open an app's config in your editor |
| `dots list` (`ls`) | List known apps and their config paths |
| `dots list --check` | Also mark whether each config file exists on disk (✓/✗) |
| `dots list --custom` | Show only the apps you registered yourself |
| `dots add <app> <path>` | Register a new app → path mapping |
| `dots update <app> <path>` | Change an app's path (also works for apps not registered yet) |
| `dots clear <app>` | Remove an app's mapping (restores built-in default if one exists) |
| `dots dots` | Open dots' own config file |
| `dots --version` | Print the version |

`dots` ships with best-guess defaults for 90+ well-known tools — shells (`zsh`,
`fish`), editors (`nvim`, `helix`), terminals (`kitty`, `ghostty`), multiplexers
(`tmux`, `zellij`), window managers (`hyprland`, `sway`, `yabai`), and a pile of CLI
utilities. Run `dots list` to see them all, or `dots list --check` to see which ones
actually exist on your machine. Anything you add or override takes precedence, and
you can drop the whole set with `use_defaults = false` (see below).

### Flags

These apply to `dots <app>`:

| Flag | Description |
| --- | --- |
| `-p`, `--print` | Print the resolved path instead of opening it |
| `--dir` | Open the containing folder instead of the file itself |
| `-e`, `--editor <cmd>` | Use a different editor, just this once |

```console
$ dots -p nvim
/home/you/.config/nvim/init.lua

$ dots -p nvim --dir
/home/you/.config/nvim

$ dots nvim --dir            # opens the folder rather than init.lua
$ dots zsh -e code           # ignore the configured editor for one run
$ wc -l "$(dots -p git)"     # --print composes with anything
```

`--print` writes the fully expanded absolute path, not the `~/…` form that
`dots list` displays — that's what makes it safe to hand to other commands.

### Example

```console
$ dots add hyprpaper ~/.config/hypr/hyprpaper.conf
added hyprpaper -> ~/.config/hypr/hyprpaper.conf

$ dots list --custom
hyprpaper    ~/.config/hypr/hyprpaper.conf

$ dots hyprpaper   # opens it in your editor
```

## Configuration

`dots` stores its config as TOML at `~/.config/dots/config.toml`. It respects
`$XDG_CONFIG_HOME`, or you can point it somewhere else entirely with `$DOTS_CONFIG_DIR`.

```toml
editor = "nvim"
use_defaults = true   # set to false to drop the built-in registry entirely

[apps.hyprpaper]
paths = ["~/.config/hypr/hyprpaper.conf"]
```

**Editor command.** Set `editor` to whatever opens your files. Use `{}` as a
placeholder for the file path; if you leave it out, the path is appended at the end:

```toml
editor = "code"                    # runs: code <path>
editor = "emacsclient -c {} -n"    # runs: emacsclient -c <path> -n
```

If `editor` is empty, `dots` falls back to `$VISUAL`, then `$EDITOR`, then `vi`.

**Built-in defaults.** With `use_defaults = true` (the default, also when the key is
absent), the built-in registry is layered underneath your own entries. Set it to
`false` to keep only the apps you registered yourself. To change just a single
default, `dots update <app> <path>` is enough — your entry wins.

Paths support `~` and environment variables (for example `$XDG_CONFIG_HOME/foo`).

### Shell completion

`dots` completes app names, so `dots nv<TAB>` → `dots nvim`.

Homebrew installs completions automatically — nothing to do.

With `go install`, set them up once:

```console
# zsh
$ mkdir -p ~/.zsh/completions
$ dots completion zsh > ~/.zsh/completions/_dots
# then in ~/.zshrc, above `compinit`:
#   fpath=(~/.zsh/completions $fpath)

# bash — requires the bash-completion package
$ mkdir -p ~/.local/share/bash-completion/completions
$ dots completion bash > ~/.local/share/bash-completion/completions/dots

# fish
$ dots completion fish > ~/.config/fish/completions/dots.fish
```

Run `dots completion <shell> --help` for other options, including system-wide paths.

## Roadmap

<details>
<summary><strong>v0.1.0 — First public release</strong></summary>

Get the foundation and safety right.

- [x] Core: app registry with built-in defaults; `dots <app>`, `list` (`--check` / `--custom`), `add`, `update`, `clear`, `dots dots`, `--version`
- [x] Safety: reserved `dots` name guard, atomic config writes, clean error printing
- [x] Quality: unit tests, CI (fmt, vet, test, build), typo sweep
- [x] Homebrew tap with a `dots` formula

</details>

<details>
<summary><strong>v0.2.0 — Ergonomics</strong></summary>

Make daily use frictionless; no architectural changes.

- [x] `-p` / `--print` — resolve the path without opening (scriptable)
- [x] Shell completion + dynamic app-name completion
- [x] Better unknown-app error with did-you-mean suggestion
- [x] `-e <editor>` one-shot override and `--dir` (open containing folder)
- [x] Root `Long` description + examples; minimal, `NO_COLOR`-aware styling

</details>

<details>
<summary><strong>v0.3.0 — Comment-safe writes & richer defaults</strong></summary>

Fix the write path before anything writes more, and go batteries-included.

- [ ] Surgical edits in `add` / `update` / `clear` that preserve comments and formatting
- [ ] Expanded curated default catalog, with per-OS paths
- [ ] `dots list --json`

</details>

<details>
<summary><strong>v0.4.0 — Discovery</strong></summary>

Depends on search paths and the catalog to match against.

- [ ] Configurable search paths
- [ ] `dots scan` / `discover` — surface configs found on disk and register them interactively

</details>

<details>
<summary><strong>v0.5.0 — Multiple configs & selection</strong></summary>

Reuses comment-safe writes (appending a second path) and adds a picker primitive.

- [ ] Multiple paths per app, with a picker when more than one exists
- [ ] `dots <app> --list` — choose among an app's configs
- [ ] `dots` with no args (`-i` / `pick`) — fuzzy-find across all apps

</details>

<details>
<summary><strong>v0.6.0 — Fetch & install</strong></summary>

- [ ] `dots get <url> [dest]` — download a config file from a URL
- [ ] `--as <app>` — download, register, and back up any existing file first

</details>

<details>
<summary><strong>v0.7.0+ — TUI</strong></summary>

- [ ] `dots tui` — full-screen interactive browser (reuses the v0.5 picker)

</details>

## Built with

- [Cobra](https://github.com/spf13/cobra) — command-line framework
- [go-toml](https://github.com/pelletier/go-toml) — TOML parsing

## Contributing

This is an early personal project, but issues, feedback, and ideas are welcome.
If you hit a bug or have a suggestion, open an issue.

## License

[MIT](LICENSE) © Navid Rajaei
