
ipfn_test(
  UNIT_TEST
  BINARY_NAME codecs
  SOURCES ${CMAKE_CURRENT_SOURCE_DIR}/codecs.cc
  SOURCE_DIRECTORY ${TESTS_ROOT}/codecs
  LIBRARIES gtest cppcodec
)
