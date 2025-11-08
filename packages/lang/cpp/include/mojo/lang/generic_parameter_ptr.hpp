#ifndef LANG_GENERIC_PARAMETER_PTR_HPP
#define LANG_GENERIC_PARAMETER_PTR_HPP

#include <memory>

namespace mojo {
namespace lang {

struct GenericParameter;
using GenericParameterPtr = std::shared_ptr<GenericParameter>;

}  // namespace lang
}  // namespace mojo

#endif  // LANG_GENERIC_PARAMETER_PTR_HPP
