#ifndef MOJO_PARSER_EXPRESSION_ARRAY_LITERAL_HPP
#define MOJO_PARSER_EXPRESSION_ARRAY_LITERAL_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace expression {

struct array_state : array_term_state {
    array_state() : array_term_state("array") {
    }
};
}
template <>
struct control<grammar::array_literal_content>
    : pegtl::change_state_and_action<grammar::array_literal_content,
                                     expression::array_state,
                                     array_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_EXPRESSION_ARRAY_LITERAL_HPP