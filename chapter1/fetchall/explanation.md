This program "fetchall", does the same as fetch of a URL’s contents as the previous example, but it fetches many URLs, all concurrently, so that the pro ess will take no longer than the longest fetch rather than the sum of all t e fetch times. This version of fetchall dis c ards t he resp ons es but rep or ts t he size and elapsed time for each one:

```
go run main.go https://golang.org http://gopl.io https://godoc.org
```
**Exercise 1.10:** Find a web site that produces a large amount of data. Investigate caching by running fetchall twice in succession to see whether the reported time changes much. Do you get the same content each time? Modify fetchall to print its output to a file so it can be examined.
- execute the above command - the output will be written to a file into the same directory.
**Exercise 1.11:** Try fetchall with longer argument lists, such as samples from the top million web sites available at alexa.com. How does the program behave if a web site just doesn’t respond?

- Start the server
```
go run server.go
```

Run the above command to server that doesn't respond - it will just waits for the response without time out.
```
go run main.go http://localhost:8000/noresponse
```