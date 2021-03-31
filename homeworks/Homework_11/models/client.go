package models

import "fmt"

type Client struct {
	Id       int
	FullName string
	Balance  int
}

func (cl *Client) PrintAll() {
	fmt.Println(cl.Id, cl.FullName, cl.Balance)
}
func (cl *Client) IncreaseBalance(a int) {
	cl.Balance = cl.Balance + a
}

func (cl *Client) DecreaseBalance(b int) {
	cl.Balance = cl.Balance - b
}
