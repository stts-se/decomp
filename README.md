# decomp
decomp.Decompounder is used for guessing compound boundaries of compound words, given a list of possible word parts.

Clone repo.
cd decomp
go get ./...
go test
go run 

Command line

 `go run cmd/decomper/decomper.go
 decomper <DECOMP FILE> <words...>|<STDIN>`

`go run cmd/decomper/decomper.go decompserver/decomp_files/sv_nst.txt zebrafink
Lines read: 41887
Lines skipped: 0
Lines added: 41887
Lines removed: 0
zebrafink	zebra+fink`

HTTP Server

`go run decompserver/decompserver.go 
decompserver <DECOMPFILES DIR>`

`go run decompserver/decompserver.go decompserver/decomp_files/
Lines read: 41887
Lines skipped: 0
Lines added: 41887
Lines removed: 0
decomper: loaded file 'decompserver/decomp_files/sv_nst.txt'
2018/09/19 21:04:05 starting decomp server at port :6778`



Go to localhost:6778 to see the API calls:

`curl http://localhost:6778/
/
/ping
/decomp
/decomp/list_decompers
/decomp/{decomper_name}/{word}
/decomp/{decomper_name}/add_prefix/{prefix}
/decomp/{decomper_name}/remove_prefix/{prefix}
/decomp/{decomper_name}/add_suffix/{suffix}
/decomp/{decomper_name}/remove_suffix/{suffix}

curl http://localhost:6778/decomp/list_decompers
["sv_nst"]

curl http://localhost:6778/decomp/sv_nst/zebrafinkar
[{"parts":["zebra","finkar"]}]`



[![GoDoc](https://godoc.org/github.com/stts-se/decomp?status.svg)](https://godoc.org/github.com/stts-se/decomp) [![Go Report Card](https://goreportcard.com/badge/github.com/stts-se/decomp)](https://goreportcard.com/report/github.com/stts-se/decomp) [![Build Status](https://travis-ci.org/stts-se/decomp.svg?branch=master)](https://travis-ci.org/stts-se/decomp)
