
add_library(base-x STATIC ${CMAKE_CURRENT_SOURCE_DIR}/base-x.hpp)
target_include_directories(base-x PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})
set_target_properties(base-x PROPERTIES LINKER_LANGUAGE CXX)

if(BUILD_TESTS)
  add_subdirectory(tests)
endif()
