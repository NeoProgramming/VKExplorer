package core

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func getTasks(db *sqlx.DB, page int, pageSize int) ([]Task, error) {
	var tasks []Task
	var err error
	query := fmt.Sprintf("SELECT id, name FROM tasks")
	if page > 0 {
		offset := (page - 1) * pageSize
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)
	}
	query += ";"

	fmt.Println("getTasks: ", query)
	err = db.Select(&tasks, query)
	if err != nil {
		fmt.Println("getTasks error", err)
		return nil, err
	}
	return tasks, nil
}

func getUsers(db *sqlx.DB, page int, pageSize int, search string, filters string, order string, desc bool) ([]User, error) {
	var users []User
	var err error
	query := fmt.Sprintf("SELECT uid, name, attrs, type, oldest, newest FROM users")

	if search != "" {
		query += fmt.Sprintf(" WHERE Name LIKE '%%%s%%'", search)
	}
	if filters != "" {
		m, im := decodeFilterMasks(filters)
		if search != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += fmt.Sprintf(" attrs & %d = %d AND attrs & %d = 0", m, m, im)
	}
	if order != "" {
		query += fmt.Sprintf(" ORDER BY %s", order)
		if desc {
			query += " DESC"
		}
	}
	if page > 0 {
		offset := (page - 1) * pageSize
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)
	}
	query += ";"

	fmt.Println("getUsers: ", query)
	err = db.Select(&users, query)
	if err != nil {
		fmt.Println("getUsers error", err)
		return nil, err
	}
	return users, nil
}

func getGroups(db *sqlx.DB, page int, pageSize int, search string, order string, desc bool) ([]Group, error) {
	var groups []Group
	var err error
	query := fmt.Sprintf("SELECT gid, name, attrs, type, oldest, newest FROM groups")

	if search != "" {
		query += fmt.Sprintf(" WHERE Name LIKE '%%%s%%'", search)
	}
	if order != "" {
		query += fmt.Sprintf(" ORDER BY %s", order)
		if desc {
			query += " DESC"
		}
	}
	if page > 0 {
		offset := (page - 1) * pageSize
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)
	}
	query += ";"

	fmt.Println("getGroups: ", query)
	err = db.Select(&groups, query)
	if err != nil {
		fmt.Println("getGroups error", err)
		return nil, err
	}
	return groups, nil
}

func getPosts(db *sqlx.DB, page int, pageSize int, search string, order string, desc bool) ([]Post, error) {
	var posts []Post
	var err error
	query := fmt.Sprintf("SELECT pid, oid, fid, date, text FROM posts")
	if search != "" {
		query += fmt.Sprintf(" WHERE Text LIKE '%%%s%%'", search)
	}
	if order != "" {
		query += fmt.Sprintf(" ORDER BY %s", order)
		if desc {
			query += " DESC"
		}
	}
	if page > 0 {
		offset := (page - 1) * pageSize
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)
	}
	query += ";"

	fmt.Println("getPosts: ", query)
	err = db.Select(&posts, query)
	if err != nil {
		fmt.Println("getPosts error", err)
		return nil, err
	}
	return posts, nil
}

func getFriends(db *sqlx.DB, uid int, page int, pageSize int, search string, order string, desc bool) ([]User, error) {
	var friends []User
	var err error
	query := fmt.Sprintf("SELECT friends.uid2 AS uid, Name, Attrs, Type FROM users JOIN friends ON users.uid = friends.uid2 WHERE friends.uid1 = %d", uid)

	if search != "" {
		query += fmt.Sprintf(" AND Name LIKE '%%%s%%'", search)
	}
	if order != "" {
		query += fmt.Sprintf(" ORDER BY %s", order)
		if desc {
			query += " DESC"
		}
	}
	if page > 0 {
		offset := (page - 1) * pageSize
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)
	}

	fmt.Println("getFriends: ", query)
	err = db.Select(&friends, query)
	if err != nil {
		fmt.Println("getFriends error: ", err)
		return nil, err
	}
	return friends, nil
}

func getMemberships(db *sqlx.DB, uid int, page int, pageSize int, search string, order string, desc bool) ([]Group, error) {
	var groups []Group
	var err error
	query := fmt.Sprintf("SELECT groups.gid AS gid, name, attrs, type FROM groups JOIN members ON groups.gid = members.gid WHERE members.uid = %d", uid)
	if search != "" {
		query += fmt.Sprintf(" AND Name LIKE '%%%s%%'", search)
	}
	if order != "" {
		query += fmt.Sprintf(" ORDER BY %s", order)
		if desc {
			query += " DESC"
		}
	}
	if page > 0 {
		offset := (page - 1) * pageSize
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)
	}

	fmt.Println("getMemberships: ", query)
	err = db.Select(&groups, query)
	if err != nil {
		fmt.Println("getMemberships error: ", err)
		return nil, err
	}
	return groups, nil
}

func getMembers(db *sqlx.DB, gid int, page int, pageSize int, search string, order string, desc bool) ([]User, error) {
	var members []User
	var err error
	query := fmt.Sprintf("SELECT members.uid AS uid, name FROM users JOIN members ON users.uid = members.uid WHERE members.gid = %d", gid)
	if search != "" {
		query += fmt.Sprintf(" AND Name LIKE '%%%s%%'", search)
	}
	if order != "" {
		query += fmt.Sprintf(" ORDER BY %s", order)
		if desc {
			query += " DESC"
		}
	}
	if page > 0 {
		offset := (page - 1) * pageSize
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)
	}

	fmt.Println("getMembers: ", query)
	err = db.Select(&members, query)
	if err != nil {
		fmt.Println("getMembers error: ", err)
		return nil, err
	}
	return members, nil
}

func getWall(db *sqlx.DB, gid int, page int, pageSize int, search string, order string, desc bool) ([]PostWithUsername, error) {
	var wall []PostWithUsername
	var err error
	query := fmt.Sprintf("SELECT pid, oid, fid, date, text, name FROM posts LEFT OUTER JOIN users ON users.uid = posts.fid WHERE posts.oid = %d", gid)
	if search != "" {
		query += fmt.Sprintf(" AND Text LIKE '%%%s%%'", search)
	}
	if order != "" {
		query += fmt.Sprintf(" ORDER BY %s", order)
		if desc {
			query += " DESC"
		}
	}
	if page > 0 {
		offset := (page - 1) * pageSize
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)
	}

	fmt.Println("getWall: ", query)
	err = db.Select(&wall, query)
	if err != nil {
		fmt.Println("getWall error: ", err)
		return nil, err
	}
	return wall, nil
}

func getCommonFriends(db *sqlx.DB, uid int) ([]User, error) {
	var users []User
	query := fmt.Sprintf(
		`SELECT uid, name FROM users JOIN friends ON 
		(users.uid = friends.uid1 AND %d = friends.uid2) OR
 		(users.uid = friends.uid2 AND %d = friends.uid1) 
		 WHERE users.uid = %d`, App.config.MyID, App.config.MyID, uid)
	err := db.Select(&users, query)
	if err != nil {
		fmt.Println("getMembers error: ", err)
		return nil, err
	}
	return users, nil
}

func getCommonGroups(db *sqlx.DB, uid int) ([]Group, error) {
	return nil, nil
}
