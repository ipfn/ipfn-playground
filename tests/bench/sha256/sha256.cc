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
#include <cstring>
#include <memory>

#include <benchmark/benchmark.h>
#include <cryptopp/sha.h>

#include <isa-l_crypto/mh_sha256.h>

static void
BM_sha256_cryptopp(benchmark::State &state) {
  size_t length = state.range(0) * sizeof(unsigned char);
  unsigned char *payload = reinterpret_cast<unsigned char *>(malloc(length));
  std::memset(payload, 'x', state.range(0));

  for (auto _ : state) {
    CryptoPP::SHA256 hash{};
    std::array<uint8_t, 32> digest{};
    hash.Update(payload, length);
    hash.Final(reinterpret_cast<unsigned char *>(digest.data()));
    benchmark::DoNotOptimize(digest);
  }
  free(payload);
}

static void
BM_sha256_ica_l_crypto(benchmark::State &state) {
  size_t length = state.range(0) * sizeof(unsigned char);
  unsigned char *payload = reinterpret_cast<unsigned char *>(malloc(length));
  std::memset(payload, 'x', state.range(0));
  struct mh_sha256_ctx *ctx =
    reinterpret_cast<struct mh_sha256_ctx *>(malloc(sizeof(mh_sha256_ctx)));

  for (auto _ : state) {
    std::array<uint8_t, 32> digest{};
    mh_sha256_init(ctx);
    mh_sha256_update(ctx, payload, length);
    mh_sha256_finalize(ctx, digest.data());
    benchmark::DoNotOptimize(digest);
  }
  free(payload);
}

static void
BM_sha512_cryptopp(benchmark::State &state) {
  size_t length = state.range(0) * sizeof(unsigned char);
  unsigned char *payload = reinterpret_cast<unsigned char *>(malloc(length));
  std::memset(payload, 'x', state.range(0));

  for (auto _ : state) {
    CryptoPP::SHA512 hash{};
    std::array<uint8_t, 32> digest{};
    hash.Update(payload, length);
    hash.Final(reinterpret_cast<unsigned char *>(digest.data()));
    benchmark::DoNotOptimize(digest);
  }
  free(payload);
}

BENCHMARK(BM_sha256_cryptopp)
  ->Range(32, 1e6)
  ->RangeMultiplier(2)
  ->Arg(1024)
  ->Arg(2048);

BENCHMARK(BM_sha256_ica_l_crypto)
  ->Range(32, 1e6)
  ->RangeMultiplier(2)
  ->Arg(1024)
  ->Arg(2048);

BENCHMARK(BM_sha512_cryptopp)
  ->Range(32, 1e6)
  ->RangeMultiplier(2)
  ->Arg(1024)
  ->Arg(2048);

BENCHMARK_MAIN();
