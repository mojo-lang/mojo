#include <mojo/mojo_test.hpp>

namespace {

using namespace mojo;

template <typename Rule>
struct action : pegtl::nothing<Rule> {};

template <>
struct action<grammar::pattern> {
    static void apply(const pegtl::action_input& in, std::string& name) {
        name = in.string();
    }
};

TEST_CASE("pattern_test", "[pattern]") {
    simple_success_check<grammar::pattern, action>("_");
    simple_success_check<grammar::pattern, action>("_:Int");
    simple_success_check<grammar::pattern, action>("_ : String");
    simple_success_check<grammar::pattern, action>("str");
    simple_success_check<grammar::pattern, action>("str:String");
    simple_success_check<grammar::pattern, action>("(_, int)");
    simple_success_check<grammar::pattern, action>("(_, int):(Int, Int)");
    simple_success_check<grammar::pattern, action>("(_, int:int1)");
    simple_success_check<grammar::pattern, action>("Foo.bar");
}
}
