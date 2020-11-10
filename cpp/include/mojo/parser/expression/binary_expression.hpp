#ifndef MOJO_PARSER_EXPRESSION_BINARY_EXPRESSION_HPP
#define MOJO_PARSER_EXPRESSION_BINARY_EXPRESSION_HPP

#include <mojo/parser/errors.hpp>
#include <mojo/parser/expression/operator_expression_state.hpp>
#include <mojo/parser/term_state.hpp>

namespace mojo {
namespace parser {
namespace expression {

using binary_operator_precedence = std::map<std::string, int>;

static const binary_operator_precedence const_precedence =
    {{"=", 1},    {"*=", 1},    {"/=", 1},          {"^=", 1},  {"%=", 1},  {"+=", 1},
     {"-=", 1},   {"|", 2},     {"conditional", 3}, {"||", 4},  {"&&", 5},  {"==", 6},
     {"!=", 6},   {">", 7},     {">=", 7},          {"<", 7},   {"<=", 7},  {"in", 7},
     {"|||", 8},  {"xor", 9},   {"&&&", 10},        {"<<", 11}, {">>", 12}, {"+", 13},
     {"-", 13},   {"*", 14},    {"/", 14},          {"%", 14},  {"..", 15}, {"..<", 15},
     {"<..", 15}, {"<..<", 15}, {"^", 16}};

static const int const_highest_precedence = 16;

struct binary_processor {
    static TermPtr success(TermPtr&& base, operator_expressions&& expressions) {
        TermPtr base_expr = std::move(base);
        operator_expressions exprs = std::move(expressions);

        reduce(base_expr, exprs);

        for (auto& expr : exprs) {
            auto term = std::move(make_term(expr.type));
            term->terms.push_back(std::move(base_expr));
            if (expr.param) {
                term->terms.push_back(std::move(expr.param));
            }
            term->terms.push_back(std::move(expr.expression));

            base_expr = std::move(term);
        }
        return std::move(base_expr);
    }

    static void reduce(TermPtr& base, operator_expressions& expressions) {
        int p = highest_precedence(expressions);
        while (p > 0) {
            bool need_process = true;
            while (need_process) {
                need_process = false;
                operator_expressions::iterator itr = expressions.begin();
                for (; itr != expressions.end(); ++itr) {
                    if (precedence(*itr) == p) {
                        TermPtr term = make_term(itr->type);

                        if (itr == expressions.begin()) {
                            term->terms.push_back(std::move(base));

                            if (itr->param) {
                                term->terms.push_back(std::move(itr->param));
                            }
                            term->terms.push_back(std::move(itr->expression));

                            base = std::move(term);
                            expressions.erase(itr);
                        }
                        else {
                            operator_expressions::iterator before = itr - 1;
                            term->terms.push_back(std::move(before->expression));

                            if (itr->param) {
                                term->terms.push_back(std::move(itr->param));
                            }
                            term->terms.push_back(std::move(itr->expression));

                            before->expression = std::move(term);
                            expressions.erase(itr);
                        }
                        need_process = true;
                        break;
                    }
                    else {
                        continue;
                    }
                }
            }

            p = highest_precedence(expressions);
        }
    }

    static int precedence(const operator_expression& expr) {
        binary_operator_precedence::const_iterator itr = const_precedence.find(expr.type);
        if (itr != const_precedence.end()) {
            return itr->second;
        }
        else {
            return 0;
        }
    }

    static int highest_precedence(const operator_expressions& expressions) {
        int higher = 0;
        int lower = const_highest_precedence;
        for (auto const& expr : expressions) {
            int p = precedence(expr);
            if (p > higher) {
                higher = p;
            }
            if (p < lower) {
                lower = p;
            }
        }
        return higher > lower ? higher : 0;
    }
};

using binary_expression_state = operator_expression_state<binary_processor>;

template <typename Rule>
struct binary_expression_action : pegtl::nothing<Rule> {};

template <>
struct binary_expression_action<grammar::binary_operator> {
    template <typename Input>
    static void apply(const Input& in, binary_expression_state& state) {
        state.add_operator(in.string());
    }
};

template <>
struct binary_expression_action<grammar::conditional_operator::begin> {
    template <typename Input>
    static void apply(const Input& in, binary_expression_state& state) {
        state.add_operator("conditional");
    }
};

template <>
struct binary_expression_action<grammar::conditional_operator::end> {
    template <typename Input>
    static void apply(const Input& in, binary_expression_state& state) {
        state.add_operator_param();
    }
};
}
template <>
struct control<grammar::binary_expression>
    : pegtl::change_state_and_action<grammar::binary_expression,
                                     expression::binary_expression_state,
                                     expression::binary_expression_action,
                                     errors> {};
}
}

#endif  // MOJO_PARSER_EXPRESSION_BINARY_EXPRESSION_HPP