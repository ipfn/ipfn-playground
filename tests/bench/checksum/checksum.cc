//
// Copyright © 2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2018 Alexander Gallego. All Rights Reserved.
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
#include <thread>

#include <benchmark/benchmark.h>
#include <xxhash.h>

static constexpr uint32_t kPayloadSize = 1 << 29;
static char kPayload[kPayloadSize]{};

static void
BM_hash64(benchmark::State &state) {
  for (auto _ : state) {
    state.PauseTiming();
    std::memset(kPayload, 'x', state.range(0));
    state.ResumeTiming();
    benchmark::DoNotOptimize(std::numeric_limits<uint32_t>::max() &
                             XXH64(kPayload, state.range(0), 0));
  }
}

BENCHMARK(BM_hash64)->RangeMultiplier(2)->Range(256, 8 << 10);

BENCHMARK_MAIN();
