package views

// common control data

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
	Title string
	Id    int
	Name  string
	Items []NameRec
	Pagination
	Search
	Tags
}

type PostRec struct {
	Pid  int
	Fid  int
	Name string
	Text string
}

type PostsList struct {
	Title string
	Id    int
	Name  string
	Items []PostRec
	Pagination
	Search
	Tags
}

// special data for db entities

type UserData struct {
	Id    int
	Name  string
	Title string
}

type GroupData struct {
	Id    int
	Name  string
	Title string
}
