#ifndef MOJO_GRAMMAR_STATEMENT_HPP
#define MOJO_GRAMMAR_STATEMENT_HPP

#include <pegtl.hh>
#include <mojo/grammar/attribute.hpp>
#include <mojo/grammar/document.hpp>
#include <mojo/grammar/expression.hpp>
#include <mojo/grammar/pattern.hpp>
#include <mojo/grammar/type.hpp>

namespace mojo {
namespace grammar {

struct declaration;
struct loop_statement;
struct branch_statement;
struct control_transfer_statement;

struct statement_separator
    : pegtl::seq<inline_seps,
                 pegtl::sor<pegtl::eolf, pegtl::one<';'>, pegtl::at<pegtl::one<'}'>>>> {};

/**
 * GRAMMAR OF A STATEMENT
 */
struct statement : pegtl::sor<loop_statement,
                              branch_statement,
                              control_transfer_statement,
                              declaration,
                              expression> {};
struct statements : pegtl::seq<statement,
                               pegtl::star<statement_separator, seps, statement>> {
};

/**
 * GRAMMAR OF A CODE BLOCK
 */
struct code_block_begin : pegtl::one<'{'> {};
struct code_block_end : pegtl::one<'}'> {};
struct code_block
    : pegtl::seq<code_block_begin, pegtl::pad_opt<statements, sep>, code_block_end> {
    using begin = code_block_begin;
    using end = code_block_end;
};

/**
 * GRAMMAR OF A LOOP STATEMENT
 */
struct for_in_statement;
struct while_statement;
struct repeat_while_statement;

struct loop_statement
    : pegtl::sor<for_in_statement, while_statement, repeat_while_statement> {};

/**
 * GRAMMAR OF A FOR_IN STATEMENT
 */
struct for_in_statement_clause
    : pegtl::
          seq<destructuring_pattern, seps, key_in, seps, expression, seps, code_block> {};
struct for_in_statement : pegtl::seq<key_for, seps, for_in_statement_clause> {};

/**
 * GRAMMAR OF A WHILE STATEMENT
 */
struct condition : expression {};
struct condition_list : pegtl::list<condition, pegtl::one<','>, sep> {};
struct while_statement_clause : pegtl::seq<condition_list, seps, code_block> {};
struct while_statement : pegtl::seq<key_while, seps, while_statement_clause> {};

/**
 * GRAMMAR OF A REPEAT_WHILE STATEMENT
 */
struct repeat_while_statement_clause
    : pegtl::seq<code_block, seps, key_while, seps, expression> {};
struct repeat_while_statement
    : pegtl::seq<key_repeat, seps, repeat_while_statement_clause> {};

/**
 * GRAMMAR OF A BRANCH STATEMENT
 */
struct if_statement;
struct match_statement;
struct branch_statement : pegtl::sor<if_statement, match_statement> {};

/**
 * GRAMMAR OF AN IF STATEMENT
 */
struct else_statement;
struct if_statement_clause
    : pegtl::seq<condition_list, seps, code_block, pre_pad_opt<else_statement, sep>> {};
struct if_statement : pegtl::seq<key_if, seps, if_statement_clause> {};
struct else_statement_clause : pegtl::sor<code_block, if_statement> {};
struct else_statement : pegtl::seq<key_else, seps, else_statement_clause> {};

struct match_arrow : pegtl_string_t("=>") {};
struct match_case
    : pegtl::seq<pattern, seps, match_arrow, seps, pegtl::sor<code_block, expression>> {};
struct match_cases : pegtl::plus<pegtl::seq<seps, match_case, statement_separator>> {};

/**
 * GRAMMAR OF A SWITCH STATEMENT
 */
struct match_cases_begin : pegtl::one<'{'> {};
struct match_cases_end : pegtl::one<'}'> {};
struct match_statement_clause : pegtl::seq<expression,
                                           seps,
                                           match_cases_begin,
                                           seps,
                                           match_cases,
                                           seps,
                                           match_cases_end> {
    using begin = match_cases_begin;
    using end = match_cases_end;
};
struct match_statement : pegtl::seq<key_match, seps, match_statement_clause> {};

/**
 * GRAMMAR OF A BREAK STATEMENT
 */
struct break_statement : key_break {};

/**
 * GRAMMAR OF A CONTINUE STATEMENT
 */
struct continue_statement : key_continue {};

/**
 * GRAMMAR OF A RETURN STATEMENT
 */
struct return_statement_clause : pre_pad_opt<expression, sep> {};
struct return_statement : pegtl::seq<key_return, return_statement_clause> {};

/**
 * GRAMMAR OF A CONTROL TRANSFER STATEMENT
 */
struct control_transfer_statement
    : pegtl::sor<break_statement, continue_statement, return_statement> {};
}
}

#endif  // MOJO_GRAMMAR_STATEMENT_HPP