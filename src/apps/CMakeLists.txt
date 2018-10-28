
find_package(Boost COMPONENTS
  chrono
  thread
  system
  REQUIRED)

add_executable(hello_binary main.cc)

target_link_libraries(hello_binary
  Secp256k1
  ipfn_core
  Boost::chrono
  Boost::thread
  Boost::system
  pthread)
