#ifndef LANG_DATA_DECL_PTR_HPP
#define LANG_DATA_DECL_PTR_HPP

#include <memory>

namespace mojo {
namespace lang {

struct DataDecl;
using DataDeclPtr = std::shared_ptr<DataDecl>;

}
}

#endif //LANG_DATA_DECL_PTR_HPP
