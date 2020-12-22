cmake_minimum_required(VERSION 3.17)
project(gnu_attr C)

set(CMAKE_C_STANDARD 11)
aux_source_directory(src SRC_LIST)

add_executable(gnu_attr ${SRC_LIST} )