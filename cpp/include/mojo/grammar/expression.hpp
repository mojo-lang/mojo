#ifndef MOJO_GRAMMAR_EXPRESSION_HPP
#define MOJO_GRAMMAR_EXPRESSION_HPP

#include <mojo/grammar/generics.hpp>
#include <mojo/grammar/identifier.hpp>
#include <mojo/grammar/keys.hpp>
#include <mojo/grammar/literal.hpp>
#include <mojo/grammar/operators.hpp>
#include <mojo/grammar/rules.hpp>
#include <mojo/grammar/string_literal.hpp>
#include <pegtl.hh>

namespace mojo {
namespace grammar {

struct literal_expression;
struct parenthesized_expression;
struct wildcard_expression;

struct generic_identifier : pegtl::seq<identifier, pegtl::opt<generic_argument_clause>> {
};

struct expression_element_separator : pegtl::one<':'> {};
struct expression_element
    : pegtl::sor<
          expression,
          pegtl::seq<identifier, seps, expression_element_separator, seps, expression>> {
    using separator = expression_element_separator;
};
struct expression_element_list : pegtl::list<expression_element, pegtl::one<','>, sep> {};

struct parenthesized_expression_begin : pegtl::one<'('> {};
struct parenthesized_expression_end : pegtl::one<')'> {};
/**
 * GRAMMAR OF A PARENTHESIZED EXPRESSION
 */
struct parenthesized_expression : pegtl::seq<parenthesized_expression_begin,
                                             pegtl::pad_opt<expression_element_list, sep>,
                                             parenthesized_expression_end> {
    using begin = parenthesized_expression_begin;
    using end = parenthesized_expression_end;
};

/**
 * GRAMMAR OF A WILDCARD EXPRESSION
 */
struct wildcard_expression : key_wildcard {};

struct lambda_expression;
struct match_expression;
struct function_call_expression_tail;
struct explicit_member_expression_tail;
struct subscript_expression_tail;
struct object_literal;

struct object_type_construction : pegtl::seq<type_name, seps, object_literal> {};

struct postfix_expression_clause
    : pegtl::seq<pegtl::sor<parenthesized_expression,
                            object_type_construction,
                            match_expression,
                            type_name,
                            generic_identifier,
                            literal_expression>,
                 pegtl::plus<seps,
                             pegtl::sor<function_call_expression_tail,
                                        explicit_member_expression_tail,
                                        subscript_expression_tail>>> {};

/**
 * GRAMMAR OF A PRIMARY EXPRESSION
 */
struct primary_expression : pegtl::sor<object_type_construction,
                                       match_expression,
                                       literal_expression,
                                       wildcard_expression,
                                       parenthesized_expression,
                                       identifier> {};

/**
 * GRAMMAR OF A POSTFIX EXPRESSION
 */
struct postfix_expression : pegtl::sor<postfix_expression_clause, primary_expression> {};

/**
 * GRAMMAR OF A FUNCTION CALL EXPRESSION
 */
struct function_call_expression_tail : parenthesized_expression {
    using begin = parenthesized_expression::begin;
    using end = parenthesized_expression::end;
};
struct function_call_expression
    : pegtl::seq<postfix_expression, seps, function_call_expression_tail> {};

/**
 * GRAMMAR OF AN EXPLICIT MEMBER EXPRESSION
 */
struct explicit_member_expression_begin : pegtl::one<'.'> {};
struct explicit_member_expression_tail : pegtl::seq<explicit_member_expression_begin,
                                                    pegtl::sor<decimal_digits,
                                                               object_type_construction,
                                                               match_expression,
                                                               type_name,
                                                               generic_identifier>> {
    using begin = explicit_member_expression_begin;
};

struct explicit_member_expression
    : pegtl::seq<postfix_expression, seps, explicit_member_expression_tail> {};

/**
 * GRAMMAR OF A SUBSCRIPT EXPRESSION
 */
struct subscript_expression_begin : pegtl::one<'['> {};
struct subscript_expression_end : pegtl::one<']'> {};
struct subscript_expression_tail : pegtl::seq<subscript_expression_begin,
                                              seps,
                                              expression,
                                              seps,
                                              subscript_expression_end> {
    using begin = subscript_expression_begin;
    using end = subscript_expression_end;
};
struct subscript_expression
    : pegtl::seq<postfix_expression, seps, subscript_expression_tail> {};

/**
 * GRAMMAR OF A PREFIX EXPRESSION
 */
struct prefix_expression : pegtl::seq<prefix_operator, seps, postfix_expression> {};

/**
 * GRAMMAR OF A UNARY EXPRESSION
 */
struct unary_expression : pegtl::sor<prefix_expression, postfix_expression> {};

/**
 * GRAMMAR OF A BINARY EXPRESSION
 */
struct binary_expression_clause
    : pegtl::
          seq<pegtl::sor<binary_operator, conditional_operator>, seps, unary_expression> {
};
struct binary_expression
    : pegtl::seq<unary_expression,
                 pegtl::plus<pegtl::seq<seps, binary_expression_clause>>> {};

/**
 * GRAMMAR OF AN EXPRESSION
 */
struct expression : pegtl::sor<lambda_expression, binary_expression, unary_expression> {};
struct expression_list : left_assoc<expression, pegtl::one<','>> {};

struct array_literal_item : expression {};
struct array_literal_items : list_tail<array_literal_item, value_separator, sep> {};

struct array_literal_content : pegtl::pad_opt<array_literal_items, sep> {};

/**
 * GRAMMAR OF A ARRAY LITERAL EXPRESSION
 */
struct array_literal
    : pegtl::seq<pegtl::one<'['>, array_literal_content, pegtl::one<']'>> {};

struct path_identifier_separator : pegtl::one<'.'> {};
struct path_identifier : pegtl::list<identifier, path_identifier_separator> {};

struct generic_path_identifier
    : pegtl::seq<path_identifier, pegtl::opt<generic_argument_clause>> {};

struct object_literal_key
    : pegtl::sor<path_identifier, integer_literal, static_string_literal> {};
struct object_literal_value_separator : pegtl::one<':'> {};
struct object_literal_item_separator : value_separator {};
struct object_literal_item
    : pegtl::seq<object_literal_key,
                 seps,
                 pegtl::opt<object_literal_value_separator, seps, expression>> {};
struct object_literal_items
    : list_tail<object_literal_item, object_literal_item_separator, sep> {};
struct object_literal_content : pegtl::pad_opt<object_literal_items, sep> {
    using item = object_literal_item;
    using items = object_literal_items;
};

/**
 * GRAMMAR OF A OBJECT LITERAL EXPRESSION
 */
struct object_literal_begin : pegtl::one<'{'> {};
struct object_literal_end : pegtl::one<'}'> {};
struct object_literal
    : pegtl::seq<object_literal_begin, object_literal_content, object_literal_end> {
    using items = object_literal_content::items;
    using item = object_literal_content::item;
    using begin = object_literal_begin;
    using end = object_literal_end;
};

struct string_prefix : pegtl::seq<pegtl::alpha, pegtl::rep_max<7, pegtl::alnum>> {};
struct string_prefix_literal : pegtl::seq<string_prefix, static_string_literal> {};

/**
 * GRAMMAR OF A LITERAL EXPRESSION
 */
struct literal_expression
    : pegtl::sor<string_prefix_literal, literal, array_literal, object_literal> {};

}  // namespace grammar
}  // namespace mojo

#endif  // MOJO_GRAMMAR_EXPRESSION_HPP