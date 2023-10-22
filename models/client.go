package models

type Client struct {
	name    string
	cpf     string
	account Account
}

func NewClient(name string, cpf string) *Client {
	return &Client{name: name, cpf: cpf}
}

func (c *Client) Name() string {
	return c.name
}

func (c *Client) SetName(name string) {
	c.name = name
}

func (c *Client) Cpf() string {
	return c.cpf
}

func (c *Client) SetCpf(cpf string) {
	c.cpf = cpf
}

func (c *Client) Account() Account {
	return c.account
}

func (c *Client) SetAccount(account Account) {
	c.account = account
}
