#ifndef MOJO_GENERIC_DATA_TYPE_HPP
#define MOJO_GENERIC_DATA_TYPE_HPP

#include <memory>
#include <mojo/core/array.hpp>
#include <mojo/core/string.hpp>
#include <mojo/lang/data_type.hpp>

namespace mojo {
namespace lang {

template<typename Base = DataType>
class GenericType : public Base {
public:
    /**
     *
     */
    struct Generic {
        String name;
        const DataType* type = nullptr;
        const DataType* constraint = nullptr;
    };

    Array<::std::unique_ptr<Generic>> generics;
};

}
}

#endif //MOJO_GENERIC_DATA_TYPE_HPP
