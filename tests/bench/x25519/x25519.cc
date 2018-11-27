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
#include <benchmark/benchmark.h>
#include <ed25519.h>
#include <x25519.h>

#include <limits>

static void
BM_x25519(benchmark::State &state) {
  for (auto _ : state) {
    state.PauseTiming();
    ed25519_secret_key sk1, sk2;
    ed25519_randombytes_unsafe(sk1, sizeof(ed25519_secret_key));
    ed25519_randombytes_unsafe(sk2, sizeof(ed25519_secret_key));
    ed25519_public_key pk2;
    ed25519_publickey(sk2, pk2);
    state.ResumeTiming();
    ed25519_secret_key shared;
    x25519(shared, sk1, pk2);
    benchmark::DoNotOptimize(shared);
  }
}

BENCHMARK(BM_x25519)->Arg(32)->Iterations(10000);

BENCHMARK_MAIN();
