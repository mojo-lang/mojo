#ifndef MOJO_PARSER_STATEMENT_MATCH_STATEMENT_HPP
#define MOJO_PARSER_STATEMENT_MATCH_STATEMENT_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace statement {

struct match_state : array_term_state {
    match_state() : array_term_state("match_statement") {
    }
};

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
struct match_action<grammar::match_statement_clause::begin> {
    template <typename Input>
    static void apply(const Input&, match_state& state) {
        state.push_element();
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
struct control<grammar::match_statement_clause>
    : pegtl::change_state_and_action<grammar::match_statement_clause,
                                     statement::match_state,
                                     statement::match_action,
                                     errors> {};
template <>
struct control<grammar::match_cases>
    : pegtl::change_state_and_action<grammar::match_cases,
                                     statement::match_cases_state,
                                     statement::match_action,
                                     errors> {};
template <>
struct control<grammar::match_case>
    : pegtl::change_state_and_action<grammar::match_case,
                                     statement::match_case_state,
                                     statement::match_action,
                                     errors> {};
}
}
#endif  // MOJO_PARSER_STATEMENT_MATCH_STATEMENT_HPP
