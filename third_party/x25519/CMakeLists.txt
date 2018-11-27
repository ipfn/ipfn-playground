
set(x25519_srcs ${CMAKE_CURRENT_SOURCE_DIR}/x25519-donna-c64.c)

if(NOT "${CMAKE_SIZEOF_VOID_P}" EQUAL "8")
  set(x25519_srcs ${CMAKE_CURRENT_SOURCE_DIR}/x25519-donna.c)
endif()

add_library(x25519 SHARED ${x25519_srcs})
target_include_directories(x25519 PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})
target_compile_options(x25519 PRIVATE -Wmissing-prototypes -Wdeclaration-after-statement -O2 -Wall)
set_target_properties(x25519 PROPERTIES LINKER_LANGUAGE CXX)

if(NOT "${CMAKE_SIZEOF_VOID_P}" EQUAL "8")
  target_compile_options(x25519 PRIVATE -m32)
endif()

include_directories(${CMAKE_CURRENT_SOURCE_DIR})

install(FILES ${CMAKE_CURRENT_SOURCE_DIR}/x25519.h DESTINATION include)

# if(BUILD_TESTS)
#   add_executable(test-noncanon)
# endif()
