# ttt
[![Build Status](https://travis-ci.com/081888thLite/ttt.svg?branch=master)](https://travis-ci.com/081888thLite/ttt)

## Tic Tac Toe Game Written in Go

## To Play the game:
### If you wish to run this package, you'll first need to [DOWNLOAD](https://golang.org/dl/) and [INSTALL](https://golang.org/doc/install) Go.
### If you have an older version of Go on your machine (<Go 1.10) you will also need to update your Go version by deleting your Go workspace and then downloading the newest version from the downloads page linked above.
1. **Change directories into your $GOPATH...**</br>
`cd $GOPATH`</br>
if that doesn't work you need to check that you have go properly installed and the $GOPATH set...</br>
2. **Go get the game by running...**</br>
`go get github.com/081888thLite/ttt`</br>
3. **Change directories into the ttt root by running...**</br>
`cd go/src/github.com/081888thLite/ttt`</br>
4. **Build the application by running...**</br>
`go install .`</br>
5. **Play the game by running...**</br>
`go run ./cmd/play-ttt/main.go`</br>

### You can then run the test from the same directory w/ the following command:
1. `go test -v`

###### For a highly detailed html code quality report and overview of this Go package, I strongly recommend using [goreporter](https://github.com/360EntSecGroup-Skylar/goreporter) which can be quickly installed and run w/ the following two commands from the same directory (~/go/src/github.com/081888thLite/ttt)...
1. `go get -u github.com/360EntSecGroup-Skylar/goreporter`
2. `goreporter -p .`
