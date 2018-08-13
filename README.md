# Facebook-Go

Facebook-Go is a [Facebook Graph API](https://developers.facebook.com/docs/graph-api) client written in Go. Currently it can:

- Get Me

## Installation

```bash
$ go get github.com/swensonhe/facebook-go
```

## Examples

### Get Me

```go
package main

import (
    "github.com/swensonhe/facebook-go"
    "fmt"
)

func main() {
    client := facebook.NewClient()
    user, err := client.GetMe("token", "email", "first_name", "last_name", "id")
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Println(user)
}

```
