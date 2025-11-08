#ifndef MOJO_DATA_TYPE_OF_STRING_HPP
#define MOJO_DATA_TYPE_OF_STRING_HPP

#include <mojo/core/string.hpp>
#include <mojo/lang/data_type_of.hpp>
#include <mojo/lang/data_visitor.hpp>
#include <mojo/lang/string_adaptor.hpp>

namespace mojo {
namespace lang {

class StringType : public DataTypeOf<String> {
public:
    StringType() {
        DataType::name = "String";
    }

    virtual ~StringType() = default;

    bool nullable() const noexcept final {
        return false;
    }

    bool hasValue(const String& value) const noexcept final {
        return true;
    }

    void visit(String& value, DataVisitor& visitor) const noexcept final {
        visitor(value);
    }

    const DataAdaptor* adaptor() const noexcept final {
        static const DataAdaptor* dataAdaptor = new StringAdaptor;
        return dataAdaptor;
    }
};

const StringType* buildType(const TypeIdentifier<String>*);
}
}

#endif  // MOJO_DATA_TYPE_OF_STRING_HPP
