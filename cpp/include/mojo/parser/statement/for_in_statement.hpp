#ifndef MOJO_PARSER_STATEMENT_FOR_IN_STATEMENT_HPP
#define MOJO_PARSER_STATEMENT_FOR_IN_STATEMENT_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace statement {

struct for_in_state : array_term_state {
    for_in_state() : array_term_state("for_in_statement") {
    }
};

template <typename Rule>
struct for_in_action : pegtl::nothing<Rule> {};

template <>
struct for_in_action<grammar::key_in> {
    template <typename Input>
    static void apply(const Input&, for_in_state& state) {
        state.push_element();
    }
};

template <>
struct for_in_action<grammar::code_block::begin> {
    template <typename Input>
    static void apply(const Input&, for_in_state& state) {
        state.push_element();
    }
};
}
template <>
struct control<grammar::for_in_statement_clause>
    : pegtl::change_state_and_action<grammar::for_in_statement_clause,
                                     statement::for_in_state,
                                     statement::for_in_action,
                                     errors> {};
}
}
#endif  // MOJO_PARSER_STATEMENT_FOR_IN_STATEMENT_HPP
