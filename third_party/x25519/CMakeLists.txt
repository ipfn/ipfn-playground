
if(NOT "${CMAKE_SIZEOF_VOID_P}" EQUAL "8")
  set(x25519_srcs ${CMAKE_CURRENT_SOURCE_DIR}/x25519.c)
else()
  set(x25519_srcs ${CMAKE_CURRENT_SOURCE_DIR}/x25519-c64.c)
endif()

add_library(x25519 STATIC ${CMAKE_CURRENT_SOURCE_DIR}/x25519.h ${x25519_srcs})
target_include_directories(x25519 PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})
target_compile_options(x25519 PRIVATE -Wmissing-prototypes -Wdeclaration-after-statement -O2 -Wall)
set_target_properties(x25519 PROPERTIES LINKER_LANGUAGE CXX)

if(NOT "${CMAKE_SIZEOF_VOID_P}" EQUAL "8")
  target_compile_options(x25519 PRIVATE -m32)
endif()

# install(FILES ${CMAKE_CURRENT_SOURCE_DIR}/x25519.h DESTINATION include)

if(BUILD_TESTS)
  ipfn_test(
    UNIT_TEST
    BINARY_NAME x25519-sc
    SOURCES ${CMAKE_CURRENT_SOURCE_DIR}/test-x25519-sc.c ${CMAKE_CURRENT_SOURCE_DIR}/test-x25519-sc.s
    SOURCE_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    LIBRARIES x25519
    PROPERTIES LINKER_LANGUAGE CXX
  )
endif()
