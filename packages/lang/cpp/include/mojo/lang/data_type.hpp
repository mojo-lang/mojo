#ifndef MOJO_LANG_DATA_TYPE_HPP
#define MOJO_LANG_DATA_TYPE_HPP

#include <ncraft/data/data_type.hpp>
#include <mojo/core/map.hpp>

namespace mojo {
namespace lang {

class Package;

//using Attributes = Map<String, Attribute>;

class DataType : public ::ncraft::data::DataType {
public:

};

}  // namespace lang

using DataType = lang::DataType;

}  // namespace mojo

#endif  // MOJO_LANG_DATA_TYPE_HPP
