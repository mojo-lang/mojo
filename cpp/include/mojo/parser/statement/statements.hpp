#ifndef MOJO_PARSER_STATEMENT_STATEMENTS_HPP
#define MOJO_PARSER_STATEMENT_STATEMENTS_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/errors.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace statement {

struct statements_state : array_term_state {
    statements_state() : array_term_state("statements") {
    }
};

template <typename Rule>
struct statements_action : pegtl::nothing<Rule> {};

template <>
struct statements_action<grammar::statement_separator> {
    template <typename Input>
    static void apply(const Input&, statements_state& state) {
        state.push_element();
    }
};
}

template <>
struct control<grammar::statements>
    : pegtl::change_state_and_action<grammar::statements,
                                     statement::statements_state,
                                     statement::statements_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_STATEMENT_STATEMENTS_HPP
