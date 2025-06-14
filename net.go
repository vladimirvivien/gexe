package gexe

import (
	"github.com/vladimirvivien/gexe/net"
)

func (e *Session) AddressUsable(addr string) error {
	return net.AddrUsable(e.Eval(addr))
}
