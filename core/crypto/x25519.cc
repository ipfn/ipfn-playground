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
#include <ipfn/crypto/key_pair.hpp>
#include <ipfn/crypto/x25519.hpp>

#include <ed25519.h>
#include <x25519.h>

#include <cassert>

namespace ipfn {
namespace crypto {

private_key
x25519_shared(const private_key &priv, const public_key &pub) {
  private_key result;
  int res = x25519(result.wptr(), priv.ptr(), pub.ptr());
  assert(res == 0);
  return result;
}

private_key
x25519_shared(const key_pair &priv, const public_key &pub) {
  return x25519_shared(priv.key_private(), pub);
}

public_key
x25519_pubkey(const private_key &seed) {
  public_key result;
  curved25519_scalarmult_basepoint(result.wptr(), seed.ptr());
  return result;
}

}  // namespace crypto
}  // namespace ipfn
