package views

type Pagination struct {
	Count       int
	CurrentPage int
	TotalPages  int
	PrevPage    int
	NextPage    int
}

type UserRec struct {
	Uid        int
	Name       string
	UpdateTime string
}

type UsersList struct {
	Title string
	Items []UserRec
	Pagination
}

type GroupRec struct {
	Gid        int
	Name       string
	UpdateTime string
}

type GroupsList struct {
	Title string
	Items []GroupRec
	Pagination
}

type PostRec struct {
	Pid  int
	Fid  int
	Name string
	Text string
}

type PostsList struct {
	Title string
	Items []PostRec
	Pagination
}

type UserData struct {
	Uid     int
	Name    string
	Title   string
	Friends []UserRec
	Groups  []GroupRec
}

type GroupData struct {
	Gid     int
	Name    string
	Title   string
	Members []UserRec
	Wall    []PostRec
}
