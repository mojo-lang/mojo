#ifndef MOJO_PARSER_EXPRESSION_LAMBDA_EXPRESSION_HPP
#define MOJO_PARSER_EXPRESSION_LAMBDA_EXPRESSION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace expression {

struct lambda_state : array_term_state {
    lambda_state() : array_term_state("lambda") {
    }
};

struct lambda_parameter_state : array_term_state {
    lambda_parameter_state() : array_term_state("lambda_parameters") {
    }
};

struct lambda_parameters_state : array_term_state {
    lambda_parameters_state() : array_term_state("lambda_parameters") {
    }
};

template <typename Rule>
struct lambda_action : pegtl::nothing<Rule> {};

template <>
struct lambda_action<grammar::lambda_expression::arrow> {
    template <typename Input>
    static void apply(const Input&, lambda_state& state) {
        state.push_element();
    }
};

template <>
struct lambda_action<grammar::type_annotation::predication> {
    template <typename Input>
    static void apply(const Input&, lambda_parameter_state& state) {
        state.push_element();
    }
};

template <>
struct lambda_action<grammar::lambda_three_dot> {
    template <typename Input>
    static void apply(const Input&, lambda_parameter_state& state) {
        state.push_element();

        state.term = make_term("three_dot");
        state.push_element();
    }
};
}

template <>
struct control<grammar::lambda_expression>
    : pegtl::change_state_and_action<grammar::lambda_expression,
                                     expression::lambda_state,
                                     expression::lambda_action,
                                     errors> {};

template <>
struct control<grammar::lambda_parameters>
    : pegtl::change_state_and_action<grammar::lambda_parameters,
                                     expression::lambda_parameters_state,
                                     array_action,
                                     errors> {};

template <>
struct control<grammar::lambda_parameter>
    : pegtl::change_state_and_action<grammar::lambda_parameter,
                                     expression::lambda_parameter_state,
                                     expression::lambda_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_EXPRESSION_LAMBDA_EXPRESSION_HPP
