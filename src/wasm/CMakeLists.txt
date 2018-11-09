
if (CMAKE_BUILD_TYPE MATCHES Debug)
  add_definitions(-g)
  add_definitions(-g4) # Generates emscripten C++ sourcemaps
endif()

# set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/build.emscripten)

file(GLOB_RECURSE IPFN_CORE_WASM_SRC *.cc ../core/*.cc ../core/**/*.cc)
file(GLOB_RECURSE IPFN_CORE_WASM_HDR *.h ../include/*.h ../include/**/*.h)

include(Emscripten)
add_executable(client_wasm ${IPFN_CORE_WASM_SRC} ${IPFN_CORE_WASM_HDR})
set_target_properties(client_wasm PROPERTIES LINK_FLAGS "-s DEMANGLE_SUPPORT=1 --bind")
# set_target_properties(client_wasm PROPERTIES LINK_FLAGS "-s DEMANGLE_SUPPORT=1 --preload-file ${CMAKE_SOURCE_DIR}/src/wasm/assets --bind")

add_custom_target(
  client_wasm_web
  ALL
  COMMAND ${CMAKE_COMMAND} -E make_directory ${CMAKE_CURRENT_BINARY_DIR}/html/
  COMMAND ${CMAKE_COMMAND} -E copy ${CMAKE_CURRENT_SOURCE_DIR}/web/index.html   ${CMAKE_CURRENT_BINARY_DIR}/html/
  WORKING_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}
)
add_dependencies(client_wasm client_wasm_web)
