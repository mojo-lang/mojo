#include <mojo/mojo_test.hpp>

namespace {

using namespace mojo;

template <typename Rule>
struct action : pegtl::nothing<Rule> {};

template <>
struct action<grammar::generic_parameter_clause> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

template <>
struct action<grammar::generic_argument_clause> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

TEST_CASE("generic_parameter_test", "[generic]") {
    simple_success_check<grammar::generic_parameter_clause, action>("<Foo:Bar,Baz>");
}

TEST_CASE("generic_argument_test", "[generic]") {
    simple_success_check<grammar::generic_argument_clause, action>("<Foo, Bar>");
    simple_success_check<grammar::generic_argument_clause, action>("<Foo, Bar<Baz>>");
    simple_success_check<grammar::generic_argument_clause, action>("<Array<Array<Dict<Foo, Bar>>>, Bar<Baz>>");
}
}

