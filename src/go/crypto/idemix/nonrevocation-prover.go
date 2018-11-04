// Copyright © 2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2016-2018 IBM Corp. All Rights Reserved.
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

package idemix

import (
	"github.com/hyperledger/fabric-amcl/amcl"
	"github.com/hyperledger/fabric-amcl/amcl/FP256BN"
	"github.com/pkg/errors"
)

type nonRevokedProver interface {
	getFSContribution(rh *FP256BN.BIG, rRh *FP256BN.BIG, cri *CredentialRevocationInformation, rng *amcl.RAND) ([]byte, error)
	getNonRevokedProof(chal *FP256BN.BIG) (*NonRevocationProof, error)
}
type nopNonRevokedProver struct{}

func (prover *nopNonRevokedProver) getFSContribution(rh *FP256BN.BIG, rRh *FP256BN.BIG, cri *CredentialRevocationInformation, rng *amcl.RAND) ([]byte, error) {
	return nil, nil
}
func (prover *nopNonRevokedProver) getNonRevokedProof(chal *FP256BN.BIG) (*NonRevocationProof, error) {
	ret := &NonRevocationProof{}
	ret.RevocationAlg = int32(ALG_NO_REVOCATION)
	return ret, nil
}

func getNonRevocationProver(algorithm RevocationAlgorithm) (nonRevokedProver, error) {
	switch algorithm {
	case ALG_NO_REVOCATION:
		return &nopNonRevokedProver{}, nil
	default:
		// unknown revocation algorithm
		return nil, errors.Errorf("unknown revocation algorithm %d", algorithm)
	}
}
