# gum

Gum is a tool for managing many git repositories at once. It's ideal for 
making mass edits, retrieving information from a common place in many repos, and
other large scale scripting tasks.

## Contents

- [Installation](#installation)
- [Basic Usage](#basic-usage)
- [Commands](#commands)
  - [Configuration Commands](#configuration-commands)
    - [prefix](#prefix)
    - [make](#make)
    - [group](#group)
    - [register](#register)
    - [unregister](#unregister)
    - [list](#list)
  - [Git Commands](#git-commands)
    - [clone](#clone)
    - [status](#status)
    - [add](#add)


## Installation

```sh
    git clone
    go install
```

## Basic Usage

The commands below show a basic workflow of:

- setting a prefix
- cloning and registering repos
- deleting a file from the repos
- committing and pushing changes

```sh
gum prefix git@github.com:zemberdotnet/

gum clone gum
gum register gum

gum clone blog
gum register blog

gum sh rm README.md
gum add README.md
gum commit -m "removing readme"
gum push
```

## Commands

### Configuration Commands

### prefix

- used to set a prefix for use in `gum clone`
- example usage:
  ```sh
  gum prefix git@github.com:zembderdotnet/
  ```

### make

`make` is a convience command to clone all the repos in the current group. Given
the following `.gum` file:

```json
{
  "Prefix": "git@github.com:zemberdotnet/",
  "CurrentGroup": "default",
  "Groups": {
    "default": ["hnreduce", "gum"],
    "otherGroup": ["blog"]
  }
}
```

make would clone `git@github.com:zemberdotnet/hnreduce` and `git@github.com:zemberdotnet/gum` to the local working directory. This command can be useful when setting up a working environment.

### group

`group` changes the active repo group. The default group is `default`. This allows 
switching between different groups of repos contained in the same directory.

### register

`register` adds a local repo to the current group. Given the active group is `default` and a directory like the following
```sh
ls -a
blog/ .gum gum/
```
Doing the `gum register blog` would add the `blog` repo to the current group, `default`.


### unregister

`unregister` removes a local repo from the current group.

### list

`list` prints a new-line delimited list of the repos registered in the current group.


### Git Commands

### clone
`clone` works the same way as `git clone <repo>`. It works in combination with
your current `prefix`, if you've set one.

Example usage with `gum prefix`:
```sh
gum prefix git@github.com:zemberdotnet/
gum clone gum
gum register gum
ls -a
.gum gum/
```
Example usage without using `gum prefix`:
```sh
gum clone git@github.com:zemberdotnet/gum
gum register gum
ls -a
.gum gum/
```

### status

`status` works the same `git status` does. It will print the `git status` for each
repo registered in the current group. It accepts the same flags that `git status`
accepts. Example usage to print the short-form status for all registered repos:
```sh
gum status -s
#target:gum
M README.md
#target:blog
M README.md
```

### add
`add` works the same as `git add <changed file>`. It will attempt to add the file
in every repository that is registered in the current group.

`add` also accepts the same flags that `git add` does. The following shows a common
example of `add`:
```sh
gum sh rm README.md
gum status
#target:gum
On branch main
Your branch is up to date with 'origin/main'.

Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        deleted:    README.md

no changes added to commit (use "git add" and/or "git commit -a")
#target:blog
On branch main
Your branch is up to date with 'origin/main'.

Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        deleted:    README.md

no changes added to commit (use "git add" and/or "git commit -a")
gum add -A
gum commit -m "remove README.md"
gum push
```


- commit
- push
- sh
- checkout
- pull
- switch
- stash
