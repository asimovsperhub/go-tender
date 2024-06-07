package desk_individual

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

type info struct {
	Name string
	Icon *os.File
}

func Test_add(t *testing.T) {
	f, err := os.Open("/Users/apple/Desktop/1.png")
	if err != nil {
		log.Println("read file fail", err)
	}
	defer f.Close()
	//fd, err := ioutil.ReadAll(f)
	//if err != nil {
	//	log.Println("read to fd fail", err)
	//}
	//fmt.Println(string(fd))
	da, _ := json.Marshal(&info{Name: "zzw", Icon: f})
	resp, err := http.NewRequest("POST", "http://127.0.0.1/api/v1/desk/enterprise/add", strings.NewReader(string(da)))
	if err != nil {
		panic(err)
	}
	resp.Header.Set("Authorization", "Bearer lQt0hsj5/IfgIzRkq0jcWBu2hZdn5eGiUCxLAl1+AB5+IUsgC54K1UVV9t7qiqdrLEdV5AaZXAcWeLMYyl91gdqFz5NrclAK84iCgCZshSgVELyFImqxWoCuI17c0GMbl/Bxu31l6QjhDhutLmVA9Q==")
	// defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	fmt.Println(resp.Header)
}
