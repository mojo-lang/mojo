#ifndef LANG_FUNC_TYPE_PTR_HPP
#define LANG_FUNC_TYPE_PTR_HPP

#include <memory>

namespace mojo {
namespace lang {

struct FuncType;
using FuncTypePtr = std::shared_ptr<FuncType>;

}
}

#endif //LANG_FUNC_TYPE_PTR_HPP
