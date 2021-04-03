# up – upload files to anywhere

![](up.png)

up is the command line client for [Goupfile](https://goupfile.com) and other
file upload services.

File upload sites are prone to abuse and are not sustainable to run in the long
term. They often disappear, and every year there are new file upload sites on
the first page of Google.

up automatically chooses among popular, available services so that you don't
have to.

## Usage

```sh
go run up.go myfile.txt
```

or

```sh
go build
./up myfile.txt
```

## Status

up currently only supports https://filebin.net and is pre-alpha. Support for
many more sites, along with new CLI features, is coming soon!

## License

MIT
