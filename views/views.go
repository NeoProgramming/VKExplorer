package views

// common control data

type Menu struct {
	MainMenu	int
	SubMenu		int
}

type Pagination struct {
	Count       int
	CurrentPage int
	TotalPages  int
	PrevPage    int
	NextPage    int
}

type Search struct {
	SearchStr	string
	AndOr		int
	SearchArg   string
}

type Tags struct {
	TagsStr		string
}

// common data

type NameRec struct {
	Id         int
	Name       string
	UpdateTime string
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
}

type PostRec struct {
	Pid  int
	Fid  int
	Name string
	Text string
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


type UserData struct {
	Menu
	Id    int
	Name  string
	Title string
}

type GroupData struct {
	Menu
	Id    int
	Name  string
	Title string
}
