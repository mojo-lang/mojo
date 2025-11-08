#ifndef MOJO_LANG_SERVICE_TYPE_HPP
#define MOJO_LANG_SERVICE_TYPE_HPP

#include <ncraft/core/noncopyable.hpp>
#include <mojo/lang/data_type.hpp>
#include <mojo/lang/func_decl.mojo.hpp>
#include <mojo/lang/struct_type_builder.hpp>

namespace mojo {
namespace lang {

/**
 *
 */
class InterfaceType : ncraft::NonCopyable {
public:
    /**
     *
     */
    String name;

    /**
     *
     */
    Array<::std::unique_ptr<NominalType>> inherits;

    /**
     *
     */
    Array<FuncDeclPtr> methods;
};

}  // namespace lang
}  // namespace mojo

MOJO_BUILD_STRUCT_TYPE(mojo, lang, InterfaceType) {
    field(&::mojo::lang::InterfaceType::name, "name", 1);
    field(&::mojo::lang::InterfaceType::inherits, "inherits", 2);
    field(&::mojo::lang::InterfaceType::methods, "methods", 3);
}

#endif  // MOJO_LANG_SERVICE_TYPE_HPP
