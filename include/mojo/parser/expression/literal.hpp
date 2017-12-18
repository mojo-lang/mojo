#ifndef MOJO_PARSER_EXPRESSION_LITERAL_HPP
#define MOJO_PARSER_EXPRESSION_LITERAL_HPP

#include <list>
#include <map>
#include <pegtl/contrib/changes.hh>

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/errors.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace expression {

template <typename Rule>
struct literal_action : pegtl::nothing<Rule> {};

template <>
struct literal_action<grammar::null_literal> {
    template <typename Input>
    static void apply(const Input&, term_state& state) {
        state.term = make_term("null_literal");
    }
};

template <>
struct literal_action<grammar::bool_literal> {
    template <typename Input>
    static void apply(const Input& input, term_state& state) {
        state.term = make_term("bool_literal", input.string());
    }
};

template <>
struct literal_action<grammar::decimal_digits> {
    template <typename Input>
    static void apply(const Input& input, term_state& state) {
        state.term = make_term("integer_literal", input.string());
    }
};

template <>
struct literal_action<grammar::integer_literal> {
    template <typename Input>
    static void apply(const Input& input, term_state& state) {
        state.term = make_term("integer_literal", input.string());
    }
};

template <>
struct literal_action<grammar::floating_point_literal> {
    template <typename Input>
    static void apply(const Input& input, term_state& state) {
        state.term = make_term("floating_point_literal", input.string());
    }
};

template <>
struct literal_action<grammar::double_quoted_static_string::content> {
    template <typename Input>
    static void apply(const Input& input, term_state& state) {
        state.term = make_term("string_literal", input.string());
    }
};

template <>
struct literal_action<grammar::single_quoted_static_string::content> {
    template <typename Input>
    static void apply(const Input& input, term_state& state) {
        state.term = make_term("string_literal", input.string());
    }
};

template <>
struct literal_action<grammar::interpolated_string_literal> {
    template <typename Input>
    static void apply(const Input& input, term_state& state) {
        state.term = make_term("interpolated_string_literal", input.string());
    }
};
}

template <>
struct control<grammar::literal>
    : pegtl::change_action<grammar::literal, expression::literal_action, errors> {};

template <>
struct control<grammar::decimal_digits>
    : pegtl::change_action<grammar::decimal_digits, expression::literal_action, errors> {
};

template <>
struct control<grammar::static_string_literal>
    : pegtl::change_action<grammar::static_string_literal,
                           expression::literal_action,
                           errors> {};
}
}

#endif  // MOJO_PARSER_EXPRESSION_LITERAL_HPP