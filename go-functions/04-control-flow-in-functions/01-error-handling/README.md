# Error Handling

Whenever we have a function which can return an error like so:

```
_, err := r.Read([]byte("test something"));
```

We can handle this error by checking for the err value:

```
if err != nil {
    fmt.Println("An error occured!", err)
    return err
}
```
