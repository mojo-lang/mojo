#include <mojo/mojo_test.hpp>

namespace {

using namespace mojo;

template <typename Rule>
struct action : pegtl::nothing<Rule> {};

template <>
struct action<grammar::comment> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

TEST_CASE("comment_test", "[lexical, comment]") {
    simple_success_check<grammar::comment, action>("/*3333\r\n*/");
}
}
