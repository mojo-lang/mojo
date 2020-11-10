#ifndef MOJO_PARSER_EXPRESSION_WILDCARD_EXPRESSION_HPP
#define MOJO_PARSER_EXPRESSION_WILDCARD_EXPRESSION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace expression {

template <typename Rule>
struct wildcard_action : pegtl::nothing<Rule> {};

template <>
struct wildcard_action<grammar::wildcard_expression> {
    template <typename Input>
    static void apply(const Input&, term_state& state) {
        state.term = make_term("wildcard_expression");
    }
};
}

template <>
struct control<grammar::wildcard_expression>
    : pegtl::change_action<grammar::wildcard_expression,
                           expression::wildcard_action,
                           errors> {};
}
}

#endif  // MOJO_PARSER_EXPRESSION_WILDCARD_EXPRESSION_HPP
