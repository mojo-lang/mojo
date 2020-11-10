#ifndef MOJO_PARSER_DECLARATION_DECLARATION_PACKAGE_DECLARATION_HPP
#define MOJO_PARSER_DECLARATION_DECLARATION_PACKAGE_DECLARATION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace declaration {

struct package_state : array_term_state {
    package_state() : array_term_state("package_declaration") {
    }
};

template <typename Rule>
struct package_action : pegtl::nothing<Rule> {};

template <>
struct package_action<grammar::object_literal::begin> {
    template <typename Input>
    static void apply(const Input&, package_state& state) {
        state.push_element();
    }
};

}
template <>
struct control<grammar::package_declaration>
    : pegtl::change_state_and_action<grammar::package_declaration,
                                     declaration::package_state,
                                     declaration::package_action,
                                     errors> {};
}
}
#endif  // MOJO_PARSER_DECLARATION_DECLARATION_PACKAGE_DECLARATION_HPP
