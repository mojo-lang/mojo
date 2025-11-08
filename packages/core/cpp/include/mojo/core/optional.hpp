#ifndef MOJO_CORE_OPTIONAL_HPP
#define MOJO_CORE_OPTIONAL_HPP

#include <ncraft/core/optional.hpp>

namespace mojo {
namespace core {

template <typename T, typename Enable = void>
using Optional = ::ncraft::core::Optional<T, Enable>;

}

template <typename T, typename Enable = void>
using Optional = core::Optional<T, Enable>;

}

#endif //MOJO_CORE_OPTIONAL_HPP
