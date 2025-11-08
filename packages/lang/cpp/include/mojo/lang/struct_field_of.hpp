#ifndef MOJO_LANG_STRUCT_FIELD_OF_HPP
#define MOJO_LANG_STRUCT_FIELD_OF_HPP

#include <mojo/lang/struct_field_for.hpp>

namespace mojo {
namespace lang {

// TODO 增加通过get、set处理的虚拟成员变量，而不是实际存在的成员变量
template <typename T, typename M>
class StructFieldOf : public StructFieldFor<T> {
public:
    using MemberPtr = M T::*;
    MemberPtr ptr_;

    StructFieldOf(MemberPtr ptr, String name, int idx) : ptr_(ptr) {
        // StructType::Field::type = getType<M>();
    }

    virtual ~StructFieldOf() {
    }

    void visit(T& record, DataVisitor& visitor) const override {
        // visitor[this->column()](get_known(record));
    }

    virtual DataMutator mutator(T&) const {
        return DataMutator();
    }

    ncraft::Result<ncraft::Any> get(ncraft::AnyConstRef record) const {
        if (!record.is_a<T>()) {
            // return detail::make_type_error_for_mismatching_record_type(type<T>(),
            //                                                           record.type_info());
        }
        auto& mref = *record.get<const T&>();
        return ncraft::Any{get_known(mref)};
    }

    ncraft::Result<void> set(ncraft::AnyRef record, ncraft::AnyConstRef value) const {
        if (!record.is_a<T>()) {
            // return detail::make_type_error_for_mismatching_record_type(type<T>(),
            //                                                           record.type_info());
        }
        if (!value.is_a<M>()) {
            // return detail::make_type_error_for_mismatching_value_type(type<T>(),
            //                                                          type<M>(),
            //                                                          value.type_info());
        }
        auto& mref = *record.get<T&>();
        auto vref = value.get<M>();
        get_known(mref) = *vref;
        return ncraft::Nothing;
    }

    const M& get_known(const T& record) const {
        return record.*ptr_;
    }

    M& get_known(T& record) const {
        return record.*ptr_;
    }

    bool hasValue(const T& record) const final {
        return false;
        // return getType<M>()->hasValue(get_known(record));
    }
};

}  // namespace lang
}  // namespace mojo

#endif  // MOJO_LANG_STRUCT_FIELD_OF_HPP
