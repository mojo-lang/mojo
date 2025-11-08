#ifndef LANG_NOMINAL_TYPE_PTR_HPP
#define LANG_NOMINAL_TYPE_PTR_HPP

#include <memory>

namespace mojo {
namespace lang {

struct NominalType;

/**
 *
 */
using NominalTypePtr = std::shared_ptr<NominalType>;

}
}

#endif //LANG_NOMINAL_TYPE_PTR_HPP
