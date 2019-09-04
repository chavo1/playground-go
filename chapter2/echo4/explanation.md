- Test:
```
$ go build gopl.io/ch2/echo4
$ ./echo4 a bc def
a bc def
$ ./echo4 -s / a bc def
a/bc/def
$ ./echo4 -n a bc def
a bc def$
$ ./echo4 -help

Usage of ./echo4:
       -n    omit trailing newline
       -s string
             separator (default " ")
```
