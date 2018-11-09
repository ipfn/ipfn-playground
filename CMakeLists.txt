cmake_minimum_required(VERSION 3.9)

project(ipfn VERSION "0.0.0" LANGUAGES C CXX)

enable_language(CXX)

# https://cmake.org/cmake/help/v3.4/policy/CMP0065.html
cmake_policy(SET CMP0065 OLD)
set(CMAKE_VERBOSE_MAKEFILE ON)
set(CMAKE_EXPORT_COMPILE_COMMANDS 1)

if(NOT CMAKE_BUILD_TYPE)
  set(CMAKE_BUILD_TYPE Release)
endif()

if(CMAKE_CXX_COMPILER_ID STREQUAL "GNU")
  set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=gnu++14")
else()
  set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++14")
endif()

set(CMAKE_MODULE_PATH
    ${CMAKE_MODULE_PATH}
    "${CMAKE_SOURCE_DIR}/cmake"
    "${CMAKE_SOURCE_DIR}/cmake/modules"
    "${CMAKE_SOURCE_DIR}/cmake/rust"
    "${CMAKE_SOURCE_DIR}/cmake/toolchain"
    "${CMAKE_SOURCE_DIR}/cmake/util")

include(Util)

# Configuration defaults
if(EXISTS ${CMAKE_CURRENT_BINARY_DIR}/config.cmake)
  include(${CMAKE_CURRENT_BINARY_DIR}/config.cmake)
elseif(EXISTS ${CMAKE_CURRENT_SOURCE_DIR}/config.cmake)
  include(${CMAKE_CURRENT_SOURCE_DIR}/config.cmake)
endif()

include(ccache)
include(tests)

if(BUILD_RUST_SDK)
  enable_language(Rust)
  include(CMakeCargo)
  include(FindRust)
endif()

set(IPFN_LINKER_LIBS "")
set(IPFN_RUNTIME_LINKER_LIBS "")

if(BUILD_DOCS)
  find_package(Doxygen)

  if(DOXYGEN_FOUND AND BUILD_DEPS_DOCS)
    set(BUILD_DOCUMENTATION ON)
  else()
    set(BUILD_DOCUMENTATION OFF)
  endif()

  add_subdirectory(docs)
endif()

add_subdirectory(src)
