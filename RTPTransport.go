package rtp

import (
	"net"
)

////////////////////Interface//////////////////////////////

type Transport interface {
	GetNetwork() string
	GetInterface() *net.Interface
	GetUDPAddr() *net.UDPAddr
	IsMulticast() bool

	SetConn(*net.UDPConn)
	GetConn() *net.UDPConn
}

////////////////////Implementation////////////////////////

type transport struct {
	net       string
	ifi       *net.Interface
	addr      *net.UDPAddr
	multicast bool

	conn *net.UDPConn
}

func NewTransport(net string, ifi *net.Interface, addr *net.UDPAddr, multicast bool) *transport {
	this := &transport{}

	this.net = net
	this.ifi = ifi
	this.addr = addr
	this.multicast = multicast

	return this
}

func (this *transport) GetNetwork() string {
	return this.net
}

func (this *transport) GetInterface() *net.Interface {
	return this.ifi
}

func (this *transport) GetUDPAddr() *net.UDPAddr {
	return this.addr
}

func (this *transport) IsMulticast() bool {
	return this.multicast
}

func (this *transport) SetConn(c *net.UDPConn) {
	this.conn = c
}

func (this *transport) GetConn() *net.UDPConn {
	return this.conn
}
