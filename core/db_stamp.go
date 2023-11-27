package core

import (
	"fmt"
	"time"
)

func (app *Application) StampUserFirends(uid int) {
	user := User{Uid: uid, FriendsUpdated: time.Now()}
	app.db.Update(&user)
	app.db.Save(&user)
	//	app.db.Model(&User{}).Where("id = ?", uid).Update("FriendsUpdated", time.Now())
	fmt.Println("StampUserFirends ", user)
}

func (app *Application) StampUserGroups(uid int) {
	app.db.Model(&User{}).Where("id = ?", uid).Update("GroupsUpdated", time.Now())
	fmt.Println("StampUserGroups")
}

func (app *Application) StampUserWall(uid int) {
	app.db.Model(&User{}).Where("id = ?", uid).Update("WallUpdated", time.Now())
	fmt.Println("StampUserWall")
}

func (app *Application) StampGroupMembers(gid int) {
	app.db.Model(&Group{}).Where("id = ?", gid).Update("MembersUpdated", time.Now())
	fmt.Println("StampGroupMembers")
}

func (app *Application) StampGroupWall(gid int) {
	app.db.Model(&Group{}).Where("id = ?", gid).Update("WallUpdated", time.Now())
	fmt.Println("StampGroupWall")
}
