package core

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"log"
	"net/http"
	"net/url"
	"time"
)

type RecordAttrs int

const (
	RA_MY        = 0x1  // my friend or my group
	RA_FAV       = 0x2  // bookmarked user of group
	RA_FRIEND    = 0x4  // my friend's friend; or my friend's group
	RA_MEMBER    = 0x8  // member of some group
	RA_LIKER     = 0x10 // liker of some record
	RA_COMMENTER = 0x20 // commenter of some record
)

func InitVK() {
	fmt.Println("VK API library initializing...")
	App.vk = api.NewVK(App.config.AccessToken)
	App.defVkClient = App.vk.Client

	if App.config.ProxyUse && App.config.ProxyAddr != "" {
		ActivateProxy()
	} else {
		fmt.Println("Working without Proxy")
	}
	if App.currentClient != nil {
		App.vk.Client = App.currentClient
	}
	fmt.Println("VK API library initialized")
}

func ActivateProxy() {
	fmt.Println("Working with Proxy...")
	fmt.Println("config.proxyAddr = ", App.config.ProxyAddr)
	proxyUrl, err := url.Parse(App.config.ProxyAddr) // "socks5://proxyIp:proxyPort"
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("proxyUrl = ", proxyUrl)
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	if transport != nil {
		fmt.Println("Transport ok")
	}
	App.proxyClient = &http.Client{
		Transport: transport,
	}
	if App.proxyClient != nil {
		fmt.Println("Proxy client prepared")
	}

	App.currentClient = App.proxyClient
	fmt.Println("Proxy enabled")
}

func DeactivateProxy() {
	App.proxyClient = nil
	App.currentClient = App.defVkClient
	fmt.Println("Proxy disabled")
}

func extractAccessToken(urlStr string) string {
	u, _ := url.Parse(urlStr)
	parameters, _ := url.ParseQuery(u.Fragment)
	accessToken := parameters.Get("access_token")
	return accessToken
}

func (app *Application) loadUserDataByName(task *Task) int {
	fmt.Println("Request for user data ")
	data, err := app.vk.UsersGet(api.Params{
		"user_ids": task.Name,
		"fields":   "id,first_name,last_name,about,city,domain,photo_max_orig,is_closed,can_access_closed",
	})
	if err != nil {
		fmt.Println(err)
		return 0
	}
	size := len(data)
	fmt.Println("Received: ", size)
	if size == 0 {
		return 0
	}
	app.UpsertUser(data[0].ID, data[0].FirstName+" "+data[0].LastName, data[0].About, data[0].City.Title, data[0].Domain, data[0].PhotoMaxOrig, 0)
	fmt.Println("User updated")
	return data[0].ID
}

func (app *Application) loadGroupDataByName(task *Task) int {
	fmt.Println("Request for group data ")
	data, err := app.vk.GroupsGetByID(api.Params{
		"group_id": task.Name,
		"fields":   "id,name,is_closed",
	})
	if err != nil {
		fmt.Println(err)
		return 0
	}
	size := len(data)
	fmt.Println("Received: ", size)
	if size == 0 {
		return 0
	}
	app.UpsertGroup(data[0].ID, data[0].Name, 0)
	fmt.Println("Group updated")
	return data[0].ID
}

func (app *Application) loadMyFriends(task *Task) {
	offset := 0
	count := 1000 // or any other value you want to set
	for {
		// make a request to the site
		fmt.Println("Request for users: ", offset)
		friends, err := app.vk.FriendsGetFields(api.Params{
			"offset": offset,
			"count":  count,
			"fields": "first_name,last_name,about,city,domain,photo_max_orig",
		})
		// if there is an error - for now, just exit to delete the task ...
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Received: ", len(friends.Items), " Total: ", friends.Count)
		// add downloaded users to the database
		for _, friend := range friends.Items {
			name := friend.FirstName + " " + friend.LastName
			fmt.Println(name)
			app.UpsertUser(friend.ID, name, friend.About, friend.City.Title, friend.Domain, friend.PhotoMaxOrig, RA_MY)
		}
		// next offset
		offset += count
		// if the received number of elements is less than the number in the package, then the package is the last
		if len(friends.Items) < count {
			break
		}
	}
	fmt.Println("Users updated")
}

func (app *Application) loadMyGroups(task *Task) {
	offset := 0
	count := 1000 // or any other value you want to set
	for {
		// make a request to the site
		fmt.Println("Request for groups: ", offset)
		groups, err := app.vk.GroupsGetExtended(api.Params{
			"offset": offset,
			"count":  count,
			"fields": "name",
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Received: ", len(groups.Items), " Total: ", groups.Count)
		// adding downloaded groups to the database
		for _, group := range groups.Items {
			name := group.Name
			fmt.Println(name)
			app.UpsertGroup(group.ID, name, RA_MY)
		}
		// next offset
		offset += count
		// if the received number of elements is less than the number in the package, then the package is the last
		if len(groups.Items) < count {
			break
		}
	}
	fmt.Println("Groups updated")
}

func (app *Application) loadMyBookmarks(task *Task) {

}

func (app *Application) loadGroupMembers(task *Task) {
	fmt.Println("loadGroupMembers for group ", task.Xid)
	offset := 0
	count := 1000 // or any other value you want to set
	for {
		// make a request to the site
		fmt.Println("Request for group members: offset=", offset)
		members, err := app.vk.GroupsGetMembersFields(api.Params{
			"group_id": task.Xid,
			"offset":   offset,
			"count":    count,
			"fields":   "first_name,last_name,about,city,domain,photo_max_orig",
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Received: ", len(members.Items), " Total: ", members.Count)
		// add downloaded users to the database
		//app.db.Begin()
		start := time.Now()
		for _, member := range members.Items {
			name := member.FirstName + " " + member.LastName
			fmt.Println(name)
			app.UpsertUser(member.ID, name, member.About, member.City.Title, member.Domain, member.PhotoMaxOrig, RA_MEMBER)
			//app.UpsertMembership(member.ID, task.Xid)
		}
		elapsed := time.Since(start)
		fmt.Println("UpsertUser: ", elapsed)
		//app.db.Commit()

		//app.db.Begin()
		start = time.Now()
		for _, member := range members.Items {
			//name := member.FirstName + " " + member.LastName
			//fmt.Println(name)
			//app.UpsertUser(member.ID, name, RA_MEMBER)
			app.UpsertMembership(member.ID, task.Xid)
		}
		elapsed = time.Since(start)
		fmt.Println("UpsertMembership: ", elapsed)
		//app.db.Commit()

		// next offset
		offset += count
		// if the received number of elements is less than the number in the package, then the package is the last
		if len(members.Items) < count {
			break
		}
	}
	fmt.Println("GroupMembers updated")
}

func (app *Application) loadUserFriends(task *Task) {
	offset := 0
	count := 1000 // or any other value you want to set
	for {
		// make a request to the site
		fmt.Println("Request for user friends: ", offset)
		friends, err := app.vk.FriendsGetFields(api.Params{
			"user_id": task.Xid,
			"offset":  offset,
			"count":   count,
			"fields":  "first_name,last_name,about,city,domain,photo_max_orig",
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Received: ", len(friends.Items), " Total: ", friends.Count)
		// add downloaded users to the database
		for _, friend := range friends.Items {
			name := friend.FirstName + " " + friend.LastName
			fmt.Println(name)
			app.UpsertUser(friend.ID, name, friend.About, friend.City.Title, friend.Domain, friend.PhotoMaxOrig, RA_MEMBER)
			app.UpsertFriendship(task.Xid, friend.ID)
		}
		// next offset
		offset += count
		// if the received number of elements is less than the number in the package, then the package is the last
		if len(friends.Items) < count {
			break
		}
	}
	// update time
	app.StampUserFirends(task.Xid)
	fmt.Println("UserFriends updated")
}

func (app *Application) loadUserGroups(task *Task) {
	offset := 0
	count := 1000 // or any other value you want to set
	for {
		// make a request to the site
		fmt.Println("Request for user groups: ", offset)
		groups, err := app.vk.GroupsGetExtended(api.Params{
			"user_id": task.Xid,
			"offset":  offset,
			"count":   count,
			"fields":  "name",
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Received: ", len(groups.Items), " Total: ", groups.Count)
		// adding downloaded groups to the database
		for _, group := range groups.Items {
			name := group.Name
			fmt.Println(name)
			app.UpsertGroup(group.ID, name, RA_FRIEND)
			app.UpsertMembership(task.Xid, group.ID)
		}
		// next offset
		offset += count
		// if the received number of elements is less than the number in the package, then the package is the last
		if len(groups.Items) < count {
			break
		}
	}
	app.StampUserGroups(task.Xid)
	fmt.Println("UserGroups updated")
}

func (app *Application) loadGroupWall(task *Task) {
	offset := 0
	totalCount := 0

	// logging
	//	logFile, err := os.OpenFile("groupwall.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer logFile.Close()
	//	log.SetOutput(logFile)

	fmt.Println("--------------------------------------------------")
	fmt.Println("loadGroupWall")
	for {
		// make a request to the site
		fmt.Println("Request for group wall: offset=", offset)

		wall, err := app.vk.WallGet(api.Params{
			"owner_id": -task.Xid,
			"offset":   offset,
			"count":    100, // 1000 is too big?,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		received := len(wall.Items)
		totalCount += received
		offset += received
		fmt.Println("Received: ", received, " RTotal: ", wall.Count, "CTotal: ", totalCount)

		// adding downloaded posts to the database
		for _, post := range wall.Items {
			text := post.Text
			fmt.Println(text)
			app.UpsertPost(post.ID, post.OwnerID, post.FromID, post.Date, post.Text,
				post.Comments.Count, post.Likes.Count, post.Reposts.Count, post.Views.Count)
			// attachments
			fmt.Println("===Att for ID=", post.ID, " OwnerID=", post.OwnerID, " FromID=", post.FromID)
			for _, att := range post.Attachments {
				if att.Type == "photo" {
					fmt.Println("Photo: ID=", att.Photo.ID, " AlbumID=", att.Photo.AlbumID, " OwnerId=", att.Photo.OwnerID,
						" UserID=", att.Photo.UserID, " Text=", att.Photo.Text)
				} else if att.Type == "posted_photo" {
					fmt.Println("PostedPhoto: ID=", att.PostedPhoto.ID, " OwnerID=", att.PostedPhoto.OwnerID)
				} else if att.Type == "video" {
					fmt.Println("Video: ID=", att.Video.ID, " OwnerID=", att.Video.OwnerID, " Title=",
						att.Video.Title, " Descr=", att.Video.Description)
				}
			}
		}

		// if the received number of elements is less than the number in the package, then the package is the last
		if totalCount >= wall.Count {
			break
		}
	}

	fmt.Println("GroupWall updated")
}

func (app *Application) loadUserWall(task *Task) {
	offset := 0
	totalCount := 0
	for {
		// make a request to the site
		fmt.Println("Request for user wall: ", offset)

		wall, err := app.vk.WallGet(api.Params{
			"owner_id": task.Xid,
			"offset":   offset,
			"count":    100,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		received := len(wall.Items)
		totalCount += received
		offset += received
		fmt.Println("Received: ", received, " RTotal: ", wall.Count, "CTotal: ", totalCount)

		// adding downloaded posts to the database
		for _, post := range wall.Items {
			//fmt.Println(post.text)
			app.UpsertPost(post.ID, post.OwnerID, post.FromID, post.Date, post.Text,
				post.Comments.Count, post.Likes.Count, post.Reposts.Count, post.Views.Count)
			// attachments
			fmt.Println("===Att for ID=", post.ID, " OwnerID=", post.OwnerID, " FromID=", post.FromID)
			for _, att := range post.Attachments {
				if att.Type == "photo" {
					fmt.Println("Photo: ID=", att.Photo.ID, " AlbumID=", att.Photo.AlbumID, " OwnerId=", att.Photo.OwnerID,
						" UserID=", att.Photo.UserID, " Text=", att.Photo.Text)
				} else if att.Type == "posted_photo" {
					fmt.Println("PostedPhoto: ID=", att.PostedPhoto.ID, " OwnerID=", att.PostedPhoto.OwnerID)
				} else if att.Type == "video" {
					fmt.Println("Video: ID=", att.Video.ID, " OwnerID=", att.Video.OwnerID, " Title=",
						att.Video.Title, " Descr=", att.Video.Description)
				}
			}
		}

		if totalCount >= wall.Count {
			break
		}
	}
	app.StampUserWall(task.Xid)
	fmt.Println("UserWall updated")
}

func (app *Application) loadUserFriendsByName(task *Task) {
	uid := app.loadUserDataByName(task)
	if uid != 0 {
		app.QueueById(TT_UserFriends, uid, "UserFriends")
	}
}

func (app *Application) loadUserGroupsByName(task *Task) {
	uid := app.loadUserDataByName(task)
	if uid != 0 {
		app.QueueById(TT_UserGroups, uid, "UserGroups")
	}
}

func (app *Application) loadUserWallByName(task *Task) {
	uid := app.loadUserDataByName(task)
	if uid != 0 {
		app.QueueById(TT_UserWall, uid, "UserWall")
	}
}

func (app *Application) loadGroupMembersByName(task *Task) {
	gid := app.loadGroupDataByName(task)
	if gid != 0 {
		app.QueueById(TT_GroupMembers, gid, "GroupMembers")
	}
}

func (app *Application) loadGroupWallByName(task *Task) {
	gid := app.loadGroupDataByName(task)
	if gid != 0 {
		app.QueueById(TT_GroupWall, gid, "GroupWall")
	}
}
