#ifndef MOJO_PARSER_STATEMENT_WHILE_STATEMENT_HPP
#define MOJO_PARSER_STATEMENT_WHILE_STATEMENT_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace statement {

struct condition_list_state : array_term_state {
    condition_list_state() : array_term_state("condition_list") {
    }
};

struct while_state : array_term_state {
    while_state() : array_term_state("while_statement") {
    }
};

template <typename Rule>
struct while_action : pegtl::nothing<Rule> {};

template <>
struct while_action<grammar::code_block::begin> {
    template <typename Input>
    static void apply(const Input&, while_state& state) {
        state.push_element();
    }
};
}
template <>
struct control<grammar::while_statement_clause>
    : pegtl::change_state_and_action<grammar::while_statement_clause,
                                     statement::while_state,
                                     statement::while_action,
                                     errors> {};

template <>
struct control<grammar::condition_list> : pegtl::change_state_and_action<grammar::condition_list,
                                                                         statement::condition_list_state,
                                                                         array_action,
                                                                         errors> {};
}
}

#endif  // MOJO_PARSER_STATEMENT_WHILE_STATEMENT_HPP
