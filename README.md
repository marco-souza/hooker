# hooker

A simple git hook manager written in go (an alternative to husky)

## Install

```sh
# install the latest version
go install github.com/marco-souza/hooker@latest
# install an specific version
go install github.com/marco-souza/hooker@v0.0.4
```

## Usage

```sh
hooker -h # help

# init your git repo
hooker init

# add pre-commit hook
hooker add pre-commit echo Hello everyone

# list all installed hooks
hooker list

# drop pre-commit hook
hooker drop pre-commit

# drop hooker
hooker drop
```
