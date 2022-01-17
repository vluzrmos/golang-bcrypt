## Golang Bcrypt

Generate bcrypt password from command line, input typing is hidden.


## Run

```bash
go run main.go
```

will be like this:

![Image of Bcrypt Sample](bcrypt-sample.png)


## Building
You can compile and make a shortcut to it by running:

```bash
go build -o bcrypt main.go
```

and move that `bcrypt` generated file to somewhere in your system path. 

> Note: Make sure to give this build the right permissions to run on your operational system.

and then you can just run:

```bash
bcrypt
```

or via shell pipe:

```bash
echo "yourpassword" | bcrypt
```

```bash
cat file-with-plain-password.txt  | bcrypt
```

To increase the cost of the hash, you should pass the argument `--cost` or `-c`, default is 10.

```bash
bcrypt -c 14
```

```bash
bcrypt --cost 14
```

# License

DBAD.