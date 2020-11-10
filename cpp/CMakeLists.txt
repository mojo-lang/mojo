cmake_minimum_required(VERSION 3.8)
project(mojo)

set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${PROJECT_SOURCE_DIR}/bin)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++17 -ftemplate-backtrace-limit=0 -fmacro-backtrace-limit=0 -ftemplate-depth-1024")

INCLUDE_DIRECTORIES(
    /usr/local/include
    ${PROJECT_SOURCE_DIR}/../../ncraft/ncraft-predefine/include
    ${PROJECT_SOURCE_DIR}/../../ncraft/ncraft-meta/include
    ${PROJECT_SOURCE_DIR}/../../ncraft/ncraft-traits/include
    ${PROJECT_SOURCE_DIR}/../../ncraft/ncraft-core/include
    ${PROJECT_SOURCE_DIR}/../../ncraft/ncraft-data/include
    ${PROJECT_SOURCE_DIR}/../lang/include
    ${PROJECT_SOURCE_DIR}/../core/include

    ${PROJECT_SOURCE_DIR}/include
    ${PROJECT_SOURCE_DIR}/test
    ${PROJECT_SOURCE_DIR}/thirdparty/PEGTL
)

#AUX_SOURCE_DIRECTORY(TEST_FILES test/mojo)
set(SOURCE_FILES
        include/mojo/grammar/rules.hpp
        src/mojo/parser/term_types.cpp
        include/mojo/parser/ast/attributes_parser.hpp
        include/mojo/parser/ast/attribute_parser.hpp
        include/mojo/parser/ast/parser.hpp
        include/mojo/parser/ast/nominal_type_parser.hpp
        include/mojo/parser/ast/type_inheritance_parser.hpp
        include/mojo/parser/ast/enum_decl_parser.hpp
        include/mojo/parser/ast/value_decl_parser.hpp
        include/mojo/parser/ast/struct_field_parser.hpp
        include/mojo/parser/ast/const_decl_parser.hpp
        include/mojo/parser/ast/struct_field_group_parser.hpp
)

AUX_SOURCE_DIRECTORY(test/mojo MOJO_TEST_SRCS)
AUX_SOURCE_DIRECTORY(test/mojo/grammar MOJO_GRAMMAR_TEST_SRCS)
AUX_SOURCE_DIRECTORY(test/mojo/parser MOJO_PARSER_TEST_SRCS)
AUX_SOURCE_DIRECTORY(test/mojo/parser/declaration MOJO_PARSER_DECLARATION_TEST_SRCS)
AUX_SOURCE_DIRECTORY(test/mojo/parser/expression MOJO_PARSER_EXPRESSION_TEST_SRCS)
AUX_SOURCE_DIRECTORY(test/mojo/parser/statement MOJO_PARSER_STATEMENT_TEST_SRCS)

set(JOSON_FILES
    thirdparty/PEGTL/examples/json_build_one.cc)

include(CMakeToolsHelpers OPTIONAL)
add_executable(mojo ${SOURCE_FILES} src/mojo.cpp include/mojo/parser/ast/path_identifier_parser.hpp include/mojo/parser/ast/generic_path_identifier_parser.hpp)
add_executable(mojo_test
        ${SOURCE_FILES}
        ${MOJO_TEST_SRCS}
        ${MOJO_GRAMMAR_TEST_SRCS}
        ${MOJO_PARSER_TEST_SRCS}
        ${MOJO_PARSER_DECLARATION_TEST_SRCS}
        ${MOJO_PARSER_EXPRESSION_TEST_SRCS}
        ${MOJO_PARSER_STATEMENT_TEST_SRCS})
target_link_libraries(mojo_test boost_system boost_filesystem)
#add_executable(json_build ${JOSON_FILES})

#add_executable(mojo_lang_test
#        ${LANG_SOURCE_FILES})
#target_link_libraries(mojo_lang_test boost_system boost_filesystem)