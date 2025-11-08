#ifndef MOJO_LANG_STRUCT_TYPE_OF_HPP
#define MOJO_LANG_STRUCT_TYPE_OF_HPP

#include <mojo/lang/struct_type.hpp>
#include <ncraft/data/data_type_of.hpp>

namespace mojo {
namespace lang {

template <typename T>
struct StructTypeBuilder;

template <class T>
class StructTypeOf : public ncraft::data::DataTypeOf<T, StructType> {
public:
    bool nullable() const noexcept final {
    }

    size_t fieldCount() const final {
        return fields_.size();
    }

    const StructType::Field* field(const ncraft::String& name) const final {
        return nullptr;
    }
    const StructType::Field* field(size_t number) const final {
        return nullptr;
    }

    //const StructType::Field* field(size_t idx) const final {
    //    return fields_[idx].get();
    //}

    void visit(T& msg, ::ncraft::data::DataVisitor& visitor) const noexcept final {
        for (auto& field : fields_) {
            field->visit(msg, visitor);
        }
    }

    bool hasValue(const T& value) const noexcept final {
        return true;
    }

    // Status set(T& msg, Field& field, Value& value) {
    //    field.set(msg, value);
    //}

private:
    template<typename U> friend struct StructTypeBuilder;
    Array<std::unique_ptr<StructField<T>>> fields_;
};
/*
template <typename T>
class StructDataAdaptor : public DataAdaptor {
public:
    virtual ~StructDataAdaptor() = default;

    Mutator mutableAt(ncraft::AnyRef value, const std::string& name) const noexcept override {
        return mutateField(value, name);
    }

    Mutator mutateField(ncraft::AnyRef,
                        const std::string&,
                        const StructType::Field* = nullptr) const noexcept override {
        const Field* field = GetType<T>()->field(name);
    }

    void clear(AnyRef) const noexcept {
    }
};*/
}
}

#endif  // MOJO_LANG_STRUCT_TYPE_OF_HPP
