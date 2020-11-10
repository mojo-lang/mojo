#include <mojo/mojo_test.hpp>

namespace {

using namespace mojo;

template <typename Rule>
struct action : pegtl::nothing<Rule> {};

template <>
struct action<grammar::document> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

TEST_CASE("document_test", "[lexical, document]") {
    simple_success_parse<grammar::document>("///3333", make_term("document", "3333"));
    simple_success_parse<grammar::following_document>("//< 3333\r\n", make_term("following_document", " 3333"));
    simple_success_parse<grammar::document>("///3333\r\n  ///mojo comment", make_term("document", "3333\nmojo comment"));
}
}
