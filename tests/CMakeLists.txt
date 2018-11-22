
set(TESTS_ROOT ${PROJECT_SOURCE_DIR}/tests)

set(GX_INSTALLED ${CMAKE_CURRENT_BINARY_DIR}/.gx-done)

add_custom_command(
  OUTPUT ${GX_INSTALLED}
  COMMAND gx install --global && touch ${GX_INSTALLED}
  WORKING_DIRECTORY ${PROJECT_SOURCE_DIR}
)

add_test(
  NAME go_ipfn_test
  COMMAND go test -failfast github.com/ipfn/ipfn/pkg/...
)

set_tests_properties(go_ipfn_test PROPERTIES
  WORKING_DIRECTORY ${PROJECT_SOURCE_DIR}
  DEPENDS ${GX_INSTALLED})
