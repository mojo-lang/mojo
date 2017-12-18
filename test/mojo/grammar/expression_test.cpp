#include <chrono>
#include <iostream>
#include <mojo/mojo_test.hpp>

namespace {

using namespace mojo;

template <typename Rule>
struct action : pegtl::nothing<Rule> {};

template <>
struct action<grammar::expression> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

TEST_CASE("expression", "[expression]") {
    simple_success_check<grammar::expression, action>("a + b");
    simple_success_check<grammar::expression, action>("!a");
    simple_success_check<grammar::expression, action>("a += b");
    simple_success_check<grammar::expression, action>("a -= b");
    simple_success_check<grammar::expression, action>("a -= 103");
    simple_success_check<grammar::expression, action>("a .. 103");
    simple_success_check<grammar::expression, action>("a ..< 103");
    simple_success_check<grammar::expression, action>("a <..< 103");
    simple_success_check<grammar::expression, action>("a <.. 103");
    simple_success_check<grammar::expression, action>("中文+english");
    simple_success_check<grammar::expression, action>("a+b*34 | 33");
}

TEST_CASE("literal-expression", "[expression, literal]") {
    simple_success_check<grammar::expression, action>("[a, b, 12, 33]");
    simple_success_check<grammar::expression, action>("[34, 34, 34, 444]");
    simple_success_check<grammar::expression, action>("[34, 34, 34, 444,]");
    simple_success_check<grammar::expression, action>("[34\n 34\n 34\n 444]");
    simple_success_check<grammar::expression, action>("[34\n \n \n, 34\n , 34, \n 444]");
    simple_success_check<grammar::expression, action>("[34\n 34\n 34\n 444\n\n  ,]");
    simple_success_check<grammar::expression, action>("[]");
    simple_success_check<grammar::expression, action>("['str1', 'string2', 'string3']");
    simple_success_check<grammar::expression, action>("['str1', 'string2', 'string3']");
    simple_success_check<grammar::expression, action>("island | category == 999 | zoom >= 1");
    simple_success_check<grammar::expression, action>("island | category == 999 | zoom in 5..10");
    simple_success_check<grammar::expression, action>("rgb(244, 243, 239)");
    simple_success_check<grammar::expression, action>("{fill.color: rgb(244, 243, 239)}");
    simple_success_check<grammar::expression, action>("{fill, color : rgb(244, 243, 239)}");
    simple_success_check<grammar::expression, action>(
        "attach_style({fill.color: rgb(244, 243, 239),"
        "fill.outline.color: rgb(246, 242, 239)})");

    const char* object_str =
        "{"
        "id: 'MapEMG',"
        "name: 'Map for EMG',"
        "sources: [{ id: '4445', host: 'localhost'}],"
        "layers: ["
        "   island | category == 999 | zoom in 1..16 | attach_style({"
        "fill.color: rgb(244, 243, 239),"
        "fill.outline.color: rgb(246, 242, 239)})"
        "]"
        "}";

    simple_success_check<grammar::expression, action>(object_str);
}

TEST_CASE("postfix-expression", "[expression]") {
    simple_success_check<grammar::expression, action>("'test'.upper");
    simple_success_check<grammar::expression, action>("Int('test')");
    simple_success_check<grammar::expression, action>("Foo{foo:'test'}");
    simple_success_check<grammar::expression, action>("foo({bar:'test'})");
    simple_success_check<grammar::expression, action>("bar.Baz.Foo{foo:'test'}");
    simple_success_check<grammar::expression, action>("ab()");
    simple_success_check<grammar::expression, action>("'123'.to<Int>().to<String>().bar['t']");
    simple_success_check<grammar::expression, action>("'123'.to<Int>.to<String>");
    simple_success_check<grammar::expression, action>("'123'.trim.to<Int>()");
    simple_success_check<grammar::expression, action>("{bar:$ + 'test'}()");
    simple_success_check<grammar::expression, action>("54.minute");
    simple_success_check<grammar::expression, action>("$1 + $2");
}

}
