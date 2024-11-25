package omiserver

import "github.com/go-redis/redis/v8"

type Client struct {
	opts *redis.Options
}

func NewClient(opts *redis.Options) *Client {
	return &Client{opts: opts}
}

func (c *Client) NewOmiServer(serverName, address string) *OmiServer {
	return newOmiServer(c.opts, serverName, address)
}
