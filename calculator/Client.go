package calculator

// Client is a toy implementation of the Calculatorer interface that can't be used in unit tests.
//
// In the real application, this would use whatever protocol the calculator service uses to
// communicate.
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
