
ipfn_test(
  UNIT_TEST
  BINARY_NAME codecs
  SOURCES ${CMAKE_CURRENT_SOURCE_DIR}/codecs.cc
  SOURCE_DIRECTORY ${TESTS_ROOT}/codecs
  LIBRARIES gtest cppcodec
)

ipfn_test(
  UNIT_TEST
  BINARY_NAME varint
  SOURCES ${CMAKE_CURRENT_SOURCE_DIR}/varint.cc
  SOURCE_DIRECTORY ${TESTS_ROOT}/varint
  LIBRARIES gtest cppcodec
)
