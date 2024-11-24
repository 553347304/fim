package flags

import (
	"fim_server/fim_chat/chat_models"
	"fim_server/fim_group/group_models"
	"fim_server/fim_user/user_models"
	"fim_server/global"
	"fmt"
)

func MigrationTable() {
	var err error
	// global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})

	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&user_models.User{},
			&user_models.Friend{},
			&user_models.FriendAuth{},
			&user_models.UserConfig{},

			&chat_models.Chat{},
			&group_models.Group{},
			&group_models.GroupMember{},
			&group_models.GroupMessage{},
			&group_models.GroupAuth{},
		)
	if err != nil {
		fmt.Println("[生成数据库表结构失败]")
		return
	}
	fmt.Println("[生成数据库表结构成功]")
}
