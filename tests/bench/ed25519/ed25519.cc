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
#include <cryptopp/eccrypto.h>
#include <ed25519.h>

static void
BM_ed25519_donna_sign(benchmark::State &state) {
  size_t length = state.range(0) * sizeof(unsigned char);
  unsigned char *payload = reinterpret_cast<unsigned char *>(malloc(length));
  std::memset(payload, 'x', state.range(0));
  for (auto _ : state) {
    state.PauseTiming();
    ed25519_public_key pk;
    ed25519_secret_key sk;
    ed25519_randombytes_unsafe(sk, sizeof(ed25519_secret_key));
    ed25519_publickey(sk, pk);
    state.ResumeTiming();
    ed25519_signature sig;
    ed25519_sign(payload, length, sk, pk, sig);
    benchmark::DoNotOptimize(sig);
  }
}

static void
BM_ed25519_donna_verify(benchmark::State &state) {
  size_t length = state.range(0) * sizeof(unsigned char);
  unsigned char *payload = reinterpret_cast<unsigned char *>(malloc(length));
  std::memset(payload, 'x', state.range(0));
  for (auto _ : state) {
    state.PauseTiming();
    ed25519_public_key pk;
    ed25519_secret_key sk;
    ed25519_randombytes_unsafe(sk, sizeof(ed25519_secret_key));
    ed25519_publickey(sk, pk);
    ed25519_signature sig;
    ed25519_sign(payload, length, sk, pk, sig);
    state.ResumeTiming();
    int verify = ed25519_sign_open(payload, length, pk, sig);
    assert(verify == 1);
    benchmark::DoNotOptimize(verify);
  }
}

static void
BM_ed25519_donna_sign_verify(benchmark::State &state) {
  size_t length = state.range(0) * sizeof(unsigned char);
  unsigned char *payload = reinterpret_cast<unsigned char *>(malloc(length));
  std::memset(payload, 'x', state.range(0));
  for (auto _ : state) {
    state.PauseTiming();
    ed25519_public_key pk;
    ed25519_secret_key sk;
    ed25519_randombytes_unsafe(sk, sizeof(ed25519_secret_key));
    ed25519_publickey(sk, pk);
    state.ResumeTiming();
    ed25519_signature sig;
    ed25519_sign(payload, length, sk, pk, sig);
    int verify = ed25519_sign_open(payload, length, pk, sig);
    assert(verify == 1);
    benchmark::DoNotOptimize(verify);
  }
}

static void
BM_ed25519_donna_pk(benchmark::State &state) {
  for (auto _ : state) {
    state.PauseTiming();
    ed25519_secret_key sk;
    ed25519_public_key pk;
    ed25519_randombytes_unsafe(sk, sizeof(ed25519_secret_key));
    state.ResumeTiming();
    ed25519_publickey(sk, pk);
    benchmark::DoNotOptimize(pk);
    benchmark::DoNotOptimize(sk);
  }
}

static void
BM_curved25519_pk(benchmark::State &state) {
  for (auto _ : state) {
    state.PauseTiming();
    ed25519_public_key pk;
    ed25519_secret_key sk;
    ed25519_randombytes_unsafe(sk, sizeof(ed25519_secret_key));
    state.ResumeTiming();
    curved25519_scalarmult_basepoint(pk, sk);
    benchmark::DoNotOptimize(pk);
    benchmark::DoNotOptimize(sk);
  }
}

BENCHMARK(BM_ed25519_donna_pk)->Arg(32);
BENCHMARK(BM_ed25519_donna_sign)->Arg(32);
BENCHMARK(BM_ed25519_donna_verify)->Arg(32);
BENCHMARK(BM_ed25519_donna_sign_verify)->Arg(32);
BENCHMARK(BM_curved25519_pk)->Arg(32);

BENCHMARK_MAIN();
