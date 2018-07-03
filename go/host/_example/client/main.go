// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"

	golog "github.com/ipfs/go-log"
	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	multiaddr "github.com/multiformats/go-multiaddr"

	gologging "github.com/whyrusleeping/go-logging"
)

func main() {
	golog.SetAllLoggers(gologging.DEBUG) // Change to DEBUG for extra info

	// The context governs the lifetime of the libp2p node
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		panic(err)
	}

	host, err := libp2p.New(ctx,
		// Use your own created keypair
		libp2p.Identity(priv),

		// Set your own listen address
		// The config takes an array of addresses, specify as many as you want.
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/9010"),
	)
	if err != nil {
		panic(err)
	}

	ma, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/9000")
	if err != nil {
		panic(err)
	}

	remoteID, err := peer.IDB58Decode("QmVGPB8vjnUaJYGx7KCLk8J8f15iC2WunVH4aBXKTH5Qg5")
	if err != nil {
		panic(err)
	}

	host.Peerstore().AddAddr(remoteID, ma, peerstore.PermanentAddrTTL)

	fmt.Printf("Hello World, my second hosts ID is %s\n", host.ID().Pretty())

	remoteInfo := host.Peerstore().PeerInfo(remoteID)
	if err := host.Connect(ctx, remoteInfo); err != nil {
		log.Fatal(err)
	}

	c := host.Network().ConnsToPeer(remoteID)
	if len(c) < 1 {
		log.Fatal("should have connection by now at least.")
	}

	select {}
}
