package decomp //ounder

import (
	"fmt"
	"strings"
	"testing"
)

var ts = "Wanted %#v got %#v\n"

func contains(ss []string, s string) bool {
	for _, s0 := range ss {
		if s == s0 {
			return true
		}
	}
	return false
}

func TestTree(t *testing.T) {

	tr := newtNode()

	if want, got := rune(0), tr.r; want != got {
		t.Errorf(ts, want, got)
	}

	tr = tr.addS("strut")
	tr = tr.addS("strutnos")
	tr = tr.addS("strutnosar")

	all := tr.list()
	if want, got := 3, len(all); want != got {
		t.Errorf(ts, want, got)
	}
	if want, got := true, contains(all, "strut"); want != got {
		t.Errorf(ts, want, got)
	}
	if want, got := true, contains(all, "strutnos"); want != got {
		t.Errorf(ts, want, got)
	}
	if want, got := true, contains(all, "strutnosar"); want != got {
		t.Errorf(ts, want, got)
	}

	cntns1 := tr.contains("strut")
	if want, got := true, cntns1; want != got {
		t.Errorf(ts, want, got)
	}

	rem1 := tr.removeS("strut")
	if want, got := true, rem1; want != got {
		t.Errorf(ts, want, got)
	}

	cntns2 := tr.contains("strut")
	if want, got := false, cntns2; want != got {
		t.Errorf(ts, want, got)
	}

	tr = tr.addS("strut")

	cntns3 := tr.contains("strutnosar")
	if want, got := true, cntns3; want != got {
		t.Errorf(ts, want, got)
	}
	rem2 := tr.removeS("strutnosar")
	if want, got := true, rem2; want != got {
		t.Errorf(ts, want, got)
	}
	cntns4 := tr.contains("strutnosar")
	if want, got := false, cntns4; want != got {
		t.Errorf(ts, want, got)
	}

	tr = tr.addS("strutnosar")

	// for k, v := range tr.sons {
	//     fmt.Printf("HOJSAN: %#v : %s\n", k, string(v.r))
	// }

	if want, got := rune(0), tr.r; want != got {
		t.Errorf(ts, want, got)
	}

	if want, got := 1, len(tr.sons); want != got {
		t.Errorf(ts, want, got)
	}

	s1 := "strutnosarna"
	prfs := tr.prefixesS(s1)
	//fmt.Printf("Arcs: %#v\n", prfs)
	if want, got := 3, len(prfs); want != got {
		t.Errorf(ts, want, got)
	}

	// for _, p := range prfs {
	//     fmt.Printf("PREFIX: %v\n", s1[p.start:p.end])
	// }

	pt := newPrefixTree()
	pt.Add("ap")
	pt.Add("hund")
	pt.Add("aphund")
	pt.Add("nos")

	s := "aphundar"
	arczz := pt.Prefixes(s)
	if want, got := 2, len(arczz); want != got {
		t.Errorf(ts, want, got)
	}

	prefs := map[string]bool{s[arczz[0].start:arczz[0].end]: true, s[arczz[1].start:arczz[1].end]: true}
	if _, ok := prefs["ap"]; !ok {
		t.Errorf(ts, "'ap'", "nothing")
	}
	if _, ok := prefs["aphund"]; !ok {
		t.Errorf(ts, "'aphund'", "nothing")
	}

	st := newSuffixTree()

	st.Add("rot")
	st.Add("mos")
	st.Add("nos")

	z := "skrotmos"
	arcs := st.Suffixes(z)
	if want, got := 1, len(arcs); want != got {
		t.Errorf(ts, want, got)
	}

	st.Add("rotmos")
	arcs = st.Suffixes(z)
	//fmt.Printf("ARKZ: %#v\n", arcs)
	if want, got := 2, len(arcs); want != got {
		t.Errorf(ts, want, got)
	}

	suffs := map[string]bool{z[arcs[0].start:arcs[0].end]: true, z[arcs[1].start:arcs[1].end]: true}
	if _, ok := suffs["mos"]; !ok {
		t.Errorf(ts, "'mos'", "nothing")
	}
	if _, ok := suffs["rotmos"]; !ok {
		t.Errorf(ts, "'rotmos'", "nothing")
	}

}

func TestTreeUTFx(t *testing.T) {

	tr := newtNode()

	if want, got := rune(0), tr.r; want != got {
		t.Errorf(ts, want, got)
	}

	tr = tr.addS("stråt")
	tr = tr.addS("stråtnös")
	tr = tr.addS("stråtnösar")

	all := tr.list()
	if want, got := 3, len(all); want != got {
		t.Errorf(ts, want, got)
	}
	if want, got := true, contains(all, "stråt"); want != got {
		t.Errorf(ts, want, got)
	}
	if want, got := true, contains(all, "stråtnös"); want != got {
		t.Errorf(ts, want, got)
	}
	if want, got := true, contains(all, "stråtnösar"); want != got {
		t.Errorf(ts, want, got)
	}

	cntns1 := tr.contains("stråt")
	if want, got := true, cntns1; want != got {
		t.Errorf(ts, want, got)
	}

	rem1 := tr.removeS("stråt")
	if want, got := true, rem1; want != got {
		t.Errorf(ts, want, got)
	}

	cntns2 := tr.contains("stråt")
	if want, got := false, cntns2; want != got {
		t.Errorf(ts, want, got)
	}

	tr = tr.addS("stråt")

	cntns3 := tr.contains("stråtnösar")
	if want, got := true, cntns3; want != got {
		t.Errorf(ts, want, got)
	}
	rem2 := tr.removeS("stråtnösar")
	if want, got := true, rem2; want != got {
		t.Errorf(ts, want, got)
	}
	cntns4 := tr.contains("stråtnösar")
	if want, got := false, cntns4; want != got {
		t.Errorf(ts, want, got)
	}

	tr = tr.addS("stråtnösar")

	// for k, v := range tr.sons {
	//     fmt.Printf("HOJSAN: %#v : %s\n", k, string(v.r))
	// }

	if want, got := rune(0), tr.r; want != got {
		t.Errorf(ts, want, got)
	}

	if want, got := 1, len(tr.sons); want != got {
		t.Errorf(ts, want, got)
	}

	s1 := "stråtnösarna"
	prfs := tr.prefixesS(s1)
	//fmt.Printf("Arcs: %#v\n", prfs)
	if want, got := 3, len(prfs); want != got {
		t.Errorf(ts, want, got)
	}

	// for _, p := range prfs {
	//     fmt.Printf("PREFIX: %v\n", s1[p.start:p.end])
	// }

	pt := newPrefixTree()
	pt.Add("ap")
	pt.Add("hund")
	pt.Add("aphund")
	pt.Add("nos")

	s := "aphundar"
	arczz := pt.Prefixes(s)
	if want, got := 2, len(arczz); want != got {
		t.Errorf(ts, want, got)
	}

	prefs := map[string]bool{s[arczz[0].start:arczz[0].end]: true, s[arczz[1].start:arczz[1].end]: true}
	if _, ok := prefs["ap"]; !ok {
		t.Errorf(ts, "'ap'", "nothing")
	}
	if _, ok := prefs["aphund"]; !ok {
		t.Errorf(ts, "'aphund'", "nothing")
	}

	st := newSuffixTree()

	st.Add("rot")
	st.Add("mos")
	st.Add("nos")

	z := "skrotmos"
	arcs := st.Suffixes(z)
	if want, got := 1, len(arcs); want != got {
		t.Errorf(ts, want, got)
	}

	st.Add("rotmos")
	arcs = st.Suffixes(z)
	//fmt.Printf("ARKZ: %#v\n", arcs)
	if want, got := 2, len(arcs); want != got {
		t.Errorf(ts, want, got)
	}

	suffs := map[string]bool{z[arcs[0].start:arcs[0].end]: true, z[arcs[1].start:arcs[1].end]: true}
	if _, ok := suffs["mos"]; !ok {
		t.Errorf(ts, "'mos'", "nothing")
	}
	if _, ok := suffs["rotmos"]; !ok {
		t.Errorf(ts, "'rotmos'", "nothing")
	}

}

func TestPaths(t *testing.T) {

	a1 := arc{start: 0, end: 3}
	a2 := arc{start: 3, end: 7}

	res := paths([]arc{a1, a2}, 0, 7)

	if want, got := 1, len(res); want != got {
		t.Errorf("NOOOO! %d %d", want, got)
	}
	p := res[0]
	if want, got := 2, len(p); want != got {
		t.Errorf("AAAA! %d %d", want, got)

	}
	a1b := p[0]
	if want, got := 0, a1b.start; want != got {
		t.Errorf("AAAA! %d %d", want, got)

	}
	if want, got := 3, p[1].start; want != got {
		t.Errorf("AAAA! %d %d", want, got)

	}

	a3 := arc{start: 3, end: 5}
	a4 := arc{start: 5, end: 7}
	a5 := arc{start: 3, end: 6}

	res = paths([]arc{a1, a2, a3, a4, a5}, 0, 7)
	if want, got := 2, len(res); want != got {
		t.Errorf("Suck %d %d", want, got)
	}
	//fmt.Printf("\n%#v\n", res)
}

func TestDecompounder(t *testing.T) {

	d := NewDecompounder()

	if w, g := false, d.ContainsPrefix("sylt"); w != g {
		t.Errorf(ts, w, g)
	}

	d.AddPrefix("sylt")
	if w, g := true, d.ContainsPrefix("sylt"); w != g {
		t.Errorf(ts, w, g)
	}
	d.RemovePrefix("sylt")
	if w, g := false, d.ContainsPrefix("sylt"); w != g {
		t.Errorf(ts, w, g)
	}
	d.AddPrefix("sylt")
	d.AddPrefix("syl")

	if w, g := false, d.ContainsSuffix("järn"); w != g {
		t.Errorf(ts, w, g)
	}
	d.AddSuffix("järn")
	if w, g := true, d.ContainsSuffix("järn"); w != g {
		t.Errorf(ts, w, g)
	}

	d.AddSuffix("tjärn")

	decomps := d.Decomp("syltjärn")
	//fmt.Printf("FANTASTIC PLASTIC: %#v\n\n", decomps)
	if w, g := 2, len(decomps); w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := 2, len(decomps[0]); w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := 2, len(decomps[1]); w != g {
		t.Errorf(ts, w, g)
	}

	// both first comp-parts:
	p1 := decomps[0][0]
	p2 := decomps[1][0]

	if p1 == p2 {
		t.Error("Aouch")
	}

	if p1 != "syl" && p2 != "syl" {
		t.Error("Aouch")
	}
	if p1 != "sylt" && p2 != "sylt" {
		t.Error("Aouch")
	}

	// both last comp-parts:
	p3 := decomps[0][1]
	p4 := decomps[1][1]
	if p3 != "järn" && p3 != "tjärn" {
		t.Error("Gosh: " + p3)
	}
	if p4 != "järn" && p4 != "tjärn" {
		t.Error("Gosh: " + p4)
	}
}

func TestDecompRecursivePrefixes(t *testing.T) {

	decomp := NewDecompounder()
	decomp.AddPrefix("svavel")
	decomp.AddPrefix("kanin")

	decomp.AddSuffix("förening")

	ds1 := decomp.Decomp("svavelkaninförening")
	//ds1 := decomp.Decomp("svavelförening")
	if w, g := 1, len(ds1); w != g {

		t.Errorf(ts, w, g)
	}

	decomp.AddSuffix("kanin")

	ds2 := decomp.Decomp("kaninkanin")
	if w, g := 1, len(ds2); w != g {

		t.Errorf(ts, w, g)
	}

	ds3 := decomp.Decomp("kaninkaninkaninkaninkanin")
	if w, g := 1, len(ds3); w != g {
		t.Errorf(ts, w, g)

	}
	if w, g := 5, len(ds3[0]); w != g {
		t.Errorf(ts, w, g)
	}

	ds4 := decomp.Decomp("kaninkaninsvavelkaninkanin")
	if w, g := 1, len(ds4); w != g {
		t.Errorf(ts, w, g)

	}
	if w, g := 5, len(ds4[0]); w != g {
		t.Errorf(ts, w, g)
	}

	// Oh my... the following test was made to cath an
	// over-generation error, due to the fact that a prefix
	// initially was allowed to end at the end of the input
	// string. This was changed, so that a prefix must end before
	// the end of the input string.

	decomp.AddPrefix("k")
	decomp.AddPrefix("a")
	decomp.AddPrefix("ka")
	decomp.AddPrefix("kan")
	decomp.AddPrefix("nin")
	decomp.AddPrefix("in")
	decomp.AddPrefix("i")
	decomp.AddPrefix("n")

	ds5 := decomp.Decomp("kaninkanin")
	unique := make(map[string]bool)
	for _, d0 := range ds5 {
		d := strings.Join(d0, "+")
		if unique[d] {
			fmt.Printf("DARN! %v\n", d)
		} else {
			unique[d] = true
		}
	}
	if w, g := len(unique), len(ds5); w != g {
		t.Errorf(ts, w, g)
	}

	n3 := "xnikolaj3000"

	decomp.AddPrefix(n3)
	ds6 := decomp.Decomp(n3)
	if w, g := 0, len(ds6); w != g {
		t.Errorf(ts, w, g)
	}
	p6 := decomp.prefixes.Prefixes(n3)
	if w, g := 0, len(p6); w != g {
		t.Errorf(ts, w, g)
	}
	// p6b := decomp.prefixes.RecursivePrefixes(n3)

	// fmt.Printf("SDGGDSGDG %#v\n", p6b)
	// if w, g := 0, len(p6b); w != g {
	// 	t.Errorf(ts, w, g)
	// }
	// p6b2 := decomp.prefixes.RecursivePrefixes(n3 + n3)
	// if w, g := 1, len(p6b2); w != g {
	// 	t.Errorf(ts, w, g)
	// }

	decomp.AddSuffix(n3)
	ds7 := decomp.Decomp(n3)
	if w, g := 0, len(ds7); w != g {
		t.Errorf(ts, w, g)
	}

	s7 := decomp.suffixes.Suffixes(n3)
	if w, g := 0, len(s7); w != g {
		t.Errorf(ts, w, g)
	}

	ds8 := decomp.Decomp(n3 + n3)
	if w, g := 1, len(ds8); w != g {
		t.Errorf(ts, w, g)
	}
	// TODO ??
	//fmt.Println("OBS: bugg: alen -> ale+n (stockholmsfinalen)")
}

func TestAlen(t *testing.T) {
	decomp := NewDecompounder()
	decomp.AddPrefix("ale")
	decomp.AddPrefix("n")
	decomp.AddPrefix("fin")
	decomp.AddPrefix("stockholms")

	decomp.AddSuffix("finalen")

	ds1 := decomp.Decomp("alen")
	if w, g := 0, len(ds1); w != g {
		t.Errorf(ts, w, g)
	}

	p1 := decomp.prefixes.Prefixes("alen")
	if w, g := 1, len(p1); w != g {
		t.Errorf(ts, w, g)
	}
	w := arc{start: 0, end: 3, cat: prefix}
	g := p1[0]
	if w != g {
		t.Errorf(ts, w, g)
	}

	ds2 := decomp.Decomp("stockholmsfinalen")
	if w, g := 1, len(ds2); w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := 2, len(ds2[0]); w != g {
		t.Errorf(ts, w, g)
	}

}

func TestGreek1(t *testing.T) {
	decomp := NewDecompounder()

	decomp.AddPrefix("μυκηναική")
	decomp.AddPrefix("μυκη")
	decomp.AddPrefix("ναικήραγ")
	decomp.AddSuffix("ναική")
	decomp.AddSuffix("ραγκαβά")
	decomp.AddSuffix("καβά")
	decomp.AddSuffix("βάκα")

	// fallback test for swe with non-ascii chars
	decomp.AddPrefix("snö")
	decomp.AddPrefix("snör")
	decomp.AddSuffix("bollar")
	ds0a := decomp.Decomp("snörbollar")
	if w, g := 1, len(ds0a); w != g {
		t.Errorf(ts, w, g)
	}
	ds0b := decomp.Decomp("snöbollar")
	if w, g := 1, len(ds0b); w != g {
		t.Errorf(ts, w, g)
	}
	p0 := decomp.prefixes.Prefixes("snöra")
	if w, g := 2, len(p0); w != g {
		t.Errorf(ts, w, g)
	}

	//test for greek chars
	cp := decomp.ContainsPrefix("μυκηναική")
	if w, g := true, cp; w != g {
		t.Errorf(ts, w, g)
	}

	p1 := decomp.prefixes.Prefixes("μυκηναική_")
	if w, g := 2, len(p1); w != g {
		t.Errorf(ts, w, g)
	}
	ds1 := decomp.Decomp("μυκηναικήραγκαβά")
	if w, g := 2, len(ds1); w != g {
		t.Errorf(ts, w, g)
	}
	ds2 := decomp.Decomp("μυκηβάκα")
	if w, g := 1, len(ds2); w != g {
		t.Errorf(ts, w, g)
	}

}

func TestLenSort(t *testing.T) {

	// return the versions with fewest compound parts first

	decomp := NewDecompounder()
	decomp.AddPrefix("odalbonde")
	decomp.AddPrefix("bonde")
	decomp.AddPrefix("odal")

	decomp.AddSuffix("husbil")
	decomp.AddSuffix("bil")

	decomp.AddPrefix("hus")

	ds1 := decomp.Decomp("odalbondehusbil")
	soFar := 0
	for _, d := range ds1 {
		if len(d) < soFar {
			t.Errorf("this thingy is not sorted: %#v", ds1)
			return
		}
		soFar = len(d)
	}

}

func TestTripleCons(t *testing.T) {

	decomp := NewDecompounder()
	decomp.AllowedTripleChars([]rune{'t'})
	decomp.AddPrefix("natt")
	decomp.AddSuffix("tåg")

	ds0 := decomp.Decomp("natttåg")
	//fmt.Printf("%v\n", ds0)
	if w, g := 1, len(ds0); w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := "natt", ds0[0][0]; w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := "tåg", ds0[0][1]; w != g {
		t.Errorf(ts, w, g)
	}

	ds1 := decomp.Decomp("nattåg")
	//fmt.Printf("%v\n", ds1)
	if w, g := 1, len(ds1); w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := "natt", ds1[0][0]; w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := "tåg", ds1[0][1]; w != g {
		t.Errorf(ts, w, g)
	}

	decomp.AddPrefix("glass")
	decomp.AddSuffix("strut")
	decomp.AllowedTripleChars([]rune{'s', 't'})
	ds2 := decomp.Decomp("glasstrut")
	if w, g := 1, len(ds2); w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := "glass", ds2[0][0]; w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := "strut", ds2[0][1]; w != g {
		t.Errorf(ts, w, g)
	}

	decomp.AddPrefix("glas")
	decomp.AddSuffix("trut")
	ds3 := decomp.Decomp("glasstrut")
	//fmt.Printf(">>> %#v\n", ds3)

	if w, g := 3, len(ds3); w != g {
		t.Errorf(ts, w, g)
	}

	var glasStrut bool
	var glassTrut bool
	var glassStrut bool
	for _, d := range ds3 {
		if d[0] == "glas" && d[1] == "strut" {
			glasStrut = true
		}
		if d[0] == "glass" && d[1] == "trut" {
			glassTrut = true
		}
		if d[0] == "glass" && d[1] == "strut" {
			glassStrut = true
		}
	}
	if !(glasStrut && glassTrut && glassStrut) {
		t.Errorf("failure of enormous proportions")
	}
}

func TestInfixS(t *testing.T) {

	decomp := NewDecompounder()
	decomp.AddPrefix("finland")
	decomp.AddSuffix("båt")

	ds1 := decomp.Decomp("finlandbåt")
	if w, g := 1, len(ds1); w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := 2, len(ds1[0]); w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := "finland", ds1[0][0]; w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := "båt", ds1[0][1]; w != g {
		t.Errorf(ts, w, g)
	}

	decomp.AddInfix("s")

	ds2 := decomp.Decomp("finlandsbåt")
	if w, g := 1, len(ds2); w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := 3, len(ds2[0]); w != g {
		t.Errorf(ts, w, g)
	}

	if w, g := "s", ds2[0][1]; w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := "finland", ds2[0][0]; w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := "båt", ds2[0][2]; w != g {
		t.Errorf(ts, w, g)
	}

	decomp.AddPrefix("båt")
	decomp.AddSuffix("kapten")

	ds3 := decomp.Decomp("finlandsbåtskapten")
	if w, g := 1, len(ds3); w != g {
		t.Errorf(ts, w, g)
	}

	if w, g := "kapten", ds3[0][4]; w != g {
		t.Errorf(ts, w, g)
	}

}

func TestLoadFromFile(t *testing.T) {
	d, err := NewDecompounderFromFile("test_word_parts.txt")
	if err != nil {
		t.Errorf("%v", err)
	}

	if w, g := 2, len(d.tripleChars); w != g {
		t.Errorf(ts, w, g)
	}

	if w, g := true, d.tripleChars['s']; w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := true, d.tripleChars['t']; w != g {
		t.Errorf(ts, w, g)
	}

	res1 := d.Decomp("grusvägarna")
	// One suggestion only
	if w, g := 1, len(res1); w != g {
		t.Errorf(ts, w, g)
	}
	// ... containing two word parts
	if w, g := 2, len(res1[0]); w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := "grus", res1[0][0]; w != g {
		t.Errorf(ts, w, g)
	}
	if w, g := "vägarna", res1[0][1]; w != g {
		t.Errorf(ts, w, g)
	}

	// TODO test that saving to file works:
	errf := d.SaveToFile("test_word_parts_out.txt")
	if errf != nil {
		t.Errorf("%v", errf)
	}
}

func TestRemoveBug(t *testing.T) {

	w := "prästgårdsvägen"
	infix := "s"

	p1 := "präst"
	p2 := "gård"

	s1 := "vägen"
	s2 := "svägen" // REMOVE THIS

	decomp := NewDecompounder()
	decomp.AddPrefix(p1)
	decomp.AddPrefix(p2)

	decomp.AddInfix(infix)

	decomp.AddSuffix(s1)
	decomp.AddSuffix(s2)

	res := decomp.Decomp(w)

	if w, g := 2, len(res); w != g {
		t.Errorf(ts, w, g)
	}

	removeRes := decomp.RemoveSuffix(s2)
	if !removeRes {
		t.Errorf(ts, true, removeRes)
	}

	res = decomp.Decomp(w)

	if w, g := 1, len(res); w != g {
		t.Errorf(ts, w, g)
	}

	// Wanted: präst gård s vägen
	if w, g := 4, len(res[0]); w != g {
		t.Errorf(ts, w, g)
	} else {

		if w, g := "präst", res[0][0]; w != g {
			t.Errorf(ts, w, g)
		}

		if w, g := "gård", res[0][1]; w != g {
			t.Errorf(ts, w, g)
		}

		if w, g := "s", res[0][2]; w != g {
			t.Errorf(ts, w, g)
		}

		if w, g := "vägen", res[0][3]; w != g {
			t.Errorf(ts, w, g)
		}
	}

}

func TestAllPotentialPrefixes(t *testing.T) {
	w := "abcd"

	p1 := "b"
	p2 := "c"
	p3 := "bc"

	pFt := newPrefixTree()
	pFt.Add(p1)
	pFt.Add(p2)
	pFt.Add(p3)

	res := pFt.allPotentialPrefixes([]rune(w), 0)

	if w, g := 3, len(res); w != g {
		t.Errorf(ts, w, g)
	}

}

func TestTripleConsBug(t *testing.T) {
	w := "nattågspersonal"
	s1 := "natt"
	s2 := "tåg"

	i1 := "s"

	s3 := "personal"

	decomp := NewDecompounder()
	decomp.AllowedTripleChars([]rune{'t'})

	decomp.AddPrefix(s1)
	decomp.AddPrefix(s2)

	decomp.AddInfix(i1)

	decomp.AddSuffix(s3)

	res := decomp.Decomp(w)

	if w, g := 1, len(res); w != g {
		t.Errorf(ts, w, g)
	}

	if w, g := 4, len(res[0]); w != g {
		t.Errorf(ts, w, g)
	}
}
