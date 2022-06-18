package A_basicLangFeatures

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	mappie := make(map[string]int)

	for _, v := range strings.Fields(s) {

		mappie[v] = mappie[v] + 1
	}

	return mappie
}

func Maps_main() {
	wc.Test(WordCount)
	// fmt.Println(WordCount("I am Learning Go! Learning"))
}
