#include <mojo/mojo_test.hpp>

namespace {

using namespace mojo;

template <typename Rule>
struct action : pegtl::nothing<Rule> {};

template <>
struct action<grammar::identifier> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

TEST_CASE("identifier_test", "[lexical, identifier]") {
    simple_failed_check<grammar::identifier, action>("*");
    simple_success_check<grammar::identifier, action>("foobar");
    simple_success_check<grammar::identifier, action>("foo_bar");
    simple_success_check<grammar::identifier, action>("foo_中文");
    simple_success_check<grammar::identifier, action>("都是中文");
    simple_success_check<grammar::identifier, action>("package");
    simple_success_check<grammar::identifier, action>("attribute");
    simple_success_parse<grammar::identifier>("$", make_term("identifier", "$0"));
    simple_success_parse<grammar::identifier>("$0", make_term("identifier", "$0"));
}
}
