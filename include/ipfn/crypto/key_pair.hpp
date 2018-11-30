//
// Copyright © 2018 The IPFN Developers. All Rights Reserved.
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
//
#pragma once

#include <ipfn/crypto/ed25519.hpp>
#include <ipfn/crypto/key_data.hpp>
#include <ipfn/crypto/x25519.hpp>

namespace ipfn {
namespace crypto {

class key_pair {
 public:
  /**
   * @brief  Constructs key pair from seed.
   */
  key_pair(private_key secret)
    : key_pair(secret, ed25519_pubkey(secret), x25519_pubkey(secret)){};

  /**
   * @brief  Constructs key pair from seed and public key.
   */
  key_pair(private_key secret, public_key pubkey)
    : key_pair(secret, pubkey, x25519_pubkey(secret)){};

  /**
   * @brief  Constructs key pair from seed, public key and x25519 public key.
   */
  key_pair(private_key secret, public_key pubkey, public_key x25519_pub)
    : private_(secret), public_(pubkey), x25519_(x25519_pub){};

  /**
   * @brief  Creates ed25519 key pair from 64 bytes of seed hex.
   */
  static std::optional<key_pair> from_seed_hex(const std::string &encoded);

  /**
   * @brief  Returns public part of ed25519 key.
   */
  inline const public_key &pubkey() const { return public_; }

  /**
   * @brief  Returns private part of ed25519 key.
   */
  inline const private_key &privkey() const { return private_; }

  /**
   * @brief  Returns public x25519 key derived from ed25519 key.
   */
  inline const public_key x25519_public() const { return x25519_; }

 private:
  private_key private_;
  public_key public_;
  public_key x25519_;
};

}  // namespace crypto
}  // namespace ipfn
