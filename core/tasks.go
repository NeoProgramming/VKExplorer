package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


func (app *Application) QueueByType(tt TaskType, cmt string) {
	var task Task
	err := app.db.First(&task, "Type = ?", tt).Error
	if err == gorm.ErrRecordNotFound {
		// create a new record
		task = Task{Type: tt, Name: cmt, Offset: 0}
		app.db.Create(&task)
	}
}

func (app *Application) QueueById(tt TaskType, id int, cmt string) {
	var task Task
	err := app.db.First(&task, "Type = ? AND Xid = ?", tt, id).Error
	if err == gorm.ErrRecordNotFound {
		task = Task{Type: tt, Name: cmt + getGroupName(app.db, id), Xid: id, Offset: 0}
		app.db.Create(&task)
	}
}

func (app *Application) QueueByName(tt TaskType, nm string) {
	var task Task
	err := app.db.First(&task, "Type = ? AND Name = ?", tt, nm).Error
	if err == gorm.ErrRecordNotFound {
		task = Task{Type: tt, Name: nm, Offset: 0}
		app.db.Create(&task)
	}
	// else if found: task already in queue
}
