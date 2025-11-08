#ifndef MOJO_LANG_DATA_TYPE_OF_NUMERIC_HPP
#define MOJO_LANG_DATA_TYPE_OF_NUMERIC_HPP

#include <mojo/core/numeric.hpp>
#include <mojo/lang/data_type_of.hpp>

namespace mojo {
namespace lang {

template <typename T>
class NumericType : public DataTypeOf<T> {
public:
    NumericType(String&& name, String&& fullName) {
        DataType::name = ::std::move(name);
        //DataType::fullName = ::std::move(fullName);
    }

    virtual~NumericType() = default;

    bool nullable() const noexcept final {
        return false;
    }

    size_t bits() const {
        return sizeof(T) * 8;
    }

    bool is_signed() const {
        return ::std::is_signed<T>::value;
    }

    bool is_float() const {
        return ::std::is_floating_point<T>::value;
    }

    bool hasValue(const T& value) const noexcept final {
        return true;
    }

    void visit(T& value, DataVisitor& visitor) const noexcept final {
        //visitor(value);
    }
};

const NumericType<Int8>* buildType(const TypeIdentifier<Int8>*);
const NumericType<Int16>* buildType(const TypeIdentifier<Int16>*);
const NumericType<Int32>* buildType(const TypeIdentifier<Int32>*);
const NumericType<Int64>* buildType(const TypeIdentifier<Int64>*);
const NumericType<UInt8>* buildType(const TypeIdentifier<UInt8>*);
const NumericType<UInt16>* buildType(const TypeIdentifier<UInt16>*);
const NumericType<UInt32>* buildType(const TypeIdentifier<UInt32>*);
const NumericType<UInt64>* buildType(const TypeIdentifier<UInt64>*);
const NumericType<Float>* buildType(const TypeIdentifier<Float>*);
const NumericType<Double>* buildType(const TypeIdentifier<Double>*);

}
}

#endif //MOJO_LANG_DATA_TYPE_OF_NUMERIC_HPP
