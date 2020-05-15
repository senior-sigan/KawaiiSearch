project(embedded-server)
cmake_minimum_required(VERSION 3.10)

list(APPEND CMAKE_MODULE_PATH ${PROJECT_SOURCE_DIR}/cmake)

set(GRPC_VERSION 1.28.1)
set(PROTOBUF_VERSION 3.11.4)

#find_package(gflags REQUIRED)
find_package(spdlog REQUIRED)
find_package(ProtobufWithTargets REQUIRED)
find_package(gRPC REQUIRED)

FILE(GLOB_RECURSE SOURCES_LIST ${PROJECT_SOURCE_DIR}/src/*.h ${PROJECT_SOURCE_DIR}/src/*.cpp ${PROJECT_SOURCE_DIR}/src/*.cc)

add_executable(${PROJECT_NAME} ${SOURCES_LIST})
target_link_libraries(${PROJECT_NAME} protobuf::libprotobuf spdlog::spdlog gRPC::grpc++)
