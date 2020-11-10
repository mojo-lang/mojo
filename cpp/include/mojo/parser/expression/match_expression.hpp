#ifndef MOJO_PARSER_EXPRESSION_MATCH_EXPRESSION_HPP
#define MOJO_PARSER_EXPRESSION_MATCH_EXPRESSION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace expression {

struct match_case_state : array_term_state {
    match_case_state() : array_term_state("match_case") {
    }
};

struct match_cases_state : array_term_state {
    match_cases_state() : array_term_state("match_cases") {
    }
};

template <typename Rule>
struct match_action : pegtl::nothing<Rule> {};

template <>
struct match_action<grammar::match_expression_clause> {
    template <typename Input>
    static void apply(const Input&, term_state& state) {
        state.term = make_term("match_expression");
    }
};

template <>
struct match_action<grammar::statement_separator> {
    template <typename Input>
    static void apply(const Input&, match_cases_state& state) {
        state.push_element();
    }
};

template <>
struct match_action<grammar::match_arrow> {
    template <typename Input>
    static void apply(const Input&, match_case_state& state) {
        state.push_element();
    }
};
}

template <>
struct control<grammar::match_expression_clause>
    : pegtl::change_action<grammar::match_expression_clause,
                           expression::match_action,
                           errors> {};

template <>
struct control<grammar::match_expression_cases>
    : pegtl::change_state_and_action<grammar::match_expression_cases,
                                     expression::match_cases_state,
                                     expression::match_action,
                                     errors> {};

template <>
struct control<grammar::match_expression_case>
    : pegtl::change_state_and_action<grammar::match_expression_case,
                                     expression::match_case_state,
                                     expression::match_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_EXPRESSION_MATCH_EXPRESSION_HPP
