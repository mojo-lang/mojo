#ifndef MOJO_PARSER_DECLARATION_INTERFACE_DECLARATION_HPP
#define MOJO_PARSER_DECLARATION_INTERFACE_DECLARATION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace declaration {

struct service_state : array_term_state {
    service_state() : array_term_state("service_declaration") {
    }
};

struct service_members_state : array_term_state {
    service_members_state() : array_term_state("service_members") {
    }
};

template <typename Rule>
struct service_action : pegtl::nothing<Rule> {};

template <>
struct service_action<grammar::type_inheritance_clause::predication> {
    template <typename Input>
    static void apply(const Input&, service_state& state) {
        state.push_element();
    }
};

template <>
struct service_action<grammar::service_declaration_clause::members_begin> {
    template <typename Input>
    static void apply(const Input&, service_state& state) {
        state.push_element();
    }
};
}

template <>
struct control<grammar::service_declaration_clause>
    : pegtl::change_state_and_action<grammar::service_declaration_clause,
                                     declaration::service_state,
                                     declaration::service_action,
                                     errors> {};

template <>
struct control<grammar::service_declaration_clause::members>
    : pegtl::change_state_and_action<grammar::service_declaration_clause::members,
                                     declaration::service_members_state,
                                     array_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_DECLARATION_INTERFACE_DECLARATION_HPP
