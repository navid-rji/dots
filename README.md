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
| `dots clear <app>` | Remove an app's mapping (restores built-in default if one exists) |
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

<details open>
<summary><strong>v0.1.0 — First public release</strong></summary>

Get the foundation and safety right.

- [x] App registry mapping app name → config path(s), with built-in defaults
- [x] `dots <app>` — resolve an app and open its config in your editor
- [x] `dots list` / `ls`, `add`, `update`, `clear`, `dots dots`
- [x] `--version` with ldflags + `debug.ReadBuildInfo` fallback
- [x] Fix first-run so the chosen editor is used immediately
- [x] Guard the reserved name `dots` in `add` / `update`
- [x] Atomic config writes (temp file + rename)
- [ ] `dots list` shows whether each config file exists on disk
- [ ] `SilenceUsage` + `SilenceErrors`; single error printer in `main`
- [ ] Unit tests for pure functions + minimal CI (fmt, vet, test, build)
- [ ] Homebrew tap (`homebrew-dots`) + README install / OS-support note
- [ ] Typo sweep

</details>

<details>
<summary><strong>v0.2.0 — Ergonomics</strong></summary>

Make daily use frictionless; no architectural changes.

- [ ] `-p` / `--print` — resolve the path without opening (scriptable)
- [ ] Shell completion + dynamic app-name completion
- [ ] Better unknown-app error with did-you-mean suggestion
- [ ] `-e <editor>` one-shot override and `--dir` (open containing folder)
- [ ] Root `Long` description + examples; minimal, `NO_COLOR`-aware styling
- [ ] `dots doctor` — check config, editor on `PATH`, path existence

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
