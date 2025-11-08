#ifndef LANG_TYPE_DECLARATION_MOJO_HPP
#define LANG_TYPE_DECLARATION_MOJO_HPP

#include <mojo/core/union.hpp>
#include <mojo/lang/data_decl_ptr.hpp>
#include <mojo/lang/enum_decl_ptr.hpp>
#include <mojo/lang/interface_decl_ptr.hpp>
#include <mojo/lang/generic_parameter_ptr.hpp>
#include <mojo/lang/struct_decl_ptr.hpp>
#include <mojo/lang/type_alias_decl_ptr.hpp>

namespace mojo {
namespace lang {

using TypeDeclaration =
    Union<DataDeclPtr, EnumDeclPtr, StructDeclPtr, InterfaceDeclPtr, GenericParameterPtr>;

bool isDataDecl(const TypeDeclaration& typeDeclaration) {
    auto const& decl = typeDeclaration.get<DataDeclPtr>();
    return decl && *decl;
}

}
}  // namespace mojo

#endif  // LANG_TYPE_DECLARATION_MOJO_HPP
