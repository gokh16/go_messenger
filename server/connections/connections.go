package connections

import "net"

type Clients struct {
	UserID      uint
	Flag        []string
	Connections []net.Conn
}
