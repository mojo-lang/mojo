#include <mojo/mojo_test.hpp>

namespace {

using namespace mojo;

template <typename Rule>
struct action : pegtl::nothing<Rule> {};

template <>
struct action<grammar::type> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

TEST_CASE("type_test", "[type]") {
    simple_success_parse<grammar::type>(
        "FooBar",
        make_term("type_identifier",
                  {make_term("generic_type_name", {make_term("type_name", "FooBar")})}));

    simple_success_parse<grammar::type>(
        "Foo.Bar",
        make_term("type_identifier",
                  {make_term("generic_type_name", {make_term("type_name", "Foo")}),
                   make_term("generic_type_name", {make_term("type_name", "Bar")})}));

    simple_success_parse<grammar::type>(
        "foo.Bar",
        make_term("type_identifier",
                  {make_term("package_identifier", {make_term("identifier", "foo")}),
                   make_term("generic_type_name", {make_term("type_name", "Bar")})}));

    simple_success_check<grammar::type, action>("Foo<B@1>");
    simple_success_check<grammar::type, action>("Foo<B>.Bar<Foobar>");
    simple_success_check<grammar::type, action>("[Foo]");
    simple_success_check<grammar::type, action>("{Foo: Bar}");
    simple_success_check<grammar::type, action>("(Foo, Bar)");
    simple_success_check<grammar::type, action>("(Foo@foo @cc(45) @1, Bar@bar @3)");
    simple_success_check<grammar::type, action>("(foo:Foo, bar:Bar)");
    simple_success_check<grammar::type, action>("(Foo, Bar)->Foobar");
    simple_success_check<grammar::type, action>("(Foo, Bar -> Result )->Foobar");
    simple_success_check<grammar::type, action>("()->Foobar");
    simple_success_check<grammar::type, action>("Foo->Foobar");
    simple_success_check<grammar::type, action>("Foo->Bar->Baz");
}
}
