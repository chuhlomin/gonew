# gonew

This repository contains templates for creating new Go projects with [`gonew` command](https://go.dev/blog/gonew).

## Example

Install `gonew`:

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

Create a new project: `gonew <template> <new-package>`

```bash
$ cd ~/Projects/username
$ gonew github.com/chuhlomin/gonew/server github.com/username/new-service

gonew: initialized github.com/username/new-service in ./new-service
```
