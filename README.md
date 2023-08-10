# gonew

This repository contains templates for creating new Go projects with [`gonew` command](https://go.dev/blog/gonew).

## Templates

- [server](./server) - a new Go server project.
- [library](./library) - a new Go library project.

## Usage

Install `gonew`:

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

Create a new project: `gonew <template> <new-package>`

```bash
$ cd ~/Projects/username
$ gonew github.com/chuhlomin/gonew/server github.com/username/new-server

gonew: initialized github.com/username/new-server in ./new-server
```

Or you may use [`new` script](https://github.com/chuhlomin/aliases/blob/main/new.sh):

```bash
$ cd ~/Projects/username

$ new server new-server
gonew: initialized github.com/username/new-server in ./new-server

$ new library new-library
gonew: initialized github.com/username/new-library in ./new-library
```
