cmake_minimum_required(VERSION 3.17)
project(module C)

set(CMAKE_C_STANDARD 11)

# Find kernel headers
list(APPEND CMAKE_MODULE_PATH "${CMAKE_SOURCE_DIR}/cmake")
find_package(KernelHeaders REQUIRED)

# find MODULE_LICENSE("GPL"), MODULE_AUTHOR() etc.
add_definitions(-D__KERNEL__ -DMODULE)

add_executable(module main.c)

# CLion IDE will find symbols from <linux/*>
target_include_directories("module" PRIVATE ${KERNELHEADERS_INCLUDE_DIRS})
