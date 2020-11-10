#ifndef MOJO_PARSER_DECLARATION_INTERFACE_DECLARATION_HPP
#define MOJO_PARSER_DECLARATION_INTERFACE_DECLARATION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace declaration {

struct interface_state : array_term_state {
    interface_state() : array_term_state("interface_declaration") {
    }
};

struct interface_members_state : array_term_state {
    interface_members_state() : array_term_state("interface_members") {
    }
};

template <typename Rule>
struct interface_action : pegtl::nothing<Rule> {};

template <>
struct interface_action<grammar::type_inheritance_clause::predication> {
    template <typename Input>
    static void apply(const Input&, interface_state& state) {
        state.push_element();
    }
};

template <>
struct interface_action<grammar::interface_declaration_clause::members_begin> {
    template <typename Input>
    static void apply(const Input&, interface_state& state) {
        state.push_element();
    }
};
}

template <>
struct control<grammar::interface_declaration_clause>
    : pegtl::change_state_and_action<grammar::interface_declaration_clause,
                                     declaration::interface_state,
                                     declaration::interface_action,
                                     errors> {};

template <>
struct control<grammar::interface_declaration_clause::members>
    : pegtl::change_state_and_action<grammar::interface_declaration_clause::members,
                                     declaration::interface_members_state,
                                     array_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_DECLARATION_INTERFACE_DECLARATION_HPP
