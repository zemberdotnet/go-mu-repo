# gum

Gum is a tool for managing many git repositories as once.

## Installation

```sh
    git clone 
    go install
```



## Commands
- prefix
  - used to set a prefix for use in `gum clone`
- clone
  - the same as `git clone <repo>`
- commit
- add
- push
- status
- sh
- register
- unregister
- checkout
- pull
- make
  - Will `clone` all repos in the active group that don't exist already. Will use active `prefix`.
- group
  - Changes the active repo group. The default group is `default`.

