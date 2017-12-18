#ifndef MOJO_PARSER_EXPRESSION_PREFIX_EXPRESSION_HPP
#define MOJO_PARSER_EXPRESSION_PREFIX_EXPRESSION_HPP

#include <mojo/grammar/expression.hpp>
#include <mojo/parser/expression/operator_expression_state.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace expression {

struct prefix_expression_state : term_state {
    TermPtr expression;

    void success(term_state& in) {
        if (expression) {  // really a prefix expression
            if (this->term) {
                expression->terms.push_back(std::move(this->term));
                in.term = std::move(expression);
            }
            else {

            }
        }
        else {  // will be an error
            std::cout << "lack of prefix operator" << std::endl;
        }
    }
};

template <typename Rule>
struct prefix_expression_action : pegtl::nothing<Rule> {};

template <>
struct prefix_expression_action<grammar::prefix_operator> {
    template <typename Input>
    static void apply(const Input& in, prefix_expression_state& state) {
        state.expression = make_term("prefix_expression");

        auto term = make_term("prefix_operator");
        term->value = in.string();

        state.expression->terms.push_back(std::move(term));
    }
};
}
template <>
struct control<grammar::prefix_expression>
    : pegtl::change_state_and_action<grammar::prefix_expression,
                                     expression::prefix_expression_state,
                                     expression::prefix_expression_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_EXPRESSION_PREFIX_EXPRESSION_HPP