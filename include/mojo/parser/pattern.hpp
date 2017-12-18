#ifndef MOJO_PARSER_PATTERNS_HPP
#define MOJO_PARSER_PATTERNS_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {

struct wildcard_pattern_state : array_term_state {
    wildcard_pattern_state() : array_term_state("wildcard_pattern") {
    }
};

struct identifier_pattern_state : array_term_state {
    identifier_pattern_state() : array_term_state("identifier_pattern") {
    }
};

struct enum_value_pattern_state : array_term_state {
    enum_value_pattern_state() : array_term_state("enum_value_pattern") {
    }
};

struct tuple_pattern_state : array_term_state {
    tuple_pattern_state() : array_term_state("tuple_pattern") {
    }
};

struct tuple_pattern_content_state : array_term_state {
    tuple_pattern_content_state() : array_term_state("content") {
    }
};

struct tuple_pattern_element_state : array_term_state {
    tuple_pattern_element_state() : array_term_state("tuple_pattern_element") {
    }
};

template <typename Rule>
struct wildcard_pattern_action : pegtl::nothing<Rule> {};

template <>
struct wildcard_pattern_action<grammar::type_annotation::predication> {
    template <typename Input>
    static void apply(const Input&, wildcard_pattern_state& state) {
        state.push_element();
    }
};

template <typename Rule>
struct pattern_action : pegtl::nothing<Rule> {};

template <>
struct pattern_action<grammar::type_annotation::predication> {
    template <typename Input>
    static void apply(const Input&, identifier_pattern_state& state) {
        state.push_element();
    }
};

template <>
struct pattern_action<grammar::enum_value_pattern::separator> {
    template <typename Input>
    static void apply(const Input&, enum_value_pattern_state& state) {
        state.push_element();
    }
};

template <typename Rule>
struct tuple_pattern_action : pegtl::nothing<Rule> {};

template <>
struct tuple_pattern_action<grammar::tuple_pattern_element::separator> {
    template <typename Input>
    static void apply(const Input&, tuple_pattern_element_state& state) {
        state.push_element();
    }
};

template <>
struct tuple_pattern_action<grammar::type_annotation::predication> {
    template <typename Input>
    static void apply(const Input&, tuple_pattern_state& state) {
        state.push_element();
    }
};

template <>
struct control<grammar::wildcard_pattern> : pegtl::change_state_and_action<grammar::wildcard_pattern,
                                                                           wildcard_pattern_state,
                                                                           wildcard_pattern_action,
                                                                           errors> {};

template <>
struct control<grammar::identifier_pattern> : pegtl::change_state_and_action<grammar::identifier_pattern,
                                                                             identifier_pattern_state,
                                                                             pattern_action,
                                                                             errors> {};

template <>
struct control<grammar::enum_value_pattern> : pegtl::change_state_and_action<grammar::enum_value_pattern,
                                                                             enum_value_pattern_state,
                                                                             pattern_action,
                                                                             errors> {};

template <>
struct control<grammar::tuple_pattern_element>
    : pegtl::change_state_and_action<grammar::tuple_pattern_element,
                                     tuple_pattern_element_state,
                                     tuple_pattern_action,
                                     errors> {};

template <>
struct control<grammar::tuple_pattern_content>
        : pegtl::change_state_and_action<grammar::tuple_pattern_content,
                tuple_pattern_content_state,
                array_action,
                errors> {};

template <>
struct control<grammar::tuple_pattern>
    : pegtl::
          change_state_and_action<grammar::tuple_pattern, tuple_pattern_state, tuple_pattern_action, errors> {
};
}
}

#endif  // MOJO_PARSER_PATTERNS_HPP
