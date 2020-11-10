#ifndef MOJO_GRAMMAR_DOCUMENT_HPP
#define MOJO_GRAMMAR_DOCUMENT_HPP

#include <pegtl.hh>

namespace mojo {
namespace grammar {

/**
 * GRAMMAR OF A DOCUMENT_COMMENT
 */
struct document_string : pegtl::until<pegtl::at<pegtl::eolf>> {};
struct document_line_begin : pegtl_string_t("///") {};
struct document_line : pegtl::seq<document_line_begin, document_string> {};

struct document_separator : pegtl::seq<pegtl::star<pegtl::eolf>, blanks> {};

struct document
    : pegtl::seq<document_line, pegtl::star<document_separator, document_line>> {};

struct following_document_string : document_string {};
struct following_document_begin : pegtl_string_t("//<") {};
struct following_document
    : pegtl::seq<following_document_begin, following_document_string> {
    using begin = following_document_begin;
};
}
}

#endif  // MOJO_GRAMMAR_DOCUMENT_HPP