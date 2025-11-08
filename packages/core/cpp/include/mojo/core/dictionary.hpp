#ifndef MOJO_CORE_DICTIONARY_HPP
#define MOJO_CORE_DICTIONARY_HPP

#include <unordered_map>

namespace mojo {
namespace core {

template<typename K, typename V>
using Map = ::std::unordered_map<K, V>;

}

template <typename K, typename V>
using Map = ::mojo::core::Map<K, V>;

}

#endif //MOJO_CORE_DICTIONARY_HPP
