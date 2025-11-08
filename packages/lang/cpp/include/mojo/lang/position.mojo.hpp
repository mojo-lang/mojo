#ifndef LANG_POSITION_MOJO_HPP
#define LANG_POSITION_MOJO_HPP

#include <mojo/lang/struct_type_builder.hpp>
#include <mojo/lang/type_decl.mojo.hpp>

namespace mojo {
namespace lang {

struct Position {
    //String filename;
    Int offset = 0;
    Int line = 0;
    Int column = 0;
};

}
}

MOJO_BUILD_STRUCT_TYPE(mojo, lang, Position) {
    //field(&mojo::lang::Position::filename, "filename", 1);
    field(&mojo::lang::Position::offset, "offset", 2);
    field(&mojo::lang::Position::line, "line", 3);
    field(&mojo::lang::Position::column, "column", 4);
}

#endif //LANG_POSITION_MOJO_HPP
