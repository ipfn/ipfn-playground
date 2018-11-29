cmake_minimum_required(VERSION 3.9)

project(isal_crypto VERSION "1.0.0" LANGUAGES C CXX ASM)

include(ExternalProject)

ExternalProject_Add(project_isal_crypto
  PREFIX ${CMAKE_BINARY_DIR}
  GIT_REPOSITORY "https://github.com/01org/isa-l_crypto"
  GIT_TAG 7dc8963e3b69610da78f24ea5e7fabf89c366340
  CONFIGURE_COMMAND aclocal && libtoolize --force --copy && autoreconf --force --install && ./configure --prefix=/third_party/isa-l_crypto
  BUILD_COMMAND make
  UPDATE_COMMAND ""
  INSTALL_COMMAND make install DESTDIR=${CMAKE_BINARY_DIR}
  BUILD_IN_SOURCE 1)

# add_library(isal_crypto STAIC IMPORTED)
# set_target_properties(isal_crypto PROPERTIES
#   IMPORTED_LOCATION ${CMAKE_CURRENT_BINARY_DIR}/lib/libisal_crypto.a
#   INTERFACE_INCLUDE_DIRECTORIES  )

add_library(isal_crypto SHARED IMPORTED GLOBAL)
set_target_properties(isal_crypto PROPERTIES INTERFACE_INCLUDE_DIRECTORIES ${CMAKE_CURRENT_BINARY_DIR}/include)
set_target_properties(isal_crypto PROPERTIES IMPORTED_LOCATION ${CMAKE_CURRENT_BINARY_DIR}/lib/libisal_crypto.so)
target_include_directories(isal_crypto INTERFACE ${CMAKE_CURRENT_BINARY_DIR}/include)
add_dependencies(isal_crypto project_isal_crypto)
# include_directories(${install_dir}/include/isal_crypto-2.0)
