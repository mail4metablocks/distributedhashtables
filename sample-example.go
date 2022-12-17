package main

import (
	"context"
	"fmt"

	libp2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
)

func main() {
	// Create a libp2p host with the Kademlia DHT module
	host, err := libp2p.New(
		context.Background(),
		libp2p.ListenAddrs([]multiaddr.Multiaddr{}),
		libp2p.Routing(dht.NewDHTClient(context.Background(), nil)),
	)
	if err != nil {
		panic(err)
	}

	// Connect to the DHT
	dht := host.Routing.(*dht.IpfsDHT)
	if err := dht.Bootstrap(context.Background()); err != nil {
		panic(err)
	}

	// Insert a (key, value) pair into the DHT
	key := []byte("my-key")
	value := []byte("my-value")
	if err := dht.PutValue(context.Background(), key, value); err != nil {
		panic(err)
	}

	// Look up the value for the given key
	if value, err := dht.GetValue(context.Background(), key); err == nil {
		fmt.Printf("Retrieved value: %s\n", value)
	}
}
    
