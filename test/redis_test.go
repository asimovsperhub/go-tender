package test

import (
	"log"
	"testing"
	"time"
)

func Test_Re(*testing.T) {
	//var (
	//	ctx = gctx.New()
	//)
	//v, _ := g.Redis().Do(ctx, "SET", "k", "v")
	//fmt.Println(v.String())
	////var (
	////	config = gredis.Config{
	////		Address: "42.193.247.183:6379",
	////		Pass:    "asimov@77",
	////		Db:      2,
	////	}
	////	ctx = context.Background()
	////)
	////group := "default"
	////gredis.SetConfig(&config, group)
	////
	////redis := gredis.Instance(group)
	////log.Println(redis)
	////// defer redis.Close(ctx)
	////_, err := redis.Do(ctx, "SET", "k", "v")
	////if err != nil {
	////	panic(err)
	////}
	////
	////r, err := redis.Do(ctx, "GET", "k")
	////if err != nil {
	////	panic(err)
	////}
	////fmt.Println(gconv.String(r))
	// log.Println(fmt.Sprintf("%.3f", float64(1)/float64(100)))
	tm2, _ := time.Parse("2006-01-02 15:04:05", "2023-06-02 00:00:00")
	log.Println(tm2)
	log.Println(tm2.Unix(), time.Now().Unix())
}
