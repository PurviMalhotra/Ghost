package discovery

type Peer struct {
	Name string
	Addr string
	Port int
}

var Peers = map[string]Peer{}