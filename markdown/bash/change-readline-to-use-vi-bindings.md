By default most shells will default to using the Emacs key bindings which is all well and good if you're familiar with Emacs or otherwise find them convient to use, there are however others who are more comfortable using vi-like keyboard bindings.

In order to enable vi mode you can execute the command:

```bash
set -o vi 
```

Which then, just like magic, will put you into the vi mode.  Add to your associated shell's run commands file to enable it by default (e.g. `.bashrc`, `.zshrc`)
