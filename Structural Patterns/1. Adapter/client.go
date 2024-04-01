package main

type Client struct {
}

func (c *Client) CalculateInt(adp Adapter) int {
	return adp.Calculate()
}
