# field

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Nadim147c/field?style=for-the-badge&logo=go&labelColor=11140F&color=BBE9AA)](https://pkg.go.dev/github.com/Nadim147c/field)
[![GitHub Repo stars](https://img.shields.io/github/stars/Nadim147c/field?style=for-the-badge&logo=github&labelColor=11140F&color=BBE9AA)](https://github.com/Nadim147c/field)
[![GitHub License](https://img.shields.io/github/license/Nadim147c/field?style=for-the-badge&logo=gplv3&labelColor=11140F&color=BBE9AA)](./LICENSE)
[![GitHub Tag](https://img.shields.io/github/v/tag/Nadim147c/field?include_prereleases&sort=semver&style=for-the-badge&logo=git&labelColor=11140F&color=BBE9AA)](https://github.com/Nadim147c/field/tags)

> [!IMPORTANT]
> 🔥 Found this useful? A quick star goes a long way.

`field` is a easy, elegant, ~minimal~ and flexible way to parse and access fields
from stream of text.

![Made with VHS](https://vhs.charm.sh/vhs-18AOVv7alBxt36JtlnV0h2.gif)

## Features

- Extract specific fields from each input line.
- Support for custom delimiters.
- Option to ignore empty lines.
- Limit the number of fields processed.
- Support multiple field selections.

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
- Install shell completions for Bash, Zsh, and Fish

## Usage

```bash
field [flags] ...<range>
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
