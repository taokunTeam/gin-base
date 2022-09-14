package bootstrap

import (
	"fmt"

	"github.com/taokunTeam/gin-base/global"

	"github.com/taokunTeam/go-storage/kodo"
	"github.com/taokunTeam/go-storage/local"
	"github.com/taokunTeam/go-storage/minio"
	"github.com/taokunTeam/go-storage/oss"
)

func InitializeStorage() {
	_, _ = local.Init(global.App.Config.Storage.Disks.Local)
	_, _ = kodo.Init(global.App.Config.Storage.Disks.QiNiu)
	_, _ = oss.Init(global.App.Config.Storage.Disks.AliOss)
	_, err := minio.Init(global.App.Config.Storage.Disks.Minio)
	fmt.Println("222222222")
	fmt.Println(err)
}
