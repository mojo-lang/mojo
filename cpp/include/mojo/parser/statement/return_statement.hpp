#ifndef MOJO_PARSER_STATEMENT_RETURN_STATEMENT_HPP
#define MOJO_PARSER_STATEMENT_RETURN_STATEMENT_HPP
#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace statement {

struct return_state : array_term_state {
    return_state() : array_term_state("return_statement") {
    }
};
}
template <>
struct control<grammar::return_statement_clause>
    : pegtl::change_state_and_action<grammar::return_statement_clause,
                                     statement::return_state,
                                     array_action,
                                     errors> {};
}
}
#endif  // MOJO_PARSER_STATEMENT_RETURN_STATEMENT_HPP
