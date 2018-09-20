package main

import (
	//"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/gorilla/mux"

	"github.com/stts-se/decomp"
)

type decomperMutex struct {
	// map from language name to decompounder, as read from word parts file dir.
	// lang is used as HTTP request parameter to select a decompounder
	decompers map[string]decomp.Decompounder
	// map from language name to word parts file name path.  Used
	// for appending new word parts to the original text file, to
	// keep the word parts text file in sync with the in-memory
	// Decompounder. (Maybe there is a saner way to handle this?)
	files map[string]string
	mutex *sync.RWMutex
}

var decomper = decomperMutex{
	decompers: make(map[string]decomp.Decompounder),
	files:     make(map[string]string),
	mutex:     &sync.RWMutex{},
}

// appendToWordPartsFile writes a line to a file.
// NB that it is not thread-safe, and should be called after locking.
func appendToWordPartsFile(fn string, line string) error {

	fh, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer fh.Close()

	_, err = fh.WriteString(line + "\n")
	if err != nil {
		return err
	}

	return nil
}

func addPrefix(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	decomperName := vars["decomper_name"]
	// if "" == lang {
	// 	msg := "no value for the expected 'lang' parameter"
	// 	log.Println(msg)
	// 	http.Error(w, msg, http.StatusBadRequest)
	// 	return
	// }

	prefix := strings.ToLower(vars["prefix"])
	if prefix == "" {
		return
	}

	// if "" == prefix {
	// 	msg := "no value for the expected 'prefix' parameter"
	// 	log.Println(msg)
	// 	http.Error(w, msg, http.StatusBadRequest)
	// 	return
	// }

	decomper.mutex.Lock()
	defer decomper.mutex.Unlock()
	fn, ok := decomper.files[decomperName]
	if !ok {
		msg := "unknown decomper: " + decomperName
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if decomper.decompers[decomperName].ContainsPrefix(prefix) {
		fmt.Fprintf(w, "prefix already found: '%s'\n", prefix)
		return
	}

	decomper.decompers[decomperName].AddPrefix(prefix)
	err := appendToWordPartsFile(fn, "PREFIX:"+prefix)
	if err != nil {
		msg := fmt.Sprintf("decompounder: failed to append to word parts file : %v", err)
		log.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "added '%s'\n", prefix)
}

// TODO cut-and-paste of addPrefix
func removePrefix(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	decomperName := vars["decomper_name"]

	prefix := vars["prefix"]
	if prefix == "" {
		return
	}

	decomper.mutex.Lock()
	defer decomper.mutex.Unlock()
	fn, ok := decomper.files[decomperName]
	if !ok {
		msg := "unknown decomper: " + decomperName
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	//writeToWordPartsFile("PREFIX")
	if !decomper.decompers[decomperName].ContainsPrefix(prefix) {
		fmt.Fprintf(w, "prefix not found: '%s'\n", prefix)
		return
	}
	decomper.decompers[decomperName].RemovePrefix(prefix)
	err := appendToWordPartsFile(fn, "REMOVE:PREFIX:"+prefix)
	if err != nil {
		msg := fmt.Sprintf("decompounder: failed to append to word parts file : %v", err)
		log.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "removed prefix: '%s'\n", prefix)
}

func addSuffix(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	decomperName := vars["decomper_name"]
	suffix := strings.ToLower(vars["suffix"])
	if suffix == "" {
		return
	}

	decomper.mutex.Lock()
	defer decomper.mutex.Unlock()
	fn, ok := decomper.files[decomperName]
	if !ok {
		msg := "unknown decomper: " + decomperName
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if decomper.decompers[decomperName].ContainsSuffix(suffix) {
		fmt.Fprintf(w, "sufffix already found: '%s'\n", suffix)
		return
	}

	decomper.decompers[decomperName].AddSuffix(suffix)
	err := appendToWordPartsFile(fn, "SUFFIX:"+suffix)
	if err != nil {
		msg := fmt.Sprintf("decompounder: failed to append to word parts file : %v", err)
		log.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "added '%s'\n", suffix)
}

// TODO cut-and-paste of addSuffix
func removeSuffix(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	decomperName := vars["decomper_name"]
	suffix := vars["suffix"]

	if suffix == "" {
		return
	}

	decomper.mutex.Lock()
	defer decomper.mutex.Unlock()
	fn, ok := decomper.files[decomperName]
	if !ok {
		msg := "unknown decomper: " + decomperName
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if !decomper.decompers[decomperName].ContainsSuffix(suffix) {
		fmt.Fprintf(w, "suffix not found: '%s'\n", suffix)
		return
	}
	decomper.decompers[decomperName].RemoveSuffix(suffix)
	err := appendToWordPartsFile(fn, "REMOVE:SUFFIX:"+suffix)
	if err != nil {
		msg := fmt.Sprintf("decompounder: failed to append to word parts file : %v", err)
		log.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "removed '%s'\n", suffix)
}

type decomps struct {
	Parts []string `json:"parts"`
}

// langFromFilePath returns the base file name stripped from any '.txt' extension
func langFromFilePath(p string) string {
	b := filepath.Base(p)
	if strings.HasSuffix(b, ".txt") {
		b = b[0 : len(b)-4]
	}
	return b
}

func decompWord(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	decomperName := vars["decomper_name"]
	word := vars["word"]
	word = strings.ToLower(word)

	var res []decomps
	decomper.mutex.RLock()
	defer decomper.mutex.RUnlock()
	_, ok := decomper.files[decomperName]
	if !ok {
		msg := "unknown 'decomper': " + decomperName
		var decompers []string
		for l := range decomper.decompers {
			decompers = append(decompers, l)
		}
		msg = fmt.Sprintf("%s. Known decomper names: %s", msg, strings.Join(decompers, ", "))
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	for _, d := range decomper.decompers[decomperName].Decomp(word) {
		res = append(res, decomps{Parts: d})
	}
	log.Println(res)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if len(res) == 0 {
		fmt.Fprintf(w, "[]\n")
		return
	}

	j, err := json.Marshal(res)
	if err != nil {
		msg := fmt.Sprintf("failed json marshalling : %v", err)
		log.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", string(j))
}

func listDecompers(w http.ResponseWriter, r *http.Request) {

	decomper.mutex.RLock()
	var res []string // res0 contains path to file
	for l := range decomper.decompers {
		res = append(res, l)
	}
	decomper.mutex.RUnlock()

	sort.Strings(res)
	j, err := json.Marshal(res)
	if err != nil {
		msg := fmt.Sprintf("failed json marshalling : %v", err)
		log.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s\n", string(j))
}

func decompMain(w http.ResponseWriter, r *http.Request) {
	// TODO error if file not found
	//http.ServeFile(w, r, "./src/decomp_demo.html")
	fmt.Printf("decompMain: PENDING")
	fmt.Fprintf(w, "%s\n", "TO DO")
}

func mainHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%s\n", strings.Join(walkedURLs, "\n"))
}

func ping(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "decompserver\n")
}

var walkedURLs = []string{}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "decompserver <DECOMPFILES DIR>\n")
		os.Exit(0)
	}

	// word decomp file dir. Each file in dn with .txt extension
	// is treated as a word parts file
	var dn = os.Args[1]

	files, err := ioutil.ReadDir(dn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(0)
	}

	// populate map of decompounders from word parts files.
	// The base file name minus '.txt' is the language name.
	var fn string
	for _, f := range files {
		fn = filepath.Join(dn, f.Name())
		if !strings.HasSuffix(fn, ".txt") {
			fmt.Fprintf(os.Stderr, "decompserver: skipping file: '%s'\n", fn)
			continue
		}

		dc, err := decomp.NewDecompounderFromFile(fn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			fmt.Fprintf(os.Stderr, "decompserver: skipping file: '%s'\n", fn)
			continue
		}

		lang := langFromFilePath(fn)
		decomper.mutex.Lock()
		decomper.decompers[lang] = dc
		decomper.files[lang] = fn
		decomper.mutex.Unlock()
		fmt.Fprintf(os.Stderr, "decomper: loaded file '%s'\n", fn)
	}

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", mainHandler).Methods("get", "post")
	r.HandleFunc("/ping", ping).Methods("get", "post")
	r.HandleFunc("/decomp", decompMain).Methods("get")
	r.HandleFunc("/decomp/list_decompers", listDecompers).Methods("get", "post")
	r.HandleFunc("/decomp/{decomper_name}/{word}", decompWord).Methods("get")                   //, "post")
	r.HandleFunc("/decomp/{decomper_name}/add_prefix/{prefix}", addPrefix).Methods("get")       //, "post")
	r.HandleFunc("/decomp/{decomper_name}/remove_prefix/{prefix}", removePrefix).Methods("get") //, "post")
	r.HandleFunc("/decomp/{decomper_name}/add_suffix/{suffix}", addSuffix).Methods("get")       //, "post")
	r.HandleFunc("/decomp/{decomper_name}/remove_suffix/{suffix}", removeSuffix).Methods("get") //, "post")

	// List route URLs to use as simple on-line documentation (at "/")
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		walkedURLs = append(walkedURLs, t)
		return nil
	})

	port := ":6778"
	log.Printf("starting decomp server at port %s\n", port)
	err = http.ListenAndServe(port, r)
	if err != nil {

		log.Fatalf("no fun: %v\n", err)
	}

}
