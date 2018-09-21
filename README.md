# decomp
decomp.Decompounder is used for guessing compound boundaries of compound words, given a list of possible word parts.

It reads a file of possible compound parts, that are either a "prefix" or a "suffix". A compound may consist of several prefixes, followed by exactly one suffix. There may be a single linking character (these are defined in the compound part file using the INFIX tag) between compound parts.

By using the ALLOWED_TRIPLE_CHARS tag, it's possible to enumerate a set of characters that may be collapsed into two, such as in `natt+tåg -> nattåg`. 
See the top of `sv_nst.txt`.

When there are several possible guesses for a word, all are presented in ascending order of number of word parts. Typically, guesses with the least number of word parts are better.


Examples:

    filmstjärna
    film stjärna


    fiskeskär

    fiske skär
    fiske s kär

(Above, the first guess is correct.)

    glasstrut

    glas strut
    glass trut
    glass strut
    glas s trut
    glass s trut

(Above, all first three are possible, but the third one is the "correct" one.) 


TODO: Add a way to select between guesses of equal number of word parts. This might be possible to do using word part frequencies. 

TODO: Better handle spurious linking characters, such as in  `glas s trut` and `glass s trut` above.


To try it out, clone this repo, then

    cd decomp
    go get ./...
    go test
    

Command line

    go run cmd/decomper/decomper.go
    decomper <DECOMP FILE> <words...>|<STDIN>

    go run cmd/decomper/decomper.go decompserver/decomp_files/sv_nst.txt zebrafink
    Lines read: 41887
    Lines skipped: 0
    Lines added: 41887
    Lines removed: 0
    zebrafink	zebra+fink


By adding ALLOWED_TRIPLE_CHARS to the word part file, the guesser can handle compound boundaries where a double character has been assimilated, such as `nattåg -> natt+tåg`:

    go run decomper.go ../../decompserver/decomp_files/sv_nst.txt nattåg
    Lines read: 41888
    Lines skipped: 0
    Lines added: 41887
    Lines removed: 0
    nattåg	natt+tåg


HTTP Server (must be started from its own directory in order for web demo page to work, I think)

    cd decompserver/

    go run decompserver/decompserver.go 
    decompserver <DECOMPFILES DIR>

    go run decompserver/decompserver.go decompserver/decomp_files/
    Lines read: 41887
    Lines skipped: 0
    Lines added: 41887
    Lines removed: 0
    decomper: loaded file 'decompserver/decomp_files/sv_nst.txt'
    2018/09/19 21:04:05 starting decomp server at port :6778`


Web demo at `http://localhost:6778/demo.html`


Go to localhost:6778 to see the API calls:

    curl http://localhost:6778/
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
    ["nob_nst","sv_nst"]

    curl http://localhost:6778/decomp/sv_nst/zebrafinkar
    [{"parts":["zebra","finkar"]}]



[![GoDoc](https://godoc.org/github.com/stts-se/decomp?status.svg)](https://godoc.org/github.com/stts-se/decomp) [![Go Report Card](https://goreportcard.com/badge/github.com/stts-se/decomp)](https://goreportcard.com/report/github.com/stts-se/decomp) [![Build Status](https://travis-ci.org/stts-se/decomp.svg?branch=master)](https://travis-ci.org/stts-se/decomp)
