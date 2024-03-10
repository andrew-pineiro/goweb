package controllers

type Message struct {
	Name  string
	Email string
	Body  string
}

func RecvMessage(msg Message) error {
	var err error
	//TODO: Get this stuff saved to a DB
	return err
}
