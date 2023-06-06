package model

// 游戏ID池(普通)
type GameIdNormal struct {
	SysId  int `json:"sys_id" gorm:"primaryKey; type:int unsigned auto_increment; <-:false;"`
	GameId int `json:"game_id" gorm:"uniqueIndex:uk_gameid; type:int unsigned; not null; default 0;"`
	UserId int `json:"user_id" gorm:"type:int unsigned; not null; default 0;"`
}

// 游戏ID池(靓号)
type GameIdExcellent struct {
	SysId  int `json:"sys_id" gorm:"primaryKey; type:int unsigned auto_increment; <-:false;"`
	GameId int `json:"game_id" gorm:"uniqueIndex:uk_gameid; type:int unsigned; not null; default 0;"`
	UserId int `json:"user_id" gorm:"type:int unsigned; not null; default 0;"`
}
