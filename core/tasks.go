package core

import (
	"fmt"
)

func (app *Application) QueueByType(tt TaskType, cmt string) {
	task := Task{Type: tt, Name: cmt, Offs: 0}
	query := `INSERT INTO tasks (type, name, offs) VALUES (:Type, :Name, :Offs) ON CONFLICT(type) DO NOTHING`
	_, err := app.db.NamedExec(query, task)	
	if err != nil {
		fmt.Println("QueueByType error", err)
	} else {
		app.taskCounter++
	}
}

func (app *Application) QueueById(tt TaskType, id int, cmt string) {
	task := Task{Type: tt, Name: cmt + getGroupName(app.db, id), Xid: id, Offs: 0}
	query := `INSERT INTO tasks (type, name, xid, offs) VALUES (:Type, :Name, :Xid, :Offs) ON CONFLICT(type) DO NOTHING`
	_, err := app.db.NamedExec(query, task)	
	if err != nil {
		fmt.Println("QueueById error", err)
	} else {
		app.taskCounter++
	}
}

func (app *Application) QueueByName(tt TaskType, nm string) {
	task := Task{Type: tt, Name: nm, Offs: 0}
	query := `INSERT INTO tasks (type, name) VALUES (:Type, :Name) ON CONFLICT(type, name) DO NOTHING`
	_, err := app.db.NamedExec(query, task)	
	if err != nil {
		fmt.Println("QueueById error", err)
	} else {
		app.taskCounter++
	}
}
