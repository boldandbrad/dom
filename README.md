> [!IMPORTANT] **dom** is currently in early development. This means it is
> missing key features and is highly unstable. Do not use it to manage critical
> files. Consider dropping a star to follow dom's progress.

# dom

**dom** (Polish for "home") is a configurable, symlink-based dotfile manager. It
helps you back up and reuse the most important files in your $HOME directory
with ease.

## Goals

- Scriptable single binary install via curl or homebrew
- Do not dictate the structure of the dotfile repository
- Backup system agnostic (git, dropbox, etc)

## Features

- Symlink dotfiles
- Run custom scripts
- Create directories and files
- CLI to automatically add and drop files from the configuration
- Dry-run mode to visualize changes before applying
- Conditional directives (tbd)
- Profiles for separating personal vs work or handling system specific configs
  (tbd)

## Install

Homebrew or curl

> Coming soon.

## Getting started

### Starting fresh

Create a dotfiles directory.

Add existing files to the default config.

Apply changes.

### Applying an existing config

Clone or sync your existing config from its source (git, dropbox, etc)

Apply the configuration to the system.

### Migrating from another dotfile manager

## Usage

Apply the default configuration to the system.

```sh
dom apply
```

Add `~/.zshrc` to the default configuration. This copies `~/.zshrc` to
`{dotfiles_dir}/config/zsh/zshrc`. If no such file exists, create an empty copy
in the dotfile dir. If only one argument is provided, add the file to the repo
root.

```sh
dom add ~/.zshrc config/zsh/zshrc
```

Replace the relevant symlink with the actual file and delete the repo copy.

```sh
dom drop config/zsh/zshrc
```

Rename or move the repo copy of the file and update the relevant symlink

```sh
dom move config/zsh/zshrc config/zsh/.zshrc
```

Replace dom managed symlinks with copies of actual files from the repo.

```sh
dom uninstall
```

List all dotfiles managed by dom and whether or not they have been modified from
last commit

```sh
dom status
```

## Configuration

> Coming soon.

## Examples

> Coming soon.

## Inspiration

- mackup
- dotbot
- chezmoi

## Legal

Copyright (c) 2024 Bradley Wojcik. Released under the MIT License. See
[LICENSE](LICENSE) for details.
