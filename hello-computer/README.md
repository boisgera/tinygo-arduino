
```go
package main

func main() {
    println("Computer says hello! 👋")
}
```

```bash
$ go run app.go 
Computer says hello! 👋
```

```bash
$ tinygo run app.go 
Computer says hello! 👋
```

```bash
$ go build app.go
$ ls
app  app.go  images  README.md
$ ./app
Computer says hello! 👋
$ ls -sh app
1,2M app
```

```bash
$ tinygo build app.go
$ ls
app  app.go  images  README.md
$ ./app
Computer says hello! 👋
$ ls -sh app
68K app
```

That's a ~20x size reduction!

![jsdklsjd](images/desktop_with_arduino_ide.svg)

![jsdklsjd](images/laptop_with_arduino_ide.svg)
