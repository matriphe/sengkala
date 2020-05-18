package sengkala

import (
	"fmt"
	"strconv"
	"strings"
)

type item struct {
	year     int
	err      error
	sengkala []string
	meaning  map[string]map[string]string
}

var watak = map[int][][]string{
	0: {
		{"byoma", "musna", "nis", "mlethik", "langit"},
		{"sirna", "ilang", "kombul", "awang-awang"},
		{"mesat", "muluk", "gegana", "nglès"},
		{"tumenga", "nenga", "luhur"},
		{"suwung", "sonya", "muksa", "doh", "tebih"},
		{"swarga", "tanpa", "barakan"},
		{"tan", "rusak", "brastha", "swuh"},
		{"walang", "kos", "pejah", "akasa"},
		{"tawang", "wiyat", "oncat", "windu", "widik-widik"},
		{"nir", "wuk", "sat", "surud", "sempal"},
	},
	1: {
		{"tunggal", "gusti", "sujanma", "semedi"},
		{"badan", "nabi", "rupa", "maha", "buddha"},
		{"niyata", "luwih", "pamasé"},
		{"wong", "buweng", "rat", "lèk", "iku"},
		{"surya", "candra", "kartika", "bumi"},
		{"wiji", "urip", "ron", "éka"},
		{"prabu", "kenya", "nekung"},
		{"raja", "putra", "sasa", "dhara"},
		{"peksi", "dara", "tyas", "wungkul", "sudira", "budi"},
		{"wani", "hyang", "jagad", "nata"},
	},
	2: {
		{"asta", "kalih", "ro", "nembah", "ngabekti"},
		{"nétra", "kembar", "myat", "mandeng", "nayana"},
		{"swiwi", "lar", "sikara", "gandhèng"},
		{"paksa", "apasang", "sungu"},
		{"athi-athi", "talingan", "dresthi"},
		{"carana", "tangan", "karna"},
		{"bau", "suku", "caksuh"},
		{"mata", "paningal", "locana"},
		{"ama", "nebah", "karnan", "ngrengga", "pengantèn", "dwi"},
		{"kanthi", "buja", "bujana"},
	},
	3: {
		{"bahni", "tiga", "ujwala", "kaèksi"},
		{"katon", "murub", "dahana", "payudah"},
		{"katingalan", "kaya", "bentèr"},
		{"nala", "uninga", "kawruh"},
		{"lir", "wrin", "wéda", "naut", "nauti"},
		{"teken", "siking", "pawaka"},
		{"kukus", "api", "apyu"},
		{"brama", "rana", "rananggana"},
		{"utawaka", "uta", "ujel", "kobar", "agni"},
		{"wignya", "guna", "tri", "jatha"},
	},
	4: {
		{"catur", "warna", "wahana", "pat", "warih"},
		{"waudadi", "dadya", "keblat", "papat"},
		{"toya", "suci", "udaka", "we"},
		{"woh", "nadi", "jladri", "sindu"},
		{"yoga", "gawé", "tlaga", "hèr", "wening"},
		{"udan", "bun", "tirta", "marta"},
		{"karya", "sumber", "sumur"},
		{"masuh", "marna", "karti", "karta"},
		{"jalaniddhi", "samodra", "udaya", "tasik"},
		{"tawa", "segara", "wédang"},
	},
	5: {
		{"pandhawa", "lima", "wisikan", "gati"},
		{"indri", "indriya", "warastra", "wrayang"},
		{"astra", "lungid", "sara", "saré"},
		{"guling", "raseksa", "diyu"},
		{"buta", "galak", "wil", "yaksa", "yaksi"},
		{"saya", "wisaya", "bana"},
		{"jemparing", "cakra", "hru"},
		{"tata", "hanata", "bayu", "bajra"},
		{"samirana", "pawaka", "maruta", "angin"},
		{"panca", "marga", "margana"},
	},
	6: {
		{"rasa", "nenem", "rinaras", "artati"},
		{"lona", "tikta", "madura", "sarkara"},
		{"amla", "kayasa", "karaséng"},
		{"oyag", "obah", "nem", "kayu"},
		{"wreksa", "glinggang", "prabatang", "oyig"},
		{"sad", "anggas", "anggang-anggang"},
		{"mangsa", "naya", "retu"},
		{"wayang", "winayang", "anggana"},
		{"ilat", "kilat", "lidhah", "lindhu", "carem", "manis"},
		{"tahen", "osik", "karengya"},
	},
	7: {
		{"sapta", "prawata", "acala", "giri"},
		{"ardi", "gora", "prabata", "himawan"},
		{"pandhita", "pitu", "kaswarèng"},
		{"resi", "sogata", "wiku"},
		{"yogi", "swara", "dwija", "suyati"},
		{"wulang", "weling", "wasita"},
		{"tunggang", "turangga", "gung"},
		{"swa", "aswa", "titihan", "kuda"},
		{"ajar", "arga", "sabda", "nabda", "angsa", "muni"},
		{"suka", "biksu", "biskuka"},
	},
	8: {
		{"astha", "basu", "anggusthi", "basuki"},
		{"slira", "murti", "bujangga", "manggala"},
		{"taksaka", "menyawak", "tekèk"},
		{"dwipa", "dwipangga", "bajul"},
		{"gajah", "liman", "dwirada", "èsthi"},
		{"éstha", "matengga", "brahma"},
		{"brahmana", "wewolu"},
		{"baya", "bebaya", "kunjara"},
		{"tanu", "sarpa", "samaja", "madya", "mangèsthi"},
		{"panagan", "ula", "naga"},
	},
	9: {
		{"bolong", "nawa", "dwara", "pintu", "kori"},
		{"bedhah", "lawang", "wiwara", "gapura"},
		{"rong", "song", "wilasita", "anglèng"},
		{"trustha", "trusthi", "trus", "butul"},
		{"déwa", "sanga", "jawata", "manjing"},
		{"arum", "ganda", "kusuma"},
		{"muka", "rudra", "masuk"},
		{"rago", "angrong", "guwa", "menga"},
		{"babahan", "lèng", "ambuka", "gatra", "anggangsir"},
		{"nanda", "wangi", "wadana"},
	},
}

// GetSengkala returns sengkala words
func GetSengkala(year string, randomizer Randomizer) []string {
	res := make([]string, len(year))
	for i, t := range reverse(year) {
		it, _ := strconv.Atoi(fmt.Sprintf("%c", t))
		w := watak[it]
		ch := w[randomizer.GetRandom(len(w)-1)]
		res[i] = strings.Title(ch[randomizer.GetRandom(len(ch)-1)])
	}

	return res
}

func newItem(year string) *item {
	y, err := strconv.Atoi(year)
	if err != nil {
		return &item{
			err: err,
		}
	}

	return &item{
		year:     y,
		err:      nil,
		sengkala: GetSengkala(year, newRandomizer()),
		meaning:  nil,
	}
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
