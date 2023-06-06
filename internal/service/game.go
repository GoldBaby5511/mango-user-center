package service

import (
	"mango-user-center/internal/model"
	"mango-user-center/pkg/util"
)

var (
	game = Game{}
)

type Game struct{}

// 分配游戏ID
func (g Game) DistributeGameId(uid int) int {
	var row model.GameIdNormal
	model.DB().Where("sys_id = ?", uid).Take(&row)
	model.DB().Model(&row).UpdateColumn("user_id", uid)

	go func() {
		var maxId int
		model.DB().Model(model.GameIdNormal{}).Pluck("MAX(sys_id)", &maxId)
		if maxId-uid < 100 {
			util.Robot.SendText("游戏ID要不够啦")
		}
	}()

	return row.GameId
}
