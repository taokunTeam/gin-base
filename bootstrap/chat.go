package bootstrap

import "github.com/taokunTeam/gin-base/utils"

func InitializeChat() (*utils.Hub, *utils.Single) {
	hub := utils.NewHub()
	single := utils.NewSingle()
	go hub.Run()
	go single.RunSingle()
	return hub, single
}
