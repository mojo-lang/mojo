#ifndef MOJO_TYPE_CONSTRUCTION_DECLARATION_HPP
#define MOJO_TYPE_CONSTRUCTION_DECLARATION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace declaration {

struct type_construction_state : array_term_state {
    type_construction_state() : array_term_state("type_construction_declaration") {
    }
};

template <typename Rule>
struct type_construction_action : pegtl::nothing<Rule> {};

template <>
struct type_construction_action<grammar::function_parameters_begin> {
    template <typename Input>
    static void apply(const Input&, type_construction_state& state) {
        state.push_element();
    }
};

template <>
struct type_construction_action<grammar::function_parameters_end> {
    template <typename Input>
    static void apply(const Input&, type_construction_state& state) {
        state.push_element();
    }
};

template <>
struct type_construction_action<grammar::code_block::begin> {
    template <typename Input>
    static void apply(const Input&, type_construction_state& state) {
        state.push_element();
    }
};
}

template <>
struct control<grammar::type_construction_declaration>
    : pegtl::change_state_and_action<grammar::type_construction_declaration,
                                     declaration::type_construction_state,
                                     declaration::type_construction_action,
                                     errors> {};
}
}

#endif  // MOJO_TYPE_CONSTRUCTION_DECLARATION_HPP
