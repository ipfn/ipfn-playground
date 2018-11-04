// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2017-2018 The go-ethereum Authors. All Rights Reserved.
//
// This file is part of the IPFN Project.
// This file was part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package accounts implements high level ECDSA account management.
package accounts

import (
	ethereum "github.com/ethereum/go-ethereum"
)

// Account - Account information.
type Account struct {
	Address Address `json:"address"`
	URL     URL     `json:"url"`
}

// Wallet represents a software or hardware wallet that might contain one or more
// accounts (derived from the same seed).
type Wallet interface {
	// URL retrieves the canonical path under which this wallet is reachable. It is
	// user by upper layers to define a sorting order over all wallets from multiple
	// backends.
	URL() URL

	// Status returns a textual status to aid the user in the current state of the
	// wallet. It also returns an error indicating any failure the wallet might have
	// encountered.
	Status() (string, error)

	// Open initializes access to a wallet instance. It is not meant to unlock or
	// decrypt account keys, rather simply to establish a connection to hardware
	// wallets and/or to access derivation seeds.
	//
	// The passphrase parameter may or may not be used by the implementation of a
	// particular wallet instance. The reason there is no passwordless open method
	// is to strive towards a uniform wallet handling, oblivious to the different
	// backend providers.
	//
	// Please note, if you open a wallet, you must close it to release any allocated
	// resources (especially important when working with hardware wallets).
	Open(passphrase string) error

	// Close releases any resources held by an open wallet instance.
	Close() error

	// Accounts retrieves the list of signing accounts the wallet is currently aware
	// of. For hierarchical deterministic wallets, the list will not be exhaustive,
	// rather only contain the accounts explicitly pinned during account derivation.
	Accounts() []Account

	// Contains returns whether an account is part of this particular wallet or not.
	Contains(account Account) bool

	// Derive attempts to explicitly derive a hierarchical deterministic account at
	// the specified derivation path. If requested, the derived account will be added
	// to the wallet's tracked account list.
	Derive(path DerivationPath, pin bool) (Account, error)

	// SelfDerive sets a base account derivation path from which the wallet attempts
	// to discover non zero accounts and automatically add them to list of tracked
	// accounts.
	//
	// Note, self derivaton will increment the last component of the specified path
	// opposed to decending into a child path to allow discovering accounts starting
	// from non zero components.
	//
	// You can disable automatic account discovery by calling SelfDerive with a nil
	// chain state reader.
	SelfDerive(base DerivationPath, chain ethereum.ChainStateReader)

	// // SignHash requests the wallet to sign the given hash.
	// //
	// // It looks up the account specified either solely via its address contained within,
	// // or optionally with the aid of any location metadata from the embedded URL field.
	// //
	// // If the wallet requires additional authentication to sign the request (e.g.
	// // a password to decrypt the account, or a PIN code o verify the transaction),
	// // an AuthNeededError instance will be returned, containing infos for the user
	// // about which fields or actions are needed. The user may retry by providing
	// // the needed details via SignHashWithPassphrase, or by other means (e.g. unlock
	// // the account in a keystore).
	// SignHash(account Account, hash []byte) ([]byte, error)

	// // SignTx requests the wallet to sign the given transaction.
	// //
	// // It looks up the account specified either solely via its address contained within,
	// // or optionally with the aid of any location metadata from the embedded URL field.
	// //
	// // If the wallet requires additional authentication to sign the request (e.g.
	// // a password to decrypt the account, or a PIN code to verify the transaction),
	// // an AuthNeededError instance will be returned, containing infos for the user
	// // about which fields or actions are needed. The user may retry by providing
	// // the needed details via SignTxWithPassphrase, or by other means (e.g. unlock
	// // the account in a keystore).
	// SignTx(account Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error)

	// // SignHashWithPassphrase requests the wallet to sign the given hash with the
	// // given passphrase as extra authentication information.
	// //
	// // It looks up the account specified either solely via its address contained within,
	// // or optionally with the aid of any location metadata from the embedded URL field.
	// SignHashWithPassphrase(account Account, passphrase string, hash []byte) ([]byte, error)

	// // SignTxWithPassphrase requests the wallet to sign the given transaction, with the
	// // given passphrase as extra authentication information.
	// //
	// // It looks up the account specified either solely via its address contained within,
	// // or optionally with the aid of any location metadata from the embedded URL field.
	// SignTxWithPassphrase(account Account, passphrase string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error)
}

// WalletEventType represents the different event types that can be fired by
// the wallet subscription subsystem.
type WalletEventType int

const (
	// WalletArrived is fired when a new wallet is detected either via USB or via
	// a filesystem event in the keystore.
	WalletArrived WalletEventType = iota

	// WalletOpened is fired when a wallet is successfully opened with the purpose
	// of starting any background processes such as automatic key derivation.
	WalletOpened

	// WalletDropped
	WalletDropped
)

// WalletEvent is an event fired by an account backend when a wallet arrival or
// departure is detected.
type WalletEvent struct {
	Wallet Wallet          // Wallet instance arrived or departed
	Kind   WalletEventType // Event type that happened in the system
}
