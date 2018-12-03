package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

var data2 = `abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`

var data = `ohvflkatysoimjxbunazgwcdpr
ohoflkctysmiqjxbufezgwcdpr
ohvflkatysciqwxfunezgwcdpr
fhvflyatysmiqjxbunazgwcdpr
ohvhlkatysmiqjxbunhzgwcdxr
ohvflbatykmiqjxbunezgscdpr
ohvflkatasaiqjxbbnezgwcdpr
ohvflkatyymiqjxrunetgwcdpr
ohvflkatbsmiqhxbunezgwcdpw
oheflkytysmiqjxbuntzgwcdpr
ohvflkatrsmiqjibunezgwcupr
ohvflkaiysmiqjxbunkzgwkdpr
ohvilkutysmiqjxbuoezgwcdpr
phvflkatysmkqjxbulezgwcdpr
ohvflkatnsmiqjxbznezgpcdpr
ohvylkatysriqjobunezgwcdpr
ohvflkatytmiqjxbunezrwcypr
ohvonkatysmiqjxbunezgwxdpr
ohvflkatgsmoqjxyunezgwcdpr
ohvflkbtqsmicjxbunezgwcdpr
ohvflkatysmgqjqbunezgwcdvr
ohvtlkatyrmiqjxbunezgwcdpi
ohvflkatyskovjxbunezgwcdpr
ohvflkayysmipjxbunezgwcdpu
ohvalkltysmiqjxbunezgecdpr
ohvflkatysmiqjxiunezgnndpr
ohvflkatyomiqjxbbnezgwcdpp
ohvflkatysmiqjxbuoezgncdpy
omvflkvtysmiqjxwunezgwcdpr
ohvflkatynmicjxbunezgwpdpr
ohvflkatyqmaqjxbunezvwcdpr
ohbfhkatysmiqjxbunezgwcdqr
ohvflkatesmiqjvbunezpwcdpr
ohvflkatysmsqjxiunezgwcdhr
ohvfjkatysmwqjxbunezgwcddr
ohvflkanysmiqjxbunwkgwcdpr
ohqflkatysmiqjxbuuezgwcddr
ohvflkatysmvqjxbznlzgwcdpr
ohvflkatysmiqjxbunjzwwqdpr
ohvfjkatysmiqxxbunezgwcupr
chvfxkatysmiqjxxunezgwcdpr
uhvflkatitmiqjxbunezgwcdpr
ohvflbatysmiqjxbuntzgwcdor
ohvflkmtysmmqjxbunexgwcdpr
ohvflsatysmyqjxjunezgwcdpr
ohvfskatysmiqjjbunezgwcdpg
ohvflkatysniqjxbunexgwcrpr
ohvfekatysmiqjxbunedswcdpr
ohvfltatysmjqjxbunezghcdpr
ohvflkatydmiqjxvunezggcdpr
oavflkatysmiqjxtunazgwcdpr
ohvflkltysmiqjxbuzeugwcdpr
ohbflkatysmiqjybuuezgwcdpr
ehvfzkatysmiqjxbuhezgwcdpr
odvflkatssmiqjxbunezgwcdpj
ohvflkatysmiqjzbufezgwbdpr
jhvflkdtysmiqqxbunezgwcdpr
ohvflkatysmiqjwbunengwcnpr
ohvfskatysmiqjxbxuezgwcdpr
ohvflkatysmiqjobvnezgwcrpr
ohvrlkatysmiqjxbwnezgrcdpr
ofvflkatysmiqjxbunezpwcdwr
ohvfxdatyomiqjxbunezgwcdpr
yhvflkatydmiqjxbubezgwcdpr
ohvflkatysdiqjxbuneztwcspr
ohvflkatydmiquxbunezgwcbpr
ohvflkatysmiqcxbukezgwcdwr
ohvflkntasmiqjxbunezghcdpr
lhvflkatysmiqjxbunezqwckpr
ehifikatysmiqjxbunezgwcdpr
ohvflkatysmiqjcbutezgwcdpm
ohvflkatjssiqrxbunezgwcdpr
oyvflkavysmiqjxlunezgwcdpr
orvflkgtysmiqjxbukezgwcdpr
ihvflkatysmiqaxbunpzgwcdpr
ohvflkatusmiqjxbbnezgwchpr
ohvflkatysbiqjxvuneugwcdpr
ohvflkatysmiqjcbungzgwcwpr
ovvflkatysmidjxbunezgscdpr
ohvflqatysmiljxbunfzgwcdpr
ghvfokatysmiqjxbunqzgwcdpr
nxvflkatysmxqjxbunezgwcdpr
ohvflkatysmiqjxbexezgwrdpr
ohvfrkatysmhqjxbuntzgwcdpr
ohvflkvtysmiqjxocnezgwcdpr
ohvglkgtysmiqjxnunezgwcdpr
ohvflkatysmnqjxbunecgwqdpr
oyvflkatysgiqjxbcnezgwcdpr
ofvflkatysmiqjxbunfzgwcdpg
otvflkttysmiqjxbunezgwmdpr
ohvflkvtysmiqjbbunezgzcdpr
ahvflkatysyiqjxbunezvwcdpr
ohiflkatysmydjxbunezgwcdpr
ohvfwkatysmvqjxbunezwwcdpr
ohvflkatysbiqjxbunergwodpr
hhvsdkatysmiqjxbunezgwcdpr
ihvflkwtysmiqjxbunezgacdpr
ohvfljatysmiqcxbunuzgwcdpr
ohvflkatysqiqlwbunezgwcdpr
ohvflkauysmkqjxwunezgwcdpr
ohvflkatysmoqjqbunezgwodpr
ohvslkvtysmipjxbunezgwcdpr
olvflkatysmiujxbunezgwctpr
osvflxatysmiqjxbenezgwcdpr
orvflkhtysmiqjxbinezgwcdpr
ohcflkatystiqjxbunezbwcdpr
ohcflkatyfmifjxbunezgwcdpr
ohvflkatdsmiqjxbrnezgwcdpt
ohvflkatysmiqjxbwnqzawcdpr
oevflkakysmiqjxbunezgwcdpt
ofvflkatysmiqjxbunbqgwcdpr
ohvflkatysmdqjxbunefqwcdpr
ohvklkalysmiqjxbunezgwcepr
ocvflhatysmiqjxbunezzwcdpr
uhvflkatysmiqmxbunezgwcxpr
ohvflkatyshikjhbunezgwcdpr
lbvflkatysmoqjxbunezgwcdpr
ohvflkatssmuqjxbunezgscdpr
ohvflkatysmifyxbuvezgwcdpr
ohvfikatysmiqjxbunezgwfupr
ohvmlkaiysmiqjxqunezgwcdpr
ohvflkatysmiqjxiunpzgwcdpo
lhvflkatysmpqjxbenezgwcdpr
ohvflkatysmiqjobunengwczpr
ohoflkatysniqjxbunezgccdpr
ohvfxkatysmiqjgbunyzgwcdpr
ohvflkytysmiljxbubezgwcdpr
hhvsdkatysmiqjxjunezgwcdpr
ohvflkatysmiqjtuunezgwcdpt
ohvfdkxtysmiqjubunezgwcdpr
ohxflkatysmiyjxbunezgwcdhr
ohvflkatysmiqjibunezgwcppd
ohvflkatysmihjxbunezgwcdhj
ohvflkatysmiqjxronezgwcdvr
ofrflxatysmiqjxbunezgwcdpr
ohvwlkatysmiqjxounezgscdpr
ohvflkatcodiqjxbunezgwcdpr
oqvflkatysmiqjxbunebgwmdpr
ohvflmatysmisjxbunezqwcdpr
ovvflkatysmiqjxbuxezgwcdpe
ohvflkatysmdejxbuneztwcdpr
hhvflkathsmiqjxbwnezgwcdpr
ohkflkatlsmsqjxbunezgwcdpr
ohvflkktysmizjxhunezgwcdpr
ohzflkatysmiqjrbunezgwcdpj
ohuflwatysmiqjxbunezgwcdgr
ohvflkatysmiqvxmunpzgwcdpr
xhvflkwtysmiqjxbunezgwjdpr
whvflkatysmiqjxbunezgzcopr
ohvflkayysmiqjxuznezgwcdpr
khvflkasysmiqjxbunezgwcdpv
ohvflkatylmiqjxbpnozgwcdpr
ohvflkgtysziqjxbunezgwgdpr
ohvfljaiysmiqjxbuvezgwcdpr
ohvflkxtyslizjxbunezgwcdpr
ohzflkatysmiqjxbcnezgwcdar
ohvflkatysmiqjxbisecgwcdpr
shvflkatyjmiqjkbunezgwcdpr
mhvflkatysmiqjxvunezgwcdpk
ohfflkatysmiqjxbunczgwcppr
ohvflkatysmiqjkzunezgwcdpc
ohvflkatysmifjxbuneygwctpr
ohvflkatysmimjbbunezgwcdpe
ohvflkatjsciqjxbunezgwcdpa
ohvxlkatysmitjxbunezswcdpr
ohvslkatfsmiqjxbunezgwudpr
ohvflkatysmiqexbugezgwcdnr
onvflkatysmiqjxkunezgtcdpr
fhsflkalysmiqjxbunezgwcdpr
oyvflkatysmiqjobxnezgwcdpr
ohvflkatysmiqjxbunezswgdvr
phvflkatyymiqjxvunezgwcdpr
oivflzutysmiqjxbunezgwcdpr
ohvflkftysmiqjxbunezkwcopr
ohvflkatysmwnjxbunezgwcdpp
ohvflkatysmiqkxcunezgwndpr
phvklkatysmiqjhbunezgwcdpr
ohvflrawysmiqjxbunhzgwcdpr
ohvflkatysmiqjxbunecgwcdig
ohvflpakysmiqjxbunezgwrdpr
odvflkatykmiqjxbunezglcdpr
ohtflkatysiiqjxblnezgwcdpr
lhvfpkatysmiqjxbupezgwcdpr
ohvflkatdsmiqjpbunezgwcdps
ohvflkztysmiqjxvunezgwjdpr
ohvflbatysmxqoxbunezgwcdpr
ohvklkaigsmiqjxbunezgwcdpr
ohvfgkawysmiqjxbunezgwcdur
ohvflkatyskpqjlbunezgwcdpr
ohvflkatyqmiqjhbupezgwcdpr
ohqflkatysmiqjxzonezgwcdpr
ohxfnkatyymiqjxbunezgwcdpr
ohmflkatpsmiqjxbunezgwcdpw
ohvflkatysmiqjibnnewgwcdpr
vevflkatysmiqjxbunezgwcypr
ohvflkatydmiqwxbungzgwcdpr
ohsrlkatysmiqjxbcnezgwcdpr
ohvflkptyvmiqexbunezgwcdpr
opzflkatysmiqjxrunezgwcdpr
ohvflkitysmiqjxcunezgwcmpr
ohvflkatysmhhjxblnezgwcdpr
ohvflkatysfiqjxbunrzgwmdpr
ohvflkatyamibjxbunezgwcdpf
ohvflkalysmigjxbunezggcdpr
ohvflkatwsmisjxbunezgdcdpr
dhvflkatysmlqjxbunszgwcdpr
ohvflkatysmiqjxbueeygwcbpr
ohvflkatgsmiqjnbunezhwcdpr
svvflkatysmiqjxbunezgwckpr
opvflkatysmiqpxbufezgwcdpr
ohnvlkatysmiqjxbunezglcdpr
phvflkutysjiqjxbunezgwcdpr
ohvflabtysmiqjjbunezgwcdpr
ouvflkatysmiqjsbunezgwcdpk
osvflkatysmijjxbunezgwcypr
owvflkatysmiqjxbukxzgwcdpr
ohvfliatvsmiljxbunezgwcdpr
ohvflkatysmiqjxbumezbwtdpr
ohvflkatyfcicjxbunezgwcdpr
ohvflkatysmiqldbunezgfcdpr
oqvflkatysmiqixkunezgwcdpr
ohvflkatysmiqjxbulezgicdpe
ohvflkatysmiqjxbuniegwcdpl
ohvflkatysmiqjwbunbzgwcdhr
ohvflkatysmiqjdbunezgwwdkr
ohqflkytysmiqjxbunezgwcdpc
ohvflkatysmigjxbunezqwwdpr
ohvfloatysmiqjpbumezgwcdpr
ohvklkathkmiqjxbunezgwcdpr
ohvflkstjsmiqjxbunezgwctpr
ohvvlkatysmiqjxbunewgwcdir
ohnflkatysmiqjxbunszgwcdlr
ohvflkatysmnqjxbunezgxcdlr
ohvfrkatysmiqjxbonezgwcdor
ihvflkatysmiqjxbuneogwcxpr
ohvflkatysmiqjxbunecgwcccr
owvflkatysmivjxbunezgwjdpr
ohvflkgtysmiqjxbunczhwcdpr
ohyqlkatysmiqjxbunezgwcypr
ohvflkatysmiqjvbunezuwcdpw
ohvflkathsmiqmxbuoezgwcdpr
ehvjlkajysmiqjxbunezgwcdpr
ohvflkltysmiqjxblnezgwjdpr
oovflkvtfsmiqjxbunezgwcdpr
olvfzkatysmiqjxyunezgwcdpr
ohvflkatysqitjxbunezgncdpr
yhvflkatysmkqjxbunazgwcdpr
zlvolkatysmiqjxbunezgwcdpr
ohvflpatysmiqjxbunezgwcapb
ohvflkatysmuqjxbunezgfcdur`

func main() {
	r := strings.NewReader(data)
	s := bufio.NewScanner(r)
	var ls []string
	for s.Scan() {
		ls = append(ls, s.Text())
	}
	sort.Strings(ls)
	for i := 1; i < len(ls); i++ {
		if levenshtein(ls[i-1], ls[i]) == 1 {
			fmt.Printf("%s\n%s\n", ls[i-1], ls[i])
			return
		}
	}
}

func min(is ...int) int {
	m := is[0]
	for _, i := range is {
		if i < m {
			m = i
		}
	}
	return m
}

func levenshtein(i, j string) int {
	y := len(i) + 1
	dm := make([]int, (len(i)+1)*(len(j)+1))
	for x := 0; x <= len(i); x++ {
		dm[x] = x
	}
	for x := 0; x <= len(j); x++ {
		dm[x*y] = x
	}
	for jj := 1; jj <= len(j); jj++ {
		for ii := 1; ii <= len(i); ii++ {
			var c int
			if j[jj-1] != i[ii-1] {
				c = 1
			}
			dm[jj*y+ii] = min(dm[jj*y+ii-1]+1, dm[(jj-1)*y+ii]+1, dm[(jj-1)*y+ii-1]+c)
		}
	}
	if false {
		for jj := 0; jj <= len(j); jj++ {
			for ii := 0; ii <= len(i); ii++ {
				fmt.Printf("%2d ", dm[jj*y+ii])
			}
			fmt.Println()
		}
	}
	return dm[len(j)*y+len(i)]
}
