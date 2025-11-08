#ifndef MOJO_LANG_TYPE_DECL_MOJO_HPP
#define MOJO_LANG_TYPE_DECL_MOJO_HPP

#include <mojo/lang/decl.mojo.hpp>
#include <mojo/lang/attribute.mojo.hpp>
#include <mojo/lang/struct_type_builder.hpp>
#include <mojo/lang/generic_parameter_ptr.hpp>

namespace mojo {
namespace lang {

struct TypeDecl : Decl {
    String name;

    Array<GenericParameterPtr> genericParams;

    Array<Attribute> attributes;
};

}
}

MOJO_BUILD_STRUCT_TYPE(mojo, lang, TypeDecl) {
    field(&mojo::lang::TypeDecl::name, "name", 10);
    //field(&mojo::lang::TypeDecl::type, "type", 1);
    field(&mojo::lang::TypeDecl::attributes, "attributes", 12);
}

#endif //MOJO_LANG_TYPE_DECL_MOJO_HPP
