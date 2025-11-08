#ifndef LANG_INTEGER_LITERAL_EXPR_MOJO_HPP
#define LANG_INTEGER_LITERAL_EXPR_MOJO_HPP

#include <mojo/core/numeric.hpp>
#include <mojo/lang/numeric_literal_expr.mojo.hpp>

namespace mojo {
namespace lang {

struct IntegerLiteralExpr : NumericLiteralExpr {
    UInt64 value = 0;
};

}
}

#endif //LANG_INTEGER_LITERAL_EXPR_MOJO_HPP
