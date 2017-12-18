#ifndef MOJO_GRAMMAR_LAMBDA_EXPRESSION_HPP
#define MOJO_GRAMMAR_LAMBDA_EXPRESSION_HPP

#include <mojo/grammar/statement.hpp>

namespace mojo {
namespace grammar {

struct lambda_three_dot : pegtl_string_t("...") {};
struct lambda_parameter : pegtl::seq<identifier,
                                     pre_pad_opt<type_annotation, sep>,
                                     pre_pad_opt<lambda_three_dot, sep>> {};

struct lambda_parameters : pegtl::list<lambda_parameter, pegtl::one<','>, sep> {};
struct lambda_parameter_clause
    : pegtl::sor<pegtl::seq<pegtl::one<'('>,
                            pegtl::pad_opt<lambda_parameters, sep>,
                            pegtl::one<')'>>,
                 lambda_parameter> {};

struct lambda_expression_arrow : pegtl_string_t("->") {};
/**
 * GRAMMAR OF A LAMBDA EXPRESSION
 */
struct lambda_expression : pegtl::seq<lambda_parameter_clause,
                                      seps,
                                      lambda_expression_arrow,
                                      seps,
                                      pegtl::sor<code_block, expression>> {
    using arrow = lambda_expression_arrow;
};
}
}

#endif  // MOJO_GRAMMAR_LAMBDA_EXPRESSION_HPP
