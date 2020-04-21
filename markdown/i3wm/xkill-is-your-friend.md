You can bind the `xkill` command to a key for quickly terminating a non-responsive program through X.

Add the following to your i3 `config` file:

```
bindsym $mod+x --release exec --no-startup-id xkill
```

When pressed `xkill` will be launch and whichever window you next click on will be terminated immediately, with `xkill` exiting afterwards.