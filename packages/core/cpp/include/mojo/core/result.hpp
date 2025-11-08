#ifndef MOJO_CORE_RESULT_HPP
#define MOJO_CORE_RESULT_HPP

#include <ncraft/core/result.hpp>

namespace mojo {
namespace core {

using Status = ::ncraft::core::Status;

template <typename... Ts>
using Result = ::ncraft::core::Result<Ts...>;

}

using Status = core::Status;

template <typename... Ts>
using Result = core::Result<Ts...>;

}

#endif //MOJO_CORE_RESULT_HPP
