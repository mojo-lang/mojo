#ifndef MOJO_LANG_ATTRIBUTE_MOJO_HPP
#define MOJO_LANG_ATTRIBUTE_MOJO_HPP
/*
 *
 */

#include <mojo/core/array.hpp>
#include <mojo/lang/expression.mojo.hpp>
#include <mojo/lang/package_ptr.hpp>
#include <mojo/lang/generic_argument_ptr.hpp>

namespace mojo {
namespace lang {

/**
 *
 */
struct Attribute {
    /**
     *
     */
    String name;

    /**
     *
     */
    Array<GenericArgumentPtr> generics;

    /**
     *
     */
    Array<Expression> arguments;

    /**
     *
     */
    PackagePtr package;
};

}  // namespace lang
}  // namespace mojo

#endif  // MOJO_LANG_ATTRIBUTE_MOJO_HPP
