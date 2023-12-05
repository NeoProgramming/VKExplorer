package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type TaskType int

// task types
const (
	TT_MyFriends TaskType = iota + 1
	TT_MyGroups
	TT_MyBookmarks
	TT_GroupMembers
	TT_UserFriends
	TT_UserGroups
	TT_GroupWall
	TT_UserWall

	TT_UserDataByName
	TT_UserFriendsByName
	TT_UserGroupsByName
	TT_UserWallByName
	TT_GroupDataByName
	TT_GroupMembersByName
	TT_GroupWallByName
)

type Task struct {
	gorm.Model
	Type   TaskType
	Name   string // visible task name; also user/group short_name
	Xid    int    // object id; also future tasks mask
	Offset int
	Status int
}

type User struct {
	Uid            int `gorm:"unique;index"`
	Name           string
	About          string
	City           string
	Domain         string
	Photo          string
	Attrs          int
	Type           int
	FriendsUpdated int64
	GroupsUpdated  int64
	WallUpdated    int64
	Oldest         int64
	Newest         int64
}

type Group struct {
	Gid            int `gorm:"uniqueIndex:groups_gid_uindex"`
	Name           string
	Attrs          int
	Type           int
	MembersUpdated int64
	WallUpdated    int64
	Oldest         int64
	Newest         int64
}

type Bookmark struct {
	Bid  int `gorm:"uniqueIndex"`
	Type int
}

// user-user link
type Friend struct {
	ID   uint `gorm:"primary_key"`
	Uid1 int
	Uid2 int
}

// user-group link
type Member struct {
	ID  uint `gorm:"primary_key"`
	Uid int
	Gid int
}

// Post on wall
type Post struct {
	ID         uint `gorm:"primary_key"`
	Pid        int  // local post id
	Oid        int  // wall owner (group id)
	Fid        int  // commenter (user id, "from")
	Date       int64
	Text       string // comment text
	CmntsCount int
	LikesCount int
	ReposCount int
	ViewsCount int
}

type PostWithUsername struct {
	Post
	Name string // username
}

// like to object
type Reaction struct {
	ID  uint `gorm:"primary_key"`
	Typ int  // object type (post, comment, image, video...)
	Oid int  // object owner id (user, group)
	Iid int  // liked item id (post, comment, image, video...)
	Uid int  // liker id (usually user, maybe group)
}

type Comment struct {
	ID   uint `gorm:"primary_key"`
	Oid  int  // object owner id
	Cid  int  // commenter id
	Text string
}

func InitDatabase() {
	db, err := gorm.Open("sqlite3", "./vk.db") //rename to vk.db
	if err != nil {
		App.dbaseConnected = false
		panic("failed to connect database")
	}
	App.dbaseConnected = true
	App.db = db

	App.db.AutoMigrate(&User{})
	App.db.AutoMigrate(&Task{})
	App.db.AutoMigrate(&Group{})
	App.db.AutoMigrate(&Bookmark{})
	App.db.AutoMigrate(&Friend{})
	App.db.AutoMigrate(&Member{})
	App.db.AutoMigrate(&Post{})

	App.db.Exec("PRAGMA journal_mode = WAL")
	App.db.Exec("PRAGMA synchronous = normal")
	App.db.Exec("PRAGMA temp_store = memory")
	App.db.Exec("PRAGMA mmap_size = 30000000000")

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
