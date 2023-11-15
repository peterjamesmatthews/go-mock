package calculator

type Client struct{}

func (client *Client) Add(a int32, b int32) int32 {
	panic("Can't call this in unit tests!")
}

func (client *Client) Subtract(a int32, b int32) int32 {
	panic("Can't call this in unit tests!")
}

func (client *Client) Multiply(a int32, b int32) int32 {
	panic("Can't call this in unit tests!")
}

func (client *Client) Divide(a int32, b int32) int32 {
	panic("Can't call this in unit tests!")
}
