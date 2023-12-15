package core

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

func getTableNames(db *sqlx.DB) ([]string, error) {
	// get a list of tables as an array
	var tableInfos []TableInfo
	err := db.Select(&tableInfos, "SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		fmt.Println("getTableNames error: ", err)
		return nil, err
	}

	tableNames := make([]string, len(tableInfos))
	for i, ti := range tableInfos {
		tableNames[i] = ti.Name
	}

	return tableNames, nil
}

func GetTables() string {
	// get a list of tables as a string separated by commas
	arr, err := getTableNames(App.db)
	if err != nil {
		return "<ERROR>"
	}
	return strings.Join(arr, ", ")
}

func getUsersCount(db *sqlx.DB) int {
	var count int
	err := db.Get(&count, "SELECT count(*) FROM users")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return count
}

func getGroupsCount(db *sqlx.DB) int {
	var count int
	err := db.Get(&count, "SELECT count(*) FROM groups")
	if err != nil {
		return 0
	}
	return count
}

func getPostsCount(db *sqlx.DB) int {
	var count int
	err := db.Get(&count, "SELECT count(*) FROM posts")
	if err != nil {
		return 0
	}
	return count
}

func getTasksCount(db *sqlx.DB) int {
	var count int
	err := db.Get(&count, "SELECT count(*) FROM tasks")
	if err != nil {
		return 0
	}
	return count
}

func getUserName(db *sqlx.DB, uid int) string {
	var user User
	err := db.Get(&user, "uid=?", uid)
	if err != nil {
		// Handle error
		return "! " + strconv.Itoa(uid)
	}
	return user.Name
}

func getGroupName(db *sqlx.DB, gid int) string {
	var group Group
	err := db.Get(&group, "gid=?", gid)
	if err != nil {
		// Handle error
		return "! " + strconv.Itoa(gid)
	}
	return group.Name
}

func getUserInfo(db *sqlx.DB, uid int) (User, error) {
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE uid=?", uid)
	return user, err
}

func getGroupInfo(db *sqlx.DB, gid int) (Group, error) {
	var group Group
	err := db.Get(&group, "SELECT * FROM groups WHERE gid=?", gid)
	return group, err
}


