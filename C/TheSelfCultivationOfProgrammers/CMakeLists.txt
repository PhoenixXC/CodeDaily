cmake_minimum_required(VERSION 3.17)
project(TheSelfCultivationOfProgrammers C)

set(CMAKE_C_STANDARD 99)
aux_source_directory(src/4.1 ALLOC_MEM_AND_ADDR)

add_executable(
        TheSelfCultivationOfProgrammers
        main.c
        src/4.1/b.c
        src/4.1/a.c src/4.1/a.c src/4.1/tiny_hello_world.c
)