#include <pegtl.hh>
#include <pegtl/trace.hh>
#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/parser.hpp>
#include <catch.hpp>

template <typename Rule, template <typename> class Action>
inline void simple_success_check(const char* str) {
    std::string name;
    pegtl::parse_string<Rule, Action>(str, "mojo", name);
    REQUIRE(name == str);
}

template <typename Rule, template <typename> class Action>
inline void trace_success_check(const char* str) {
    std::string name;
    pegtl::trace_string<Rule, Action>(str, "mojo", name);
    REQUIRE(name == str);
}

template <typename Rule, template <typename> class Action>
inline void simple_failed_check(const char* str) {
    std::string name;
    bool ret = pegtl::parse_string<Rule, Action>(str, "mojo", name);
    REQUIRE((!ret || name != str));
}

template <typename Rule, template <typename> class Action>
inline void trace_failed_check(const char* str) {
    std::string name;
    bool ret = pegtl::trace_string<Rule, Action>(str, "mojo", name);
    REQUIRE((!ret || name != str));
}

template <typename Rule>
inline void simple_success_parse(const char* str, mojo::TermPtr expected = mojo::TermPtr()) {
    mojo::parser::term_state state;
    state.term.reset();
    pegtl::parse_string<Rule, pegtl::nothing, mojo::parser::control>(str, "mojo", state);

    if (expected) {
        REQUIRE((state.term && *state.term == *expected));
    }
    else {
        if (state.term) {
            std::cout << *state.term << std::endl;
        }
        else {
            std::cout << "failed to parse: " << str << std::endl;
        }
    }
}