package server

type Sync struct {
	RemoveList  chan int
	CommandList chan *Command
	WriteList   chan *Reply
	ClientList  chan *Client
}

func NewSync() *Sync {
	return &Sync{
		RemoveList:  make(chan int, 128),
		CommandList: make(chan *Command, 1024),
		WriteList:   make(chan *Reply, 1024),
		ClientList:  make(chan *Client, 128),
	}
}
