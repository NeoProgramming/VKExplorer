package core

import (
	"fmt"
)

func (app *Application) QueueByType(tt int, cmt string) {
	var c int
	app.db.Get(&c, `SELECT COUNT(*) FROM tasks WHERE type = ?`, tt)
	if c == 0 {
		query := `INSERT INTO tasks (type, name, offs) VALUES (:type, :name, :offs)`
		task := Task{TType: tt, Name: cmt, Offs: 0}
		_, err := app.db.NamedExec(query, task)	
		if err != nil {
			fmt.Println("QueueByType error: ", err)
		} else {
			app.taskCounter++
		}
	}
}

func (app *Application) QueueById(tt int, id int, cmt string) {
	var c int
	app.db.Get(&c, `SELECT COUNT(*) FROM tasks WHERE type = ? AND id = ?`, tt, id)
	if c == 0 {
		task := Task{TType: tt, Name: cmt + getGroupName(app.db, id), Xid: id, Offs: 0}
		query := `INSERT INTO tasks (name, type, xid, offs) VALUES (:name, :type, :xid, :offs)`
		_, err := app.db.NamedExec(query, task)	
		if err != nil {
			fmt.Println("QueueById error: ", err)
		} else {
			app.taskCounter++
		}
	}
}

func (app *Application) QueueByName(tt int, nm string) {
	var c int
	app.db.Get(&c, `SELECT COUNT(*) FROM tasks WHERE type = ? AND name = ?`, tt, nm)
	if c == 0 {
		task := Task{TType: tt, Name: nm, Offs: 0}
		query := `INSERT INTO tasks (type, name) VALUES (:type, :name)`
		_, err := app.db.NamedExec(query, task)	
		if err != nil {
			fmt.Println("QueueByName error: ", err)
		} else {
			app.taskCounter++
		}
	}
}
