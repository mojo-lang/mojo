#ifndef MOJO_GRAMMAR_PATTERN_HPP
#define MOJO_GRAMMAR_PATTERN_HPP

#include <pegtl.hh>
#include <mojo/grammar/identifier.hpp>
#include <mojo/grammar/keys.hpp>
#include <mojo/grammar/type.hpp>

namespace mojo {
namespace grammar {

struct pattern;

/**
 * GRAMMAR OF A WILDCARD PATTERN
 */
struct wildcard_pattern : pegtl::seq<key_wildcard, pre_pad_opt<type_annotation, sep>> {};

/**
 * GRAMMAR OF AN IDENTIFIER PATTERN
 */
struct identifier_pattern : pegtl::seq<identifier, pre_pad_opt<type_annotation, sep>> {};

/**
 * GRAMMAR OF A TUPLE PATTERN
 */
struct tuple_pattern_element_separator : pegtl::one<':'> {};
struct tuple_pattern_element
    : pegtl::seq<pegtl::opt<identifier, seps, tuple_pattern_element_separator, seps>, pattern> {
    using separator = tuple_pattern_element_separator;
};

struct tuple_pattern_content : pegtl::list<tuple_pattern_element, pegtl::one<','>, sep> {};
struct tuple_pattern_clause
    : pegtl::seq<pegtl::one<'('>, pegtl::pad_opt<tuple_pattern_content, sep>, pegtl::one<')'>> {
    using content = tuple_pattern_content;
};

struct tuple_pattern : pegtl::seq<tuple_pattern_clause, pegtl::opt<type_annotation>> {
    using content = tuple_pattern_clause::content;
};

/**
 * GRAMMAR OF AN ENUMERATION CASE PATTERN
 */
struct enum_value_pattern_separator : pegtl::one<'.'> {};
struct enum_value_pattern : pegtl::seq<type_identifier, enum_value_pattern_separator, identifier> {
    using separator = enum_value_pattern_separator;
};

/**
 * GRAMMAR OF AN OPTIONAL PATTERN
 */
struct optional_pattern : pegtl::seq<identifier, pegtl::one<'?'>> {};

/**
 * GRAMMAR OF AN EXPRESSION PATTERN
 */
struct expression_pattern : expression {};

struct is_pattern : pegtl::seq<key_is, seps, type> {};

struct as_pattern
    : pegtl::seq<
          pegtl::sor<key_wildcard, identifier, tuple_pattern_clause, enum_value_pattern, expression_pattern>,
          seps,
          key_as,
          seps,
          type> {};

/**
 * GRAMMAR OF A TYPE CASTING PATTERN
 */
struct type_casting_pattern : is_pattern /* pegtl::sor<is_pattern, as_pattern>*/ {};

struct destructuring_pattern : pegtl::sor<wildcard_pattern, identifier_pattern, tuple_pattern> {};

/**
 * GRAMMAR OF A PATTERN
 */
struct pattern
    : pegtl::sor<type_casting_pattern, destructuring_pattern, enum_value_pattern /*, expression_pattern*/> {};
}
}
#endif  // MOJO_GRAMMAR_PATTERN_HPP