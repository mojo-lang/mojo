#ifndef MOJO_PARSER_EXPRESSION_OBJECT_HPP
#define MOJO_PARSER_EXPRESSION_OBJECT_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace expression {

struct object_state : dictionary_term_state {
    object_state() : dictionary_term_state("object_literal", "object_field") {
    }
};

template <typename Rule>
struct object_action : pegtl::nothing<Rule> {};

template <>
struct object_action<grammar::object_literal_value_separator> {
    template <typename Input>
    static void apply(const Input&, object_state& state) {
        state.add_key();
    }
};

template <>
struct object_action<grammar::object_literal_item> {
    template <typename Input>
    static void apply(const Input&, object_state& state) {
        state.set_value();
    }
};
}

template <>
struct control<grammar::object_literal_content>
    : pegtl::change_state_and_action<grammar::object_literal_content,
                                     expression::object_state,
                                     expression::object_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_EXPRESSION_OBJECT_HPP