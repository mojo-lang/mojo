#ifndef MOJO_PARSER_STATEMENT_IF_STATEMENT_HPP
#define MOJO_PARSER_STATEMENT_IF_STATEMENT_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace statement {

struct if_state : array_term_state {
    if_state() : array_term_state("if_statement") {
    }
};

struct else_state : array_term_state {
    else_state() : array_term_state("else_statement") {
    }
};

template <typename Rule>
struct if_action : pegtl::nothing<Rule> {};

template <>
struct if_action<grammar::code_block::begin> {
    template <typename Input>
    static void apply(const Input&, if_state& state) {
        state.push_element();
    }
};

template <>
struct if_action<grammar::key_else> {
    template <typename Input>
    static void apply(const Input&, if_state& state) {
        state.push_element();
    }
};
}
template <>
struct control<grammar::if_statement_clause>
    : pegtl::change_state_and_action<grammar::if_statement_clause,
                                     statement::if_state,
                                     statement::if_action,
                                     errors> {};
template <>
struct control<grammar::else_statement_clause>
    : pegtl::change_state_and_action<grammar::else_statement_clause,
                                     statement::if_state,
                                     array_action,
                                     errors> {};
}
}
#endif  // MOJO_PARSER_STATEMENT_IF_STATEMENT_HPP
