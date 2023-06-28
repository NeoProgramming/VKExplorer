package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

const VkCount = 50

func (app *Application) worker() {
	defer app.wg.Done()
	// loop until the execution flag is cleared
	fmt.Println("Worker running...")
	for app.running {

		// take the task from the queue (do not delete yet)
		var task Task
		if err := app.db.First(&task).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Handle "record not found" error
				fmt.Println("No tasks")
				app.running = false
				break
			} else {
				// Handle other errors
			}
		}

		// Update the "status" field of the first record
		// do not need to update anything yet, leave it for the future
		/*task.Status = 1
		if err := app.db.Model(&task).Update("Status", task.Status).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Handle "record not found" error
				continue
			} else {
				// Handle other errors
			}
		}*/

		// perform the task
		app.executeTask(&task)

		// task completed
		// Delete the first record from the "tasks" table
		if err := app.db.Delete(&task).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Handle "record not found" error
			} else {
				// Handle other errors
			}
		}

		time.Sleep(1 * time.Second)
	}
	app.running = false
	fmt.Println("Worker stopped naturally")
}

func (app *Application) executeTask(task *Task) {
	switch task.Type {
	case MyFriends:
		app.loadMyFriends(task)
	case MyGroups:
		app.loadMyGroups(task)
	case MyBookmarks:
		app.loadMyBookmarks(task)
	case GroupMembers:
		app.loadGroupMembers(task)
	case UserFriends:
		app.loadUserFriends(task)
	case UserGroups:
		app.loadUserGroups(task)
	case GroupWall:
		app.loadGroupWall(task)
	case UserWall:
		app.loadUserWall(task)
	case UserFriendsByName:
		app.loadUserFriendsByName(task)
	case UserGroupsByName:
		app.loadUserGroupsByName(task)
	case UserWallByName:
		app.loadUserWallByName(task)
	case GroupMembersByName:
		app.loadGroupMembersByName(task)
	case GroupWallByName:
		app.loadGroupWallByName(task)
	}
}

func (app *Application) GetStatus() string {
	app.counter++
	//fmt.Println("GetStatus ", app.counter)
	return strconv.Itoa(app.counter)
}
