# Set the minimum CMake version
cmake_minimum_required(VERSION 3.5) 

# Set the project name
project(hello_world)

# Specify the C++ standard
set(CMAKE_CXX_STANDARD 11)

# Add source file
add_executable(hello demo.cpp)

# Include directories for headers
target_include_directories(hello PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}) 