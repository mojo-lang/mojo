#include <chrono>
#include <iostream>
#include <mojo/mojo_test.hpp>

namespace {
using namespace mojo;

TEST_CASE("type_declaration_parser", "[declaration]") {
    const char* type_1 = "type Foo {\n"
                         "  type Bar {\n"
                         "      foo: Int @1 = 2\n"
                         "      bar: String @2 = 'test'\n"
                         "  }\n"
                         "  enum Enum {\n"
                         "      a, b, c, d\n"
                         "  }\n"
                         "  foo: String @1 = 'foo'\n"
                         "  bar: {Int: String} @2\n"
                         "}";

    simple_success_parse<grammar::declaration>("type Foo { foo: Int = 2 , bar: String @1}");
    simple_success_parse<grammar::declaration>(type_1);

    simple_success_parse<grammar::grammar>("type Highway {\n\tnumber: String @1 //< 道路国家标准编号\n\tstart : String @2\n\tend   : String @3\n}");
}
}


