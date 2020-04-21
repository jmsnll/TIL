If you ever need to quickly spin up an HTTP server to serve files from a directory you can make use the `http.server` module (or the `SimpleHTTPServer` if you're still on Python 2).  The module defaults to using port `8000` but a different port can be specified as a parameter:

```bash
$ python -m http.server 8080
Serving HTTP on 0.0.0.0 port 8080 (http://0.0.0.0:8080/) ...
127.0.0.1 - - [21/Apr/2020 01:14:26] "GET / HTTP/1.1" 200 -
127.0.0.1 - - [21/Apr/2020 01:14:26] code 404, message File not found
127.0.0.1 - - [21/Apr/2020 01:14:26] "GET /favicon.ico HTTP/1.1" 404 -
127.0.0.1 - - [21/Apr/2020 01:14:28] "GET /README.md HTTP/1.1" 200 -
```

It will serve files from the current working directory, with the default page being `index.html` otherwise falling back to a simple directory listing.