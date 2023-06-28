package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


func (app *Application) QueueMyGroups() {
	// add task "my groups"
	var task Task
	err := app.db.First(&task, "Type = ?", MyGroups).Error
	if err != nil {
		// handle error / not found?
		// create a new record
		task = Task{Type: MyGroups, Name: "MyGroups", Offset: 0}
		app.db.Create(&task)
	} else {
		// found
		// task already in queue
	}
}

func (app *Application) QueueMyFriends() {
	// add task "my friends"
	var task Task
	err := app.db.First(&task, "Type = ?", MyFriends).Error
	if err != nil {
		// handle error / not found?
		if err == gorm.ErrRecordNotFound {
			// create a new record
			task = Task{Type: MyFriends, Name: "MyFriends", Offset: 0}
			app.db.Create(&task)
		}
	} else {
		// found
		// task already in queue
	}
}

func (app *Application) QueueMyBookmarks() {

}

func (app *Application) QueueGroupMembers(gid int) {
	// add task "group members"
	var task Task
	err := app.db.First(&task, "Type = ? AND Xid = ?", GroupMembers, gid).Error
	if err != nil {
		task = Task{Type: GroupMembers, Name: "GroupMembers " + getGroupName(app.db, gid), Xid: gid, Offset: 0}
		app.db.Create(&task)
	}
}

func (app *Application) QueueGroupWall(gid int) {
	// add task "group wall"
	var task Task
	err := app.db.First(&task, "Type = ? AND Xid = ?", GroupWall, gid).Error
	if err != nil {
		task = Task{Type: GroupWall, Name: "GroupWall " + getGroupName(app.db, gid), Xid: gid, Offset: 0}
		app.db.Create(&task)
	}
}

func (app *Application) QueueUserFriends(uid int) {
	// add task "user friends"
	var task Task
	err := app.db.First(&task, "Type = ? AND Xid = ?", UserFriends, uid).Error
	if err != nil {
		task = Task{Type: UserFriends, Name: "UserFriends " + getUserName(app.db, uid), Xid: uid, Offset: 0}
		app.db.Create(&task)
	}
}

func (app *Application) QueueUserGroups(uid int) {
	// add task "user groups"
	var task Task
	err := app.db.First(&task, "Type = ? AND Xid = ?", UserGroups, uid).Error
	if err != nil {
		task = Task{Type: UserGroups, Name: "UserGroups " + getUserName(app.db, uid), Xid: uid, Offset: 0}
		app.db.Create(&task)
	}
}

func (app *Application) QueueUserWall(uid int) {
	// add task "user wall"
	var task Task
	err := app.db.First(&task, "Type = ? AND Xid = ?", UserWall, uid).Error
	if err != nil {
		task = Task{Type: UserWall, Name: "UserWall " + getUserName(app.db, uid), Xid: uid, Offset: 0}
		app.db.Create(&task)
	}
}
