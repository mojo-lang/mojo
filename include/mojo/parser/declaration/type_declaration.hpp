#ifndef MOJO_PARSER_DECLARATION_TYPE_DECLARATION_HPP
#define MOJO_PARSER_DECLARATION_TYPE_DECLARATION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>
#include <mojo/parser/term_types.hpp>

namespace mojo {
namespace parser {
namespace declaration {

struct type_state : array_term_state {
    type_state() : array_term_state(kTypeDecl) {
    }
};

struct type_members_state : array_term_state {
    type_members_state() : array_term_state(kStructMembers) {
    }
};

struct type_member_state : array_term_state {
    type_member_state() : array_term_state(kStructMember) {
    }
};

struct type_member_group_state : array_term_state {
    type_member_group_state() : array_term_state(kStructMemberGroup) {
    }
};

struct type_member_group_members_state : array_term_state {
    type_member_group_members_state() : array_term_state(kStructMembers) {
    }
};

template <typename Rule>
struct type_action : pegtl::nothing<Rule> {};

template <>
struct type_action<grammar::type_annotation::predication> {
    template <typename Input>
    static void apply(const Input&, type_member_state& state) {
        state.push_element();
    }
};

template <>
struct type_action<grammar::initializer::predication> {
    template <typename Input>
    static void apply(const Input&, type_member_state& state) {
        state.push_element();
    }
};

template <>
struct type_action<grammar::following_document::begin> {
    template <typename Input>
    static void apply(const Input&, type_member_state& state) {
        state.push_element();
    }
};

template <>
struct type_action<grammar::generic_parameter_clause::begin> {
    template <typename Input>
    static void apply(const Input&, type_state& state) {
        state.push_element();
    }
};



template <>
struct type_action<grammar::type_alias_clause::predication> {
    template <typename Input>
    static void apply(const Input&, type_state& state) {
        state.push_element();
        state.array->type = kTypeAliasDecl;
    }
};

template <>
struct type_action<grammar::type_inheritance_clause::predication> {
    template <typename Input>
    static void apply(const Input&, type_state& state) {
        state.push_element();
    }
};

template <>
struct type_action<grammar::type_members_clause::members_begin> {
    template <typename Input>
    static void apply(const Input&, type_state& state) {
        state.push_element();
    }
};

template <>
struct type_action<grammar::type_member_group_declaration::members_begin> {
    template <typename Input>
    static void apply(const Input&, type_member_group_state& state) {
        state.push_element();
    }
};
}

template <>
struct control<grammar::type_declaration_clause>
    : pegtl::change_state_and_action<grammar::type_declaration_clause,
                                     declaration::type_state,
                                     declaration::type_action,
                                     errors> {};

template <>
struct control<grammar::type_members_clause::members>
    : pegtl::change_state_and_action<grammar::type_members_clause::members,
                                     declaration::type_members_state,
                                     array_action,
                                     errors> {};

template <>
struct control<grammar::type_member_declaration>
    : pegtl::change_state_and_action<grammar::type_member_declaration,
                                     declaration::type_member_state,
                                     declaration::type_action,
                                     errors> {};

template <>
struct control<grammar::type_member_group_declaration>
    : pegtl::change_state_and_action<grammar::type_member_group_declaration,
                                     declaration::type_member_group_state,
                                     declaration::type_action,
                                     errors> {};

template <>
struct control<grammar::type_member_group_declaration::members>
    : pegtl::change_state_and_action<grammar::type_member_group_declaration::members,
                                     declaration::type_member_group_members_state,
                                     array_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_DECLARATION_TYPE_DECLARATION_HPP
