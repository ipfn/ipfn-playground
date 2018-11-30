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

#include <libhydrogen/hydro_x25519.h>

#ifdef USE_SODIUM
#include <sodium/crypto_scalarmult_curve25519.h>
#include <sodium/crypto_scalarmult_ed25519.h>
#endif

#include <limits>

#define X25519_BENCH(B_EXPR)                                                   \
  for (auto _ : state) {                                                       \
    state.PauseTiming();                                                       \
    ed25519_secret_key sk1, sk2;                                               \
    ed25519_randombytes_unsafe(sk1, sizeof(ed25519_secret_key));               \
    ed25519_randombytes_unsafe(sk2, sizeof(ed25519_secret_key));               \
    ed25519_public_key pk2;                                                    \
    ed25519_publickey(sk2, pk2);                                               \
    state.ResumeTiming();                                                      \
    ed25519_secret_key shared;                                                 \
    benchmark::DoNotOptimize(B_EXPR);                                          \
    benchmark::DoNotOptimize(shared);                                          \
  }

static void
BM_x25519(benchmark::State &state) {
  X25519_BENCH(x25519(shared, sk1, pk2));
}

static void
BM_hydro_x25519(benchmark::State &state) {
  X25519_BENCH(hydro_x25519_scalarmult(shared, sk1, pk2, true));
}

#ifdef USE_SODIUM
static void
BM_sodium_x25519_ed(benchmark::State &state) {
  X25519_BENCH(crypto_scalarmult_ed25519(shared, sk1, pk2));
}

static void
BM_sodium_x25519_ec(benchmark::State &state) {
  X25519_BENCH(crypto_scalarmult_curve25519(shared, sk1, pk2));
}
#endif

static void
BM_curved25519_pk(benchmark::State &state) {
  for (auto _ : state) {
    state.PauseTiming();
    ed25519_public_key pk;
    ed25519_secret_key sk;
    ed25519_randombytes_unsafe(sk, sizeof(ed25519_secret_key));
    state.ResumeTiming();
    curved25519_scalarmult_basepoint(pk, sk);
  }
}

BENCHMARK(BM_curved25519_pk)->Arg(32);

BENCHMARK(BM_x25519)->Arg(32)->Iterations(10000);
BENCHMARK(BM_hydro_x25519)->Arg(32)->Iterations(10000);
#ifdef USE_SODIUM
BENCHMARK(BM_sodium_x25519_ed)->Arg(32)->Iterations(10000);
BENCHMARK(BM_sodium_x25519_ec)->Arg(32)->Iterations(10000);
#endif

BENCHMARK_MAIN();
