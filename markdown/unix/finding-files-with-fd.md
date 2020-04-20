[`sharkdp/fd`](https://github.com/sharkdp/fd) is a simple, fast and user-friendly alternative to the `find` command typically found on unix systems.

While it does not seek to mirror all of find's powerful functionality, it provides sensible (opinionated) defaults for 80% of the use cases.

Find files matching the given pattern in the current directory, `fd` always performs a regex search and any regex expression can be used:

```
fd {{pattern}}
```

Find files matching a pattern with a specific extension:

```
fd -e {{extension}} {{pattern}}
```

Find files in a specific directory:

```
fd {{pattern}} {{path/to/dir}}
```

### Integration with `fzf` 

Using fd with fzf

You can use fd to generate input for the command-line fuzzy finder fzf:

```
export FZF_DEFAULT_COMMAND='fd --type file'
export FZF_CTRL_T_COMMAND="$FZF_DEFAULT_COMMAND"
```

Then, you can type vim <Ctrl-T> on your terminal to open fzf and search through the fd-results.

Alternatively, you might like to follow symbolic links and include hidden files (but exclude .git folders):

```
export FZF_DEFAULT_COMMAND='fd --type file --follow --hidden --exclude .git'
```

You can even use fd's colored output inside fzf by setting:

```
export FZF_DEFAULT_COMMAND="fd --type file --color=always"
export FZF_DEFAULT_OPTS="--ansi"
```

For more details, see the Tips section of the fzf README.