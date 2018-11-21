
set(TESTS_ROOT ${PROJECT_SOURCE_DIR}/tests)

add_custom_target(go_ipfn_gx)
add_custom_command(
  TARGET go_ipfn_gx
  COMMAND cd ${PROJECT_SOURCE_DIR} && gx install --global
)

add_custom_target(go_ipfn_gx_all ALL DEPENDS go_ipfn_gx)

add_test(
  NAME go_ipfn_test
  COMMAND go test -failfast github.com/ipfn/ipfn/pkg/...
)

set_tests_properties(go_ipfn_test PROPERTIES
  WORKING_DIRECTORY ${PROJECT_SOURCE_DIR}
  DEPENDS go_ipfn_gx)
