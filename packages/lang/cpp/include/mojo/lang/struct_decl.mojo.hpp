#ifndef MOJO_LANG_STRUCT_DECL_MOJO_HPP
#define MOJO_LANG_STRUCT_DECL_MOJO_HPP

#include <mojo/lang/struct_type_builder.hpp>
#include <mojo/lang/type_decl.mojo.hpp>
#include <mojo/lang/enum_decl.mojo.hpp>
#include <mojo/lang/nominal_type.mojo.hpp>
#include <mojo/lang/type_alias_decl.mojo.hpp>
#include <mojo/lang/struct_decl_ptr.hpp>

namespace mojo {
namespace lang {

/**
 *
 */
struct StructDecl : TypeDecl {
    StructType* type;

    StructDeclPtr enclosingDecl;//包含此struct的外部struct

    Array<TypeAliasDecl> typeAliasDecls;

    Array<StructDeclPtr> structDecls;

    Array<std::shared_ptr<EnumDecl>> enumDecls;
};

}
}

MOJO_BUILD_STRUCT_TYPE(mojo, lang, StructDecl) {
    field(&mojo::lang::StructDecl::type, "type", 1);
    field(&mojo::lang::StructDecl::enclosingDecl, "enclosing_decl", "enclosingDecl", 2);
    field(&mojo::lang::StructDecl::typeAliasDecls, "type_alias_decls", "typeAliasDecls", 3);
    field(&mojo::lang::StructDecl::structDecls, "struct_decls", "structDecls", 4);
    field(&mojo::lang::StructDecl::enumDecls, "enum_decls", "enumDecls", 5);
}

#endif //MOJO_LANG_STRUCT_DECL_MOJO_HPP
