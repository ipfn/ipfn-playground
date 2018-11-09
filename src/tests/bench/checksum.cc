// Copyright 2018 The IPFN Developers
// Copyright 2018 Alexander Gallego
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
