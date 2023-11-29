package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

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

func getUsers(db *gorm.DB, page int, pageSize int, search string, order string, desc bool) ([]User, error) {
	var users []User
	var err error

	if page > 0 {
		offset := (page - 1) * pageSize
		if search == "" {
			err = db.Offset(offset).Limit(pageSize).Find(&users).Error
		} else {
			search = "%" + search + "%"
			err = db.Where("Name LIKE ?", search).Offset(offset).Limit(pageSize).Find(&users).Error
		}
	} else {
		if search == "" {
			err = db.Find(&users).Error
		} else {
			search = "%" + search + "%"
			err = db.Where("Name LIKE ?", search).Find(&users).Error
		}
	}
	if err != nil {
		return nil, err
	}
	return users, nil
}

func getGroups(db *gorm.DB, page int, pageSize int, search string, order string, desc bool) ([]Group, error) {
	var groups []Group
	var err error

	if page > 0 {
		offset := (page - 1) * pageSize
		if search == "" {
			err = db.Offset(offset).Limit(pageSize).Find(&groups).Error
		} else {
			search = "%" + search + "%"
			err = db.Where("Name LIKE ?", search).Offset(offset).Limit(pageSize).Find(&groups).Error
		}
	} else {
		if search == "" {
			err = db.Find(&groups).Error
		} else {
			search = "%" + search + "%"
			err = db.Where("Name LIKE ?", search).Find(&groups).Error
		}
	}
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func getPosts(db *gorm.DB, page int, pageSize int, search string, order string, desc bool) ([]Post, error) {
	var posts []Post
	var err error
	if page > 0 {
		offset := (page - 1) * pageSize
		if search == "" {
			err = db.Offset(offset).Limit(pageSize).Find(&posts).Error
		} else {
			search = "%" + search + "%"
			err = db.Where("Text LIKE ?", search).Offset(offset).Limit(pageSize).Find(&posts).Error
		}
	} else {
		if search == "" {
			err = db.Find(&posts).Error
		} else {
			search = "%" + search + "%"
			err = db.Where("Text LIKE ?", search).Find(&posts).Error
		}
	}
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func getFriends(db *gorm.DB, uid int, page int, pageSize int, search string, order string, desc bool) ([]User, error) {
	var friends []User
	var err error
		
	query := fmt.Sprintf("SELECT friends.uid2 AS uid, Name, Attrs, Type FROM users JOIN friends ON users.uid = friends.uid2 WHERE friends.uid1 = %d", uid)
		
	if search != "" {
		query += fmt.Sprintf(" AND Name LIKE '%%%s%%'", search)
	}
	if page > 0 {
		offset := (page - 1) * pageSize
		query += fmt.Sprintf(" OFFSET %d LIMIT %d", offset, pageSize);
	} 
	if order != "" {
		query += fmt.Sprintf(" ORDER BY %s", order)
		if desc {
			query += " DESC"
		}
	}
			
	fmt.Println("getFriends: ", uid, " = ", query)
	err = db.Raw(query).Scan(&friends).Error
	
	if err != nil {
		fmt.Println("getFriends error")
		return nil, err
	}
	return friends, nil
}

func getMemberships(db *gorm.DB, uid int, page int, pageSize int, search string, order string, desc bool) ([]Group, error) {
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

func getMembers(db *gorm.DB, gid int, page int, pageSize int, search string, order string, desc bool) ([]User, error) {
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

func getWall(db *gorm.DB, gid int, page int, pageSize int, search string, order string, desc bool) ([]PostWithUsername, error) {
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
