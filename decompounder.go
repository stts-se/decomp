package decomp //ounder

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// TODO:
// compute freqs
// completion?
// bestGuess (heuristics)

// TNode is kind of a trie-structure representing strings (words).
// A path trough the TNode that ends with leaf = true represents a
// string/word.
type tNode struct {
	// the current character in a string
	r rune
	// child nodes of this node
	sons map[rune]*tNode
	// true if this character ends an input string
	leaf bool
}

// newtNode initialises and returns a tNode
func newtNode() *tNode {
	return &tNode{sons: make(map[rune]*tNode)}
}

// wrappers
func (t *tNode) addS(s string) *tNode {
	return t.add([]rune(s))
}
func (t *tNode) removeS(s string) bool {
	return t.remove([]rune(s))
}
func (t *tNode) prefixesS(s string) []arc {
	return t.prefixes([]rune(s))
}

// add inserts a strings into the tNode and builds up sub-nodes as it
// goes
func (t *tNode) add(rs []rune) *tNode {

	if len(rs) == 0 {
		return t
	}

	// Pick off first rune in string
	r := rs[0]

	// This path so far already exists.
	// Recursively keep adding
	if son, ok := t.sons[r]; ok {
		if len(rs) == 1 {
			son.leaf = true
			// This is where you could increment a frequency counter.
			// You'd want to add a frequency field to both tNode and arc.
		}
		son.add(rs[1:])

	} else { // new path
		son := newtNode()
		son.r = r
		if len(rs) == 1 {
			son.leaf = true
		}
		t.sons[r] = son
		son.add(rs[1:])
	}

	return t
}

// contains returns true iff s is a leaf
func (t *tNode) contains(s string) bool {

	if s == "" {
		return false
	}

	res := false
	sons := t.sons
	for _, r := range s {
		v, ok := sons[r]
		if !ok { // s not a path in t
			return false
		}
		res = v.leaf
		sons = v.sons
	}

	return res
}

// remove sets leaf = false, but does not actually remove the
// character sequence of the string from the tree (since it may be a
// substring of another string).
// If the string is not in t, nothing happens.
// Returns true if the string was 'removed' from the tree, false otherwise.
// TODO Purge paths that do not lead to a leaf = true.
func (t *tNode) remove(rs []rune) bool {

	if len(rs) == 0 {
		return false
	}

	sons := t.sons
	for i, r := range rs {
		if v, ok := sons[r]; ok {
			if i == len(rs)-1 { // last rune
				v.leaf = false
				return true
			}
			// keep following the path
			sons = v.sons

		} else {
			return false // rs not a path in t
		}
	}

	return false
}

// list returns all paths ending in a leaf, that is all strings added
// to the top tNode (if not subsequently removed).
func (t *tNode) list() []string {
	var res []string

	for _, s := range t.sons {
		listAccu(s, "", &res)
	}

	return res
}

// listAccu is an accumulator helper function to list() above
func listAccu(t *tNode, soFar string, accu *[]string) {
	if t.leaf {
		*accu = append(*accu, soFar+string(t.r))
	}
	for _, k := range t.sons {
		listAccu(k, soFar+string(t.r), accu)
	}
}

// arcType, the type of an arc (long name since 'type' is a reserved word in Go)
type arcType int

const (
	prefix arcType = iota
	infix
	suffix
)

// arc represents a substring of a string, with a start and end index
// of the string.
type arc struct {
	start int
	end   int
	cat   arcType // used to eliminate unwanted sequences of arcs
	strn  string  // The actual string is needed in order to handle triple consonant compounds
}

func (a arc) len() int {
	return len(a.runes())
}

func (a arc) runes() []rune {
	return []rune(a.strn)
}

// Returns the matching prefix substrings of s that exist in t in the
// form of arcs. A prefix must be shorter than the input string.
func (t *tNode) prefixes(rs []rune) []arc {
	var res []arc

	sons := t.sons
	for i, r := range rs {
		// path in tree

		if v, ok := sons[r]; ok {
			sons = v.sons
			// '&& i < len(s)-1' ensures that the prefix is shorter than s

			if v.leaf && i < len(rs)-1 {
				res = append(res, arc{end: i + 1, cat: prefix})
			}
		} else { // not a path in tree
			return res
		}
	}

	return res
}

type prefixTree struct {
	tree *tNode // TODO rename 'tree' to 'prefixes'?
	// infixes are gluing parts that may appear once after a prefix
	infixes *tNode
}

func newPrefixTree() prefixTree {
	return prefixTree{tree: newtNode(), infixes: newtNode()}
}

func (t prefixTree) Add(s string) {
	t.tree.add([]rune(s))
}
func (t prefixTree) Remove(s string) bool {
	return t.tree.remove([]rune(s))
}
func (t prefixTree) AddInfix(s string) {
	t.infixes.add([]rune(s))
}
func (t prefixTree) RemoveInfix(s string) bool {
	return t.infixes.remove([]rune(s))
}

func (t prefixTree) Prefixes(s string) []arc {
	return t.tree.prefixes([]rune(s))
}

func (t prefixTree) Infixes(s string) []arc {
	return t.infixes.prefixes([]rune(s))
}

func (t prefixTree) allPotentialPrefixes(rs []rune, index int) []arc {

	var res []arc

	// Chop one rune at the time, and consume rs until the bitter
	// end, and collect all prefixes as you go

	for i := index; i < len(rs); i++ {
		pFixes := t.Prefixes(string(rs[i:]))

		for _, a := range pFixes {
			a.start = a.start + i
			a.end = a.end + i
			a.strn = string(rs[a.start:a.end])
			a.cat = prefix
			res = append(res, a)
		}

		iFixes := t.Infixes(string(rs[i:]))

		// infixes cannot start a string
		if i > 0 {
			for _, a := range iFixes {
				a.start = a.start + i
				a.end = a.end + i
				a.strn = string(rs[a.start:a.end])
				a.cat = infix
				res = append(res, a)
			}
		}
	}

	return res
}

type suffixTree struct {
	tree *tNode
}

func newSuffixTree() suffixTree {
	return suffixTree{tree: newtNode()}
}

// Reverse returns its argument string reversed rune-wise left to right.
// Lifted from https://github.com/golang/example/blob/master/stringutil/reverse.go
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func (t suffixTree) Add(s string) {
	r := reverse(s)
	t.tree.add([]rune(r))
}

func (t suffixTree) Remove(s string) bool {
	r := reverse(s)
	return t.tree.remove([]rune(r))
}

// Suffixes returns the arc for suffixes of s in t. A suffix may not
// span the complete s.
func (t suffixTree) Suffixes(s string) []arc {
	r := reverse(s)

	// the arcs going from right to left (strings inverted)
	suffArcs := t.tree.prefixes([]rune(r))

	// invert arcs to go from left to right
	end := len([]rune(r))
	var res []arc
	for _, a := range suffArcs {
		res = append(res, arc{start: end - a.end, end: end - a.start, cat: suffix})
	}

	return res
}

// Decompounder is a struct that holds potential word parts as 'prefixes' or 'suffixes'.
type Decompounder struct {
	prefixes    prefixTree
	suffixes    suffixTree
	tripleChars map[rune]bool // Chars that may be contracted into two
	// when occurring at compound boundaries:
	// 'natt'+'tåg' => 'nattåg'
}

// NewDecompounder returns a fresh instance of a Decompounder.
func NewDecompounder() Decompounder {
	return Decompounder{prefixes: newPrefixTree(), suffixes: newSuffixTree(), tripleChars: make(map[rune]bool)}
}

// AllowedTripleChars defines what characters that can be reduced when
// appearing in triples over compound boundaries. An example where 't'
// is an allowed triple consonats: 'nattåg' -> 'natt'+'tåg'.
func (d Decompounder) AllowedTripleChars(r []rune) {
	for _, r0 := range r {
		d.tripleChars[r0] = true
	}
}

func (d Decompounder) arcs(rs []rune) []arc {
	var res []arc

	res0 := append(res, d.prefixes.allPotentialPrefixes(rs, 0)...)

	res1 := append(res, d.suffixes.Suffixes(string(rs))...)

	// ensure no duplicate arcs
	found := make(map[arc]bool)
	for _, a := range res0 {
		if !found[a] {
			res = append(res, a)
			found[a] = true
		}
	}
	for _, a := range res1 {
		if !found[a] {
			res = append(res, a)
			found[a] = true
		}
	}

	// add the actual substrings
	for i, a := range res {
		res[i].strn = string(rs[a.start:a.end])
	}

	return res
}

type endChar struct {
	i int
	r rune
}

func genTripleConsonantArcs(rs []rune, tripleChars map[rune]bool, arcs []arc) []arc {
	var res []arc

	// Add new arcs for pairs of overlapping arcs where the
	// left arc ends in two tripleChars and the right arc starts
	// with the same letter. The overlap should be one character (a1.end-1 == a2.from)

	//var finalDoubleCons []arc

	// These structs contain the arcs ending in double chars (if defined in tripleChars)
	endCharsAt := make(map[endChar][]arc)

	// This loop collects all arcs ending in double letters that are defined in tripleChars
	for _, a := range arcs {
		// Strings of two characters are illegal first words of overlapping compounds
		if a.len() < 3 {
			continue
		}

		rs := a.runes()
		last := rs[len(rs)-1]
		lastButOne := rs[len(rs)-2]

		if !tripleChars[last] || last != lastButOne {
			continue
		}

		ec := endChar{i: a.end, r: last}
		endCharsAt[ec] = append(endCharsAt[ec], a)
	}

	// This loop looks for overlapping arcs starting with the same chars that ends arc in encCharAt
	for _, a := range arcs {
		// Single characters strings are not legal second words in overlapping compounds
		if a.len() < 2 {
			continue
		}
		rs := a.runes()
		r := rs[0]
		i := a.start
		ec := endChar{i: i + 1, r: r} // overlapping arcs with the same character r?
		if arcz, ok := endCharsAt[ec]; ok {
			for _, a0 := range arcz {
				// for each overlapping pair of arcs, add a new first arc that ends one char from the end of the string
				res = append(res, arc{start: a0.start, end: a0.end - 1, strn: a0.strn})
			}
		}
	}

	return res
}

// AddPrefix adds a string that can function as a word part beginning a word.
func (d Decompounder) AddPrefix(s string) {
	d.prefixes.Add(s)
}

// RemovePrefix removes a string previously added by AddPrefix.
func (d Decompounder) RemovePrefix(s string) bool {
	return d.prefixes.Remove(s)
}

// ContainsPrefix returns true if s was previously added by AddPrefix.
func (d Decompounder) ContainsPrefix(s string) bool {
	return d.prefixes.tree.contains(s)
}

// AddInfix adds a string that can function as an infix between two word parts.
func (d Decompounder) AddInfix(s string) {
	d.prefixes.AddInfix(s)
}

// RemoveInfix removes an infix added by AddInfix.
func (d Decompounder) RemoveInfix(s string) bool {
	return d.prefixes.RemoveInfix(s)
}

// ContainsInfix returns true if s was previously added by AddInfix.
func (d Decompounder) ContainsInfix(s string) bool {
	return d.prefixes.infixes.contains(s)
}

// AddSuffix adds a string that can function as a word part ending a word.
func (d Decompounder) AddSuffix(s string) {
	d.suffixes.Add(s)
}

// RemoveSuffix removes a string previously added by AddSuffix.
func (d Decompounder) RemoveSuffix(s string) bool {
	return d.suffixes.Remove(s)
}

// ContainsSuffix returns true if s was previously added by AddSuffix.
func (d Decompounder) ContainsSuffix(s string) bool {
	return d.suffixes.tree.contains(reverse(s))
}

// Er... the odd spelling was to signal that these are constants, but that they should not be exported (lower case cap)
const aLLOWEDtRIPLEcHARS, sUFFIX, pREFIX, iNFIX, rEMOVE string = "ALLOWED_TRIPLE_CHARS", "SUFFIX", "PREFIX", "INFIX", "REMOVE"

// List returns all wordparts of Decompounder prefixed with type,
// PREFIX:, INFIX: or SUFFIX:.  The strings of the different types are
// sorted alphabetically for each category. This ordering is probably
// different from the original insert order.
func (d Decompounder) List() []string {
	var res []string

	var triples []rune
	for k := range d.tripleChars {
		triples = append(triples, k)
	}

	if len(triples) > 0 {
		t := aLLOWEDtRIPLEcHARS + ":" + string(triples)
		res = append(res, t)
	}

	ps := d.prefixes.tree.list()
	sort.Strings(ps)
	for _, p := range ps {
		res = append(res, pREFIX+":"+p)
	}

	is := d.prefixes.infixes.list()
	sort.Strings(is)
	for _, i := range is {
		res = append(res, iNFIX+":"+i)
	}

	ss := d.suffixes.tree.list()
	sort.Strings(ss)
	for _, s := range ss {
		res = append(res, sUFFIX+":"+reverse(s))
	}

	return res
}

// SaveToFile saved the current state of Decompounder to text file fName.
func (d Decompounder) SaveToFile(fName string) error {
	var fh *os.File
	var err error

	fh, err = os.OpenFile(fName, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		return err
	}
	defer fh.Close()
	w := bufio.NewWriter(fh)
	defer w.Flush()
	// TODO sort lines alphabetically?
	// This should be done inside d.List() ?
	for _, s := range d.List() {
		w.WriteString(s + "\n")

	}
	return err
}

// NewDecompounderFromFile initializes a Decompounder from a text file of the following format:
// (ALLOWED_TRIPLE_CHARS:<character sequence>)?
//(REMOVE:)?<PREFIX|INFIX|SUFFIX>:<lower-case string>
//
// The optional REMOVE: command is used as a simple means to remove an entry,
// by merely append the REMOVE: tagged line to the text file. The
// REMOVE: line must occur somewhere after the original line to be removed
// (otherwise it will be added anew).
//
// # line starting with '#' is ignored
// '' empty lines are ignore
func NewDecompounderFromFile(fileName string) (Decompounder, error) {
	var err error
	res := NewDecompounder()
	fh, err := os.Open(fileName)
	if err != nil {
		return res, err
	}
	defer fh.Close()

	linesRead := 0
	linesSkipped := 0
	linesRemoved := 0
	linesAdded := 0
	s := bufio.NewScanner(fh)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		linesRead++
		if l == "" || strings.HasPrefix(l, "#") {
			continue
		}
		// TODO parse string and report mismatching strings
		// add or remove PREFIX, INFIX, SUFFIX
		// 1) print failing line to STDERR
		// 2) count nummber of failing lines and if > 0 return error
		fs := strings.SplitN(l, ":", 2)
		if len(fs) != 2 {
			//err = fmt.Errorf("invalid line skipped: %s", l)
			linesSkipped++
			continue
		}

		if fs[0] == rEMOVE {
			fsRem := strings.SplitN(fs[1], ":", 2)

			switch fsRem[0] {
			//TODO this one has not yet been tested
			case aLLOWEDtRIPLEcHARS:

				tc := strings.Replace(strings.TrimSpace(fsRem[1]), " ", "", -1)
				for _, c := range tc {
					delete(res.tripleChars, c)
				}

			case pREFIX:

				if res.RemovePrefix(strings.ToLower(fsRem[1])) {
					linesRemoved++
				}

			case iNFIX:

				if res.RemoveInfix(strings.ToLower(fsRem[1])) {
					linesRemoved++
				}

			case sUFFIX:

				if res.RemoveSuffix(strings.ToLower(fsRem[1])) {
					linesRemoved++
				}

			default:

				linesSkipped++

			}
			continue // REMOVE
		}

		// This check should not be needed? Covered by default below?
		if fs[0] != aLLOWEDtRIPLEcHARS && fs[0] != pREFIX && fs[0] != iNFIX && fs[0] != sUFFIX {
			//err = fmt.Errorf("invalid line skipped: %s", l)
			linesSkipped++
			continue
		}

		switch fs[0] {
		case aLLOWEDtRIPLEcHARS:
			tc := strings.Replace(strings.TrimSpace(fs[1]), " ", "", -1)
			res.AllowedTripleChars([]rune(tc))
		case pREFIX:
			res.AddPrefix(strings.ToLower(fs[1]))
			linesAdded++
		case iNFIX:
			res.AddInfix(strings.ToLower(fs[1]))
			linesAdded++
		case sUFFIX:
			res.AddSuffix(strings.ToLower(fs[1]))
			linesAdded++
		default:
			linesSkipped++
		}

	}

	if s.Err() != nil {
		// We've already got an error, so just append to that
		if err != nil {
			err = fmt.Errorf("%v : %v", err, s.Err())
		}
		// no previous error
		if err == nil {
			err = s.Err()
		}
	}

	// TODO if verbose:
	//fmt.Fprintf(os.Stderr, "Lines read: %d\nLines skipped: %d\nLines added: %d\nLines removed: %d\n", linesRead, linesSkipped, linesAdded, linesRemoved)
	return res, err
}

// ByLen sorts [][]string according to length
type ByLen [][]string

func (b ByLen) Len() int {
	return len(b)
}
func (b ByLen) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ByLen) Less(i, j int) bool {
	// TODO? add frequency to arcs, and sort them instead: first by
	// length, second by highest lowest freq.
	return len(b[i]) < len(b[j])
}

// Decomp tries to find potential sequences of word parts in s.
func (d Decompounder) Decomp(s string) [][]string {
	var res [][]string

	rs := []rune(s)
	arcs := d.arcs(rs)
	newArcs := genTripleConsonantArcs(rs, d.tripleChars, arcs)
	arcs = append(arcs, newArcs...)

	paths := paths(arcs, 0, len(rs))

	for _, p := range paths {
		res = append(res, pathToDecomp(p, s))
	}
	sort.Sort(ByLen(res))
	return res
}

func paths(as []arc, from, to int) [][]arc {

	// ensure that there are no duplicate arcs, since these will
	// generate multiple identical paths
	found := make(map[arc]bool)
	var uniqueAs []arc
	for _, a := range as {
		if !found[a] {
			uniqueAs = append(uniqueAs, a)
			found[a] = true
		}
	}

	arcMap := make(map[int][]arc)
	for _, a := range uniqueAs {
		v := arcMap[a.start]
		arcMap[a.start] = append(v, a)
	}

	var path []arc
	var res [][]arc
	pathsAccu(arcMap, from, to, path, &res)
	return res
}

func pathsAccu(as map[int][]arc, from, to int, currPath []arc, paths *[][]arc) {

	arcs, ok := as[from]
	if !ok { // no path from 'from'
		return
	}

	for _, arc := range arcs {
		// Yeah! Complete path!
		if arc.end == to {
			path := currPath
			path = append(path, arc)
			*paths = append(*paths, path)
		}

		// A path cannot follow consecutive 'infix' arcs
		if len(currPath) > 0 { // We are not on the first arc
			lastArc := currPath[len(currPath)-1]
			// Nope, cannot have two infix arcs in a row
			if arc.cat == infix && lastArc.cat == infix {
				continue
			}
			// An infix character cannot follow on a word
			// part ending in the same character as the infix character
			if len(lastArc.strn) > 0 { // (Don't understand how a previous arc can be the empty string, but OK...)
				rs := []rune(lastArc.strn)
				lastChar := string(rs[len(rs)-1])
				if arc.cat == infix && lastChar == arc.strn {
					continue
				}
			}
		}
		// Keep treading down the path
		newPath := currPath
		newPath = append(newPath, arc)

		pathsAccu(as, arc.end, to, newPath, paths)

	}
}

func pathToDecomp(p []arc, s string) []string {
	var res []string
	// TODO error checking
	for _, a := range p {
		s0 := a.strn
		res = append(res, s0)
	}
	return res
}
