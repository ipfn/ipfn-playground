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

#include <array>
#include <optional>

namespace ipfn {
namespace crypto {

/**
 * @brief  Key data with 32 byte size.
 */
using key_data = std::array<uint8_t, 32>;

/**
 * @brief  Key data wrapper class.
 */
class secret_key {
 public:
  secret_key() : key_data_(){};
  secret_key(key_data data) : key_data_(data){};

  /**
   * @brief  Creates ed25519 key from 64 bytes of hex representation.
   */
  static std::optional<secret_key> from_hex(const std::string &encoded);

  /**
   * @brief  Returns ed25519 key data.
   */
  inline const key_data &data() const { return key_data_; }

  /**
   * @brief  Returns direct access to ed25519 key data.
   */
  inline const unsigned char *ptr() const { return key_data_.data(); }

  /**
   * @brief  Returns direct write access to ed25519 key data.
   */
  inline unsigned char *wptr() { return key_data_.data(); }

  /**
   * @brief  Returns hex encoded ed25519 key data.
   */
  std::string encode_hex() const;

 private:
  key_data key_data_;
};

/**
 * @brief  Private ed25519 key.
 */
using private_key = secret_key;

/**
 * @brief  Public ed25519 key.
 */
using public_key = secret_key;

/**
 * @brief  Creates ed25519 key data from hex.
 * @note   It will return null if size is below 64 bytes.
 */
std::optional<key_data> key_data_from_hex(const std::string &encoded);

}  // namespace crypto
}  // namespace ipfn
