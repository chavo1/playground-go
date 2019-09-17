Within a Go program, every package is identified by a unique string called its import path. 
These are the strings that appear in an import declaration like "../tempconv".
The language specification doesn’t define where these strings come from or what they mean; it’s up to the tools to interpret them.
```
$ go run main.go 32
32°F = 0°C, 32°C = 89.6°F
$ go run main.go 212
212°F = 100°C, 212°C = 413.6°F
$ go run main.go -40
-40°F = -40°C, -40°C = -40°F
```