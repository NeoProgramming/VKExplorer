package core

import (
	"fmt"
	"strings"
)

func (app *Application) UpsertUser(uid int, name string, about string, city string, domain string, photo string, attrs int) {
	//	fmt.Println("UpsertUser")
	var query string
	if strings.TrimSpace(name) == "DELETED" {
		query = `INSERT INTO users (uid, name, about, city, domain, photo, attrs) 
	VALUES (:uid, :name, :about, :city, :domain, :photo, :attrs) ON CONFLICT(uid) DO UPDATE 
	SET attrs=attrs|:attrs;`
	} else {
		query = `INSERT INTO users (uid, name, about, city, domain, photo, attrs) 
	VALUES (:uid, :name, :about, :city, :domain, :photo, :attrs) ON CONFLICT(uid) DO UPDATE 
	SET name=:name, about=:about, city=:city, domain=:domain, photo=:photo, attrs=attrs|:attrs;`
	}
	user := User{Uid: uid, About: about, City: city, Domain: domain, Photo: photo}
	_, err := app.db.NamedExec(query, user)
	if err != nil {
		fmt.Println("UpsertUser error", err)
	}
}

func (app *Application) UpsertGroup(gid int, name string, attrs int) {
	//	fmt.Println("UpsertGroup")
	query := `INSERT INTO groups (gid, name, attrs) 
	VALUES (:uid, :name, :attrs) ON CONFLICT(gid) DO UPDATE 
	SET name=:name, attrs=attrs|:attrs;`
	group := Group{Gid: gid, Name: name, Attrs: attrs}
	_, err := app.db.NamedExec(query, group)
	if err != nil {
		fmt.Println("UpsertGroup error", err)
	}
}

func (app *Application) UpsertMembership(uid int, gid int) {
	//	fmt.Println("UpsertMembership")
	query := `INSERT INTO members (uid, gid) 
	VALUES (:uid, :gid) ON CONFLICT(uid, gid) DO NOTHING`
	member := Member{Uid: uid, Gid: gid}
	_, err := app.db.NamedExec(query, member)
	if err != nil {
		fmt.Println("UpsertMembership error", err)
	}
}

func (app *Application) UpsertFriendship(uid1 int, uid2 int) {
	//	fmt.Println("UpsertFriendship")
	query := `INSERT INTO friends (uid1, uid2) 
	VALUES (:uid1, :uid2) ON CONFLICT(uid1, uid2) DO NOTHING`
	friend := Friend{Uid1: uid1, Uid2: uid2}
	_, err := app.db.NamedExec(query, friend)
	if err != nil {
		fmt.Println("UpsertFriendship error", err)
	}
}

func (app *Application) UpsertPost(pid int, oid int, fid int, date int64, text string, cc int, lc int, rc int, vc int) {
	//	fmt.Println("UpsertPost")
	query := `INSERT INTO posts (pid, oid, fid, date, text, cmnts_count, likes_count, repos_count, views_count) 
	VALUES (:Pid, :Oid, :Fid, :Date, :Text, :CmntsCount, :LikesCount, :LikesCount, :ViewsCount) 
	ON CONFLICT(pid, oid, fid) DO UPDATE 
	SET date=:Date, text=:Text, cmnts_count=:CmntsCount, likes_count=:LikesCount, repos_count=:ReposCount, views_count=:ViewsCount;`
	post := Post{Pid: pid, Oid: oid, Fid: fid, Date: date, Text: text, CmntsCount: cc, LikesCount: lc, ReposCount: rc, ViewsCount: vc}
	_, err := app.db.NamedExec(query, post)
	if err != nil {
		fmt.Println("UpsertPost error", err)
	}
}

func (app *Application) UpsertLike(typ int, oid int, iid int, uid int) {

}
