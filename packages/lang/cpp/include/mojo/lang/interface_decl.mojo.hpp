#ifndef LANG_INTERFACE_DECL_MOJO_HPP
#define LANG_INTERFACE_DECL_MOJO_HPP

#include <mojo/lang/enum_decl.mojo.hpp>
#include <mojo/lang/interface_type.hpp>
#include <mojo/lang/nominal_type.mojo.hpp>
#include <mojo/lang/struct_type_builder.hpp>
#include <mojo/lang/type_alias_decl.mojo.hpp>
#include <mojo/lang/type_decl.mojo.hpp>

namespace mojo {
namespace lang {

/**
 *
 */
struct InterfaceDecl : TypeDecl {
    InterfaceType* type;
    Array<TypeAliasDecl> typeAliasDecls;
};

}  // namespace lang
}  // namespace mojo

MOJO_BUILD_STRUCT_TYPE(mojo, lang, InterfaceDecl) {
    field(&mojo::lang::InterfaceDecl::type, "type", 1);
    field(&mojo::lang::InterfaceDecl::typeAliasDecls,
          "type_alias_decls",
          "typeAliasDecls",
          1);
}

#endif  // LANG_INTERFACE_DECL_MOJO_HPP
