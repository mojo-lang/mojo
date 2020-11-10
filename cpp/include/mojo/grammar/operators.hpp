#ifndef MOJO_GRAMMAR_OPERATOR_HPP
#define MOJO_GRAMMAR_OPERATOR_HPP

#include <pegtl.hh>
#include <mojo/grammar/keys.hpp>

namespace mojo {
namespace grammar {

template <char O, char... N>
struct op_one : pegtl::seq<pegtl::one<O>, pegtl::at<pegtl::not_one<N...>>> {};

template <char O, char P, char... N>
struct op_two : pegtl::seq<pegtl::string<O, P>, pegtl::at<pegtl::not_one<N...>>> {};

template <char O, char P, char Q, char... N>
struct op_three : pegtl::seq<pegtl::string<O, P, Q>, pegtl::at<pegtl::not_one<N...>>> {};

/**
 * GRAMMAR OF PREFIX OPERATOR
 */
struct prefix_operator : pegtl::sor<op_one<'!', '='>, op_one<'-', '='>, key_not> {};

struct pow_operator : op_one<'^', '^', '='> {};
struct multiplicative_operator : pegtl::sor<op_one<'*', '*', '='>, op_one<'/', '/', '='>, op_one<'%', '%', '='>> {};
struct additive_operator : pegtl::sor<op_one<'+', '+', '='>, op_one<'-', '-', '='>> {};

struct bitwise_shift_operator : pegtl::sor<pegtl_string_t("<<"), pegtl_string_t(">>")> {};
struct bitwise_and_operator : pegtl_string_t("&&&") {};
struct bitwise_inclusive_or_operator : pegtl_string_t("|||") {};
struct bitwise_exclusive_or_operator : key_xor {};

struct relational_operator
    : pegtl::sor<pegtl_string_t("<="), pegtl_string_t(">="), op_one<'<', '<', '.'>, op_one<'>', '>'>> {};
struct equality_operator : pegtl::sor<pegtl_string_t("=="), pegtl_string_t("!="), key_in> {};

struct logical_and_operator : op_two<'&', '&', '&'> {};
struct logical_or_operator : op_two<'|', '|', '|'> {};

struct pipe_operator : pegtl::one<'|'> {};

struct assignment_operator : pegtl::sor<op_one<'=', '='>,
                                        pegtl_string_t("*="),
                                        pegtl_string_t("/="),
                                        pegtl_string_t("^="),
                                        pegtl_string_t("%="),
                                        pegtl_string_t("+="),
                                        pegtl_string_t("-=")> {};

struct close_range_operator : op_two<'.', '.', '<'> {};
struct open_range_operator : pegtl_string_t("<..<") {};
struct left_open_range_operator : op_three<'<', '.', '.', '<'> {};
struct right_open_range_operator : pegtl_string_t("..<") {};

struct range_operator : pegtl::sor<right_open_range_operator,
                                   close_range_operator,
                                   open_range_operator,
                                   left_open_range_operator> {};

struct range_in_operator : key_in {};

/**
 * GRAMMAR OF PREFIX OPERATOR
 */
struct binary_operator : pegtl::sor<pow_operator,
                                    multiplicative_operator,
                                    additive_operator,
                                    bitwise_shift_operator,
                                    bitwise_and_operator,
                                    bitwise_inclusive_or_operator,
                                    bitwise_exclusive_or_operator,
                                    relational_operator,
                                    equality_operator,
                                    logical_and_operator,
                                    logical_or_operator,
                                    pipe_operator,
                                    assignment_operator,
                                    range_operator,
                                    range_in_operator> {};

struct three_dots : pegtl_string_t("...") {};

struct expression;

/**
 * GRAMMAR OF A CONDITIONAL OPERATOR
 */
struct conditional_operator_begin : pegtl::one<'?'> {};
struct conditional_operator_end : pegtl::one<':'> {};
struct conditional_operator
    : pegtl::seq<conditional_operator_begin, seps, expression, seps, conditional_operator_end> {
    using begin = conditional_operator_begin;
    using end = conditional_operator_end;
};
}
}

#endif  // MOJO_GRAMMAR_OPERATOR_HPP