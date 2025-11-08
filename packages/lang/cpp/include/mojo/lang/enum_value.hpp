#ifndef MOJO_LANG_ENUM_VALUE_HPP
#define MOJO_LANG_ENUM_VALUE_HPP

#include <mojo/lang/enum_type.hpp>

namespace mojo {
namespace lang {

template <typename T>
class EnumValue : public EnumType::Value {
public:
    virtual ~EnumValue() {
    }

    //virtual bool hasValue(const T& msg) const = 0;
    //virtual void visit(T&, ObjectVisitor& visitor) const = 0;
};

}
}

#endif //MOJO_LANG_ENUM_VALUE_HPP
