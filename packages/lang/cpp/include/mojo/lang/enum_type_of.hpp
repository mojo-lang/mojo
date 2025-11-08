#ifndef MOJO_LANG_ENUM_TYPE_OF_HPP
#define MOJO_LANG_ENUM_TYPE_OF_HPP

#include <mojo/lang/enum_type.hpp>
#include <ncraft/data/data_type_of.hpp>

namespace mojo {
namespace lang {

template <typename T>
class EnumTypeOf : public ncraft::data::DataTypeOf<T, EnumType> {
public:
};

}  // namespace lang
}  // namespace mojo

#endif  // MOJO_LANG_ENUM_TYPE_OF_HPP
