package core

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

func getTableNames(db *gorm.DB) ([]string, error) {
	// get a list of tables as an array
	var tableInfos []TableInfo
	//db = db.Debug()
	result := db.Raw("SELECT name FROM sqlite_master WHERE type='table';").Scan(&tableInfos)
	//fmt.Println(tableInfos)
	if result.Error != nil {
		return nil, result.Error
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

func getUsersCount(db *gorm.DB) int {
	var count int64
	result := db.Model(&User{}).Count(&count)
	if result.Error != nil {
		return 0
	}
	return int(count)
}

func getGroupsCount(db *gorm.DB) int {
	var count int64
	result := db.Model(&Group{}).Count(&count)
	if result.Error != nil {
		return 0
	}
	return int(count)
}

func getPostsCount(db *gorm.DB) int {
	var count int64
	result := db.Model(&Post{}).Count(&count)
	if result.Error != nil {
		return 0
	}
	return int(count)
}

func getTasksCount(db *gorm.DB) int {
	var count int64
	result := db.Model(&Task{}).Count(&count)
	if result.Error != nil {
		return 0
	}
	return int(count)
}

func getUserName(db *gorm.DB, uid int) string {
	var user User
	result := db.First(&user, "uid=?", uid)
	if result.Error != nil {
		// Handle error
		return "! " + strconv.Itoa(uid)
	}
	return user.Name
}

func getGroupName(db *gorm.DB, gid int) string {
	var group Group
	result := db.First(&group, "gid=?", gid)
	if result.Error != nil {
		// Handle error
		return "! " + strconv.Itoa(gid)
	}
	return group.Name
}

func getUserInfo(db *gorm.DB, uid int) (User, error) {
	var user User
	result := db.First(&user, "uid=?", uid)
	return user, result.Error
}

func getGroupInfo(db *gorm.DB, gid int) (Group, error) {
	var group Group
	result := db.First(&group, "gid=?", gid)
	return group, result.Error
}


