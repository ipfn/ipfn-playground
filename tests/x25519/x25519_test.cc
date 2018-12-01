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

#include <x25519.h>

#include <gtest/gtest.h>
#include <utility>

using namespace ipfn::crypto;

TEST(x25519, shared_key) {
  auto expected_shared_hex =
    "42dedd506f22f8bbe71c2dbfc31e50e2db53861a6f55a2cc77e07e4e271f9807";
  auto expected_pk1_hex =
    "4652486ebc271520d844e5bdda9ac243c05dcbe7bc9b93807073a32177a6f73d";
  auto expected_pk2_hex =
    "ffbc7ba2e4c43be03f8a7f020d0651f582ad1901c254eebb4ec2ecb73148e50d";
  auto seed1 = *key_pair::from_seed_hex(
    "1b4f0e9851971998e732078544c96b36c3d01cedf7caa332359d6f1d83567014");
  auto seed2 = *key_pair::from_seed_hex(
    "60303ae22b998861bce3b28f33eec1be758a213c86c93c076dbe9f558c11c752");

  ASSERT_EQ(seed1.x25519_public().encode_hex(), expected_pk1_hex);
  ASSERT_EQ(seed2.x25519_public().encode_hex(), expected_pk2_hex);

  auto shared1 = x25519_shared(seed1, seed2.x25519_public());
  auto shared2 = x25519_shared(seed2, seed1.x25519_public());
  ASSERT_EQ(shared1.encode_hex(), shared2.encode_hex());
  ASSERT_EQ(shared1.encode_hex(), expected_shared_hex);
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

TEST(x25519, 1k) {
  unsigned char e1k[32];
  unsigned char e2k[32];
  unsigned char e1e2k[32];
  unsigned char e2e1k[32];
  unsigned char e1[32] = {3};
  unsigned char e2[32] = {5};
  unsigned char k[32] = {9};

  int loop;
  int i;

  for (loop = 0; loop < 1000; ++loop) {
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
