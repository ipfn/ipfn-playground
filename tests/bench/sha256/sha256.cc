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

static void
BM_sha256_cryptopp(benchmark::State &state) {
  for (auto _ : state) {
    state.PauseTiming();
    size_t length = state.range(0) * sizeof(unsigned char);
    unsigned char *payload = reinterpret_cast<unsigned char *>(malloc(length));
    std::memset(payload, 'x', state.range(0));
    std::array<uint8_t, 32> digest{};
    state.ResumeTiming();
    CryptoPP::SHA256 hash{};
    hash.Update(payload, length);
    hash.Final(reinterpret_cast<unsigned char *>(digest.data()));
    benchmark::DoNotOptimize(digest);
    state.PauseTiming();
    free(payload);
    state.ResumeTiming();
  }
}

BENCHMARK(BM_sha256_cryptopp)->RangeMultiplier(2)->Range(8, 8 << 10);

BENCHMARK_MAIN();
