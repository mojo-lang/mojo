#ifndef LANG_NOMINAL_TYPE_MOJO_HPP
#define LANG_NOMINAL_TYPE_MOJO_HPP

#include <mojo/core/array.hpp>
#include <mojo/core/string.hpp>
#include <mojo/lang/attribute.mojo.hpp>
#include <mojo/lang/nominal_type_ptr.hpp>
#include <mojo/lang/type_declaration.mojo.hpp>

namespace mojo {
namespace lang {

/**
 *
 */
struct NominalType {
    /**
     *
     */
    String name;

    /**
     *
     */
    Union<String, PackagePtr> package;

    /**
     *
     */
    TypeDeclaration typeDeclaration;

    /**
     *
     */
    Array<NominalTypePtr> genericArguments;

    /**
     *
     */
    Array<Attribute> attributes;

    /**
     *
     */
    NominalTypePtr enclosingType;

    /**
     *
     */
    NominalType() : typeDeclaration(DataDeclPtr{}) {
    }
};

}  // namespace lang
}  // namespace mojo

#endif  // LANG_NOMINAL_TYPE_MOJO_HPP