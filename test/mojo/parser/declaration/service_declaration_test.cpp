#include <chrono>
#include <iostream>
#include <mojo/mojo_test.hpp>

namespace {
using namespace mojo;

TEST_CASE("service_declaration_parser", "[declaration]") {
    const char* service_1 =
        "/// this is \n"
        "/// comment \n"
        "@api('testt')\n"
        "service FooService {\n"
        "  /// this is \n"
        "  /// comment \n"
        "  @path('get')\n"
        "  echo(str:String)->String"
        "}";

    simple_success_parse<grammar::grammar>(service_1);
}
}



