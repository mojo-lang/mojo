#ifndef MOJO_PARSER_STATEMENT_CONTINUE_STATEMENT_HPP
#define MOJO_PARSER_STATEMENT_CONTINUE_STATEMENT_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace statement {

template <typename Rule>
struct continue_action : pegtl::nothing<Rule> {};

template <>
struct continue_action<grammar::continue_statement> {
    template <typename Input>
    static void apply(const Input&, term_state& state) {
        state.term = make_term("continue_statement");
    }
};
}

template <>
struct control<grammar::continue_statement>
    : pegtl::change_action<grammar::continue_statement, statement::continue_action, errors> {};
}
}
#endif  // MOJO_PARSER_STATEMENT_CONTINUE_STATEMENT_HPP
