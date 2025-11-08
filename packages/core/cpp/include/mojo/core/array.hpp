#ifndef MOJO_CORE_ARRAY_HPP
#define MOJO_CORE_ARRAY_HPP

#include <vector>

namespace mojo {
namespace core {

template<typename T>
using Array = ::std::vector<T>;

}

template<typename T>
using Array = ::mojo::core::Array<T>;

}

#endif //MOJO_CORE_ARRAY_HPP
