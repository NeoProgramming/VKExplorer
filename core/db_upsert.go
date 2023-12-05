package core

import (
	"fmt"
	"strings"
)

type Rec struct {
	uid   int
	name  string
	attrs int
}

func (app *Application) UpsertUser1(uid int, name string, about string, city string, domain string, photo string, attrs int) {
	var c []Rec
	app.db.Raw("SELECT uid, name, attrs FROM users WHERE uid = ?", uid).Scan(&c)
	if len(c) == 0 {
		query := fmt.Sprintf("INSERT INTO users(uid, name, about, city, domain, photo, attrs) VALUES(%d, '%s', '%s', '%s', '%s', '%s', %d);",
			uid, name, about, city, domain, photo, attrs)
		app.db.Exec(query)
	} else {
		if strings.TrimSpace(name) == "DELETED" {
			name = c[0].name
		}
		attrs |= c[0].attrs
		query := fmt.Sprintf("UPDATE users SET name = '%s', about = '%s', city = '%s', domain = '%s', photo = '%s', attrs = %d WHERE uid = %d;",
			name, about, city, domain, photo, attrs, uid)
		app.db.Exec(query)
	}
}

func (app *Application) UpsertUser(uid int, name string, about string, city string, domain string, photo string, attrs int) {
	user := User{Uid: uid, About: about, City: city, Domain: domain, Photo: photo}
	app.db.FirstOrCreate(&user, User{Uid: uid})
	if !(strings.TrimSpace(name) == "DELETED" && user.Name != "DELETED") {
		user.Name = name
	}
	user.Attrs |= attrs
	app.db.Save(&user)
}

func (app *Application) UpsertGroup(gid int, name string, attrs int) {
	fmt.Println("UpsertGroup")
	group := Group{Gid: gid, Name: name, Attrs: attrs}
	result := app.db.FirstOrCreate(&group, Group{Gid: gid})
	if result.Error != nil {
		fmt.Println("ERROR1: ", result.Error)
	}
	group.Attrs |= attrs
	res2 := app.db.Save(&group)
	if res2.Error != nil {
		fmt.Println("ERROR2: ", res2.Error)
	}
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

func (app *Application) UpsertPost(pid int, oid int, fid int, date int64, text string, cc int, lc int, rc int, vc int) {
	post := Post{Pid: pid, Oid: oid, Fid: fid, Date: date, Text: text, CmntsCount: cc, LikesCount: lc, ReposCount: rc, ViewsCount: vc}
	// find by pair of ID's
	app.db.FirstOrCreate(&post, Post{Pid: pid, Oid: oid})
	// modify
	// ...
	// save
	app.db.Save(&post)
}
func (app *Application) UpsertLike(typ int, oid int, iid int, uid int) {
	like := Reaction{Typ: typ, Oid: oid, Iid: iid, Uid: uid}
	app.db.FirstOrCreate(&like, Reaction{Typ: typ, Oid: oid, Iid: iid, Uid: uid})
	// modify
	// ...
	// save
	app.db.Save(&like)
}
