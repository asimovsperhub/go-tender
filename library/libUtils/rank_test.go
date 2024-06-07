package libUtils

import (
	"log"
	"strings"
	"testing"
)

func Test_Rank(t *testing.T) {
	// w54264
	prev := "w50856"
	next := "w66391"
	cc, _ := Rank(prev, next)
	log.Println(cc)
}

func Test_Pank(t *testing.T) {
	//log.Println("w50856" < "w66391")
	//if "ah" < "b" {
	//	log.Println("ah" < "b")
	//}
	log.Println(strings.Replace("广东省", "省", "", -1))
}
