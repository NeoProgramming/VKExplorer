package core

import (
	"fmt"
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

func getTasks(db *gorm.DB, page int, pageSize int) ([]Task, error) {
	var tasks []Task
	var err error
	if page > 0 {
		offset := (page - 1) * pageSize
		err = db.Offset(offset).Limit(pageSize).Find(&tasks).Error
	} else {
		err = db.Find(&tasks).Error
	}
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func getUsers(db *gorm.DB, page int, pageSize int) ([]User, error) {
	var users []User
	var err error
	if page > 0 {
		offset := (page - 1) * pageSize
		err = db.Offset(offset).Limit(pageSize).Find(&users).Error
	} else {
		err = db.Find(&users).Error
	}
	if err != nil {
		return nil, err
	}
	return users, nil
}

func getGroups(db *gorm.DB, page int, pageSize int) ([]Group, error) {
	var groups []Group
	var err error
	if page > 0 {
		offset := (page - 1) * pageSize
		err = db.Offset(offset).Limit(pageSize).Find(&groups).Error
	} else {
		err = db.Find(&groups).Error
	}
	if err != nil {
		return nil, err
	}
	return groups, nil
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

func getUserData(db *gorm.DB, uid int) (string, error) {
	return getUserName(db, uid), nil
}

func getGroupData(db *gorm.DB, gid int) (string, error) {
	return getGroupName(db, gid), nil
}

func getFriends(db *gorm.DB, uid int, page int, pageSize int) ([]User, error) {
	var friends []User
	var err error
	fmt.Println("getFriends: ", uid)
	if page > 0 {
		offset := (page - 1) * pageSize
		err = db.Raw("SELECT friends.uid2 AS uid, Name, Attrs, Age, Type FROM users JOIN friends ON users.uid = friends.uid2 WHERE friends.uid1 =  ? OFFSET ? LIMIT ?",
			uid, offset, pageSize).Scan(&friends).Error
	} else {
		err = db.Raw("SELECT friends.uid2 AS uid, Name, Attrs, Age, Type FROM users JOIN friends ON users.uid = friends.uid2 WHERE friends.uid1 =  ?",
			uid).Scan(&friends).Error
	}
	if err != nil {
		fmt.Println("getFriends error")
		return nil, err
	}
	return friends, nil
}

func getMemberships(db *gorm.DB, uid int, page int, pageSize int) ([]Group, error) {
	var groups []Group
	var err error
	fmt.Println("getMemberships: ", uid)
	if page > 0 {
		offset := (page - 1) * pageSize
		err = db.Raw("SELECT groups.gid AS gid, name, attrs, type FROM groups JOIN members ON groups.gid = members.gid WHERE members.uid = ? OFFSET ? LIMIT ?",
			uid, offset, pageSize).Scan(&groups).Error
	} else {
		err = db.Raw("SELECT groups.gid AS gid, name, attrs, type FROM groups JOIN members ON groups.gid = members.gid WHERE members.uid = ?",
			uid).Scan(&groups).Error
	}
	if err != nil {
		fmt.Println("getMemberships error")
		return nil, err
	}
	fmt.Println("groups len == ", len(groups))
	return groups, nil
}

func getMembers(db *gorm.DB, gid int, page int, pageSize int) ([]User, error) {
	var members []User
	var err error
	query := "SELECT members.uid AS uid, name FROM users JOIN members ON users.uid = members.uid WHERE members.gid = ?"
	if page > 0 {
		offset := (page - 1) * pageSize
		query += " OFFSET ? LIMIT ?"
		err = db.Raw(query, gid, offset, pageSize).Scan(&members).Error
	} else {
		err = db.Raw(query, gid).Scan(&members).Error
	}

	if err != nil {
		return nil, err
	}
	return members, nil
}

func getWall(db *gorm.DB, gid int, page int, pageSize int) ([]PostWithUsername, error) {
	var wall []PostWithUsername
	var err error
	query := "SELECT pid, oid, fid, date, text, name FROM posts LEFT OUTER JOIN users ON users.uid = posts.fid WHERE posts.oid = ?"
	if page > 0 {
		offset := (page - 1) * pageSize
		query += " OFFSET ? LIMIT ?"
		err = db.Raw(query, gid, offset, pageSize).Scan(&wall).Error
	} else {
		err = db.Raw(query, gid).Scan(&wall).Error
	}
	if err != nil {
		return nil, err
	}
	return wall, nil
}
