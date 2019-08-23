- Simply execute:
```
$ go run main.go
```
- Than add the following in stdin:
```
apples
coconuts
apples
bananas
apples
bananas
```
- Press Ctrl + D this tells the terminal that it should register a EOF on standard input, which bash interprets as a desire to exit.
- The result:
```
$ go run main.go
apples
coconuts
apples
bananas
apples
bananas
3D      apples
2       bananas
```