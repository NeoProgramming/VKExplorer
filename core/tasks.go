package core

import (
	"strconv"
)

func (app *Application) GetStatus() string {
	app.counter++
	//fmt.Println("GetStatus ", app.counter)
	return strconv.Itoa(app.counter)
}
