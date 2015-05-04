# gen-validator

Generate validation methods using `go:generate`.

# Installation

```
$ go get github.com/bluele/gen-validator
```

# Usage

```
$ gen-validator
Usage of gen-validator:
        validator [flags] -model T [directory]
        validator [flags[ -model T files... # Must be a single package
For more information, see:
Flags:
  -model="": comma-separated list of model names; must be set
  -output="": output file name; default srcdir/<type>_validator.go
```

# Example

```
// model.go
package models

import (
  "errors"
  "fmt"
)

//go:generate gen-validator -model=User
type User struct {
  Name string `validator:"min=3,max=12"`
  Age  int    `validator:"min=18"`
}

func (obj *User) ErrorMaxName() error {
  return errors.New("Ensure this value is less than or equal to 12")
}

func (obj *User) ErrorMinAge() error {
  return errors.New("Ensure this value is greater than or equal to 18")
}
```

```
$ ls
models.go

$ go generate ./...

$ ls
models.go models_validator.go
```

```
// models_validator.go
// generated by validator -model=User; DO NOT EDIT

package models

import "github.com/bluele/gen-validator/libs"

func (obj *User) Validate() []error {
  var err error
  var errors []error

  err = libs.Min(obj.Name, "3")
  if err != nil {
    errors = append(errors, err)
  }
  err = libs.Max(obj.Name, "12")
  if err != nil {
    errors = append(errors, obj.ErrorMaxName())
  }
  err = libs.Min(obj.Age, "21")
  if err != nil {
    errors = append(errors, obj.ErrorMinAge())
  }

  return errors
}
```

# Author

**Jun Kimura**

* <http://github.com/bluele>
* <junkxdev@gmail.com>