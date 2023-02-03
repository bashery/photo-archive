package main

import "net"

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast(sender *client, msg string) {
	for addr, m := range r.members {
		// do not send message to origin sender
		if sender.conn.RemoteAddr() != addr {
			m.msg(msg)
		}
	}
}
