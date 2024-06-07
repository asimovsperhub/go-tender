package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "tender/internal/app/common/logic"
	_ "tender/internal/app/desk/logic"
	_ "tender/internal/app/system/logic"
	_ "tender/internal/packed"

	"tender/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
