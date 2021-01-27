cmake_minimum_required(VERSION 3.17)
project(module C)

set(CMAKE_C_STANDARD 11)

# Find the kernel release
execute_process(
        COMMAND uname -r
        OUTPUT_VARIABLE KERNEL_RELEASE
        OUTPUT_STRIP_TRAILING_WHITESPACE
)

set(KERNELHEADERS_DIR /usr/lib/modules/${KERNEL_RELEASE}/build)
set(KERNELHEADERS_INCLUDE_DIRS ${KERNELHEADERS_DIR}/include)
set(KERNELHEADERS_ARCH_INCLUDE_DIRS ${KERNELHEADERS_DIR}/arch/x86/include/)

# Find kernel headers
include_directories(${KERNELHEADERS_INCLUDE_DIRS} ${KERNELHEADERS_ARCH_INCLUDE_DIRS})

# find MODULE_LICENSE("GPL"), MODULE_AUTHOR() etc.
add_definitions(-D__KERNEL__ -DMODULE)

add_executable(module main.c)

# CLion IDE will find symbols from <linux/*>
target_include_directories("module" PRIVATE ${KERNELHEADERS_INCLUDE_DIRS})
