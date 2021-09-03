package main

import (
	"fmt"
	"github.com/brenordv/go-monitor-internet-connection/internal/clients"
	"github.com/brenordv/go-monitor-internet-connection/internal/core"
	"github.com/brenordv/go-monitor-internet-connection/internal/handlers"
	"github.com/brenordv/go-monitor-internet-connection/internal/utils"
	"log"
	"time"
)

func main() {
	var runtimeCfg *core.RuntimeConfig
	var err error
	var urls []string
	var noConnInfo *core.NoConnectionInfo
	count := 0
	runtimeCfg, err = utils.LoadRuntimeConfig()
	urls, err = utils.LoadWatchlistUrls()
	handlers.PanicOnError(err)

	err = clients.Notify("Starting to monitor internet connection.", runtimeCfg.TelegramBotToken, runtimeCfg.TelegramChatId)
	handlers.PanicOnError(err)

	for {
		for _, url := range urls {
			count++
			log.Printf("[Request #%010d] Using '%s' to check for internet connection...\n", count, url)


			req := clients.MakeClient(url, runtimeCfg.Headers)
			isConnected, err := req.CheckConnection()

			if !isConnected && noConnInfo == nil {

				noConnInfo = &core.NoConnectionInfo{
					Start: time.Now(),
					Url:   url,
					Error: err,
				}
				log.Printf("Lost internet connection at '%v'. Url: %s\n", noConnInfo.Start, noConnInfo.Url)

			} else if isConnected && noConnInfo != nil {
				noConnInfo.End = time.Now()
				noConnInfo.DownTime = time.Since(noConnInfo.Start)

				msg := fmt.Sprintf("Internet connection restored! Time without internet: %s", noConnInfo.DownTime)
				log.Printf("%s\n",msg)

				key, content, sErr := noConnInfo.Serialize()
				handlers.PanicOnError(sErr)

				err = clients.Persist(key, content, runtimeCfg.NoConnInfoTTLinHours)
				handlers.PanicOnError(err)

				err = clients.Notify(msg, runtimeCfg.TelegramBotToken, runtimeCfg.TelegramChatId)
				handlers.PanicOnError(err)

				noConnInfo = nil
			}

			time.Sleep(time.Duration(runtimeCfg.DelayInSeconds) * time.Second)
		}
	}
}
