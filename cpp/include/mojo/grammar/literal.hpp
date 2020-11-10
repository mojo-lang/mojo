#ifndef MOJO_GRAMMAR_LITERAL_HPP
#define MOJO_GRAMMAR_LITERAL_HPP

#include <pegtl/rules.hh>
#include <pegtl/utf8.hh>

namespace mojo {
namespace grammar {

struct decimal_digits;
struct integer_literal;
struct floating_point_literal;
struct string_literal;

/**
 * GRAMMAR OF A NUMERIC LITERAL
 */
struct numeric_literal
    : pegtl::seq<pegtl::opt<pegtl::one<'-'>>, pegtl::sor<integer_literal, floating_point_literal>> {};

/**
 * GRAMMAR OF A BOOL LITERAL
 */
struct bool_literal : pegtl::sor<key_true, key_false> {};

/**
 * GRAMMAR OF A NULL LITERAL
 */
struct null_literal : key_null {};

/**
 * GRAMMAR OF A LITERAL
 */
struct literal : pegtl::sor<numeric_literal, string_literal, bool_literal, null_literal> {};

struct binary_literal;
struct octal_literal;
struct decimal_literal;
struct hexadecimal_literal;

/**
 * GRAMMAR OF AN INTEGER LITERAL
 */
struct integer_literal : pegtl::sor<binary_literal,
                                    octal_literal,
                                    hexadecimal_literal,
                                    pegtl::seq<decimal_literal, pegtl::not_at<pegtl::one<'.', 'e', 'E'>>>> {};

struct binary_digit : pegtl::ascii::range<'0', '1'> {};
struct binary_literal_character : pegtl::sor<binary_digit, pegtl::one<'_'>> {};

/**
 * GRAMMAR OF AN BINARY INTEGER LITERAL
 */
struct binary_literal
    : pegtl::seq<pegtl_string_t("0b"), binary_digit, pegtl::star<binary_literal_character>> {};

struct octal_digit : pegtl::ascii::range<'0', '7'> {};
struct octal_digit_character : pegtl::sor<octal_digit, pegtl::one<'_'>> {};

/**
 * GRAMMAR OF AN OCTAL INTEGER LITERAL
 */
struct octal_literal : pegtl::seq<pegtl_string_t("0o"), octal_digit, pegtl::star<octal_digit_character>> {};

struct decimal_digit : pegtl::ascii::digit {};
struct decimal_digits : pegtl::plus<decimal_digit> {};
struct decimal_literal_character : pegtl::sor<decimal_digit, pegtl::one<'_'>> {};

/**
 * GRAMMAR OF AN HEXADECIMAL INTEGER LITERAL
 */
struct decimal_literal : pegtl::seq<decimal_digit, pegtl::star<decimal_literal_character>> {};

struct hexadecimal_literal_character : pegtl::sor<pegtl::ascii::xdigit, pegtl::one<'_'>> {};
struct hexadecimal_literal
    : pegtl::seq<pegtl_string_t("0x"), pegtl::ascii::xdigit, pegtl::star<hexadecimal_literal_character>> {};

/**
 * GRAMMAR OF A FLOATING-POINT LITERAL
 */
struct sign : pegtl::one<'+', '-'> {};
struct floating_point_e : pegtl::one<'e', 'E'> {};
struct decimal_exponent : pegtl::seq<floating_point_e, pegtl::opt<sign>, decimal_literal> {};
struct decimal_fraction : pegtl::seq<pegtl::one<'.'>, decimal_literal> {};

struct floating_point_literal
    : pegtl::seq<decimal_literal, pegtl::opt<decimal_fraction>, pegtl::opt<decimal_exponent>> {};

}
}

#endif  // MOJO_GRAMMAR_LITERAL_HPP