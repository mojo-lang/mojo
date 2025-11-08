#ifndef LANG_STATEMENT_MOJO_HPP
#define LANG_STATEMENT_MOJO_HPP

#include <mojo/core/union.hpp>
#include <mojo/lang/declaration.mojo.hpp>
#include <mojo/lang/expression.mojo.hpp>

namespace mojo {
namespace lang {

using Statement = Union<Expression, Declaration>;
}
}  // namespace mojo

#endif  // LANG_STATEMENT_MOJO_HPP
