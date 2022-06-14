# gum

Gum is a tool for managing many git repositories at once.

## Installation

```sh
    git clone
    go install
```

## Commands

### Configuration Commands

- prefix
  - used to set a prefix for use in `gum clone`
- make
  - Will `clone` all repos in the active group that don't exist already. Will use active `prefix`.
- group
  - Changes the active repo group. The default group is `default`.
- register
  - adds a local repo to the current group
- unregister
  - removes a local repo from the current group

### Git Commands
- clone
  - the same as `git clone <repo>`
- commit
- add
- push
- status
- sh
- checkout
- pull
- switch
- stash
