# minigo
A mini golang interpreter based on [yaegi-template](https://github.com/Eun/yaegi-template) and [yaegi](https://github.com/containous/yaegi).


## Usage
```go
#!/usr/bin/env minigo
fmt.Println("Hello World")
```

> Note that _minigo_ comes with a go runtime, you don't have to install go!

## Installation
Download and install from the [Relases Page](https://github.com/Eun/minigo/releases).
Or compile it yourself using `go`

## Importing
You can import other packages by specifying `$GOPATH` and placing them in `$GOPATH/src`.  
However [yaegi](https://github.com/containous/yaegi) is (currently) limited in functionality so don't expect all packages are working.
