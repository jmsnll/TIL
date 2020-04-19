The `syscall.Exec` function can be used to execute an external program. Instead of forking a child process though, it runs the external command in place of the current process. You need to give the function three pieces of information: the location of the binary, the arguments to be executed and the relevant environment.

```go
package main

import (
    "fmt"
    "os"
    "syscall"
)

func main() {
    // get the system's environment variables
    environment := os.Environ()

    // get a slice of the pieces of the command 
    command := []string{"tmux", "new-session", "-s", "burrito"}

    err := syscall.Exec("/usr/local/bin/tmux", command, environment)
    if err != nil {
        panic(err)
    }
}
```
