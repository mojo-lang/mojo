#ifndef MOJO_PARSER_DECLARATION_CONST_DECLARATION_HPP
#define MOJO_PARSER_DECLARATION_CONST_DECLARATION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace declaration {

struct constant_state : array_term_state {
    constant_state() : array_term_state("const_declaration") {
    }
};

template <typename Rule>
struct constant_action : pegtl::nothing<Rule> {};

template <>
struct constant_action<grammar::type_annotation::predication> {
    template <typename Input>
    static void apply(const Input&, constant_state& state) {
        state.push_element();
    }
};

template <>
struct constant_action<grammar::initializer::predication> {
    template <typename Input>
    static void apply(const Input&, constant_state& state) {
        state.push_element();
    }
};

}
template <>
struct control<grammar::constant_declaration>
    : pegtl::change_state_and_action<grammar::constant_declaration,
                                     declaration::constant_state,
                                     declaration::constant_action,
                                     errors> {};
}
}

#endif //MOJO_PARSER_DECLARATION_CONST_DECLARATION_HPP
