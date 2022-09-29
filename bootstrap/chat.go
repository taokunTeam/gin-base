package bootstrap

import "github.com/taokunTeam/gin-base/utils"

func InitializeChat() *utils.Hub {
	hub := utils.NewHub()
	go hub.Run()
	return hub
}
