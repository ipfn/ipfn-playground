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
#include <ipfn/crypto/key_data.hpp>

#include <cppcodec/hex_default_lower.hpp>

namespace ipfn {
namespace crypto {

std::optional<secret_key>
secret_key::from_hex(const std::string &encoded) {
  auto data = key_data_from_hex(encoded);
  if (!data) { return std::nullopt; }
  return secret_key(*data);
};

std::string
secret_key::encode_hex() const {
  return hex::encode(key_data_);
}

std::optional<key_data>
key_data_from_hex(const std::string &encoded) {
  // decoded size must be 32 bytes
  if (encoded.size() != 64) { return std::nullopt; }
  return hex::decode<key_data>(encoded);
};

}  // namespace crypto
}  // namespace ipfn
