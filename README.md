# dom

dom - polish for home.

## Features

- Symlink dotfiles
- Run commands
- Create directories and files


## Usage

Install symlinks for all files managed by dom.

```sh
dom apply
```

Move `~/.zshrc` to `{dotfiles_repo}/config/zsh/zshrc` and replace it with a symlink to the new location. If no such file exists, create it first.

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

List all dotfiles managed by dom and whether or not they have been modified from last commit

```sh
dom status
```

## LICENSE

MIT
