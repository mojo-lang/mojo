#ifndef MOJO_PARSER_STATEMENT_REPEAT_WHILE_STATEMENT_HPP
#define MOJO_PARSER_STATEMENT_REPEAT_WHILE_STATEMENT_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace statement {

struct repeat_while_state : array_term_state {
    repeat_while_state() : array_term_state("repeat_while_statement") {
    }
};

template <typename Rule>
struct repeat_while_action : pegtl::nothing<Rule> {};

template <>
struct repeat_while_action<grammar::key_while> {
    template <typename Input>
    static void apply(const Input&, repeat_while_state& state) {
        state.push_element();
    }
};
}
template <>
struct control<grammar::repeat_while_statement_clause>
    : pegtl::change_state_and_action<grammar::repeat_while_statement_clause,
                                     statement::repeat_while_state,
                                     statement::repeat_while_action,
                                     errors> {};
}
}
#endif  // MOJO_PARSER_STATEMENT_REPEAT_WHILE_STATEMENT_HPP
