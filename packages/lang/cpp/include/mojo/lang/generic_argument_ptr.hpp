#ifndef LANG_GENERIC_ARGUMENT_PTR_HPP
#define LANG_GENERIC_ARGUMENT_PTR_HPP

#include <memory>

namespace mojo {
    namespace lang {

        struct NominalType;
        using GenericArgumentPtr = std::shared_ptr<NominalType>;

    }  // namespace lang
}  // namespace mojo

#endif  // LANG_GENERIC_ARGUMENT_PTR_HPP
