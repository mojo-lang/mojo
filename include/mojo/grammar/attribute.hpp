#ifndef MOJO_GRAMMAR_ATTRIBUTE_HPP
#define MOJO_GRAMMAR_ATTRIBUTE_HPP

#include <pegtl.hh>
#include <mojo/grammar/at.hpp>
#include <mojo/grammar/expression.hpp>
#include <mojo/grammar/identifier.hpp>

namespace mojo {
namespace grammar {

struct attribute_argument_begin : pegtl::one<'('> {};
struct attribute_argument_end : pegtl::one<')'> {};
struct attribute_argument
    : pegtl::seq<attribute_argument_begin, seps, expression, seps, attribute_argument_end> {
    using begin = attribute_argument_begin;
    using end = attribute_argument_end;
};

struct attribute_clause
    : pegtl::sor<decimal_digits, pegtl::seq<path_identifier, pre_pad_opt<attribute_argument, pegtl::ascii::blank>>> {};

/**
 * GRAMMAR OF AN ATTRIBUTE
 */
struct attribute : pegtl::seq<pegtl::one<'@'>, attribute_clause> {
    using predication = pegtl::one<'@'>;
    using content = attribute_clause;
};

struct group_attribute_element_separator : pegtl::one<':'> {};
struct group_attribute_element
    : pegtl::seq<path_identifier, seps, group_attribute_element_separator, seps, expression> {
    using separator = group_attribute_element_separator;
};

struct group_attribute_clause : list_tail<group_attribute_element, value_separator, sep> {
    using element = group_attribute_element;
    using separator = value_separator;
};

/**
 * GRAMMAR OF AN ATTRIBUTE GROUP
 */
struct group_attribute
    : pegtl::seq<pegtl_string_t("@{"), pegtl::pad_opt<group_attribute_clause, sep>, pegtl::one<'}'>> {
    using content = group_attribute_clause;
};

struct attribute_separator
    : pegtl::seq<inline_seps, pegtl::star<pegtl::eolf>, blanks> {};

using attributes_clause = pegtl::seq<pegtl::sor<group_attribute, attribute>,
                                     pegtl::star<attribute_separator, pegtl::sor<group_attribute, attribute>>>;

/**
 * GRAMMAR OF AN ATTRIBUTES
 */
struct attributes : pegtl::seq<at<pegtl::one<'@'>>, attributes_clause> {
    using predication = at<pegtl::one<'@'>>;
    using content = attributes_clause;
};
}
}

#endif  // MOJO_GRAMMAR_ATTRIBUTE_HPP