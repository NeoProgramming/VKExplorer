package views

// common control data

type Menu struct {
	MainMenu int
	SubMenu  int
}

type Pagination struct {
	Count       int
	CurrentPage int
	TotalPages  int
	PrevPage    int
	NextPage    int
}

type Search struct {
	SearchStr string
	AndOr     int
	SearchArg string
}

type Tags struct {
	TagsStr string
}

type Order struct {
	Sort string
	Desc int
}

// common data

type NameRec struct {
	Id               int
	Name             string
	OldestUpdateTime string
	NewestUpdateTime string
}

type Column struct {
	Name  string
	Desc  int
	Title string
	Query *string
}

type NameList struct {
	Menu
	Pagination
	Search
	Tags
	Title string
	Id    int
	Name  string
	Items []NameRec
	Columns []Column
}

type PostRec struct {
	Pid  int
	Fid  int
	Name string
	Text string
	Date string
}

type PostsList struct {
	Menu
	Pagination
	Search
	Tags
	Title string
	Id    int
	Name  string
	Items []PostRec
}

// special data for db entities

type UserInfo struct {
	Menu
	Id             int
	Name           string
	Title          string
	FriendsUpdated string
	GroupsUpdated  string
	WallUpdated    string
}

type GroupData struct {
	Menu
	Id             int
	Name           string
	Title          string
	MembersUpdated string
	WallUpdated    string
}
