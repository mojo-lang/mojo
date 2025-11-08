#ifndef MOJO_LANG_DATA_VISITOR_HANDLER_HPP
#define MOJO_LANG_DATA_VISITOR_HANDLER_HPP

#include <mojo/lang/data_visitor.hpp>

namespace mojo {
namespace lang {

/**
 * concept Handler {
 *      void objectStart();
 *      void objectEnd();
 *      void arrayStart();
 *      void arrayEnd();
 *      void field(const StringView& key);
 *      void value(const T&);
 *      void error(Error&& error);
 * };
 */
template <typename Handler>
struct DataVisitHandlerAdaptor : public DataVisitor {

    void visitNull() final {
        handler_.value(Nonthing);
    }

    void visitBool(bool& value) noexcept final {
        handler_.value(value);
    }

    void visitInt8(std::int8_t& value) noexcept final {
        handler_.value(value);
    }

    void visitInt16(std::int16_t& value) noexcept final {
        handler_.value(value);
    }

    void visitInt32(std::int32_t& value) noexcept final {
        handler_.value(value);
    }

    void visitInt64(std::int64_t& value);

    void visitUInt8(std::uint8_t& value);

    void visitUInt16(std::uint16_t& value);

    void visitUInt32(std::uint32_t& value);

    void visitUInt64(std::uint64_t& value);

    void visitFloat(float& value);

    void visitDouble(double& value);

    void visitString(std::string& value);

    void visitValue(AnyRef data, const DataType* type);

    void visitKey(const std::string& key) {
        if (!inMap) {
            handler_.objectStart();
        }

        handler_.field(key);
    }

    void visitIndex(std::int64_t idx) {
    }

    void visitField(const StructType::Field* field) {
    }

    bool modifiable() const {
        return false;
    }

    bool nullCurrently() const {
        return false;
    }

private:
    enum class State {
        kNone,
        kMap,
        kMapKey,
        kMapValue,
        kArray,
        kStructType
    };

    State state_ = State.kNone;
    Handler& handler_;
};

}  // namespace lang
}  // namespace mojo

#endif  // MOJO_LANG_DATA_VISITOR_HANDLER_HPP