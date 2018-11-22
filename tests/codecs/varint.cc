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
#include <ipfn/util/coding/varint.h>

#include <gtest/gtest.h>
#include <utility>

template <typename T>
inline const char *
data_finish(const T &t) {
  return reinterpret_cast<const char *>(&(t.back())) + 1;
}

static std::vector<std::pair<uint64_t, std::vector<uint8_t>>> varint_test_data =
  {{2U, {0x2}},
   {4U, {0x4}},
   {8U, {0x8}},
   {16U, {0x10}},
   {32U, {0x20}},
   {64U, {0x40}},
   {128U, {0x80, 0x1}},
   {256U, {0x80, 0x2}},
   {512U, {0x80, 0x4}},
   {1024U, {0x80, 0x8}},
   {2048U, {0x80, 0x10}},
   {4096U, {0x80, 0x20}},
   {8192U, {0x80, 0x40}},
   {5577006791947779410U,
    {0xd2, 0xfa, 0xf3, 0xbf, 0x90, 0xc4, 0xe0, 0xb2, 0x4d}},
   {8674665223082153551U,
    {0xcf, 0xac, 0xfc, 0xf9, 0xf5, 0xc1, 0xa6, 0xb1, 0x78}},
   {15352856648520921629U,
    {0x9d, 0xe4, 0xd5, 0xb4, 0xf6, 0xb8, 0x93, 0x88, 0xd5, 0x1}},
   {13260572831089785859U,
    {0x83, 0xf8, 0xb5, 0xda, 0xb7, 0x97, 0xc1, 0x83, 0xb8, 0x1}},
   {3916589616287113937U,
    {0xd1, 0xc5, 0x9b, 0xce, 0x94, 0xb0, 0xa1, 0xad, 0x36}},
   {6334824724549167320U,
    {0xd8, 0xd1, 0xf5, 0xe8, 0xe0, 0xb0, 0xf4, 0xf4, 0x57}},
   {9828766684487745566U,
    {0x9e, 0x80, 0xd8, 0xc8, 0x97, 0xe7, 0xb2, 0xb3, 0x88, 0x1}},
   {10667007354186551956U,
    {0x94, 0xa5, 0x93, 0x96, 0xc2, 0xd5, 0xb4, 0x84, 0x94, 0x1}},
   {894385949183117216U, {0xa0, 0x8f, 0xa4, 0xc9, 0x83, 0xe9, 0xdf, 0xb4, 0xc}},
   {11998794077335055257U,
    {0x99, 0xd7, 0xf7, 0xc4, 0xc1, 0xf4, 0x91, 0xc2, 0xa6, 0x1}}};

TEST(varint, parse) {
  for (auto &test_pair : varint_test_data) {
    uint64_t result;
    const char *end = varint::parse64(test_pair.second, &result);
    ASSERT_EQ(result, test_pair.first);
    ASSERT_EQ(end, data_finish(test_pair.second));
  }
}

TEST(varint, encode) {
  const size_t size = 10;
  char *buf = static_cast<char *>(malloc(size * sizeof(char)));
  for (auto &test_pair : varint_test_data) {
    char *end = varint::encode64(buf, test_pair.first);
    auto result = std::vector<uint8_t>(buf, end);
    ASSERT_EQ(result, test_pair.second);
  }
  free(buf);
}

int
main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
