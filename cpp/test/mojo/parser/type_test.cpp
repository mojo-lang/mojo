#include <chrono>
#include <iostream>
#include <mojo/mojo_test.hpp>

namespace {
using namespace mojo;

TEST_CASE("types_parser", "[type]") {
    simple_success_parse<grammar::type>("Int",
                                         make_term("type_identifier",
                                                   {make_term("generic_type_name",
                                                              {make_term("type_name", "Int")})}));
    simple_success_parse<grammar::type>(
        "f.b.Foo.Bar",
        make_term("type_identifier",
                  {make_term("package_identifier",
                             {make_term("identifier", "f"), make_term("identifier", "b")}),
                   make_term("generic_type_name", {make_term("type_name", "Foo")}),
                   make_term("generic_type_name", {make_term("type_name", "Bar")})}));

    simple_success_parse<grammar::type>("Either<Int, String>");
    simple_success_parse<grammar::type>("Either<Int@1, String@2>");
    simple_success_parse<grammar::type>("[Int]");
    simple_success_parse<grammar::type>("(Int, String, (Int, String))");
    simple_success_parse<grammar::type>("(Int@1, String@2)");
    simple_success_parse<grammar::type>("(int:Int@1, string:String@2)");
    simple_success_parse<grammar::type>("(Int, String) -> String");
    simple_success_parse<grammar::type>("Int -> String -> Int");
}
}
