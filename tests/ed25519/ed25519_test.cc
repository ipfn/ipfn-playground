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
#include <cppcodec/hex_default_lower.hpp>

#include <ed25519.h>
#include <utility>

#include <gtest/gtest.h>

TEST(ed25519, public_key) {
  std::vector<uint8_t> seed1 = hex::decode(
    "1b4f0e9851971998e732078544c96b36c3d01cedf7caa332359d6f1d83567014");
  std::vector<uint8_t> seed2 = hex::decode(
    "60303ae22b998861bce3b28f33eec1be758a213c86c93c076dbe9f558c11c752");

  ed25519_public_key pk1, pk2;
  ed25519_publickey(seed1.data(), pk1);
  ed25519_publickey(seed2.data(), pk2);

  ASSERT_EQ(hex::encode(pk1),
            "f33235d17f08fe3301747e873d83cdf37c317cb448e4b65d3fdd00c08d57a24e");
  ASSERT_EQ(hex::encode(pk2),
            "667e390ba5dcb5b79e371654027807459b1ab7becb4e778f73e9eec090205b10");
}

int
main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
