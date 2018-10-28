cmake_minimum_required(VERSION 3.5)

project(ipfn LANGUAGES C CXX)

enable_language(CXX)

set(CMAKE_VERBOSE_MAKEFILE on)

set(CMAKE_MODULE_PATH
    ${CMAKE_MODULE_PATH}
    "${CMAKE_SOURCE_DIR}/cmake"
    "${CMAKE_SOURCE_DIR}/cmake/modules"
    "${CMAKE_SOURCE_DIR}/cmake/rust"
    "${CMAKE_SOURCE_DIR}/cmake/toolchain"
    "${CMAKE_SOURCE_DIR}/cmake/util")

include(ccache)
include(generic)
include(Util)

# Configuration defaults
if(EXISTS ${CMAKE_CURRENT_BINARY_DIR}/config.cmake)
  include(${CMAKE_CURRENT_BINARY_DIR}/config.cmake)
elseif(EXISTS ${CMAKE_CURRENT_SOURCE_DIR}/config.cmake)
  include(${CMAKE_CURRENT_SOURCE_DIR}/config.cmake)
endif()

if(USE_LLVM)
  include(FindLLVM)
  include(LLVM)
endif()

if(BUILD_RUST)
  enable_language(Rust)
  include(CMakeCargo)
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

include(ProjectSecp256k1)

add_subdirectory(src)
