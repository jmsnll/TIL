[Brace Expansions](https://www.gnu.org/software/bash/manual/html_node/Brace-Expansion.html) can be very helpful whenever typing to manipulate paths inside a terminal.

- Batch creation of nested directories:

```bash 
mkdir -p 1.0.0-{release,src}/{package,dist,logs}
mkdir -p {1.0.0-{release,src},0.9.1-{beta,src}}
``` 

- Quickly replacing parts of a files name or path:

```bash 
# shorthand
mv worker-20200222103944.log{.bak,} # replaces the `.bak` extension with ``, removing it
mv version-info-{1.2.43,1.3.0}.md   # replaces `1.2.43` with `1.3.0`

# expands to 
mv worker-20200222103944.log.bak worker-20200222103944.log
mv version-info-1.2.43.md version-info-1.3.0.md
``` 
