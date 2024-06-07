package libmessage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func Test_open(t *testing.T) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://open.api.tianyancha.com/services/open/search/2.0", nil)
	req.Header.Add("Authorization", token)
	q := req.URL.Query()
	q.Add("word", "熙锐")
	q.Add("pageSize", "40")
	q.Add("pageNum", "1")
	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	var data interface{}
	json.Unmarshal(body, &data)
	log.Println(data)
	var mapdata map[string]map[string][]map[string]string
	json.Unmarshal(body, &mapdata)
	item := mapdata["result"]["items"]
	var name_list []string
	for _, v := range item {
		// log.Println(v["name"])
		name_list = append(name_list, v["name"])
	}
	// log.Println(data["result"])
}

// 工商信息和主要人员
func Test_IndustryCommerce(t *testing.T) {
	// log.Println(keywords, pageSize, pageNum)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://open.api.tianyancha.com/services/open/ic/baseinfoV3/2.0", nil)
	req.Header.Add("Authorization", token)
	q := req.URL.Query()
	// pageSize
	// 不要分页
	q.Add("keyword", "深圳市熙锐科技有限公司")
	//q.Add("pageSize", "20")
	//q.Add("pageNum", "1")
	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	log.Println(data)
}

//主要人员
// https://open.api.tianyancha.com/services/open/ic/staff/2.0
//func Test_Personnel(t *testing.T) {
//	// log.Println(keywords, pageSize, pageNum)
//	client := &http.Client{}
//	req, _ := http.NewRequest("GET", "https://open.api.tianyancha.com/services/open/ic/staff/2.0", nil)
//	req.Header.Add("Authorization", token)
//	q := req.URL.Query()
//	// pageSize
//	// 要分页
//	q.Add("keyword", "深圳市熙锐科技有限公司")
//	q.Add("pageSize", "20")
//	q.Add("pageNum", "1")
//	req.URL.RawQuery = q.Encode()
//	resp, _ := client.Do(req)
//	body, _ := ioutil.ReadAll(resp.Body)
//	var data interface{}
//	json.Unmarshal(body, &data)
//	log.Println(data)
//}

// 建筑资质

func Test_Qualification(t *testing.T) {
	// log.Println(keywords, pageSize, pageNum)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://open.api.tianyancha.com/services/open/bq/qualification/2.0", nil)
	req.Header.Add("Authorization", token)
	q := req.URL.Query()
	// pageSize
	// 要分页
	q.Add("keyword", "安徽省公路工程建设监理有限责任公司")
	q.Add("pageSize", "20")
	q.Add("pageNum", "1")
	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	log.Println(data)
}

// 司法风险
func Test_lawSuit(t *testing.T) {
	// log.Println(keywords, pageSize, pageNum)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://open.api.tianyancha.com/services/open/jr/lawSuit/3.0", nil)
	req.Header.Add("Authorization", token)
	q := req.URL.Query()
	// pageSize
	// 要分页
	q.Add("keyword", "安徽省公路工程建设监理有限责任公司")
	q.Add("pageSize", "20")
	q.Add("pageNum", "1")
	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	log.Println(data)
}

// 经营风险
//punishmentInfo

func Test_PunishmentInfo(t *testing.T) {
	// log.Println(keywords, pageSize, pageNum)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://open.api.tianyancha.com/services/open/mr/punishmentInfo/3.0", nil)
	req.Header.Add("Authorization", token)
	q := req.URL.Query()
	// pageSize
	// 要分页
	q.Add("keyword", "安徽省公路工程建设监理有限责任公司")
	q.Add("pageSize", "20")
	q.Add("pageNum", "1")
	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	log.Println(data)
}
