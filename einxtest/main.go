package main

import (
	"github.com/Cyinx/einx"
	"github.com/Cyinx/einx/slog"
	"github.com/golangstudy0/einxtest/clientmgr"
)

var logic = einx.GetModule("logic")

func main() {
	slog.SetLogPath("log/game_server/")
	logic.AddTcpServer(":2345", clientmgr.Instance)
	slog.LogInfo("game_server", "start server...")
	einx.Run()
	einx.Close()
}
