#ifndef MOJO_LANG_DATA_ADAPTOR_PTR_HPP
#define MOJO_LANG_DATA_ADAPTOR_PTR_HPP

#include <memory>

namespace mojo {
namespace lang {

class DataAdaptor;
using DataAdaptorPtr = std::shared_ptr<DataAdaptor>;

}
}

#endif //MOJO_LANG_DATA_ADAPTOR_PTR_HPP
