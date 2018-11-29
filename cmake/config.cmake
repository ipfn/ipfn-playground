
ipfn_option(WITH_CCACHE "Build with ccache" OFF)

ipfn_option(BUILD_SEASTAR "Build with seastar" OFF)

ipfn_option(BUILD_TESTS "Build with tests" OFF)
ipfn_option(BUILD_TEST_COVERAGE "Build tests with coverage" OFF)
ipfn_option(BUILD_BENCHMARKS "Build with benchmarks" OFF)

ipfn_option(BUILD_INTEGRATION_TESTS "Build integration tests" OFF)

ipfn_option(BUILD_RUST_SDK "Build Rust target" OFF)
ipfn_option(BUILD_PYTHON_SDK "Build Python target" OFF)
ipfn_option(BUILD_WASM_TARGET "Build WebAssembly target" OFF)

ipfn_option(BUILD_DOCS "Build documentation"  OFF)
ipfn_option(BUILD_DEPS_DOCS "Build dependencies with documentation" OFF)
ipfn_option(BUILD_DEPS_DEMOS "Build dependencies with demos" OFF)
ipfn_option(BUILD_DEPS_TESTS "Build dependencies with tests" OFF)
ipfn_option(BUILD_DEPS_BENCH "Build dependencies with benchmarks" OFF)

ipfn_option(USE_SSE "Use SSE instructions for ica-l_crypto" OFF)
ipfn_option(USE_SSE2 "Use SSE2 instructions for ica-l_crypto and/or ed25519-donna" OFF)
ipfn_option(USE_AVX "Use AVX instructions for ica-l_crypto" OFF)
ipfn_option(USE_AVX2 "Use AVX2 instructions for ica-l_crypto" OFF)
ipfn_option(USE_AVX512 "Use AVX512 instructions for ica-l_crypto" OFF)
