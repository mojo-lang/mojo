#ifndef MOJO_PARSER_DECLARATION_FUNCTION_DECLARATION_HPP
#define MOJO_PARSER_DECLARATION_FUNCTION_DECLARATION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace declaration {

struct function_parameter_state : array_term_state {
    function_parameter_state() : array_term_state("function_parameter") {
    }
};

struct function_parameters_state : array_term_state {
    function_parameters_state() : array_term_state("function_parameters") {
    }
};

struct function_state : array_term_state {
    function_state() : array_term_state("function_declaration") {
    }
};

template <typename Rule>
struct function_action : pegtl::nothing<Rule> {};

template <>
struct function_action<grammar::type_annotation::predication> {
    template <typename Input>
    static void apply(const Input&, function_parameter_state& state) {
        state.push_element();
    }
};

template <>
struct function_action<grammar::initializer::predication> {
    template <typename Input>
    static void apply(const Input&, function_parameter_state& state) {
        state.push_element();
    }
};

template <>
struct function_action<grammar::generic_parameter_clause::begin> {
    template <typename Input>
    static void apply(const Input&, function_state& state) {
        state.push_element();
    }
};

template <>
struct function_action<grammar::function_parameters_begin> {
    template <typename Input>
    static void apply(const Input&, function_state& state) {
        state.push_element();
    }
};

template <>
struct function_action<grammar::function_parameters_end> {
    template <typename Input>
    static void apply(const Input&, function_state& state) {
        state.push_element();
    }
};

template <>
struct function_action<grammar::code_block::begin> {
    template <typename Input>
    static void apply(const Input&, function_state& state) {
        state.push_element();
    }
};

}
template <>
struct control<grammar::function_declaration_clause>
    : pegtl::change_state_and_action<grammar::function_declaration_clause,
                                     declaration::function_state,
                                     declaration::function_action,
                                     errors> {};

template <>
struct control<grammar::function_declaration_clause::parameters>
    : pegtl::change_state_and_action<grammar::function_declaration_clause::parameters,
                                     declaration::function_parameters_state,
                                     array_action,
                                     errors> {};

template <>
struct control<grammar::function_declaration_clause::parameter>
    : pegtl::change_state_and_action<grammar::function_declaration_clause::parameter,
                                     declaration::function_parameter_state,
                                     declaration::function_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_DECLARATION_FUNCTION_DECLARATION_HPP
