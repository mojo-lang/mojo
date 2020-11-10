#ifndef MOJO_GRAMMAR_KEYS_HPP
#define MOJO_GRAMMAR_KEYS_HPP

#include <pegtl.hh>
#include <mojo/grammar/rules.hpp>

namespace mojo {
namespace grammar {

/**
 *
 */
struct str_and : pegtl_string_t("and") {};
struct str_as : pegtl_string_t("as") {};
struct str_attribute : pegtl_string_t("attribute") {};
struct str_break : pegtl_string_t("break") {};
struct str_const : pegtl_string_t("const") {};
struct str_continue : pegtl_string_t("continue") {};
struct str_else : pegtl_string_t("else") {};
struct str_enum : pegtl_string_t("enum") {};
struct str_false : pegtl_string_t("false") {};
struct str_for : pegtl_string_t("for") {};
struct str_func : pegtl_string_t("func") {};
struct str_if : pegtl_string_t("if") {};
struct str_import : pegtl_string_t("import") {};
struct str_in : pegtl_string_t("in") {};
struct str_interface : pegtl_string_t("interface") {};
struct str_is : pegtl_string_t("is") {};
struct str_match : pegtl_string_t("match") {};
struct str_not : pegtl_string_t("not") {};
struct str_null : pegtl_string_t("null") {};
struct str_or : pegtl_string_t("or") {};
struct str_package : pegtl_string_t("package") {};
struct str_repeat : pegtl_string_t("repeat") {};
struct str_return : pegtl_string_t("return") {};
struct str_true : pegtl_string_t("true") {};
struct str_type : pegtl_string_t("type") {};
struct str_var : pegtl_string_t("var") {};
struct str_while : pegtl_string_t("while") {};
struct str_xor : pegtl_string_t("xor") {};
struct str_wildcard : pegtl_string_t("_") {};

struct str_keyword : pegtl::sor<str_and,
                                str_as,
                                str_break,
                                str_const,
                                str_continue,
                                str_else,
                                str_false,
                                str_for,
                                str_if,
                                str_import,
                                str_is,
                                str_match,
                                str_not,
                                str_null,
                                str_or,
                                str_repeat,
                                str_return,
                                str_true,
                                str_var,
                                str_while,
                                str_xor> {};

struct identifier_character;

template <typename Key>
struct key : pegtl::seq<Key, pegtl::not_at<identifier_character>> {};

struct key_and : key<str_and> {};
struct key_attribute : key<str_attribute> {};
struct key_as : key<str_as> {};
struct key_break : key<str_break> {};
struct key_const : key<str_const> {};
struct key_continue : key<str_continue> {};
struct key_else : key<str_else> {};
struct key_enum : key<str_enum> {};
struct key_false : key<str_false> {};
struct key_for : key<str_for> {};
struct key_func : key<str_func> {};
struct key_if : key<str_if> {};
struct key_import : key<str_import> {};
struct key_in : key<str_in> {};
struct key_interface : key<str_interface> {};
struct key_is : key<str_is> {};
struct key_match : key<str_match> {};
struct key_not : key<str_not> {};
struct key_null : key<str_null> {};
struct key_or : key<str_or> {};
struct key_package : key<str_package> {};
struct key_repeat : key<str_repeat> {};
struct key_return : key<str_return> {};
struct key_true : key<str_true> {};
struct key_type : key<str_type> {};
struct key_var : key<str_var> {};
struct key_while : key<str_while> {};
struct key_wildcard : key<str_wildcard> {};
struct key_xor : key<str_xor> {};

struct keyword : key<str_keyword> {};
}
}
#endif  // MOJO_GRAMMAR_KEYS_HPP
