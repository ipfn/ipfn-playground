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
#include <cppcodec/base32_crockford.hpp>
#include <cppcodec/base64_rfc4648.hpp>
#include <ipfn/util/coding/base32i.h>

#include <gtest/gtest.h>

using base64 = cppcodec::base64_rfc4648;
using base32 = cppcodec::base32_crockford;
using base32i = cppcodec::base32_ipfn;

static std::vector<std::pair<std::string, std::string>> base_test_data = {
  {{"Ag=="}, {"0b"}},
  {{"BA=="}, {"0s"}},
  {{"CA=="}, {"p0"}},
  {{"EA=="}, {"z0"}},
  {{"IA=="}, {"y0"}},
  {{"QA=="}, {"b0"}},
  {{"gAE="}, {"s00s"}},
  {{"gAI="}, {"s0p0"}},
  {{"gAQ="}, {"s0z0"}},
  {{"gAg="}, {"s0y0"}},
  {{"gBA="}, {"s0b0"}},
  {{"gCA="}, {"s0s0"}},
  {{"gEA="}, {"sp00"}},
  {{"0vrzv5DE4LJN"}, {"6tad8duscnstynb"}},
  {{"z6z8+fXBprF4"}, {"e7kde7d4cxntz70"}},
  {{"neTVtPa4k4jVAQ=="}, {"nhjrtr8khzfc34bp"}},
  {{"g/i12reXwYO4AQ=="}, {"sduttk4hjl0c8w0p"}},
  {{"0cWbzpSwoa02"}, {"68zehn55kzs66rs"}},
  {{"2NH16OCw9PRX"}, {"mqblt680kq6db4c"}},
  {{"noDYyJfnsrOIAQ=="}, {"n60r3jyhu7et8z0p"}},
  {{"lKWTlsLVtISUAQ=="}, {"jjje89kz6k6bf90p"}},
  {{"oI+kyYPp37QM"}, {"5z86fjvqa8dmbq0"}},
  {{"mdf3xMH0kcKmAQ=="}, {"n8tld3xp7jbu9fsp"}},
  {{"YW55IGNhcm5hbCBwbGVhc3VyZQ=="}, {"v9h8jbqqv9exuctvypcxcetpwr6hyeb"}},
};

TEST(cppcodec, base32) {
  auto decoded = base64::decode("YW55IGNhcm5hbCBwbGVhc3VyZQ==");
  auto encoded = base32::encode(decoded);
  ASSERT_EQ(encoded, "C5Q7J833C5S6WRBC41R6RSB1EDTQ4S8");
}

TEST(cppcodec, base32i_encode) {
  for (auto &test_pair : base_test_data) {
    auto decoded = base64::decode(test_pair.first);
    auto encoded = base32i::encode(decoded);
    ASSERT_EQ(encoded, test_pair.second);
  }
}

TEST(cppcodec, base32i_decode) {
  for (auto &test_pair : base_test_data) {
    auto decoded = base32i::decode(test_pair.second);
    auto encoded = base64::encode(decoded);
    ASSERT_EQ(encoded, test_pair.first);
  }
}

int
main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
