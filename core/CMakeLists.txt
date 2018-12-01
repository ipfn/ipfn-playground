
# IPFN core

file(GLOB_RECURSE IPFN_CORE_SRC
  ${CMAKE_CURRENT_SOURCE_DIR}/*.cc
  ${CMAKE_CURRENT_SOURCE_DIR}/**/*.cc
)
file(GLOB_RECURSE IPFN_CORE_HDR
  ${CMAKE_CURRENT_SOURCE_DIR}/../include/*.h
  ${CMAKE_CURRENT_SOURCE_DIR}/../include/**/*.h
  ${CMAKE_CURRENT_SOURCE_DIR}/../include/*.hpp
  ${CMAKE_CURRENT_SOURCE_DIR}/../include/**/*.hpp
)

add_library(ipfn_core SHARED ${IPFN_CORE_SRC} ${IPFN_CORE_HDR})

target_link_libraries(ipfn_core
    ${IPFN_LINKER_LIBS}
    ${IPFN_RUNTIME_LINKER_LIBS}
    x25519
    ed25519-donna
    base-x
    cppcodec)

target_include_directories(ipfn_core PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/../include)
set_target_properties(ipfn_core PROPERTIES LINKER_LANGUAGE CXX)
