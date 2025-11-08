#ifndef MOJO_STRUCT_FIELD_VIRTUAL_OF_HPP
#define MOJO_STRUCT_FIELD_VIRTUAL_OF_HPP

#include <mojo/lang/struct_field_for.hpp>

namespace mojo {
namespace lang {

template <typename T, typename M>
class StructVirtualField : public StructFieldFor<T> {
public:
    using MemberGetPtr = M (T::*)();
    using MemberSetPtr = void (T::*)(const M&);

public:
    StructVirtualField(MemberGetPtr getter, MemberSetPtr setter, String&& name, int idx)
        : getter_(getter), setter_(setter) {
        StructType::Field::type = ::ncraft::data::dataType<M>();
    }

private:
    MemberGetPtr getter_;
    MemberSetPtr setter_;
};
}
}

#endif  // MOJO_STRUCT_FIELD_VIRTUAL_OF_HPP
