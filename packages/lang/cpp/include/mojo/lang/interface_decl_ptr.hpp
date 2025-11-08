#ifndef LANG_INTERFACE_DECL_PTR_HPP
#define LANG_INTERFACE_DECL_PTR_HPP

#include <memory>

namespace mojo {
namespace lang {

struct InterfaceDecl;
using InterfaceDeclPtr = std::shared_ptr<InterfaceDecl>;

}
}

#endif //LANG_INTERFACE_DECL_PTR_HPP
