#ifndef MOJO_LANG_DATA_TYPES_HPP
#define MOJO_LANG_DATA_TYPES_HPP

#include <mojo/lang/data_type.hpp>
#include <mojo/lang/data_visitor.hpp>
#include <ncraft/core/maybe.hpp>

#include <mojo/core/string.hpp>

namespace mojo {
namespace lang {

class NothingTypeType : public DataType {
public:
    bool isNullable() const final {
        return true;
    }

    const std::string& name() const final {
        return name_;
    }

    const TypeInfo& typeInfo() const {
        return ::ncraft::core::GetTypeInfo<NothingType>::Value;
    }

    void visitData(AnyRef data, DataVisitor& visitor) const {
        visitor.visit_nil();
    }

    void visitData(AnyConstRef data, DataVisitor& visitor) const {
        visitor.visit_nil();
    }

private:
    const std::string name_{"std.Null"};
};

const NothingTypeType* buildType(const TypeIdentifier<NothingType>*);

namespace detail {
std::string maybe_type_name(const DataType* inner);
}

template <typename T>
class MaybeType : public DataTypeFor<::ncraft::Maybe<T>> {
public:
    MaybeType() {
    }

    const std::string& name() const final {
        return detail::maybe_type_name(getType<T>());
    }

    bool isNullable() const {
        return true;
    }

    bool hasValue(const Maybe<T>& value) const final {
        return static_cast<bool>(value);
    }

    void visit(Maybe<T>& value, DataVisitor& visitor) const final {
        if (visitor.can_modify()) {
            if (visitor.is_nil_at_current()) {
                value = Nothing;
            }
            else {
                value = T{};
                visitor(*value);
            }
        }
        else {
            if (value) {
                getType<T>()->visitData(*value, visitor);
            }
            else {
                visitor.visit_nil();
            }
        }
    }
};

template <typename T>
const MaybeType<T>* buildType(const TypeIdentifier<Maybe<T>>*) {
    static const MaybeType<T>* p = new MaybeType<T>{};
    return p;
}
}
}

#endif //MOJO_LANG_DATA_TYPES_HPP
