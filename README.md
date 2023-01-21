# Delbra

Interactive CLI tool that makes it comfortable to delete multiple local Git branches at once.

## Demo:

![Demo](delbra.gif 'Demo')

## Install:

Using Brew:

```bash
brew tap okradze/delbra --custom-remote https://github.com/okradze/delbra.git
brew install okradze/delbra/delbra
```

## Usage:

Run `delbra` in working directory of your Git repository.

### Keybindings:

- `j` or `ArrowDown` - move selection down
- `k` or `ArrowUp` - move selection up
- `Space` - select/deselect branch
- `Enter` - delete selected branches
- `q`, `esc`, `ctrl+c` - quit
