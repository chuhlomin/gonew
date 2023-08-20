# gonew

This repository contains templates for creating new Go projects with [`gonew` command](https://go.dev/blog/gonew).

## Templates

- [server](./server) - Go server
- [library](./library) - Go library
- [action](./action) - Dockerized GitHub Action

## Usage

Install `gonew`:

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

Create a new project: `gonew <template> <new-package>`

```bash
cd ~/Projects/username
gonew github.com/chuhlomin/gonew/server github.com/username/new-server
gonew github.com/chuhlomin/gonew/library github.com/username/new-library
gonew github.com/chuhlomin/gonew/action github.com/username/new-action
```

Or you may use [`new` script](https://github.com/chuhlomin/aliases/blob/main/new.sh):

```bash
cd ~/Projects/username
new server new-server
new library new-library
new action new-action
```
