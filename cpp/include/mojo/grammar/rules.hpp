#ifndef MOJO_GRAMMAR_RULES_HPP
#define MOJO_GRAMMAR_RULES_HPP

#include <pegtl.hh>
#include <pegtl/ascii.hh>
#include <mojo/grammar/comment.hpp>

namespace mojo {
namespace grammar {

struct blanks : pegtl::star<pegtl::ascii::blank> {};

struct sep : pegtl::sor<pegtl::ascii::space, comment> {};
struct seps : pegtl::star<sep> {};

struct inline_sep : pegtl::sor<pegtl::ascii::blank, block_comment> {};
struct inline_seps : pegtl::star<inline_sep> {};

template <char Char, char AtChar>
struct separator
    : pegtl::seq<inline_seps, pegtl::sor<pegtl::seq<seps, pegtl::one<Char>>, pegtl::eolf, pegtl::at<pegtl::one<AtChar>>>> {
};

template <char Char>
struct separator<Char, '_'>
    : pegtl::seq<inline_seps, pegtl::sor<pegtl::seq<seps, pegtl::one<Char>>, pegtl::eolf>> {
};

// support line break in separator
template <typename Rule, char Char, char AtChar = '_'>
using list = pegtl::seq<Rule, pegtl::star<separator<Char, AtChar>, seps, Rule>>;

using value_separator = separator<',', '_'>;
using statement_separator = separator<';', '}'>;

template <typename Rule, typename Pad>
using pre_pad_opt = pegtl::opt<pegtl::star<Pad>, Rule>;

template <typename Rule, typename Sep, typename Pad>
using list_tail =
    pegtl::seq<pegtl::list<Rule, pegtl::seq<Sep, pegtl::star<Pad>>>, pegtl::opt<pegtl::star<Pad>, Sep>>;

template <typename S, typename O>
struct left_assoc : pegtl::seq<S, seps, pegtl::star<pegtl::if_must<O, seps, S, seps>>> {};

template <typename S, typename O>
struct right_assoc : pegtl::seq<S, seps, pegtl::opt<pegtl::if_must<O, seps, right_assoc<S, O>>>> {};
}
}

#endif  // MOJO_GRAMMAR_RULES_HPP
