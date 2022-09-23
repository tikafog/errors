# Errors #

## Errors is an object for modular error management package ##

### Useages ###

```go
var err = errors.RegisterModule("name")
var (
ErrFirstErr = err.New("first err")
ErrSecondErr = err.Errorf("%v", "second err")
)

func doSomething() error {
//...
return ErrFirstErr
}

func doSomething() error {
//...
return err.New("third err")
}

func doSomething() error {
//this will only return an error wrapped from errors
return errors.New("fourth err")
}

func doSomething() error {
//...
err := doAnything()
return err.Wrap(err, "fifth err")
}

func doSomething() error {
//...
err := doAnything()
return err.WrapIndex(err, ErrFirstErr) Index
}


```

if used module errors, the error will print like this:

```shell
    Module[name]: first err
```

otherwise,the error will print like this:

```shell
    fourth err
```

### Other ###

there two functions will register errors to the unknown module

```
func IndexNew(str string) Index 
func IndexErrorf(format string, args ...interface{}) 
```

others are used same as errors

