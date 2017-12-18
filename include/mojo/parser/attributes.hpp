#ifndef MOJO_PARSER_ATTRIBUTES_HPP
#define MOJO_PARSER_ATTRIBUTES_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {

struct attribute_state : array_term_state {
    attribute_state() : array_term_state("attribute") {}
};

struct group_attribute_state : dictionary_term_state {
    group_attribute_state() : dictionary_term_state("group_attribute", "attribute") {}
};

struct attributes_state : array_term_state {
    attributes_state() : array_term_state("attributes") {}
};

template <typename Rule>
struct attributes_action : pegtl::nothing<Rule> {};

template <>
struct attributes_action<grammar::attribute_argument::begin> {
    template <typename Input>
    static void apply(const Input&, attribute_state& state) {
        state.push_element();
    }
};

template <>
struct attributes_action<grammar::group_attribute::content::element::separator> {
    template <typename Input>
    static void apply(const Input&, group_attribute_state& state) {
        state.add_key();
    }
};

template <>
struct attributes_action<grammar::group_attribute::content::element> {
    template <typename Input>
    static void apply(const Input&, group_attribute_state& state) {
        state.set_value();
    }
};

template <>
struct attributes_action<grammar::attribute> {
    template <typename Input>
    static void apply(const Input&, attributes_state& state) {
        state.push_element();
    }
};

template <>
struct attributes_action<grammar::group_attribute> {
    template <typename Input>
    static void apply(const Input&, attributes_state& state) {
        state.push_element();
    }
};

template <>
struct control<grammar::attribute::content>
    : pegtl::change_state_and_action<grammar::attribute::content, attribute_state, attributes_action, errors> {
};

template <>
struct control<grammar::group_attribute::content>
    : pegtl::change_state_and_action<grammar::group_attribute::content,
                                     group_attribute_state,
                                     attributes_action,
                                     errors> {};

template <>
struct control<grammar::attributes::content>
    : pegtl::change_state_and_action<grammar::attributes::content, attributes_state, attributes_action, errors> {};
}
}

#endif  // MOJO_PARSER_ATTRIBUTES_HPP
