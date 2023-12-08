package core

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//type TaskType int

// task types
const (
	TT_MyFriends  = iota + 1
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
	Id     int      `db:"id"`
	TType  int 		`db:"type"`
	Name   string   `db:"name"`
	Xid    int      `db:"xid"`
	Offs   int      `db:"offs"`
	Status int      `db:"status"`
}
const SQLITE_SCHEMA_Tasks string = 
`CREATE TABLE IF NOT EXISTS "tasks" (
	"id"	integer PRIMARY KEY AUTOINCREMENT,
	"type"	integer,
	"name"	varchar(255),
	"xid"	integer,
	"offs"	integer,
	"status"	integer
)`

type User struct {
	Uid            int		`db:"uid"`
	Name           string	`db:"name"`
	About          string	`db:"about"`
	City           string	`db:"city"`
	Domain         string	`db:"domain"`
	Photo          string	`db:"photo"`
	Attrs          int		`db:"attrs"`
	Type           int		`db:"type"`
	FriendsUpdated int64	`db:"friends_updated"`
	GroupsUpdated  int64	`db:"groups_updated"`
	WallUpdated    int64	`db:"wall_updated"`
	Oldest         int64	`db:"oldest"`
	Newest         int64	`db:"newest"`
}
const SQLITE_SCHEMA_Users string = 
`CREATE TABLE IF NOT EXISTS "users" (
	"uid"	integer,
	"name"	varchar(255),
	"about"	varchar(255),
	"city"	varchar(255),
	"domain"	varchar(255),
	"photo"	varchar(255),
	"attrs"	integer,
	"type"	integer,
	"friends_updated"	bigint,
	"groups_updated"	bigint,
	"wall_updated"	bigint,
	PRIMARY KEY("uid")
)`

type Group struct {
	Gid            int		`db:"gid"`
	Name           string	`db:"name"`
	Attrs          int		`db:"attrs"`
	Type           int		`db:"type"`
	MembersUpdated int64	`db:"members_updated"`
	WallUpdated    int64	`db:"wall_updated"`
	Oldest         int64	`db:"oldest"`
	Newest         int64	`db:"newest"`
}
const SQLITE_SCHEMA_Groups string = 
`CREATE TABLE IF NOT EXISTS "groups" (
	"gid"	integer,
	"name"	varchar(255),
	"attrs"	integer,
	"type"	integer,
	"members_updated"	bigint,
	"wall_updated"	bigint,
	"oldest"	bigint,
	"newest"	bigint,
	PRIMARY KEY("gid")
)`

type Bookmark struct {
	Bid  int	`db:"bid"`
	Type int	`db:"type"`
}
const SQLITE_SCHEMA_Bookmarks string = 
`CREATE TABLE IF NOT EXISTS "bookmarks" (
	"bid"	integer,
	"type"	integer
)`

// user-user link
type Friend struct {
	ID   uint	`db:"id"`
	Uid1 int	`db:"uid1"`
	Uid2 int	`db:"uid2"`
}
const SQLITE_SCHEMA_Friends string = 
`CREATE TABLE IF NOT EXISTS "friends" (
	"id"	integer PRIMARY KEY AUTOINCREMENT,
	"uid1"	integer,
	"uid2"	integer
)`

// user-group link
type Member struct {
	ID  uint	`db:"id"`
	Uid int		`db:"uid"`
	Gid int		`db:"gid"`
}
const SQLITE_SCHEMA_Members string = 
`CREATE TABLE IF NOT EXISTS "members" (
	"id"	integer PRIMARY KEY AUTOINCREMENT,
	"uid"	integer,
	"gid"	integer
)`

// Post on wall
type Post struct {
	ID         uint		`db:"id"`
	Pid        int  	`db:"pid"`	// local post id
	Oid        int  	`db:"oid"`	// wall owner (group id)
	Fid        int  	`db:"fid"`	// commenter (user id, "from")
	Date       int64	`db:"date"`
	Text       string 	`db:"text"`	// comment text
	CmntsCount int		`db:"cmnts_count"`
	LikesCount int		`db:"likes_count"`
	ReposCount int		`db:"repos_count"`
	ViewsCount int		`db:"views_count"`
}
const SQLITE_SCHEMA_Posts string = 
`CREATE TABLE IF NOT EXISTS "posts" (
	"id"	integer PRIMARY KEY AUTOINCREMENT,
	"pid"	integer,
	"oid"	integer,
	"fid"	integer,
	"date"	bigint,
	"text"	varchar(255),
	"cmnts_count"	integer,
	"likes_count"	integer,
	"repos_count"	integer,
	"views_count"	integer
)`

// indexes
const SQLITE_SCHEMA_FriendsIndex string = 
`CREATE UNIQUE INDEX IF NOT EXISTS idx_friends_uid1_uid2 ON Friends (uid1, uid2);`
const SQLITE_SCHEMA_MembersIndex string = 
`CREATE UNIQUE INDEX IF NOT EXISTS idx_members_uid_gid ON Members (uid, gid);`
const SQLITE_SCHEMA_PostsIndex string = 
`CREATE UNIQUE INDEX IF NOT EXISTS idx_members_pid_oid_fid ON Posts (pid, oid, fid);`

// join(posts, users)
type PostWithUsername struct {
	Post
	Name string `db:"name"`
}

// like to object
type Reaction struct {
	ID  uint 
	Type int  // object type (post, comment, image, video...)
	Oid int  // object owner id (user, group)
	Iid int  // liked item id (post, comment, image, video...)
	Uid int  // liker id (usually user, maybe group)
}

type Comment struct {
	ID   uint 
	Oid  int  // object owner id
	Cid  int  // commenter id
	Text string
}

func InitDatabase() {
	db, err := sqlx.Open("sqlite3", "./vk.db") //rename to vk.db
	if err != nil {
		App.dbaseConnected = false
		panic("failed to connect database")
	}
	App.dbaseConnected = true
	App.db = db

	App.db.Exec(SQLITE_SCHEMA_FriendsIndex);
	App.db.Exec(SQLITE_SCHEMA_MembersIndex);
	App.db.Exec(SQLITE_SCHEMA_PostsIndex);

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
