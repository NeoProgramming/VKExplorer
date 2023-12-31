package core

import (
	"database/sql"
	"fmt"
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
		err := app.db.Get(&task, "SELECT * FROM tasks")
		if err != nil {
			if err == sql.ErrNoRows {
				// Handle "record not found" error
				fmt.Println("No tasks")
				app.running = false
				break
			} else {
				// Handle other errors
				fmt.Println("Task extracting error: ", err)
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
		fmt.Println("Executing task: ", task.Name)
		app.executeTask(&task)

		// task completed; delete the record from the "tasks" table
		fmt.Println("Deleting task: ", task.Name)
		_, err = app.db.Exec("DELETE FROM tasks WHERE id = ?", task.Id)
		if err != nil {
			fmt.Println("worker delete task error: ", err)
		} else {
			app.completeCounter++
		}

		time.Sleep(1 * time.Second)
	}
	app.running = false
	fmt.Println("Worker stopped naturally")
}

func (app *Application) executeTask(task *Task) {
	switch task.TType {
	case TT_MyFriends:
		app.loadMyFriends(task)
	case TT_MyGroups:
		app.loadMyGroups(task)
	case TT_MyBookmarks:
		app.loadMyBookmarks(task)
	case TT_GroupMembers:
		app.loadGroupMembers(task)
	case TT_UserFriends:
		app.loadUserFriends(task)
	case TT_UserGroups:
		app.loadUserGroups(task)
	case TT_GroupWall:
		app.loadGroupWall(task)
	case TT_UserWall:
		app.loadUserWall(task)

	case TT_MyDataByName:
		app.loadMyData(task)
	case TT_UserDataByName:
		app.loadUserDataByName(task)
	case TT_UserFriendsByName:
		app.loadUserFriendsByName(task)
	case TT_UserGroupsByName:
		app.loadUserGroupsByName(task)
	case TT_UserWallByName:
		app.loadUserWallByName(task)
	case TT_GroupDataByName:
		app.loadGroupDataByName(task)
	case TT_GroupMembersByName:
		app.loadGroupMembersByName(task)
	case TT_GroupWallByName:
		app.loadGroupWallByName(task)
	}
}

func (app *Application) GetStatus() string {
	app.counter++
	return fmt.Sprintf("C=%d I=%d D=%d", app.counter, app.taskCounter, app.completeCounter)
}
