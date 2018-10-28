
file(GLOB_RECURSE IPFN_PYTHON_SRC *.cc)

pybind11_add_module(ipfn_python ${IPFN_PYTHON_SRC})
