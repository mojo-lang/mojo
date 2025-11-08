#ifndef MOJO_LANG_OBJECT_VISIT_HANDLER_HPP
#define MOJO_LANG_OBJECT_VISIT_HANDLER_HPP

#include <mojo/lang/object.hpp>

namespace mojo {
namespace lang {

class ObjectVisitHandler {
public:
    ObjectVisitHandler() {}


    Object object() && {
        return std::move(object_);
    }

private:
    Object object_;
};

}
}

#endif //MOJO_LANG_OBJECT_VISIT_HANDLER_HPP
