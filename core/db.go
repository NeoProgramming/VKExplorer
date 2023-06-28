package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strconv"
	"strings"
)

type TaskType int

// task types
// add 
const (
	MyFriends TaskType = iota + 1
	MyGroups
	MyBookmarks
	GroupMembers
	UserFriends
	UserGroups
	GroupWall
	UserWall
)

type Task struct {
	gorm.Model
	Type   TaskType
	Name   string
	Xid    int
	Offset int
	Status int
}

type User struct {
	gorm.Model
	Uid   int
	Name  string
	Attrs int
	Age   int
	Type  int
}

type Group struct {
	gorm.Model
	Gid   int
	Name  string
	Attrs int
	Type  int
}

type Bookmark struct {
	gorm.Model
	Bid int
}

// user-user link
type Friend struct {
	gorm.Model
	Uid1 int
	Uid2 int
}

// user-group link
type Member struct {
	gorm.Model
	Uid int
	Gid int
}

// post on wall
type Post struct {
	gorm.Model
	Pid int
	Oid int		// wall owner
	Fid int		// commenter 
	Date int
	Text string
	Rpid int	// reply pid
}

func InitDatabase() {
	db, err := gorm.Open("sqlite3", "./vk.db") //rename to vk.db
	if err != nil {
		App.dbaseConnected = false
		panic("failed to connect database")
	}
	App.dbaseConnected = true
	App.db = db

	App.db.AutoMigrate(&User{}, &Task{}, &Group{}, &Bookmark{}, &Friend{}, &Member{}, &Post{})
	fmt.Println("Database vk.db opened")
}

func quitDatabase() {
	if App.db != nil {
		App.db.Close()
		fmt.Println("Database vk.db closed")
	}
}

type TableInfo struct {
	Name string `gorm:"column:name"`
}

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

func getUserName(db *gorm.DB, uid int) string {
	var user User
	result := db.First(&user, "uid=?", uid)
	if result.Error != nil {
		// Handle error
		return "!NOTFOUND: " + strconv.Itoa(uid)
	}
	return user.Name
}

func getGroupName(db *gorm.DB, gid int) string {
	var group Group
	result := db.First(&group, "gid=?", gid)
	if result.Error != nil {
		// Handle error
		return "!NOTFOUND: " + strconv.Itoa(gid)
	}
	return group.Name
}


func (app *Application) UpsertUser(uid int, name string, attrs int) {
	// update or add User record
	user := User{Uid: uid}
	// find by ID
	app.db.FirstOrCreate(&user, User{Uid: uid})
	// change name ; fill fields
	if !(name == "DELETED" && user.Name != "DELETED") {
		user.Name = name
	}
	user.Attrs |= attrs
	app.db.Save(&user)
}

func (app *Application) UpsertGroup(gid int, name string, attrs int) {
	// update or add Group record
	group := Group{Gid: gid}
	// find by ID
	app.db.FirstOrCreate(&group, Group{Gid: gid})
	// change name ; fill fields
	group.Name = name
	group.Attrs |= attrs
	app.db.Save(&group)
}

func (app *Application) UpsertMembership(uid int, gid int) {
	// update or add Member record
	member := Member{Uid: uid, Gid: gid}
	// find by pair of ID's
	app.db.FirstOrCreate(&member, Member{Uid: uid, Gid: gid})
	// modify
	// ...
	// save
	app.db.Save(&member)
}

func (app *Application) UpsertFriendship(uid1 int, uid2 int) {
	// update or add Friend record
	friend := Friend{Uid1: uid1, Uid2: uid2}
	// find by pair of ID's
	app.db.FirstOrCreate(&friend, Friend{Uid1: uid1, Uid2: uid2})
	// modify
	// ...
	// save
	app.db.Save(&friend)
}