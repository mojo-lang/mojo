#ifndef MOJO_DICTIONARY_TYPE_HPP
#define MOJO_DICTIONARY_TYPE_HPP

#include <mojo/lang/data_type_of.hpp>
#include <mojo/lang/generic_type.hpp>

namespace mojo {
namespace lang {

template<typename K, typename V>
class MapType : public DataTypeOf<GenericType<>> {
public:

};

template <typename K, typename V> inline
const MapType<K,V>* buildType(const TypeIdentifier<MapType<K,V>>*) {
    static const MapType<K,V>* p = new MapType<K,V>{};
    return p;
}

}
}

#endif //MOJO_DICTIONARY_TYPE_HPP
