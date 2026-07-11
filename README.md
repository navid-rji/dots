# dots

## Roadmap

### Phase 0 (MVP):

- [x] App registry: internal model mapping app name -> config path(s) + defaults
- [x] dots <app> : resolves the app and opens the config in the editor
- [x] editor: cmd to open the config in editor of choice can be set by user
- [x] dots own config file: ~/.config/dots/config.toml
- [x] dots list / ls: show all known apps, their resolved paths, and whether each file actually exists
- [x] dots add <app> <path>: register a custom mapping
- [x] dots dots / dots edit: open dot's own config
- [ ] dots remove <app>: remove a mapping -> if there is a default exsiting resetoring it
- [ ] protect the dots config file from accidental deletion / overwriting

### Phase 1 (Discovery and custom directories):

- [ ] Configurable search paths
- [ ] dots scan / discover: walk the search paths + known locations and show what it finds -> let the user register discovered configs into the registry

### Phase 2 (Interactive picker):

- [ ] dots <app> : if multiple configs are found, show an interactive picker to choose which one to open
- [ ] dots <app> --list: show all known configs for the app and let the user choose which one to open
- [ ] dots without args (or -i / dots pick): an interactive fuzzy finder over all registered apps, filter by typing, hit enter to open. (maybe vim bindings / arrow keys for navigating the list)

### Phase 3 (Downloads):

- [ ] dots get <url> [dest]: fetch file from URL
- [ ] --as <app>: download and register as an app. Backup before overwrite

### Phase 4 (TUI):
- [ ] dots tui: full screen interface



## Libraries recommended by claude:

- CLI framework: Cobra (spf13/cobra) paired with Viper (spf13/viper) for config management
- TUI: Bubble Tea (charmbracelet/bubbletea) plus Lip Gloss (charmbracelet/lipgloss) for styling and Bubbles (charmbracelet/bubbles) for common components like lists, tables, and text input
- Downloads: Go's built-in net/http package



