#ifndef LANG_FUNC_TYPE_MOJO_HPP
#define LANG_FUNC_TYPE_MOJO_HPP

#include <mojo/core/array.hpp>
#include <mojo/lang/value_decl.mojo.hpp>
#include <mojo/lang/func_decl_ptr.hpp>

namespace mojo {
namespace lang {

/**
 *
 */
struct FuncType {
    /**
     *
     */
    Array<ValueDecl> parameters;

    /**
     *
     */
    std::unique_ptr<NominalType> returnType;
};

}
}

#endif //LANG_FUNC_TYPE_MOJO_HPP
