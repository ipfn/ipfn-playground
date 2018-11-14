cargo_build(NAME ipfn_rust)

if(BUILD_TESTS)
  add_subdirectory(tests)
endif()
