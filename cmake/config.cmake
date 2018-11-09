
ipfn_option(WITH_CCACHE "Build with ccache" OFF)

ipfn_option(BUILD_SMF "Build with SMF" OFF)

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

ipfn_option(USE_TVM_RUNTIME "Build with TVM runtime" OFF)
ipfn_option(USE_TVM_COMPILER "Build with TVM compiler (includes TVM runtime)" OFF)
