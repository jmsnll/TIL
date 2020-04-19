The `lsof` command is used to **l**i**s**t **o**pen **f**iles, which includes network connections.  So whenever your docker container or service fails to bind to a port because it is in use you can now find which process is using that port.

- Find the process that opened a local internet port:
`lsof -i :{{port}}`

- Find the process that opened a local TCP port:
`lsof -i TCP:{{port}}`