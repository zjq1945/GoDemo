package jackymysql

type Employee struct {
	ID     int
	Name   string
	Depart string
}

type Consumer struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Age      int    `json:"Age"`
	Location string `json:"Location"`
}
