#ifndef MOJO_PARSER_EXPRESSION_POSTFIX_EXPRESSION_HPP
#define MOJO_PARSER_EXPRESSION_POSTFIX_EXPRESSION_HPP

#include <mojo/grammar/grammar.hpp>
#include <mojo/parser/expression/operator_expression_state.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace expression {

struct postfix_processor {
    static TermPtr success(TermPtr&& base, operator_expressions&& expressions) {
        TermPtr first = std::move(base);

        for (auto& expr : expressions) {
            auto term = std::move(make_term(expr.type));
            term->terms.push_back(std::move(first));

            if (expr.expression) {
                term->terms.push_back(std::move(expr.expression));
            }

            first = std::move(term);
        }
        return std::move(first);
    }
};

using postfix_expression_state = operator_expression_state<postfix_processor>;

template <typename Rule>
struct postfix_expression_action : pegtl::nothing<Rule> {};

template <>
struct postfix_expression_action<grammar::function_call_expression_tail::begin> {
    template <typename Input>
    static void apply(const Input&, postfix_expression_state& state) {
        state.add_operator("function_call");
    }
};

template <>
struct postfix_expression_action<grammar::explicit_member_expression_tail::begin> {
    template <typename Input>
    static void apply(const Input&, postfix_expression_state& state) {
        state.add_operator("explicit_member");
    }
};

template <>
struct postfix_expression_action<grammar::subscript_expression_tail::begin> {
    template <typename Input>
    static void apply(const Input&, postfix_expression_state& state) {
        state.add_operator("subscript");
    }
};
}

template <>
struct control<grammar::postfix_expression_clause>
    : pegtl::change_state_and_action<grammar::postfix_expression_clause,
                                     expression::postfix_expression_state,
                                     expression::postfix_expression_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_EXPRESSION_POSTFIX_EXPRESSION_HPP