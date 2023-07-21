package views

type Pagination struct {
	Count       int
	CurrentPage int
	TotalPages  int
	PrevPage    int
	NextPage    int
}

type UserRec struct {
	Uid  int
	Name string
}

type UsersList struct {
	Title string
	Items []UserRec
	Pagination
}

type GroupRec struct {
	Gid  int
	Name string
}

type GroupsList struct {
	Title string
	Items []GroupRec
	Pagination
}

type GroupData struct {
	Gid     int
	Name    string
	Title   string
	Members []UserRec
}

type UserData struct {
	Uid     int
	Name    string
	Title   string
	Friends []UserRec
	Groups  []GroupRec
}
