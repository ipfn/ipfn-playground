
include_directories(include)

add_subdirectory(third_party)
add_subdirectory(core)

if(BUILD_APPS)
  add_subdirectory(python)
endif()

add_subdirectory(apps)

if(BUILD_PYTHON)
  add_subdirectory(python)
endif()

if(BUILD_RUST)
  add_subdirectory(rust)
endif()

if(BUILD_WASM)
  add_subdirectory(wasm)
endif()
