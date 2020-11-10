#ifndef MOJO_PARSER_EXPRESSION_PARENTHESIZED_EXPRESSION_HPP
#define MOJO_PARSER_EXPRESSION_PARENTHESIZED_EXPRESSION_HPP

#include <iostream>
#include <mojo/grammar/grammar.hpp>

namespace mojo {
namespace parser {
namespace expression {

struct parenthesized_expression_state : array_term_state {
    parenthesized_expression_state() : array_term_state("parenthesized_expression") {
    }
};

struct expression_element_state : array_term_state {
    expression_element_state() : array_term_state("expression_element") {
    }
};

template <typename Rule>
struct expression_element_action : pegtl::nothing<Rule> {};

template <>
struct expression_element_action<grammar::expression_element::separator> {
    template <typename Input>
    static void apply(const Input&, expression_element_state& state) {
        state.push_element();
    }
};
}

template <>
struct control<grammar::expression_element>
    : pegtl::change_state_and_action<grammar::expression_element,
                                     expression::expression_element_state,
                                     expression::expression_element_action,
                                     errors> {};

template <>
struct control<grammar::expression_element_list>
    : pegtl::change_state_and_action<grammar::expression_element_list,
                                     expression::parenthesized_expression_state,
                                     array_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_EXPRESSION_PARENTHESIZED_EXPRESSION_HPP