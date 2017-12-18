#ifndef MOJO_PARSER_EXPRESSION_IDENTIFER_HPP
#define MOJO_PARSER_EXPRESSION_IDENTIFER_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/control.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace expression {

struct path_identifier_state : array_term_state {
    path_identifier_state() : array_term_state("path_identifier") {
    }
};

struct generic_identifier_state : array_term_state {
    generic_identifier_state() : array_term_state("generic_identifier") {
    }
};

template <typename Rule>
struct identifier_action : pegtl::nothing<Rule> {};

template <>
struct identifier_action<grammar::identifier> {
    template <typename Input>
    static void apply(const Input& in, term_state& state) {
        auto term = make_term("identifier");

        if (in.string()[0] == '$') {
            if (in.string().size() > 1) {
                term->value = in.string();
            }
            else {
                term->value = "$0";
            }
        }
        else {
            term->value = in.string();
        }

        state.term = std::move(term);
    }
};

template <>
struct identifier_action<grammar::generic_argument_clause::begin> {
    template <typename Input>
    static void apply(const Input&, generic_identifier_state& state) {
        state.push_element();
    }
};

template <>
struct identifier_action<grammar::path_identifier_separator> {
    template <typename Input>
    static void apply(const Input&, path_identifier_state& state) {
        state.push_element();
    }
};
}

template <>
struct control<grammar::identifier>
    : pegtl::change_action<grammar::identifier, expression::identifier_action, errors> {};

template <>
struct control<grammar::generic_identifier>
    : pegtl::change_state_and_action<grammar::generic_identifier,
                                     expression::generic_identifier_state,
                                     expression::identifier_action,
                                     errors> {};

template <>
struct control<grammar::path_identifier>
    : pegtl::change_state_and_action<grammar::path_identifier,
                                     expression::path_identifier_state,
                                     expression::identifier_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_EXPRESSION_IDENTIFER_HPP
