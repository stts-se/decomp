# decomp
decomp.Decompounder is used for guessing compound boundaries of compound words, given a list of possible word parts.

It reads a file of possible compound parts, that are either a "prefix" or a "suffix". A compound may consist of several prefixes, followed by exactly one suffix.

There are two compound word files available in this repo, one for Swedish, `decompserver/decomp_files/sv_nst.tst`, and one for Norwegian Bokmål, `decompserver/decomp_files/nob_nst.tst`. You will most likely need to add word parts to these files for your own application or domain (see web API below).

[Go docs for decomp](https://godoc.org/github.com/stts-se/decomp).

## Linking -s-

There may be a single linking character between compound parts, as in `tidningsartikel -> tidning s artikel`. These are defined in the words parts file using the INFIX label.


## Triple identical characters collapsed into two at compound boundaries

By using the ALLOWED_TRIPLE_CHARS tag, it is possible to enumerate a set of characters that may be collapsed into two, such as in `natt+tåg -> nattåg`. 
See the top of `sv_nst.txt`.

## All possible guesses are generated

When there are several possible guesses for a word, all are presented in ascending order of number of word parts. Typically, guesses with the least number of word parts are better.


## Examples:

All examples below are in Swedish.

    filmstjärna
    film stjärna

    operasångaren
    opera sångaren

    dansbandssångare
    dans band s sångare

    tidningsartikel
    tidning s artikel    

The word part `tidning` is listed as  PREFIX, `s` is defined INFIX and `artikel` is listed as a SUFFIX in the word parts file.


## Examples of competing guesses:

    fiskeskär

    fiske skär
    fiske s kär

Above, the first guess, with the least number of compound parts, is correct.

    glasstrut

    glas strut
    glass trut
    glass strut
    glas s trut
    glass s trut

Above, all first three are possible, but the third one, `glass strut`, is the "correct" one. (Notice that three `s` have become two in the compound.)


TODO: Add a way to select between guesses of equal number of word parts. This might be possible to do using word part frequencies. 

TODO: Better handle spurious linking characters, such as in  `glas s trut` and `glass s trut` above.

TODO: Describe the REMOVE label of the word parts file, etc.

# Running it

## To try it out, clone this repo, then

    cd decomp
    go get ./...
    go test
    

## Command line

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


## HTTP Server

    cd decompserver/

    go run decompserver.go 
    decompserver <DECOMPFILES DIR>

    go run decompserver.go decomp_files/
    Lines read: 30275
    Lines skipped: 0
    Lines added: 30274
    Lines removed: 0
    decomper: loaded file 'decomp_files/nob_nst.txt'
    Lines read: 41888
    Lines skipped: 0
    Lines added: 41887
    Lines removed: 0
    decomper: loaded file 'decomp_files/sv_nst.txt'
    2018/09/22 11:22:29 starting decomp server at port :6778

 

Go to localhost:6778 to see the API calls, either in your browser or using e.g. curl:

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


List the current decompers:

    curl http://localhost:6778/decomp/list_decompers
    ["nob_nst","sv_nst"]


Try a decomper (sv_nst) and a word (zebrafinkar):

    curl http://localhost:6778/decomp/sv_nst/zebrafinkar
    [{"parts":["zebra","finkar"]}]



### Web demo at `http://localhost:6778/demo.html`

The server must currently be started from its own directory, as above, in order for web demo page to work (the `static` directory is hard wired).



[![GoDoc](https://godoc.org/github.com/stts-se/decomp?status.svg)](https://godoc.org/github.com/stts-se/decomp) [![Go Report Card](https://goreportcard.com/badge/github.com/stts-se/decomp)](https://goreportcard.com/report/github.com/stts-se/decomp) [![Build Status](https://travis-ci.org/stts-se/decomp.svg?branch=master)](https://travis-ci.org/stts-se/decomp)
