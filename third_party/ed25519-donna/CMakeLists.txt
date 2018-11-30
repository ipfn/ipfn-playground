
add_library(ed25519-donna STATIC ${CMAKE_CURRENT_SOURCE_DIR}/ed25519.c)
set_property(TARGET ed25519-donna PROPERTY POSITION_INDEPENDENT_CODE ON)
target_include_directories(ed25519-donna PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})

if(USE_SSE2)
  if("${CMAKE_SIZEOF_VOID_P}" EQUAL "8")
    target_compile_options(ed25519-donna PRIVATE -m64 -O3 -c -DED25519_SSE2)
  else("${CMAKE_SIZEOF_VOID_P}" EQUAL "8")
    target_compile_options(ed25519-donna PRIVATE -m32 -O3 -c -DED25519_SSE2 -msse2)
  endif()
endif()

# TODO: hash function
target_link_libraries(ed25519-donna PUBLIC OpenSSL::SSL)
