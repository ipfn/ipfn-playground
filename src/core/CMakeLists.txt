
file(GLOB_RECURSE IPFN_CORE_SRC *.cc **/*.cc)
file(GLOB_RECURSE IPFN_CORE_HDR ../include/*.h ../include/**/*.h)

add_library(ipfn_core SHARED ${IPFN_CORE_SRC} ${IPFN_CORE_HDR})

target_link_libraries(ipfn_core
    ${IPFN_LINKER_LIBS}
    ${IPFN_RUNTIME_LINKER_LIBS})

target_include_directories(ipfn_core PUBLIC ".")
set_target_properties(ipfn_core PROPERTIES LINKER_LANGUAGE CXX)
