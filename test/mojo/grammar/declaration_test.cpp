#include <mojo/mojo_test.hpp>

namespace {

using namespace mojo;

template <typename Rule>
struct action : pegtl::nothing<Rule> {};

template <>
struct action<grammar::declaration> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

template <>
struct action<grammar::declarations> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

TEST_CASE("declaration_test", "[declaration]") {
    simple_success_check<grammar::declaration, action>("package foo.bar {name: 'bar', version: '1.1.2' }");
    simple_success_check<grammar::declaration, action>("import f.*");
    simple_success_check<grammar::declaration, action>("import foobar as fb");
    simple_success_check<grammar::declaration, action>("import foo.bar as fb");
    simple_success_check<grammar::declaration, action>("import foo.bar.Foo, Bar, convert, to");
    simple_success_check<grammar::declaration, action>("import foo.bar.Foo as FbFoo, Bar as FbBar, convert, to as fb_to");

    simple_success_check<grammar::declaration, action>("var a = 0");
    simple_success_check<grammar::declaration, action>("a: Int = 0");
    simple_success_check<grammar::declaration, action>("var a : Int = 0");

    simple_success_check<grammar::declaration, action>("enum Foo { bar1, bar2 }");
    simple_success_check<grammar::declaration, action>("enum Foo : Int { bar1 = 2, bar2 = 4}");
    simple_success_check<grammar::declaration, action>("enum Foo : Int {\n bar1 = 1  \n  bar2  \n  }");

    simple_success_check<grammar::declaration, action>("@ff(true)\nfunc foo()->Int {\n var bar1 = 1  \n  return bar2  \n  }");

    simple_success_check<grammar::declaration, action>("Foo(a:Int)");
    simple_success_check<grammar::declaration, action>("Foo(a:Int) {\n var bar1 = 1  \n  return bar2  \n }");

    const char* type_1 = "type Foo {\n"
            "  type Bar {\n"
            "      foo: Int @1 = 2\n"
            "      bar: String @2 = 'test'\n"
            "  }\n"
            "  enum Enum {\n"
            "      a, b, c, d\n"
            "  }\n"
            "  foo: String @1 = 'foo'\n"
            "  bar: {Int: String} @2\n"
            "}";
    simple_success_check<grammar::declaration, action>(type_1);

    const char* type_2 = "type Foo<Bar> : Object {\n"
                        " type Bar {\n"
                        "       foo: Int @1 = 3\n"
                        " }\n"
                        " foo: String @1 = 'foo'\n"
                        " @skip(true) {\n"
                        "   bar: {Int: String} @2\n"
                        " }\n"
                        " }";

    simple_success_check<grammar::declaration, action>(type_2);

    const char* service_1 =
            "/// that is \n"
            "/// comment \n"
            "@api('testt')\n"
            "service FooService {\n"
            "  /// this is \n"
            "  /// comment"
            "  @path('get')\n"
            "  echo(str:String)->String"
            "}";

    simple_success_check<grammar::declaration, action>(service_1);

    const char* service_2 =
            "service FooService {\n"
            "  /// this is \n"
            "  /// comment\n"
            "  @path('get')\n"
            "  echo(str:String)->String"
            "}";
    simple_success_check<grammar::declaration, action>(service_2);

    simple_success_check<grammar::grammar, action>("type Highway {\n\tnumber: String @1 //< 道路国家标准编号\n\tstart : String @2\n\tend   : String @3\n}");
}
}
