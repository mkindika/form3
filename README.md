
# form3-account-sdk #

form3-account-sdk is a Go client library for accessing the [Form3 API v1][].
Form3 Take Home Exercise - Indika Munaweera

## Disclaimer ##

Please note that this is my first Go project and I have been working to improve my Go language skills. 
This client libary provides [Fetch][], [Create][] and [Delete][] operations on the account resource.

## Installation ##

form3-account-sdk is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/mkindika/form3-account-sdk/v1
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/mkindika/form3-account-sdk/v1"
```

and run `go get` without parameters.

## Usage ##

```go
import "github.com/mkindika/form3-account-sdk/v1/form3"	
import "github.com/google/form3-account-sdk/form3" // with go modules disabled
```

Construct a new Form3 API client, then use the account service on the client to
fetch, create and delete an account. For example:

```go
client := form3.NewClient()

// Create a new account
response, _, _ := c.Account.Create(ctx, accountData)
```

You can set the base URL, agent and provide a http client. For example:

```go
\\ Update base URL. Default base url is http://localhost:8080
\\ Do not provide trailing \
client := form3.NewClient().WithBaseURL("http://localhost:9090")

\\ Update user agent. Default is form3.
client := form3.NewClient().WithUserAgent("http://localhost:9090")

\\ Also you can provide a http client with own configurations
httpClient := &http.Client{}
client := form3.NewClient().WithHttpClient(httpClient)
```

### Integration Tests ###
These test are excuted against the dockerized mock account API. 

[Form3 API v1]: https://api-docs.form3.tech/api.html
[Fetch]: https://api-docs.form3.tech/api.html#organisation-accounts-fetch
[Create]: https://api-docs.form3.tech/api.html#organisation-accounts-create
[Delete]: https://api-docs.form3.tech/api.html#organisation-accounts-delete