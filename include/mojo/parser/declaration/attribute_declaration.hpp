#ifndef MOJO_PARSER_DECLARATION_ATTRIBUTE_DECLARATION_HPP
#define MOJO_PARSER_DECLARATION_ATTRIBUTE_DECLARATION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace declaration {

struct attribute_state : array_term_state {
    attribute_state() : array_term_state("attribute_declaration") {
    }
};

template <typename Rule>
struct attribute_action : pegtl::nothing<Rule> {};

template <>
struct attribute_action<grammar::type_annotation::predication> {
    template <typename Input>
    static void apply(const Input&, attribute_state& state) {
        state.push_element();
    }
};

template <>
struct attribute_action<grammar::initializer::predication> {
    template <typename Input>
    static void apply(const Input&, attribute_state& state) {
        state.push_element();
    }
};
}
template <>
struct control<grammar::attribute_declaration_clause>
    : pegtl::change_state_and_action<grammar::attribute_declaration_clause,
                                     declaration::attribute_state,
                                     declaration::attribute_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_DECLARATION_ATTRIBUTE_DECLARATION_HPP
