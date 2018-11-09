
set(TESTS_ROOT ${PROJECT_SOURCE_DIR}/src/tests)

ipfn_test(
  UNIT_TEST
  BINARY_NAME example_test
  SOURCES ${TESTS_ROOT}/example_test.cc
  SOURCE_DIRECTORY ${TESTS_ROOT}
  LIBRARIES gtest
)
