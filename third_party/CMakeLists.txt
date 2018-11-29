# Contains source code from smf.
#
# Copyright (c) 2018 IPFN Developers. All rights reserved.
# Copyright (c) 2017-2018 Alexander Gallego. All rights reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

set(BUILD_TESTING ${BUILD_DEPS_TESTS} CACHE BOOL "cppcodec tests")
add_subdirectory(cppcodec)
target_include_directories(cppcodec SYSTEM INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/cppcodec)

if(BUILD_PYTHON_SDK)
  add_subdirectory(pybind11)
endif()

set(SMF_THIRD_PARTY_DIR "${CMAKE_CURRENT_SOURCE_DIR}/smf/src/third_party")
set(XXHASH_SOURCE_DIR   "${SMF_THIRD_PARTY_DIR}/xxHash" CACHE "" EXTERNAL FORCE)
set(XXHASH_SOURCE_FILE  "${SMF_THIRD_PARTY_DIR}/xxHash/xxhash.c" CACHE "" EXTERNAL FORCE)

if(BUILD_BENCHMARKS)
  set_option_force(BENCHMARK_ENABLE_GTEST_TESTS OFF)
  set_option_force(BENCHMARK_ENABLE_TESTING     OFF)
endif()

if(BUILD_SEASTAR)
  # seastar
  set(SEASTAR_CXX_DIALECT gnu++14)
  set(SEASTAR_ENABLE_TESTS ${BUILD_DEPS_TESTS} CACHE BOOL "disable seastar tests")
  set(SEASTAR_ENABLE_WERROR OFF CACHE BOOL "disable seastar -Werror")
  # smf
  set(SMF_ENABLE_TESTS ${BUILD_DEPS_TESTS} CACHE BOOL "disable tests for smf")
  set(SMF_ENABLE_UNIT_TESTS ${BUILD_DEPS_TESTS} CACHE BOOL "disable seastar tests")
  set(SMF_ENABLE_BENCHMARK_TESTS ${BUILD_DEPS_BENCH} CACHE BOOL "disable seastar tests")
  set(SMF_BUILD_PROGRAMS ${BUILD_DEPS_DEMOS} CACHE BOOL "disable demo programs for smf")
  add_subdirectory(smf)
  target_include_directories(seastar SYSTEM INTERFACE
    ${CMAKE_CURRENT_SOURCE_DIR}/smf/third_party/seastar
    ${seastar_INCLUDE_DIRS})
endif()

if(BUILD_TESTS AND NOT SMF_ENABLE_TESTS)
  add_subdirectory(${SMF_THIRD_PARTY_DIR}/googletest/googletest)
endif()

if(BUILD_BENCHMARKS)
  add_subdirectory(${SMF_THIRD_PARTY_DIR}/google_benchmark)
endif()

add_subdirectory(ed25519-donna)
add_subdirectory(x25519)
add_subdirectory(isa-l_crypto)
