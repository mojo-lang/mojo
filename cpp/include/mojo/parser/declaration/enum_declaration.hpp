#ifndef MOJO_PARSER_DECLARATION_ENUM_DECLARATION_HPP
#define MOJO_PARSER_DECLARATION_ENUM_DECLARATION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>
#include <mojo/parser/term_types.hpp>

namespace mojo {
namespace parser {
namespace declaration {

struct enum_state : array_term_state {
    enum_state() : array_term_state(kEnumDecl) {
    }
};

struct enum_members_state : array_term_state {
    enum_members_state() : array_term_state(kEnumMembers) {
    }
};

struct enum_member_state : array_term_state {
    enum_member_state() : array_term_state(kEnumMember) {
    }
};

template <typename Rule>
struct enum_action : pegtl::nothing<Rule> {};

template <>
struct enum_action<grammar::document_separator> {
    template <typename Input>
    static void apply(const Input&, enum_member_state& state) {
        state.push_element();
    }
};

template <>
struct enum_action<grammar::attributes::predication> {
    template <typename Input>
    static void apply(const Input&, enum_member_state& state) {
        state.push_element();
    }
};

template <>
struct enum_action<grammar::initializer::predication> {
    template <typename Input>
    static void apply(const Input&, enum_member_state& state) {
        state.push_element();
    }
};

template <>
struct enum_action<grammar::following_document::begin> {
    template <typename Input>
    static void apply(const Input&, enum_member_state& state) {
        state.push_element();
    }
};

template <>
struct enum_action<grammar::type_inheritance_clause::predication> {
    template <typename Input>
    static void apply(const Input&, enum_state& state) {
        state.push_element();
    }
};

template <>
struct enum_action<grammar::enum_declaration_clause::members_begin> {
    template <typename Input>
    static void apply(const Input&, enum_state& state) {
        state.push_element();
    }
};
}
template <>
struct control<grammar::enum_declaration_clause>
    : pegtl::change_state_and_action<grammar::enum_declaration_clause,
                                     declaration::enum_state,
                                     declaration::enum_action,
                                     errors> {};

template <>
struct control<grammar::enum_declaration_clause::members>
    : pegtl::change_state_and_action<grammar::enum_declaration_clause::members,
                                     declaration::enum_members_state,
                                     array_action,
                                     errors> {};

template <>
struct control<grammar::enum_declaration_clause::member>
    : pegtl::change_state_and_action<grammar::enum_declaration_clause::member,
                                     declaration::enum_member_state,
                                     declaration::enum_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_DECLARATION_ENUM_DECLARATION_HPP
