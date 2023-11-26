package core

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

func (app *Application) UpsertPost(pid int, oid int, fid int, date int, text string, cc int, lc int, rc int, vc int) {
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
