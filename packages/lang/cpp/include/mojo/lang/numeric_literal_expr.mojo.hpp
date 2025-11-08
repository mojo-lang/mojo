#ifndef LANG_NUMERIC_LITERAL_EXPR_MOJO_HPP
#define LANG_NUMERIC_LITERAL_EXPR_MOJO_HPP

#include <mojo/core/bool.hpp>
#include <mojo/lang/literal_expr.mojo.hpp>

namespace mojo {
namespace lang {

struct NumericLiteralExpr : LiteralExpr {
    Bool isNegative = false;
};

}
}

#endif //LANG_NUMERIC_LITERAL_EXPR_MOJO_HPP
