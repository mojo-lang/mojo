#ifndef LANG_STRING_LITERAL_EXPR_MOJO_HPP
#define LANG_STRING_LITERAL_EXPR_MOJO_HPP

#include <mojo/core/string.hpp>
#include <mojo/lang/literal_expr.mojo.hpp>

namespace mojo {
namespace lang {

struct StringLiteralExpr : LiteralExpr {
    String value;
};

}
}

#endif //LANG_STRING_LITERAL_EXPR_MOJO_HPP
