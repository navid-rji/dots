package registry

import "github.com/navid-rji/dots/internal/config"

// Disclaimer: The defaults in this file were created by Claude and not checked

// defaults is the built-in app registry: a best-guess config path for each
// well-known program. Anything the user registers themselves overrides these,
// and `use_defaults = false` drops the whole set.
//
// Rules of thumb for entries here:
//   - one path per app, the most common location on a modern Linux/macOS box
//   - prefer the XDG location when upstream has actually moved to it
//   - the name is what the user types, so keep it short and obvious
func defaults() map[string]config.App {
	return map[string]config.App{
		// --- Shells & shell tooling ---------------------------------------
		"bash":     {Paths: []string{"~/.bashrc"}},
		"zsh":      {Paths: []string{"~/.zshrc"}},
		"fish":     {Paths: []string{"~/.config/fish/config.fish"}},
		"nushell":  {Paths: []string{"~/.config/nushell/config.nu"}},
		"elvish":   {Paths: []string{"~/.config/elvish/rc.elv"}},
		"xonsh":    {Paths: []string{"~/.xonshrc"}},
		"profile":  {Paths: []string{"~/.profile"}},
		"inputrc":  {Paths: []string{"~/.inputrc"}},
		"starship": {Paths: []string{"~/.config/starship.toml"}},
		"atuin":    {Paths: []string{"~/.config/atuin/config.toml"}},
		"direnv":   {Paths: []string{"~/.config/direnv/direnv.toml"}},

		// --- Terminal emulators -------------------------------------------
		"alacritty": {Paths: []string{"~/.config/alacritty/alacritty.toml"}},
		"kitty":     {Paths: []string{"~/.config/kitty/kitty.conf"}},
		"ghostty":   {Paths: []string{"~/.config/ghostty/config"}},
		"wezterm":   {Paths: []string{"~/.config/wezterm/wezterm.lua"}},
		"foot":      {Paths: []string{"~/.config/foot/foot.ini"}},
		"rio":       {Paths: []string{"~/.config/rio/config.toml"}},

		// --- Multiplexers -------------------------------------------------
		"tmux":   {Paths: []string{"~/.tmux.conf"}},
		"zellij": {Paths: []string{"~/.config/zellij/config.kdl"}},
		"screen": {Paths: []string{"~/.screenrc"}},

		// --- Editors ------------------------------------------------------
		"nvim":    {Paths: []string{"~/.config/nvim/init.lua"}},
		"vim":     {Paths: []string{"~/.vimrc"}},
		"helix":   {Paths: []string{"~/.config/helix/config.toml"}},
		"emacs":   {Paths: []string{"~/.emacs.d/init.el"}},
		"kakoune": {Paths: []string{"~/.config/kak/kakrc"}},
		"micro":   {Paths: []string{"~/.config/micro/settings.json"}},
		"zed":     {Paths: []string{"~/.config/zed/settings.json"}},
		"nano":    {Paths: []string{"~/.nanorc"}},

		// --- Version control & dev tooling --------------------------------
		"git":       {Paths: []string{"~/.gitconfig"}},
		"gitignore": {Paths: []string{"~/.config/git/ignore"}},
		"lazygit":   {Paths: []string{"~/.config/lazygit/config.yml"}},
		"tig":       {Paths: []string{"~/.tigrc"}},
		"gh":        {Paths: []string{"~/.config/gh/config.yml"}},
		"ssh":       {Paths: []string{"~/.ssh/config"}},
		"gpg":       {Paths: []string{"~/.gnupg/gpg.conf"}},
		"docker":    {Paths: []string{"~/.docker/config.json"}},
		"kube":      {Paths: []string{"~/.kube/config"}},
		"k9s":       {Paths: []string{"~/.config/k9s/config.yaml"}},
		"aws":       {Paths: []string{"~/.aws/config"}},
		"npm":       {Paths: []string{"~/.npmrc"}},
		"cargo":     {Paths: []string{"~/.cargo/config.toml"}},
		"pip":       {Paths: []string{"~/.config/pip/pip.conf"}},
		"go":        {Paths: []string{"~/.config/go/env"}},
		"gradle":    {Paths: []string{"~/.gradle/gradle.properties"}},
		"mise":      {Paths: []string{"~/.config/mise/config.toml"}},
		"asdf":      {Paths: []string{"~/.asdfrc"}},
		"curl":      {Paths: []string{"~/.curlrc"}},
		"wget":      {Paths: []string{"~/.wgetrc"}},
		"tealdeer":  {Paths: []string{"~/.config/tealdeer/config.toml"}},

		// --- CLI utilities ------------------------------------------------
		"bat":       {Paths: []string{"~/.config/bat/config"}},
		"btop":      {Paths: []string{"~/.config/btop/btop.conf"}},
		"htop":      {Paths: []string{"~/.config/htop/htoprc"}},
		"ripgrep":   {Paths: []string{"~/.config/ripgrep/config"}},
		"yazi":      {Paths: []string{"~/.config/yazi/yazi.toml"}},
		"ranger":    {Paths: []string{"~/.config/ranger/rc.conf"}},
		"lf":        {Paths: []string{"~/.config/lf/lfrc"}},
		"vifm":      {Paths: []string{"~/.config/vifm/vifmrc"}},
		"fastfetch": {Paths: []string{"~/.config/fastfetch/config.jsonc"}},
		"neofetch":  {Paths: []string{"~/.config/neofetch/config.conf"}},
		"newsboat":  {Paths: []string{"~/.config/newsboat/config"}},
		"zathura":   {Paths: []string{"~/.config/zathura/zathurarc"}},
		"mpv":       {Paths: []string{"~/.config/mpv/mpv.conf"}},
		"mpd":       {Paths: []string{"~/.config/mpd/mpd.conf"}},
		"ncmpcpp":   {Paths: []string{"~/.config/ncmpcpp/config"}},

		// --- Window managers & compositors --------------------------------
		"hyprland": {Paths: []string{"~/.config/hypr/hyprland.conf"}},
		"hyprlock": {Paths: []string{"~/.config/hypr/hyprlock.conf"}},
		"sway":     {Paths: []string{"~/.config/sway/config"}},
		"i3":       {Paths: []string{"~/.config/i3/config"}},
		"niri":     {Paths: []string{"~/.config/niri/config.kdl"}},
		"river":    {Paths: []string{"~/.config/river/init"}},
		"awesome":  {Paths: []string{"~/.config/awesome/rc.lua"}},
		"qtile":    {Paths: []string{"~/.config/qtile/config.py"}},
		"bspwm":    {Paths: []string{"~/.config/bspwm/bspwmrc"}},
		"openbox":  {Paths: []string{"~/.config/openbox/rc.xml"}},
		"xmonad":   {Paths: []string{"~/.xmonad/xmonad.hs"}},

		// --- Bars, launchers, notifications, X ----------------------------
		"waybar":     {Paths: []string{"~/.config/waybar/config"}},
		"polybar":    {Paths: []string{"~/.config/polybar/config.ini"}},
		"rofi":       {Paths: []string{"~/.config/rofi/config.rasi"}},
		"wofi":       {Paths: []string{"~/.config/wofi/config"}},
		"dunst":      {Paths: []string{"~/.config/dunst/dunstrc"}},
		"mako":       {Paths: []string{"~/.config/mako/config"}},
		"sxhkd":      {Paths: []string{"~/.config/sxhkd/sxhkdrc"}},
		"picom":      {Paths: []string{"~/.config/picom/picom.conf"}},
		"swaylock":   {Paths: []string{"~/.config/swaylock/config"}},
		"xinit":      {Paths: []string{"~/.xinitrc"}},
		"xresources": {Paths: []string{"~/.Xresources"}},
		"gtk3":       {Paths: []string{"~/.config/gtk-3.0/settings.ini"}},
		"gtk4":       {Paths: []string{"~/.config/gtk-4.0/settings.ini"}},

		// --- macOS --------------------------------------------------------
		"yabai":     {Paths: []string{"~/.config/yabai/yabairc"}},
		"skhd":      {Paths: []string{"~/.config/skhd/skhdrc"}},
		"aerospace": {Paths: []string{"~/.config/aerospace/aerospace.toml"}},
		"karabiner": {Paths: []string{"~/.config/karabiner/karabiner.json"}},
	}
}
