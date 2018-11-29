## Framework for displaying a web view within a Go application

Currently supports macOS and Windows. Before this can be used, the CEF
bindings must be installed:

```
go get -u -d github.com/richardwilkes/cef
cd $GOPATH/src/github.com/richardwilkes/cef
./setup.sh
```

An example use of this library can be found at
https://github.com/richardwilkes/webapp-example
