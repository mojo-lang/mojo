#ifndef LANG_TYPE_ALIAS_DECL_MOJO_HPP
#define LANG_TYPE_ALIAS_DECL_MOJO_HPP

#include <mojo/lang/nominal_type.mojo.hpp>
#include <mojo/lang/type_decl.mojo.hpp>
#include <mojo/lang/type_alias_decl_ptr.hpp>

namespace mojo {
namespace lang {

/**
 *
 */
struct TypeAliasDecl : TypeDecl {
    NominalType type;
};

}  // namespace lang
}  // namespace mojo

#endif  // LANG_TYPE_ALIAS_DECL_MOJO_HPP
