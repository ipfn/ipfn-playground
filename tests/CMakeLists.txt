
set(TESTS_ROOT ${PROJECT_SOURCE_DIR}/tests)

add_test(
  NAME go_ipfn
  COMMAND go test -v -failfast github.com/ipfn/ipfn/pkg/...
  WORKING_DIRECTORY ${TESTS_ROOT})

# set_tests_properties(go_ipfn PROPERTIES ENVIRONMENT GOPATH=${CMAKE_SOURCE_DIR})
