[![Build Status](https://travis-ci.org/iamalirezaj/go-respond.svg?branch=develop)](https://travis-ci.org/iamalirezaj/go-respond)
[![License](https://img.shields.io/:license-mit-blue.svg?style=flat-square)](#license)


# Go Respond

This package is provided to be used on golang and it gives clean methods to handle json response with specific predetermined messages.

# Requirement
* Golang 1.10

**The package is in process.**

## Install

Run this commands

    $ go get github.com/iamalirezaj/go-respond

## Usage

You can use these methods in deffernt ways:

There are hot ones for quick usage, besides some provided to manage outputs on your own way

import package
```go
import "github.com/iamalirezaj/go-respond"
```

create respond instance
```go
var respond josh.Respond
```

**Some are shown below:**

When request succeeds and contains data to return as a result:
```go
respond.Succeed(map[string] interface{} {
    "some_key": "some_data"
})
```

When deletion action succeeds:
```go
respond.DeleteSucceeded()
```

When updating succeeds:
```go
respond.UpdateSucceeded()
```

When insertion succeeds:
```go
respond.InsertSucceeded()
```

When deletion action fails:
```go
respond.DeleteFaild()
```

When updating fails:
```go
respond.UpdateFaild()
```

when insertion fails:
```go
respond.InsertFaild()
```

Not Found Error:
```go
respond.NotFound()
```

When parameters entered are wrong:
```go
respond.WrongParameters()
```

When requested method is not allowed:
```go
respond.MethodNotAllowed()
```

```go
respond.RequestFieldNotFound()
```

Validation errors:
```go
respond.ValidationErrors(map[string] interface{} {
    "some_key": "some_validation_errors_data"
})
```

###customization
You can do more:
```go
respond.SetStatusCode(200).setStatusText('Success.').RespondWithMessage('Your custom message')
```

## License
The MIT License (MIT). Please see [License File](LICENSE.md) for more information.

