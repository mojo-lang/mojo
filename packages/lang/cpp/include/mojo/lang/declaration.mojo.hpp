#ifndef LANG_DECLARATION_MOJO_HPP
#define LANG_DECLARATION_MOJO_HPP

#include <mojo/core/union.hpp>
#include <mojo/lang/data_decl_ptr.hpp>
#include <mojo/lang/enum_decl_ptr.hpp>
#include <mojo/lang/generic_parameter_ptr.hpp>
#include <mojo/lang/interface_decl_ptr.hpp>
#include <mojo/lang/struct_decl_ptr.hpp>
#include <mojo/lang/type_alias_decl_ptr.hpp>

namespace mojo {
namespace lang {

/**
 *
 */
using Declaration = Union<DataDeclPtr,
                          EnumDeclPtr,
                          StructDeclPtr,
                          InterfaceDeclPtr,
                          TypeAliasDeclPtr,
                          GenericParameterPtr>;

}  // namespace lang
}  // namespace mojo

#endif  // LANG_DECLARATION_MOJO_HPP
