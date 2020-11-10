#ifndef MOJO_GRAMMAR_MATCH_EXPRESSION_HPP
#define MOJO_GRAMMAR_MATCH_EXPRESSION_HPP

#include <mojo/grammar/statement.hpp>

namespace mojo {
namespace grammar {

struct match_expression_case : match_case {};
struct match_expression_cases : list<match_expression_case, ';', '_'> {};

/**
 * GRAMMAR OF A MATCH EXPRESSION
 */
struct match_expression_clause
    : pegtl::seq<match_cases_begin, seps, match_expression_cases, seps, match_cases_end> {
    using begin = match_cases_begin;
    using end = match_cases_end;
};

struct match_expression : pegtl::seq<key_match, seps, match_expression_clause> {};
}
}

#endif  // MOJO_GRAMMAR_MATCH_EXPRESSION_HPP
