package controller

import (
	"github.com/freshtoastman/imaptool/models"
	"github.com/freshtoastman/imaptool/tools"
)

func CheckKefuPass(username string, password string) (models.User, models.User_role, bool) {
	info := models.FindUser(username)
	var uRole models.User_role
	if info.Name == "" || info.Password != tools.Md5(password) {
		return info, uRole, false
	}
	uRole = models.FindRoleByUserId(info.ID)

	return info, uRole, true
}
