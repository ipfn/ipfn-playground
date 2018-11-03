# Contains source code from smf.
#
# Copyright (c) 2018 IPFN Developers. All rights reserved.
# Copyright (c) 2017-2018 Alexander Gallego. All rights reserved.
#

if(BUILD_PYTHON_SDK)
  add_subdirectory(pybind11)
endif()

# smf + seastar
set(SEASTAR_CXX_DIALECT gnu++14)
set(SMF_ENABLE_TESTS ${BUILD_DEPS_TESTS}  CACHE BOOL "disable tests for smf")
set(SEASTAR_ENABLE_TESTS ${BUILD_DEPS_TESTS} CACHE BOOL "disable seastar tests")
set(SMF_BUILD_PROGRAMS ${BUILD_DEPS_DEMOS} CACHE BOOL "disable demo programs for smf")
set(SEASTAR_ENABLE_WERROR OFF CACHE BOOL "disable seastar -Werror")
add_subdirectory(smf)

# TVM Compiler stack
if(USE_TVM_RUNTIME OR USE_TVM_COMPILER)
  set(USE_SORT ON CACHE BOOL "Build with sort")
  set(USE_RANDOM ON CACHE BOOL "Build with random")
  add_subdirectory(tvm)
endif()

# secp256k1 is included in root CMake
