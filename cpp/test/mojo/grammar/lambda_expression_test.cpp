#include <chrono>
#include <iostream>
#include <mojo/mojo_test.hpp>

namespace {

using namespace mojo;

template<typename Rule>
struct action : pegtl::nothing<Rule> {
};

template<>
struct action<grammar::expression> {
    static void apply(const pegtl::action_input &in, std::string &name) {
        name = in.string();
    }
};

template<>
struct action<grammar::lambda_expression> {
    static void apply(const pegtl::action_input &in, std::string &name) {
        name = in.string();
    }
};

TEST_CASE("lambda-expression", "[expression]") {
    simple_success_check<grammar::lambda_expression, action>("(a:Int, b:Int)->a*b");
    simple_success_check<grammar::expression, action>("(a:Int, b:Int)->{a*b}");
    simple_success_check<grammar::expression, action>("(a, b)->a*b");
    simple_success_check<grammar::expression, action>("(a, b)->{a*b}");
    simple_success_check<grammar::expression, action>("a->{a*3}");
    simple_success_check<grammar::expression, action>("a->a*3");
}
}
