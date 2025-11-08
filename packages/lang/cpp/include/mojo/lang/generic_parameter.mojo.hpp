#ifndef LANG_GENERIC_PARAMETER_MOJO_HPP
#define LANG_GENERIC_PARAMETER_MOJO_HPP

#include <mojo/lang/generic_parameter_ptr.hpp>
#include <mojo/lang/nominal_type_ptr.hpp>
#include <mojo/lang/struct_type_builder.hpp>
#include <mojo/lang/type_decl.mojo.hpp>

namespace mojo {
namespace lang {

/**
 *
 */
struct GenericParameter : TypeDecl {
    /**
     *
     */
    NominalTypePtr constraint;
};

}  // namespace lang
}  // namespace mojo

MOJO_BUILD_STRUCT_TYPE(mojo, lang, GenericParameter) {
    field(&::mojo::lang::GenericParameter::constraint, "constraint", 1);
}

#endif  // LANG_GENERIC_PARAMETER_MOJO_HPP
