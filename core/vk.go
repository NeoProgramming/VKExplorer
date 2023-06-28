package core

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"net/url"
)

type RecordAttrs int

const (
	RA_MY     = 0x1  // my friend or my group
	RA_FAV    = 0x2  // bookmarked user of group
	RA_FRIEND = 0x4  // my friend's friend; or my friend's group
	RA_MEMBER = 0x8  // member of my group
	RA_LIKER  = 0x10 // liker of some record
)

func InitVK() {
	App.vk = api.NewVK(App.config.AccessToken)
	fmt.Println("VK API library initialized")
}

func extractAccessToken(urlStr string) string {
	u, _ := url.Parse(urlStr)
	parameters, _ := url.ParseQuery(u.Fragment)
	accessToken := parameters.Get("access_token")
	return accessToken
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
			"fields": "first_name,last_name",
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
			app.UpsertUser(friend.ID, name, RA_MY)
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
	offset := 0
	count := 1000 // or any other value you want to set
	for {
		// make a request to the site
		fmt.Println("Request for group members: ", offset)
		members, err := app.vk.GroupsGetMembersFields(api.Params{
			"group_id": task.Xid,
			"offset":   offset,
			"count":    count,
			"fields":   "first_name,last_name",
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Received: ", len(members.Items), " Total: ", members.Count)
		// add downloaded users to the database
		for _, member := range members.Items {
			name := member.FirstName + " " + member.LastName
			fmt.Println(name)
			app.UpsertUser(member.ID, name, RA_MEMBER)
			app.UpsertMembership(member.ID, task.Xid)
		}
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
			"fields":  "first_name,last_name",
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
			app.UpsertUser(friend.ID, name, RA_MEMBER)
			app.UpsertFriendship(task.Xid, friend.ID)
		}
		// next offset
		offset += count
		// if the received number of elements is less than the number in the package, then the package is the last
		if len(friends.Items) < count {
			break
		}
	}
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
	fmt.Println("UserGroups updated")
}

func (app *Application) loadGroupWall(task *Task) {
	offset := 0
	count := 1000
	for {
		// make a request to the site
		fmt.Println("Request for group wall: ", offset)
		
		wall, err := app.vk.WallGet(api.Params{
			"owner_id": -task.Xid,
			"offset":  offset,
			"count":   count,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Received: ", len(wall.Items), " Total: ", wall.Count)
		
		// adding downloaded groups to the database
		
		
		// next offset
		offset += count
		// if the received number of elements is less than the number in the package, then the package is the last
		if len(wall.Items) < count {
			break
		}
	}
	fmt.Println("GroupWall updated")
}

func (app *Application) loadUserWall(task *Task) {
	offset := 0
	count := 1000
	for {
		// make a request to the site
		fmt.Println("Request for user wall: ", offset)
		
		wall, err := app.vk.WallGet(api.Params{
			"owner_id": task.Xid,
			"offset":  offset,
			"count":   count,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Received: ", len(wall.Items), " Total: ", wall.Count)
		
		// adding downloaded groups to the database
		
		
		// next offset
		offset += count
		// if the received number of elements is less than the number in the package, then the package is the last
		if len(wall.Items) < count {
			break
		}
	}
	fmt.Println("UserWall updated")
}

