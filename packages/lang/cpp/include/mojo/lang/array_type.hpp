#ifndef MOJO_ARRAY_TYPE_HPP
#define MOJO_ARRAY_TYPE_HPP

#include <mojo/lang/data_type_of.hpp>
#include <mojo/lang/generic_type.hpp>

namespace mojo {
namespace lang {

template<typename T>
class ArrayType : public DataTypeOf<GenericType<>> {
public:

};

template <typename T> inline
const ArrayType<T>* buildType(const TypeIdentifier<ArrayType<T>>*) {
    static const ArrayType<T>* p = new ArrayType<T>{};
    return p;
}

}
}

#endif //MOJO_ARRAY_TYPE_HPP
