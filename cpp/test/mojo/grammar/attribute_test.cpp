#include <mojo/mojo_test.hpp>
#include <mojo/grammar/grammar.hpp>

namespace {

using namespace mojo;

template <typename Rule>
struct action : pegtl::nothing<Rule> {};

template <>
struct action<grammar::attribute> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

template <>
struct action<grammar::attributes> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

TEST_CASE("grammar_attribute_test", "[attribute]") {
    simple_success_check<grammar::attribute, action>("@1");
    simple_success_check<grammar::attribute, action>("@foo");
    simple_success_check<grammar::attribute, action>("@foo(true)");
    simple_success_check<grammar::attribute, action>("@foo({1:23,3:34})");
    simple_success_check<grammar::attributes, action>("@foo(12)@bar('456777')");
    simple_success_check<grammar::attributes, action>("@foo(12)    @bar('456777')");
    simple_success_check<grammar::attributes, action>("@foo(12)\n@bar('456777')");
    simple_success_check<grammar::attributes, action>("@foo\n@2\n    @bar('456777')");
}

TEST_CASE("grammar_group_attribute_test", "[attribute]") {
    simple_success_check<grammar::attributes, action>("@{foo:true,bar:'3456'}");
    simple_success_check<grammar::attributes, action>("@foo({1:23,3:34})");
}

TEST_CASE("grammar_attribute_sep_test", "[attribute]") {
    simple_failed_check<grammar::attributes, action>("@foo\n@2//comment...\n\n   @bar('456777')");
}
}
