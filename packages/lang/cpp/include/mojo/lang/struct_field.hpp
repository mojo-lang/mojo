#ifndef MOJO_LANG_STRUCT_FIELD_HPP
#define MOJO_LANG_STRUCT_FIELD_HPP

#include <mojo/lang/value_decl.mojo.hpp>

namespace mojo {
namespace lang {

struct StructField : ValueDecl {
    UInt index = 0;

    StructField() {
    }
};

}  // namespace lang
}  // namespace mojo

#endif  // MOJO_LANG_STRUCT_FIELD_HPP
