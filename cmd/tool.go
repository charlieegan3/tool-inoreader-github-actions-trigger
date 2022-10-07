package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	inoGHTool "github.com/charlieegan3/tool-inoreader-github-actions-trigger/pkg/tool"
	"github.com/charlieegan3/toolbelt/pkg/tool"
)

func main() {
	viper.SetConfigFile(os.Args[1])
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
		os.Exit(1)
	}

	tb := tool.NewBelt()

	toolCfg := viper.Get("tools").(map[string]interface{})
	tb.SetConfig(toolCfg)

	err = tb.AddTool(&inoGHTool.InoreaderGithubActions{})
	if err != nil {
		log.Fatalf("failed to add tool: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case <-c:
			cancel()
		}
	}()

	tb.RunServer(ctx, "0.0.0.0", "3000")
}
