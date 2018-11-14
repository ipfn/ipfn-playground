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

package host

import (
	"context"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
	libp2phost "github.com/libp2p/go-libp2p-host"
)

// Host - Node host interface.
type Host interface {
	libp2phost.Host

	// RecoverPublicKey - Recovers public key from peer ID.
	RecoverPublicKey() (*btcec.PublicKey, error)
}

// New - Creates a new node host.
func New(ctx context.Context, opts ...libp2p.Option) (_ Host, err error) {
	// opts = append([]libp2p.Option{}, opts...)
	host, err := libp2p.New(ctx, opts...)
	return &p2pHost{
		Host: host,
	}, nil
}

type p2pHost struct {
	libp2phost.Host
}

func (h *p2pHost) RecoverPublicKey() (_ *btcec.PublicKey, err error) {
	pk, err := h.Host.ID().ExtractPublicKey()
	if err != nil {
		return
	}
	if pk == nil {
		return nil, nil
	}
	pubkey, ok := pk.(*crypto.Secp256k1PublicKey)
	if !ok {
		return nil, fmt.Errorf("invalid pubkey type %v", pk)
	}
	return (*btcec.PublicKey)(pubkey), nil
}
