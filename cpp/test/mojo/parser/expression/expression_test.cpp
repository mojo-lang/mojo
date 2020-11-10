#include <chrono>
#include <iostream>
#include <mojo/mojo_test.hpp>

namespace {
using namespace mojo;

TEST_CASE("expression_parser", "[expression]") {
    const char* expr_str_1 = "island | category == 999 | zoom in 1..16 | attach_style";
    simple_success_parse<grammar::expression>(expr_str_1);
}
}
