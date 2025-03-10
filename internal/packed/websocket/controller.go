package websocket

import (
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// LoginController 用户登录
func LoginController(client *Client, req *request) {

	// userId := gconv.Uint64(0)

	uid, ok := req.Data["user_id"].(json.Number)

	uidT, _ := uid.Int64()
	fmt.Println(uid, ok, uidT)

	client.UserId = uint64(uidT)
	// 用户登录
	login := &login{
		UserId: uint64(uidT),
		Client: client,
	}
	clientManager.Login <- login
	client.SendMsg(&WResponse{
		Event: Login,
		Data:  "success",
	})
}

func IsAppController(client *Client) {
	client.isApp = true
}

// JoinController 加入
func JoinController(client *Client, req *request) {
	name := gconv.String(req.Data["name"])

	if !client.tags.Contains(name) {
		client.tags.Append(name)
	}
	client.SendMsg(&WResponse{
		Event: Join,
		Data:  client.tags.Slice(),
	})
}

// QuitController 退出
func QuitController(client *Client, req *request) {
	name := gconv.String(req.Data["name"])
	if client.tags.Contains(name) {
		client.tags.RemoveValue(name)
	}
	client.SendMsg(&WResponse{
		Event: Quit,
		Data:  client.tags.Slice(),
	})
}
func PingController(client *Client) {
	currentTime := uint64(gtime.Now().Unix())
	client.Heartbeat(currentTime)
}
