#include <mojo/mojo_test.hpp>

namespace {

using namespace mojo;

template <typename Rule>
struct action : pegtl::nothing<Rule> {};

template <>
struct action<grammar::literal> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

TEST_CASE("literal_test", "[lexical, literal]") {
    simple_success_check<grammar::literal, action>("12345");
    simple_success_check<grammar::literal, action>("012345");
    simple_success_check<grammar::literal, action>("0x012345BFFCD");
    simple_success_check<grammar::literal, action>("12345.333");
    simple_success_check<grammar::literal, action>("1.333e5");
    simple_success_check<grammar::literal, action>("\"string\"");
    simple_success_check<grammar::literal, action>("\'string\'");
    simple_success_check<grammar::literal, action>("\'str\\tin\1g\'");
}
}
