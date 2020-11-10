#include <chrono>
#include <iostream>
#include <mojo/mojo_test.hpp>

namespace {
using namespace mojo;

TEST_CASE("enum_declaration_parser", "[declaration]") {
    simple_success_parse<grammar::declaration>("enum Foo { bar1, bar2}");
    simple_success_parse<grammar::declaration>("enum Foo : Int { bar1 = 2, bar2 = 4}");
    simple_success_parse<grammar::declaration>("enum Foo : Int {\n bar1 = 1\n bar2\n }");
}
}

