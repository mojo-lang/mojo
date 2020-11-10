#include <chrono>
#include <iostream>
#include <mojo/mojo_test.hpp>

namespace {
using namespace mojo;

TEST_CASE("interface_declaration_parser", "[declaration]") {
    const char* interface_1 =
        "/// this is \n"
        "/// comment \n"
        "@api('testt')\n"
        "interface FooService {\n"
        "  /// this is \n"
        "  /// comment \n"
        "  @path('get')\n"
        "  echo(str:String)->String"
        "}";

    simple_success_parse<grammar::grammar>(interface_1);
}
}



