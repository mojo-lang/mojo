#ifndef NCRAFT_SDATA_OBJECT_HPP
#define NCRAFT_SDATA_OBJECT_HPP

#include <map>
#include <ncraft/core/cloning_ptr.hpp>
#include <ncraft/core/either.hpp>
#include <ncraft/core/error.hpp>
#include <ncraft/core/maybe.hpp>
#include <ncraft/format/formatters.hpp>
#include <mojo/lang/data_reader_of.hpp>
#include <mojo/lang/data_writer_of.hpp>

namespace mojo {
namespace lang {

/**
 * An "object" is a piece of structured data that can be modified.
 *
 * Think JSON builder.
 */
struct Object final : DataReaderOf<Object, const Object&>, DataWriterOf<Object, Object&> {
    Object(Null = kNull) : data_(kNull) {
    }

    Object(Boolean b) : data_(b) {
    }

    Object(Integer n) : data_(n) {
    }

    Object(Real r) : data_(r) {
    }

    Object(String s) : data_(std::move(s)) {
    }

    Object(const Object&) = default;

    Object(Object&&) = default;

    Object& operator=(Object&&) = default;

    Object& operator=(const Object&) = default;

    // Convenience:
    Object(int n) : data_((Integer)n) {
    }

    Object(const char* str) : data_(std::string{str}) {
    }

    static Object map() {
        Object o;
        o.data_ = Map{};
        return std::move(o);
    }

    static Object list() {
        Object o;
        o.data_ = List{};
        return std::move(o);
    }

    Object* clone() const {
        return new Object(*this);
    }

    SDataType type() const;

    Object& operator[](size_t idx) {
        return this->mutableSubscript(idx);
    }

    const Object& operator[](size_t idx) const {
        return this->subscript(idx);
    }

    Object& operator[](const String& key) {
        return this->mutableSubscript(key);
    }

    const Object& operator[](const String& key) const {
        return this->subscript(key);
    }

    Object& operator[](const char* key) {
        return this->mutableSubscript(key);
    }

    const Object& operator[](const char* key) const {
        return this->subscript(key);
    }

    // This sets the type to 'List':
    void reserve(size_t idx);

    // Reader required interface:
    const Object& dataReader() const {
        return *this;
    }

    Maybe<Boolean> getBoolean() const {
        return data_.get<Boolean>();
    }

    Maybe<Integer> getInteger() const {
        return data_.get<Integer>();
    }

    Maybe<Real> getReal() const {
        return data_.get<Real>();
    }

    Maybe<String> getString() const {
        return data_.get<String>();
    }

    bool has(const String& key) const;

    const Object& get(const String& key) const;

    size_t length() const;

    const Object& at(size_t idx) const;

    struct ListEnumerator;
    struct DictEnumerator;

    ReaderEnumeratorPtr enumerator() const;

    // Writer required interface:
    Object& writerInterface() {
        return *this;
    }

    bool setNothing() {
        data_ = Nothing;
        return true;
    }

    bool setBoolean(Boolean b) {
        data_ = b;
        return true;
    }

    bool setInteger(Integer n) {
        data_ = n;
        return true;
    }

    bool setReal(Real r) {
        data_ = r;
        return true;
    }

    bool setString(String str) {
        data_ = std::move(str);
        return true;
    }

    Object& mutableAtIndex(size_t idx);

    bool append(Object other);

    Object& mutableAtKey(const String& key);

    bool erase(const String& key);

private:
    using List = std::vector<CloningPtr<Object>>;
    using Map = std::map<std::string, CloningPtr<Object>>;

    Either<NothingType, Boolean, Integer, Real, String, List, Map> data_;

    static const Object g_null_object;

    friend struct Enumerator;
    friend struct Adapter<Object>;
};

struct ObjectIndexOutOfBounds : Error {
    ObjectIndexOutOfBounds(const std::string& msg)
        : Error{std::errc::result_out_of_range, msg} {
    }
};

template <>
struct AdapterOf<Object> : Adapter {
    Adapter(Object& ref, Bitflags<Options> options) : ref_(ref), options_(options) {
    }

    // IReader interface:
    SDataType type() const final {
        return ref_.type();
    }

    Maybe<Boolean> getBoolean() const final {
        return ref_.getBoolean();
    }

    Maybe<Integer> getInteger() const final {
        return ref_.getInteger();
    }

    Maybe<Real> getReal() const final {
        return ref_.getReal();
    }

    Maybe<String> getString() const final {
        return ref_.getString();
    }

    bool has(const String& key) const final {
        return ref_.has(key);
    }

    ReaderPtr get(const String& key) const final {
        return makeReader(ref_.get(key), options_);
    }

    size_t length() const final {
        return ref_.length();
    }

    ReaderPtr at(size_t idx) const final {
        return makeReader(ref_.at(idx), options_);
    }

    ReaderEnumeratorPtr enumerator() const final {
        return ref_.enumerator();
    }

    // IWriter interface:
    bool setNothing() final {
        return ref_.setNothing();
    }

    bool setBoolean(Boolean b) final {
        return ref_.setBoolean(b);
    }

    bool setInteger(Integer n) final {
        return ref_.setInteger(n);
    }

    bool setReal(Real r) final {
        return ref_.setReal(r);
    }

    bool setString(String str) final {
        return ref_.setString(std::move(str));
    }

    AdapterPtr mutableAtIndex(size_t idx) final {
        return makeAdapter(ref_.mutableAtIndex(idx), options_);
    }

    AdapterPtr append() final {
        ref_.append(Object{});
        AdapterPtr ptr;
        ref_.data_.template when<Object::List>(
            [&](Object::List& list) { ptr = makeAdapter(*list.back(), options_); });
        return std::move(ptr);
    }

    AdapterPtr mutableAtKey(const String& key) final {
        return makeAdapter(ref_.mutableAtKey(key), options_);
    }

    bool erase(const String& key) final {
        return ref_.erase(key);
    }

    Object& ref_;
    Bitflags<Options> options_;
};

}  // namespace lang
}  // namespace mojo

namespace ncraft {
namespace format {

template <>
class Formatter<Printable, mojo::lang::Object> {
public:
    explicit Formatter(const sdata::Object& val) : val_(val) {
    }

    template <class Callback>
    void operator()(Argument<>& arg, Callback& cb) const {
        switch (val_.type()) {
            case sdata::SDataType::Nothing:
                format(nullptr, arg, cb);
                break;
            case sdata::SDataType::Boolean:
                format(val_.read<sdata::Boolean>(), arg, cb);
                break;
            case sdata::SDataType::Integer:
                format(val_.read<sdata::Integer>(), arg, cb);
                break;
            case sdata::SDataType::Real:
                format(val_.read<sdata::Real>(), arg, cb);
                break;
            case sdata::SDataType::String:
                format(val_.read<sdata::String>(), arg, cb);
                break;
            case sdata::SDataType::List:
                // Formatter(val_.at(arg.splitIntKey())).format(arg, cb);
                break;
            case sdata::SDataType::Map:
                // Formatter(val_.at(arg.splitKey().toFbstring())).format(arg, cb);
                break;
        }
    }

private:
    const mojo::lang::Object& val_;
};

}  // namespace format
}  // namespace ncraft

#endif  // NCRAFT_SDATA_OBJECT_HPP
