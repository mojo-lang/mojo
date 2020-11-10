#ifndef MOJO_GRAMMAR_COMMENT_HPP
#define MOJO_GRAMMAR_COMMENT_HPP

#include <pegtl/rules.hh>

namespace mojo {
namespace grammar {

struct line_comment : pegtl::disable<pegtl_string_t("//"),
                                      pegtl::not_at<pegtl::one<'/', '<'>>,
                                      pegtl::until<pegtl::eolf>> {};

//TODO: fix nest comment
struct block_comment : pegtl::disable<pegtl_string_t("/*"), pegtl::until<pegtl_string_t("*/")>> {};

struct comment : pegtl::sor<line_comment, block_comment> {};
}
}

#endif  // MOJO_GRAMMAR_COMMENT_HPP