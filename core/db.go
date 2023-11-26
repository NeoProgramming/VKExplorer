package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
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
	gorm.Model
	Uid           int
	Name          string
	Attrs         int
	Age           int
	Type          int
	FiendsUpdated time.Time
	GroupsUpdated time.Time
	WallUpdated   time.Time
}

type Group struct {
	gorm.Model
	Gid            int
	Name           string
	Attrs          int
	Type           int
	MembersUpdated time.Time
	WallUpdated    time.Time
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

// Post on wall
type Post struct {
	gorm.Model
	Pid        int // local post id
	Oid        int // wall owner (group id)
	Fid        int // commenter (user id, "from")
	Date       int
	Text       string // comment text
	CmntsCount int
	LikesCount int
	ReposCount int
	ViewsCount int
}

type PostWithUsername struct {
	Post
	Name string
}

// like to object
type Reaction struct {
	gorm.Model
	Typ int // object type (post, comment, image, video...)
	Oid int // object owner id (user, group)
	Iid int // liked item id (post, comment, image, video...)
	Uid int // liker id (usually user, maybe group)
}

type Comment struct {
	gorm.Model
	Oid  int // object owner id
	Cid  int // commenter id
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

	App.db.AutoMigrate(&User{}, &Task{}, &Group{}, &Bookmark{}, &Friend{}, &Member{}, &Post{})

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
