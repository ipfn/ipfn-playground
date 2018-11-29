#!/bin/bash
#
# Copyright 2018 SMF Authors
# Copyright 2018 IPFN Developers
#

include(CMakeParseArguments)
enable_testing()
set(INTEGRATION_TESTS "")
set(UNIT_TESTS "")
set(BENCHMARK_TESTS "")
set(TEST_RUNNER ${PROJECT_SOURCE_DIR}/third_party/smf/src/test_runner.py)

message(STATUS "BUILD_TESTS=${BUILD_TESTS}")
message(STATUS "BUILD_BENCHMARKS=${BUILD_BENCHMARKS}")
message(STATUS "BUILD_INTEGRATION_TESTS=${BUILD_INTEGRATION_TESTS}")

function (ipfn_test)
  set(options INTEGRATION_TEST UNIT_TEST BENCHMARK_TEST)
  set(oneValueArgs BINARY_NAME SOURCE_DIRECTORY)
  set(multiValueArgs SOURCES LIBRARIES INCLUDES DEFINITIONS PROPERTIES)
  cmake_parse_arguments(IPFN_TEST "${options}" "${oneValueArgs}" "${multiValueArgs}" ${ARGN})

  if(IPFN_TEST_INTEGRATION_TEST AND BUILD_INTEGRATION_TESTS)
    set(IPFN_TEST_BINARY_NAME "ipfn_${IPFN_TEST_BINARY_NAME}_integration_test")
    set(INTEGRATION_TESTS "${INTEGRATION_TESTS} ${IPFN_TEST_BINARY_NAME}")
    add_executable(
      ${IPFN_TEST_BINARY_NAME} "${IPFN_TEST_SOURCES}")
    install(TARGETS ${IPFN_TEST_BINARY_NAME} DESTINATION bin)
    target_link_libraries(
      ${IPFN_TEST_BINARY_NAME}
      PUBLIC "${IPFN_TEST_LIBRARIES}")
    if(BUILD_TEST_COVERAGE)
      target_compile_options(
        ${IPFN_TEST_BINARY_NAME}
        PRIVATE --coverage)
    endif()
    add_test(
      NAME ${IPFN_TEST_BINARY_NAME}
      COMMAND
      ${TEST_RUNNER}
      --type integration
      --binary $<TARGET_FILE:${IPFN_TEST_BINARY_NAME}>
      --directory ${IPFN_TEST_SOURCE_DIRECTORY}
      )
  endif()
  if(IPFN_TEST_UNIT_TEST AND BUILD_TESTS)
    set(IPFN_TEST_BINARY_NAME "ipfn_${IPFN_TEST_BINARY_NAME}_unit_test")
    set(UNIT_TESTS "${UNIT_TESTS} ${IPFN_TEST_BINARY_NAME}")
    add_executable(
      ${IPFN_TEST_BINARY_NAME} "${IPFN_TEST_SOURCES}")
    install(TARGETS ${IPFN_TEST_BINARY_NAME} DESTINATION bin)
    target_link_libraries(
      ${IPFN_TEST_BINARY_NAME} "${IPFN_TEST_LIBRARIES}")
    if(BUILD_TEST_COVERAGE)
      target_compile_options(
        ${IPFN_TEST_BINARY_NAME}
        PRIVATE --coverage)
    endif()
    add_test(
      NAME ${IPFN_TEST_BINARY_NAME}
      COMMAND
      ${TEST_RUNNER}
      --type unit
      --binary $<TARGET_FILE:${IPFN_TEST_BINARY_NAME}>
      --directory ${IPFN_TEST_SOURCE_DIRECTORY}
      )
  endif()
  if(IPFN_TEST_BENCHMARK_TEST AND BUILD_BENCHMARKS)
    set(IPFN_TEST_BINARY_NAME "ipfn_${IPFN_TEST_BINARY_NAME}_benchmark")
    set(BENCHMARK_TESTS "${BENCHMARK_TESTS} ${IPFN_TEST_BINARY_NAME}")
    add_executable(
      ${IPFN_TEST_BINARY_NAME} "${IPFN_TEST_SOURCES}")
    target_link_libraries(
      ${IPFN_TEST_BINARY_NAME} "${IPFN_TEST_LIBRARIES}")
    add_test(
      NAME ${IPFN_TEST_BINARY_NAME}
      COMMAND
      ${TEST_RUNNER}
      --type benchmark
      --binary $<TARGET_FILE:${IPFN_TEST_BINARY_NAME}>
      --directory ${IPFN_TEST_SOURCE_DIRECTORY}
      )
  endif()
  foreach(i ${IPFN_TEST_INCLUDES})
    message("including ${i}")
    target_include_directories(${IPFN_TEST_BINARY_NAME}
      PUBLIC ${i})
  endforeach()
  foreach(i ${IPFN_TEST_DEFINITIONS})
    message("target_compile_definitions ${i}")
    target_compile_definitions(${IPFN_TEST_BINARY_NAME}
      PUBLIC ${i})
  endforeach()
  if(IPFN_TEST_PROPERTIES)
    message("set_target_properties ${IPFN_TEST_PROPERTIES}")
    set_target_properties(${IPFN_TEST_BINARY_NAME} PROPERTIES ${IPFN_TEST_PROPERTIES})
  endif()
endfunction()

if(BUILD_TESTS)
  add_custom_target(check
    COMMAND ctest --output-on-failure -N -R "^IPFN"
    DEPENDS "${UNIT_TESTS} ${INTEGRATION_TESTS} ${BENCHMARK_TESTS}")
endif()
