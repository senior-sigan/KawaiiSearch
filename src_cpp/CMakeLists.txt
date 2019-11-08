cmake_minimum_required(VERSION 3.12)
project(kawaii_search)

set(CMAKE_CXX_STANDARD 17)

# TODO: somehow handle library to use. Maybe cmake flags??
set(CMAKE_PREFIX_PATH ./libs/osx/libtorch)
# set(CMAKE_PREFIX_PATH ./libs/linux_cpu/libtorch)
# set(CMAKE_PREFIX_PATH ./libs/linux_gpu/libtorch)

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -fPIC")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fPIC")

find_package(Torch REQUIRED)

add_library(${PROJECT_NAME} SHARED embedder.cpp embedder.hpp embedder_c.cpp embedder_c.h)
target_link_libraries(${PROJECT_NAME} "${TORCH_LIBRARIES}")
