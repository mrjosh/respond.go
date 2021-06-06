# Go Respond
This package is provided to be used on golang and it gives clean methods to handle json response with specific predetermined messages.

## Install
$ go get github.com/mrjosh/respond.go

## Usage
import package
```go
import "github.com/mrjosh/respond.go"
```

create respond instance
```go

// argument rw is http.ResponseWriter or 
// gin.Writer if youre using gin-gonic
var jspon = respond.NewWithWriter(rw)

// or if you want to use custom languages
var jspon = respond.NewWithWriter(rw).Language("fa")
```

**Some are shown below:**

When request succeeds and contains data to return as a result:
```go
jspon.Succeed(map[string] interface{} {
  "some_key": "some_data"
})
```

When deletion action succeeds:
```go
jspon.DeleteSucceeded()
```

When updating succeeds:
```go
jspon.UpdateSucceeded()
```

When insertion succeeds:
```go
jspon.InsertSucceeded()
```

When deletion action fails:
```go
jspon.DeleteFaild()
```

When updating fails:
```go
jspon.UpdateFaild()
```

when insertion fails:
```go
jspon.InsertFaild()
```

Not Found Error:
```go
jspon.NotFound()
```

When parameters entered are wrong:
```go
jspon.WrongParameters()
```

When requested method is not allowed:
```go
jspon.MethodNotAllowed()
```

```go
jspon.RequestFieldNotFound()
```

Validation errors:
```go
jspon.ValidationErrors(map[string] interface{} {
  "some_key": "some_validation_errors_data"
})
```

###customization
You can do more:
```go
jspon.SetStatusCode(200).setStatusText('Success.').RespondWithMessage('Your custom message')
```

## License
The MIT License (MIT). Please see [License File](LICENSE.md) for more information.

