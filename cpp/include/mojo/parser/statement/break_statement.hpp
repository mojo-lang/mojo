#ifndef MOJO_PARSER_STATEMENT_BREAK_STATEMENT_HPP
#define MOJO_PARSER_STATEMENT_BREAK_STATEMENT_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace statement {

template <typename Rule>
struct break_action : pegtl::nothing<Rule> {};

template <>
struct break_action<grammar::break_statement> {
    template <typename Input>
    static void apply(const Input&, term_state& state) {
        state.term = make_term("break_statement");
    }
};
}

template <>
struct control<grammar::break_statement>
    : pegtl::change_action<grammar::break_statement, statement::break_action, errors> {};
}
}

#endif  // MOJO_PARSER_STATEMENT_BREAK_STATEMENT_HPP
