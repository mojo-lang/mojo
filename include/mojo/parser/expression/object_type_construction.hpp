#ifndef MOJO_PARSER_EXPRESSION_OBJECT_TYPE_CONSTRUCTION_HPP
#define MOJO_PARSER_EXPRESSION_OBJECT_TYPE_CONSTRUCTION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace expression {

struct object_type_construction_state : array_term_state {
    object_type_construction_state() : array_term_state("object_type_construction") {
    }
};

template <typename Rule>
struct object_type_construction_action : pegtl::nothing<Rule> {};

template <>
struct object_type_construction_action<grammar::object_literal::begin> {
    template <typename Input>
    static void apply(const Input&, object_type_construction_state& state) {
        state.push_element();
    }
};
}

template <>
struct control<grammar::object_type_construction>
    : pegtl::change_state_and_action<grammar::object_type_construction,
                                     expression::object_type_construction_state,
                                     expression::object_type_construction_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_EXPRESSION_OBJECT_TYPE_CONSTRUCTION_HPP
