Sometimes you may want to debug a Python script in a specific environment, be it on a remote server or inside a docker container.  Without going through the hassle of setting up remote debugging through an IDE like PyCharm you can instead use the built in debugging module in the Python standard library `pdb`.

First you place a call to `pdb.set_trace()` in the location of where you would like your break point, you can simplify this by including the `pdb` import on the same line:

```python
def _add_result_to_targets(self, targets, result):
    name = self._new_name()
    self.__printer.singleton_pprinters.setdefault(
        id(result), lambda obj, p, cycle: p.text(name)
    )

    # import and attach python debugger at the current frame on the stack
    import pdb; pdb.set_trace()  

    self.names_to_values[name] = result  # <-- debug cursor begins
    for target in targets:
        self.bundles.setdefault(target, []).append(VarReference(name))
```

Once set, you can then call your program interactively through the terminal and you will now be in an interactive debugging session:

```bash
> hypothesis-python/src/hypothesis/stateful.py(358)_add_result_to_targets()
-> self.names_to_values[name] = result
(Pdb) >
```

All of the commands are available through the use of `help` and `help <topic>`.