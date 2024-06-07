package libmessage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//6887ada6-818b-4d70-8498-4190b79cacdd
const token = "b48f733a-cf5a-47a9-ae68-ab9970251811"

func OpenTanYanChaSearch(keywords string, pageSize string, pageNum string) interface{} {
	log.Println(keywords, pageSize, pageNum)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://open.api.tianyancha.com/services/open/search/2.0", nil)
	req.Header.Add("Authorization", token)
	q := req.URL.Query()
	// pageSize
	q.Add("word", keywords)
	q.Add("pageSize", pageSize)
	q.Add("pageNum", pageNum)
	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	return data
}

// 工商信息

func IndustryCommerceAndPersonnel(keywords string) (interface{}, int) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://open.api.tianyancha.com/services/open/ic/baseinfoV3/2.0", nil)
	req.Header.Add("Authorization", token)
	q := req.URL.Query()
	// pageSize
	q.Add("keyword", keywords)
	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	var res map[string]int
	json.Unmarshal(body, &res)
	code := res["error_code"]
	return data, code
}

// 主要人员
//func Personnel(keywords string) interface{} {
//	client := &http.Client{}
//	req, _ := http.NewRequest("GET", "https://open.api.tianyancha.com/services/open/ic/staff/2.0", nil)
//	req.Header.Add("Authorization", "d618c4be-eaba-43e0-bf8e-1c68ba543bc4")
//	q := req.URL.Query()
//	// pageSize
//	// 要分页
//	q.Add("keyword", keywords)
//	q.Add("pageSize", "20")
//	q.Add("pageNum", "1")
//	req.URL.RawQuery = q.Encode()
//	resp, _ := client.Do(req)
//	body, _ := ioutil.ReadAll(resp.Body)
//	var data interface{}
//	json.Unmarshal(body, &data)
//	log.Println(data)
//	return data
//}

// 建筑资质
func Qualification(keywords string, pageSize string, pageNum string) (interface{}, int) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://open.api.tianyancha.com/services/open/bq/qualification/2.0", nil)
	req.Header.Add("Authorization", token)
	q := req.URL.Query()
	// pageSize
	// 要分页
	q.Add("keyword", keywords)
	q.Add("pageSize", pageSize)
	q.Add("pageNum", pageNum)
	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	log.Println(data)
	var res map[string]int
	json.Unmarshal(body, &res)
	code := res["error_code"]
	return data, code
}

// 司法风险
func LawSuit(keywords string, pageSize string, pageNum string) (interface{}, int) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://open.api.tianyancha.com/services/open/jr/lawSuit/3.0", nil)
	req.Header.Add("Authorization", token)
	q := req.URL.Query()
	// pageSize
	// 要分页
	q.Add("keyword", keywords)
	q.Add("pageSize", pageSize)
	q.Add("pageNum", pageNum)
	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	log.Println(data)
	var res map[string]int
	json.Unmarshal(body, &res)
	code := res["error_code"]
	return data, code
}

// 经营风险
func PunishmentInfo(keywords string, pageSize string, pageNum string) (interface{}, int) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://open.api.tianyancha.com/services/open/mr/punishmentInfo/3.0", nil)
	req.Header.Add("Authorization", token)
	q := req.URL.Query()
	// pageSize
	// 要分页
	q.Add("keyword", keywords)
	q.Add("pageSize", pageSize)
	q.Add("pageNum", pageNum)
	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	log.Println(data)
	var res map[string]int
	json.Unmarshal(body, &res)
	code := res["error_code"]
	return data, code
}
