cmake_minimum_required(VERSION 3.7)
project(mojo-core)

set(CMAKE_CXX_STANDARD 17)

INCLUDE_DIRECTORIES(
        /usr/local/include
        ${PROJECT_SOURCE_DIR}/../../ncraft/ncraft/ncraft-core/include
        ${PROJECT_SOURCE_DIR}/include
)

set(SOURCE_FILES
        include/mojo/core/any.hpp
        include/mojo/core/array.hpp
        include/mojo/core/bool.hpp
        include/mojo/core/map.hpp
        include/mojo/core/error.hpp
        include/mojo/core/null.hpp
        include/mojo/core/numeric.hpp
        include/mojo/core/object.hpp
        include/mojo/core/optional.hpp
        include/mojo/core/range.hpp
        include/mojo/core/result.hpp
        include/mojo/core/string.hpp
        include/mojo/core/tuple.hpp
        include/mojo/core/union.hpp
        )

SET(MOJO_CORE_LIB mojo-core)

add_library(${MOJO_CORE_LIB} STATIC ${SOURCE_FILES} ${INCLUDE_FILES})
