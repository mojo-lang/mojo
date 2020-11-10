#include <mojo/mojo_test.hpp>

namespace {
using namespace mojo;

template <typename Rule>
struct action : pegtl::nothing<Rule> {};

template <>
struct action<grammar::statement> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

template <>
struct action<grammar::statements> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

template <>
struct action<grammar::code_block> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

TEST_CASE("statement_test", "[statement, declaration]") {
    simple_success_check<grammar::statement, action>("foo_bar");
    simple_success_check<grammar::code_block, action>("{foo_bar}");
    simple_success_check<grammar::code_block, action>("{foo * bar}");
    simple_success_check<grammar::code_block, action>("{foo * bar \n foo + bar}");
}

TEST_CASE("statement_for_test", "[statement, declaration]") {
    //trace_success_check<grammar::statement, action>("for a in 1..3 { print('hello') }");
    simple_success_check<grammar::statement, action>("for foo in bar { foo += 3 }");
}

TEST_CASE("statement_while_test", "[statement, declaration]") {
    simple_success_check<grammar::statements, action>("while foo <= 100, bar > 0 { sum += foo \n sum += bar }");
    simple_success_check<grammar::statements, action>("repeat { \n   sum += foo \n   sum += bar } \nwhile foo <= 100");
    simple_success_check<grammar::statements, action>("foo:Int = 1\n bar:Int = 2\n sum:Int = 0\n repeat { \n   sum += foo \n   sum += bar } \nwhile foo <= 100");
}

TEST_CASE("statement_if_test", "[statement, declaration]") {
    simple_success_check<grammar::statements, action>("if foo > 0 { bar += 1 }");
    simple_success_check<grammar::statements, action>("if foo > 0 { bar += 1 } else { bar - 1}");
    simple_success_check<grammar::statements, action>("if foo > 0 { bar += 1 } else if foo < -5 { bar - 1 } else { bar - 2 }");
}

TEST_CASE("statement_match_test", "[statement, declaration]") {
    simple_success_check<grammar::statements, action>("match foo { is Int => foo + 1\n is String => Int(foo) + 2\n _ => print(foo) }");
}
}