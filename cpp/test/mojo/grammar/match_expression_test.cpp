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
struct action<grammar::match_expression> {
    static void apply(const pegtl::action_input &in, std::string &name) {
        name = in.string();
    }
};

TEST_CASE("match-expression", "[expression]") {
    simple_success_check<grammar::match_expression, action>("match {a => 23\n b=> 45}");
    simple_success_check<grammar::expression, action>("4.match {a => 23\n b=> 45}");
    simple_success_check<grammar::expression, action>("match {a => 23\n b=> 45}(4).to<String>");
}

}
