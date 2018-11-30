
set(RUST_TESTS_ROOT ${PROJECT_SOURCE_DIR}/rust/tests)

ipfn_test(
  UNIT_TEST
  BINARY_NAME crust
  SOURCES ${RUST_TESTS_ROOT}/crust.cc
  SOURCE_DIRECTORY ${RUST_TESTS_ROOT}
  LIBRARIES gtest ipfn_rust ${RUST_DEPS_LIBS}
)

# Hello.cc
add_executable(ipfn_rust_ctest hello.cc)
target_link_libraries(ipfn_rust_ctest ipfn_core ipfn_rust ${RUST_DEPS_LIBS})
