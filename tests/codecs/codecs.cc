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
#include <ipfn/util/coding/base32.hpp>
#include <ipfn/util/coding/base32i.hpp>
#include <ipfn/util/coding/base58.hpp>
#include <ipfn/util/coding/base64.hpp>

#include <gtest/gtest.h>

using namespace ipfn;

static std::vector<std::tuple<std::string, std::string, std::string>>
  base_test_data = {
    {{"Ag=="}, {"0b"}, {"2"}},
    {{"BA=="}, {"0s"}, {"4"}},
    {{"CA=="}, {"p0"}, {"8"}},
    {{"EA=="}, {"z0"}, {"G"}},
    {{"IA=="}, {"y0"}, {"W"}},
    {{"QA=="}, {"b0"}, {"16"}},
    {{"gAE="}, {"s00s"}, {"9gv"}},
    {{"gAI="}, {"s0p0"}, {"9h0"}},
    {{"gAQ="}, {"s0z0"}, {"9h2"}},
    {{"gAg="}, {"s0y0"}, {"9h6"}},
    {{"gBA="}, {"s0b0"}, {"9hE"}},
    {{"gCA="}, {"s0s0"}, {"9hU"}},
    {{"gEA="}, {"sp00"}, {"9i4"}},
    {{"0vrzv5DE4LJN"}, {"6tad8duscnstynb"}, {"2di1pinM6vY5R"}},
    {{"z6z8+fXBprF4"}, {"e7kde7d4cxntz70"}, {"2bIUuvRoYVJJ6"}},
    {{"neTVtPa4k4jVAQ=="}, {"nhjrtr8khzfc34bp"}, {"8oTEh3A5PkFlIv"}},
    {{"g/i12reXwYO4AQ=="}, {"sduttk4hjl0c8w0p"}, {"7O22Mld0VISh2X"}},
    {{"0cWbzpSwoa02"}, {"68zehn55kzs66rs"}, {"2coGoah288m38"}},
    {{"2NH16OCw9PRX"}, {"mqblt680kq6db4c"}, {"2i45K8B6LC5QF"}},
    {{"noDYyJfnsrOIAQ=="}, {"n60r3jyhu7et8z0p"}, {"8qSP47PJuEDghp"}},
    {{"lKWTlsLVtISUAQ=="}, {"jjje89kz6k6bf90p"}, {"8KLYKATTtgv3Yv"}},
    {{"oI+kyYPp37QM"}, {"5z86fjvqa8dmbq0"}, {"22V6mGCAt9GKi"}},
    {{"mdf3xMH0kcKmAQ=="}, {"n8tld3xp7jbu9fsp"}, {"8bHikW1E6BL9DJ"}},
    {{"YW55IGNhcm5hbCBwbGVhc3VyZQ=="},
     {"v9h8jbqqv9exuctvypcxcetpwr6hyeb"},
     {"HmVGpLTqmKQ3BZaV3O6c5olgNl"}},
};

TEST(cppcodec, base32) {
  auto decoded = base64::decode("YW55IGNhcm5hbCBwbGVhc3VyZQ==");
  auto encoded = base32::encode(decoded);
  ASSERT_EQ(encoded, "C5Q7J833C5S6WRBC41R6RSB1EDTQ4S8");
}

TEST(cppcodec, base32i_encode) {
  for (auto &test_pair : base_test_data) {
    auto decoded = base64::decode(std::get<0>(test_pair));
    auto encoded = base32i::encode(decoded);
    ASSERT_EQ(encoded, std::get<1>(test_pair));
  }
}

TEST(cppcodec, base32i_decode) {
  for (auto &test_pair : base_test_data) {
    auto decoded = base32i::decode(std::get<1>(test_pair));
    auto encoded = base64::encode(decoded);
    ASSERT_EQ(encoded, std::get<0>(test_pair));
  }
}

TEST(cppcodec, base58_encode) {
  for (auto &test_pair : base_test_data) {
    auto decoded = base64::decode(std::get<0>(test_pair));
    auto encoded = base58::encode(decoded);
    ASSERT_EQ(encoded, std::get<2>(test_pair));
  }
}

TEST(cppcodec, base58_decode) {
  for (auto &test_pair : base_test_data) {
    auto decoded = base58::decode(std::get<2>(test_pair));
    auto encoded = base64::encode(decoded);
    ASSERT_EQ(encoded, std::get<0>(test_pair));
  }
}

int
main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
