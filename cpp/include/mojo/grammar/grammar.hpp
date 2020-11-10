#ifndef MOJO_GRAMMAR_GRAMMAR_HPP
#define MOJO_GRAMMAR_GRAMMAR_HPP

#include <pegtl.hh>
#include <mojo/grammar/statement.hpp>
#include <mojo/grammar/declaration.hpp>
#include <mojo/grammar/lamda_expression.hpp>
#include <mojo/grammar/match_expression.hpp>
#include <mojo/grammar/projection_expression.hpp>

namespace mojo {
namespace grammar {

/**
 * GRAMMAR OF THE MOJO
 */
struct grammar : pegtl::pad<statements, sep> {};
}
}

#endif  // MOJO_GRAMMAR_GRAMMAR_HPP
