This program introduces functions from two packages, net/http and io/ioutil. The http.Get function makes an HTTP request and, if there is no error, returns the result in the response struct resp. The Body field of resp contains the ser ver response as a readable stream. Next, ioutil.ReadAll reads the entire response; the result is stored in b. The Body st re am is clos e d to avoid leaking resources, and Printf writes the response to the standard output.

```
go run main.go http://gopl.io
```
- If the HTTP request fails, there will be an error message
```
$ go run main.go http://bad.gopl.io
fetch: Get http://bad.gopl.io: dial tcp: lookup bad.gopl.io: no such host
exit status 1
```
**Exercise 1.7:** The function call io.Copy(dst, src) reads from src and writes to dst. Use it instead of ioutil.ReadAll to copy the response body to os.Stdout without requiring a buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.

**Exercise 1.8:** Modify fetch to add the prefix http:// to each argument URL if it is missing. You might want to use strings.HasPrefix.

**Exercise 1.9:** Modify fetch to also print the HTTP status code, found in resp.Status.
