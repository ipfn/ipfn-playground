
add_subdirectory(ipfn-test)

if(BUILD_SEASTAR)
  add_subdirectory(server)
endif()
