#ifndef MOJO_DATA_TYPE_OF_HPP
#define MOJO_DATA_TYPE_OF_HPP

#include <ncraft/core/type_info.hpp>
#include <mojo/lang/data_type.hpp>

namespace mojo {
namespace lang {

template <class T, class Base = DataType>
class DataTypeOf : public Base {
public:
    virtual ~DataTypeOf() = default;

    const ncraft::TypeInfo& typeInfo() const noexcept final {
        return ncraft::core::GetTypeInfo<T>::Value;
    }

    void visit(ncraft::AnyConstRef data, DataVisitor& visitor) const noexcept final {
        if (&data.typeInfo() != &this->typeInfo()) {
            //std::stringstream stream;
            //stream << "visitData called with data of wrong type (got "
            //       << data.typeInfo().get_name() << ", expected " << this->name;
            //throw TypeError(stream.str());
        }
        // XXX: Erm.. Maybe we should find a better way to do this.
        this->visit(const_cast<T&>(*data.get<const T&>()), visitor);
    }

    virtual void visit(T& value, DataVisitor& visitor) const noexcept = 0;

    virtual bool hasValue(const T& value) const noexcept = 0;
};

}
}

#endif //MOJO_DATA_TYPE_OF_HPP
