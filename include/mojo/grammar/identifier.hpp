#ifndef MOJO_GRAMMAR_IDENTIFIER_HPP
#define MOJO_GRAMMAR_IDENTIFIER_HPP

#include <pegtl.hh>
#include <pegtl/utf8.hh>
#include <mojo/grammar/keys.hpp>

namespace mojo {
namespace grammar {

struct decimal_digits;

struct identifier_head
    : pegtl::sor<pegtl::ascii::lower,
                 pegtl::utf8::ranges<0xA1, 0xFF>,
                 pegtl::utf8::ranges<0x0100, 0x02FF, 0x0370, 0x167F, 0x1681, 0x180D, 0x180F, 0x1DBF, 0x1E00, 0x1FFF>,
                 pegtl::utf8::ranges<0x200B, 0x200D, 0x202A, 0x20CF, 0x2100, 0x23FF, 0x2460, 0x2FFF, 0x3001, 0xD7FF>,
                 pegtl::utf8::ranges<0xF900, 0xFD3D, 0xFD40, 0xFDCF, 0xFDF0, 0xFE1F, 0xFE30, 0xFE44, 0xFE47, 0xFFFD>,
                 pegtl::utf8::ranges<0x10000, 0x1FFFD, 0x20000, 0x2FFFD, 0x30000, 0x3FFFD, 0x40000, 0x4FFFD>,
                 pegtl::utf8::ranges<0x50000, 0x5FFFD, 0x60000, 0x6FFFD, 0x70000, 0x7FFFD, 0x80000, 0x8FFFD>,
                 pegtl::utf8::ranges<0x90000, 0x9FFFD, 0xA0000, 0xAFFFD, 0xB0000, 0xBFFFD, 0xC0000, 0xCFFFD>,
                 pegtl::utf8::ranges<0xD0000, 0xDFFFD, 0xE0000, 0xEFFFD>> {};

struct identifier_character
    : pegtl::sor<identifier_head,
                 pegtl::ascii::upper,
                 pegtl::ascii::digit,
                 pegtl::one<'_'>,
                 pegtl::utf8::ranges<0x0300, 0x036F, 0x1DC0, 0x1DFF, 0x20D0, 0x20FF, 0xFE20, 0xFE2F>> {};

struct identifier_characters : pegtl::plus<identifier_character> {};

struct implicit_parameter_name : pegtl::seq<pegtl::one<'$'>, pegtl::opt<decimal_digits>> {};

/**
 * GRAMMAR OF AN IDENTIFIER
 */
struct identifier : pegtl::seq<pegtl::not_at<keyword>,
                               pegtl::sor<pegtl::seq<identifier_head, pegtl::opt<identifier_characters>>,
                                          implicit_parameter_name>> {};
}
}

#endif  // MOJO_GRAMMAR_IDENTIFIER_HPP