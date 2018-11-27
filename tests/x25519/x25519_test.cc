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
#include <x25519.h>

#include <gtest/gtest.h>
#include <utility>

TEST(x25519, shared_key) {
  std::vector<uint8_t> seed1 = hex::decode(
    "1b4f0e9851971998e732078544c96b36c3d01cedf7caa332359d6f1d83567014");
  std::vector<uint8_t> seed2 = hex::decode(
    "60303ae22b998861bce3b28f33eec1be758a213c86c93c076dbe9f558c11c752");

  ed25519_public_key pk1;
  ed25519_publickey(seed2.data(), pk1);

  ed25519_secret_key shared;
  x25519(shared, seed1.data(), pk1);
  ASSERT_EQ(hex::encode(shared),
            "ca957a6cedf359467c060feda1eb6ac27c105c49f3d83cecf4f1d17e8bcf841c");

  ed25519_public_key shpk;
  ed25519_publickey(shared, shpk);
  ASSERT_EQ(hex::encode(shpk),
            "275cb05c798d9aba960759d59ad27aa6f2e60171e9c74474a315e5538929b187");
}

TEST(x25519, noncanon) {
  // moved from x25519 package
  static const uint8_t point1[32] = {
    0x25, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
  };
  static const uint8_t point2[32] = {
    0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
    0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
    0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
  };
  static const uint8_t scalar[32] = {1};
  uint8_t out1[32], out2[32];

  x25519(out1, scalar, point1);
  x25519(out2, scalar, point2);

  ASSERT_FALSE(memcmp(out1, out2, sizeof(out1)) == 0);
}

TEST(x25519, 10k) {
  unsigned char e1k[32];
  unsigned char e2k[32];
  unsigned char e1e2k[32];
  unsigned char e2e1k[32];
  unsigned char e1[32] = {3};
  unsigned char e2[32] = {5};
  unsigned char k[32] = {9};

  int loop;
  int i;

  for (loop = 0; loop < 10000; ++loop) {
    x25519(e1k, e1, k);
    x25519(e2e1k, e2, e1k);
    x25519(e2k, e2, k);
    x25519(e1e2k, e1, e2k);
    for (i = 0; i < 32; ++i)
      ASSERT_EQ(e1e2k[i], e2e1k[i]);
    for (i = 0; i < 32; ++i)
      e1[i] ^= e2k[i];
    for (i = 0; i < 32; ++i)
      e2[i] ^= e1k[i];
    for (i = 0; i < 32; ++i)
      k[i] ^= e1e2k[i];
  }
}

int
main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
