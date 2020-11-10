#ifndef MOJO_PARSER_DECLARATION_DECLARATIONS_HPP
#define MOJO_PARSER_DECLARATION_DECLARATIONS_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace declaration {

struct initializer_state : array_term_state {
    initializer_state() : array_term_state("initializer") {
    }
};
}
template <>
struct control<grammar::initializer::content>
    : pegtl::change_state_and_action<grammar::initializer::content,
                                     declaration::initializer_state,
                                     array_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_DECLARATION_DECLARATIONS_HPP
