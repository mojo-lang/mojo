#ifndef MOJO_LANG_SOURCE_FILE_HPP
#define MOJO_LANG_SOURCE_FILE_HPP

#include <mojo/core/array.hpp>
#include <mojo/core/map.hpp>
#include <mojo/core/string.hpp>
#include <mojo/lang/data_decl_ptr.hpp>
#include <mojo/lang/enum_decl_ptr.hpp>
#include <mojo/lang/func_decl_ptr.hpp>
#include <mojo/lang/interface_decl_ptr.hpp>
#include <mojo/lang/statement.mojo.hpp>
#include <mojo/lang/struct_decl_ptr.hpp>
#include <mojo/lang/package_ptr.hpp>

namespace mojo {
namespace lang {

struct SourceFile {
    String name;

    String suffixName;

    PackagePtr package;

    // full_name: String @2

    /**
     * statements in the source file one by one.
     */
    Array<Statement> statements;

    /**
     * decls' index
     */
    // import_decls:    [ImportDecl] @10 @linked
    // const_decls:     [ConstDecl] @11 @linked
    // attribute_decls: [AttributeDecl] @12 @linked

    Map<String, DataDeclPtr> dataDecls;
    Map<String, EnumDeclPtr> enumDecls;
    Map<String, StructDeclPtr> structDecls;
    Map<String, InterfaceDeclPtr> interfaceDecls;
    Map<String, FuncDeclPtr> funcDecls;
};

}  // namespace lang
}  // namespace mojo

#endif  // MOJO_LANG_SOURCE_FILE_HPP
