#include <chrono>
#include <iostream>
#include <mojo/mojo_test.hpp>

namespace {
using namespace mojo;

TEST_CASE("pattern_parser", "[pattern, parser]") {
    simple_success_parse<grammar::pattern>("_");
    simple_success_parse<grammar::pattern>("_:Int");
    simple_success_parse<grammar::pattern>("_ : String");
    simple_success_parse<grammar::pattern>("str");
    simple_success_parse<grammar::pattern>("str:String");
    simple_success_parse<grammar::pattern>("(_, int)");
    simple_success_parse<grammar::pattern>("(_, int):(Int, Int)");
    simple_success_parse<grammar::pattern>("(_, int: int1)");
}
}
