#ifndef MOJO_LANG_STRUCT_ADAPTOR_HPP
#define MOJO_LANG_STRUCT_ADAPTOR_HPP

#include <ncraft/data/data_adaptor_of.hpp>

namespace mojo {
namespace lang {

template<typename T>
class StructAdaptor : public ncraft::data::DataAdaptorOf<T> {
public:

};

}
}

#endif //MOJO_LANG_STRUCT_ADAPTOR_HPP
