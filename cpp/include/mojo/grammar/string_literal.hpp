#ifndef MOJO_GRAMMAR_STRING_LITERAL_HPP
#define MOJO_GRAMMAR_STRING_LITERAL_HPP

#include <pegtl/rules.hh>
#include <pegtl/utf8.hh>

namespace mojo {
namespace grammar {

struct expression;
struct decimal_digits;

struct integer_literal;
struct floating_point_literal;
struct string_literal;

struct escaped_character
    : pegtl::sor<pegtl_string_t("\\0"),
                 pegtl_string_t("\\\\"),
                 pegtl_string_t("\\t"),
                 pegtl_string_t("\\n"),
                 pegtl_string_t("\\r"),
                 pegtl_string_t("\\\""),
                 pegtl_string_t("\\'"),
                 pegtl::seq<pegtl_string_t("\\u"), pegtl::rep<4, pegtl::ascii::xdigit>>,
                 pegtl::seq<pegtl_string_t("\\u{"),
                            pegtl::rep_min_max<1, 8, pegtl::ascii::xdigit>,
                            pegtl_string_t("}")>> {};

struct double_quoted_range
    : pegtl::ascii::ranges<0x0, 0x09, 0xB, 0x0C, 0x0E, 0x21, 0x23, 0x5B> {};
struct single_quoted_range
    : pegtl::ascii::ranges<0x0, 0x09, 0xB, 0x0C, 0x0E, 0x26, 0x28, 0x5B> {};

struct quoted_text_item
    : pegtl::sor<escaped_character, pegtl::utf8::ranges<0x5D, 0xD7FF, 0xE000, 0x10FFFF>> {
};

struct interpolated_text_item
    : pegtl::sor<
          pegtl::seq<pegtl_string_t("\{"), seps, expression, seps, pegtl::one<'}'>>,
          quoted_text_item> {};

template <typename Item>
struct double_quoted_basic_string_chars
    : pegtl::star<pegtl::sor<double_quoted_range, Item>> {};

template <typename Item>
struct single_quoted_basic_string_chars
    : pegtl::star<pegtl::sor<single_quoted_range, Item>> {};

struct double_quoted_string_begin : pegtl::one<'"'> {};
struct double_quoted_string_end : pegtl::one<'"'> {};

struct single_quoted_string_begin : pegtl::one<'\''> {};
struct single_quoted_string_end : pegtl::one<'\''> {};

template <typename Item>
struct double_quoted_basic_string : pegtl::seq<double_quoted_string_begin,
                                               double_quoted_basic_string_chars<Item>,
                                               double_quoted_string_end> {
    using begin = double_quoted_string_begin;
    using end = double_quoted_string_end;
    using content = double_quoted_basic_string_chars<Item>;
};

template <typename Item>
struct single_quoted_basic_string : pegtl::seq<single_quoted_string_begin,
                                               single_quoted_basic_string_chars<Item>,
                                               single_quoted_string_end> {
    using begin = single_quoted_string_begin;
    using end = single_quoted_string_end;
    using content = single_quoted_basic_string_chars<Item>;
};

struct double_quoted_static_string : double_quoted_basic_string<quoted_text_item> {};
struct single_quoted_static_string : single_quoted_basic_string<quoted_text_item> {};

struct static_string_literal
    : pegtl::sor<double_quoted_static_string, single_quoted_static_string> {};

struct interpolated_string_literal {};

/**
 * GRAMMAR OF A STRING LITERAL
 */
struct string_literal : static_string_literal {};

// pegtl::sor<static_string_literal, interpolated_string_literal> {};
}
}

#endif  // MOJO_GRAMMAR_STRING_LITERAL_HPP
