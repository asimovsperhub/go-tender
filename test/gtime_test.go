package test

import (
	"github.com/gogf/gf/v2/os/gtime"
	"log"
	"testing"
)

func Test_gtime(t *testing.T) {
	log.Println(gtime.Now().AddDate(0, 2, 0))
}
