package core

import (
	"fmt"
	"time"
)

func (app *Application) StampUserFriends(uid int) {
	app.db.Exec("UPDATE Users SET friends_updated = ? WHERE uid = ? ", time.Now().Unix(), uid)
	fmt.Println("StampUserFirends ")
}

func (app *Application) StampUserGroups(uid int) {
	app.db.Exec("UPDATE Users SET groups_updated = ? WHERE uid = ? ", time.Now().Unix(), uid)
	fmt.Println("StampUserGroups ")
}

func (app *Application) StampUserWall(uid int) {
	app.db.Exec("UPDATE Users SET wall_updated = ? WHERE uid = ? ", time.Now().Unix(), uid)
	fmt.Println("StampUserWall ")
}

func (app *Application) StampGroupMembers(gid int) {
	app.db.Exec("UPDATE Groups SET members_updated = ? WHERE uid = ? ", time.Now().Unix(), gid)
	fmt.Println("StampGroupMembers ")
}

func (app *Application) StampGroupWall(gid int) {
	app.db.Exec("UPDATE Groups SET wall_updated = ? WHERE uid = ? ", time.Now().Unix(), gid)
	fmt.Println("StampGroupWall ")
}
