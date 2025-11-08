#ifndef MOJO_LANG_STRUCT_TYPE_JSON_FORMATTER_HPP
#define MOJO_LANG_STRUCT_TYPE_JSON_FORMATTER_HPP
/*
 *
 */

#include <ncraft/format/formatter.hpp>
#include <mojo/lang/struct_type.hpp>
#include <ncraft/json/json_writer.hpp>

namespace ncraft {
namespace format {

template <typename Tag, typename T>
class Formatter<json::Json<Tag>, T, meta::EnableIf<mojo::lang::IsStructType<T>>> {
public:
    Formatter(const T& value) : value_(value) {
    }

    template <class Callback>
    void format()(Argument<json::Json<Tag>>& arg, Callback& cb) const {
        using Appender = Appender<FormatCallback<json::Json<Tag>, Callback>>;
        auto writer = json::JsonWriter<Appender::Buffer>{Appender{cb}};

        auto type = getDataType<T>();

        writer.startObject();
        writer.endObject();
    }

private:
    const T& value_;
};

}  // namespace format
}  // namespace ncraft

#endif  // MOJO_LANG_STRUCT_TYPE_JSON_FORMATTER_HPP
