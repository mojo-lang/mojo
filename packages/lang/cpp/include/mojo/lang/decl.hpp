#ifndef LANG_DECL_HPP
#define LANG_DECL_HPP

#include <ncraft/traits/traits.hpp>
#include <mojo/lang/data_type.hpp>

namespace mojo {
namespace lang {

// Implement this:
// const Decl* buildDecl(const TypeIdentifier<T>* null);

template <typename T>
auto getDecl() -> decltype(buildDecl(::std::declval<const TypeIdentifier<ncraft::traits::RemoveCVRef<T>>*>())) {
    static const TypeIdentifier<ncraft::traits::RemoveCVRef<T>> typeIdentifier;
    static const auto t = buildDecl(&typeIdentifier);
    return t;
}

}
}

#endif //LANG_DECL_HPP
