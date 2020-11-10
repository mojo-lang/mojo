#ifndef MOJO_PARSER_DECLARATION_VARIABLE_DECLARATION_HPP
#define MOJO_PARSER_DECLARATION_VARIABLE_DECLARATION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace declaration {

struct variable_state : array_term_state {
    variable_state() : array_term_state("variable_declaration") {
    }
};

template <typename Rule>
struct variable_action : pegtl::nothing<Rule> {};

template <>
struct variable_action<grammar::type_annotation::predication> {
    template <typename Input>
    static void apply(const Input&, variable_state& state) {
        state.push_element();
    }
};

template <>
struct variable_action<grammar::initializer::predication> {
    template <typename Input>
    static void apply(const Input&, variable_state& state) {
        state.push_element();
    }
};

}
template <>
struct control<grammar::variable_declaration>
    : pegtl::change_state_and_action<grammar::variable_declaration,
                                     declaration::variable_state,
                                     declaration::variable_action,
                                     errors> {};

}
}

#endif //MOJO_PARSER_DECLARATION_VARIABLE_DECLARATION_HPP
