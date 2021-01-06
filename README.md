# minigo [![Actions Status](https://github.com/Eun/minigo/workflows/CI/badge.svg)](https://github.com/Eun/minigo/actions) [![Codecov](https://img.shields.io/codecov/c/github/Eun/minigo.svg)](https://codecov.io/gh/Eun/minigo) [![GoDoc](https://godoc.org/github.com/Eun/minigo?status.svg)](https://godoc.org/github.com/Eun/minigo) [![go-report](https://goreportcard.com/badge/github.com/Eun/minigo)](https://goreportcard.com/report/github.com/Eun/minigo)
A mini golang interpreter based on [yaegi-template](https://github.com/Eun/yaegi-template) and [yaegi](https://github.com/traefik/yaegi).


## Usage
```go
#!/usr/bin/env minigo
println("Hello World")
```

> Note that _minigo_ comes with a go runtime, you don't have to install go!

## Use docker!
```go
#!/usr/bin/env -S docker run -ti -v ${PWD}:/app/ -w /app --rm eunts/minigo:latest
println("Hello World")
```

## Installation
Download and install from the [Relases Page](https://github.com/Eun/minigo/releases).
Or compile it yourself using `go`

## Importing
You can import other packages by specifying `$GOPATH` and placing them in `$GOPATH/src`.  
However [yaegi](https://github.com/traefik/yaegi) is (currently) limited in functionality so don't expect all packages are working.
