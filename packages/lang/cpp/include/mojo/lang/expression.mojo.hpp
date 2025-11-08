#ifndef LANG_EXPRESSION_MOJO_HPP
#define LANG_EXPRESSION_MOJO_HPP

#include <mojo/core/union.hpp>
#include <mojo/lang/integer_literal_expr.mojo.hpp>
#include <mojo/lang/string_literal_expr.mojo.hpp>

namespace mojo {
namespace lang {

using Expression = Union<IntegerLiteralExpr,
                         StringLiteralExpr>;

}
}

#endif //LANG_EXPRESSION_MOJO_HPP
