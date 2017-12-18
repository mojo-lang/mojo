#include <chrono>
#include <iostream>
#include <mojo/mojo_test.hpp>

namespace {
using namespace mojo;

TEST_CASE("array_parser", "[expression]") {
    simple_success_parse<grammar::expression>("['12', '13', '34']");
    simple_success_parse<grammar::expression>("['12'\n '13'\n '34']");
    simple_success_parse<grammar::expression>("[12, 13, 34]");
    simple_success_parse<grammar::expression>("[]");
}
}
