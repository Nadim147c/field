# field

[![Go Reference](https://pkg.go.dev/badge/github.com/Nadim147c/field.svg)](https://pkg.go.dev/github.com/Nadim147c/field)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0.en.html)

`field` is a easy, elegant, ~minimal~ and flexible way to parse and access fields
from stream of text.

## Features

- Extract specific fields from each input line.
- Support for custom delimiters.
- Option to ignore empty lines.
- Limit the number of fields processed.
- Support multiple field selections.
- Auto-completion for Bash, Zsh, and Fish shells.

## Installation

### From source

Clone and build:

```bash
git clone https://github.com/Nadim147c/field.git
cd field
make
make install PREFIX=$HOME/.local
```

> You can also set `PREFIX=/usr` to install the binary to /usr/bin.

This will:

- Build the binary (`field`)
- Install the binary to `~/.local/bin` (or `$PREFIX/bin`)
- Install the LICENSE file
- Install shell completions for Bash, Zsh, and Fish

## Usage

```bash
field [flags] ...<range>
```

### Examples

```bash
# Extract the second field from ps output (PID) and kill those processes
ps aux | grep bad-process | field 2 | xargs kill

# Print only the usernames (first field) from /etc/passwd and ignore empty lines
cat /etc/passwd | field -i -d: 1

# Show just the command (limit number of field 11) from ps output
ps aux | field -n11 11

# Extract multiple fields (user and PID) and print them
ps aux | field 1 2
```

---

## Shell Completion

`field` supports shell completions via [Carapace](https://github.com/rsteube/carapace):

```bash
# Bash
source <(field _carapace bash)

# Zsh
source <(field _carapace zsh)

# Fish
field _carapace fish | source
```

## Contributing

Contributions are welcome! To contribute:

1. Fork and Clone the repository.
2. Create a new branch (`git checkout -b feature/my-feature`).
3. Make your changes.
4. Run tests and linters:

```bash
make lint
```

5. Submit a pull request.

Please follow the existing code style and write clear commit messages.

## License

`field` is licensed under the **GPL-3.0 License**. See [LICENSE](LICENSE) for details.
