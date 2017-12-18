#ifndef MOJO_PARSER_EXPRESSION_OPERATOR_EXPRESSION_STATE_HPP
#define MOJO_PARSER_EXPRESSION_OPERATOR_EXPRESSION_STATE_HPP

#include <mojo/lang/term.hpp>

namespace mojo {
namespace parser {
namespace expression {

struct operator_expression {
    std::string type;
    TermPtr param;
    TermPtr expression;
};

using operator_expressions = std::vector<operator_expression>;

template <typename Processor>
struct operator_expression_state : term_state {
    TermPtr expression;
    operator_expressions operator_expressions;

    void add_operator(const std::string& type) {
        if (operator_expressions.empty()) {  // the first operator
            expression = std::move(term);
        }
        else {
            operator_expressions.back().expression = std::move(term);
        }

        operator_expressions.push_back({type});
        term.reset();
    }

    void add_operator_param() {
        if (!operator_expressions.empty()) {
            operator_expressions.back().param = std::move(term);
        }
        else {
            // error
        }
    }

    void success(term_state& in) {
        if (operator_expressions.empty()) {
            if (expression) {  // is prefix expression
                in.term = std::move(expression);
            }
            else {
                in.term = std::move(term);
            }
        }
        else {
            operator_expressions.back().expression = std::move(term);
            TermPtr t = Processor::success(std::move(expression),
                                           std::move(operator_expressions));
            in.term = std::move(t);
        }
    }
};
}
}
}

#endif  // MOJO_PARSER_EXPRESSION_OPERATOR_EXPRESSION_STATE_HPP