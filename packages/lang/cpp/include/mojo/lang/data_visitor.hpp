#ifndef MOJO_LANG_DATA_VISITOR_HPP
#define MOJO_LANG_DATA_VISITOR_HPP

#include <mojo/core/any.hpp>
#include <mojo/core/bool.hpp>
#include <mojo/core/numeric.hpp>
#include <mojo/core/null.hpp>
#include <mojo/core/string.hpp>
#include <mojo/lang/data_type.hpp>

namespace mojo {
namespace lang {

class StructTypeField;

class DataVisitor {
public:
    enum class Action {
        kNone,
        kStop,
        kSkip
    };

    // Implement this interface:

    /**
     *
     */
    virtual void visitNull() noexcept = 0;

    /**
     *
     */
    virtual void visitBool(Bool&) noexcept = 0;

    /**
     *
     */
    virtual void visitInt8(Int8&) noexcept = 0;

    /**
     *
     */
    virtual void visitInt16(Int16&) noexcept = 0;

    /**
     *
     */
    virtual void visitInt32(Int32&) noexcept = 0;

    /**
     *
     */
    virtual void visitInt64(Int64&) noexcept = 0;

    /**
     *
     */
    virtual void visitUInt8(UInt8&) noexcept = 0;

    /**
     *
     */
    virtual void visitUInt16(UInt16&) noexcept = 0;

    /**
     *
     */
    virtual void visitUInt32(UInt32&) noexcept = 0;

    /**
     *
     */
    virtual void visitUInt64(UInt64&) noexcept = 0;

    /**
     *
     */
    virtual void visitFloat(Float&) noexcept = 0;

    /**
     *
     */
    virtual void visitDouble(Double&) noexcept = 0;

    /**
     *
     */
    virtual void visitString(String&) noexcept = 0;

    /**
     *
     * @param data
     * @param type
     */
    virtual void visitData(ncraft::AnyRef data, const DataType* type) noexcept = 0;

    /**
     *
     * @param key
     */
    virtual void key(const String& name) = 0;

    /**
     *
     * @param idx
     */
    virtual void index(Int64 idx) = 0;

    /**
     *
     * @param field
     */
    virtual void field(const StructTypeField* fld) = 0;

    /**
     *
     * @param type
     */
    virtual void begin(const DataType* type) noexcept = 0;

    /**
     *
     */
    virtual void end() noexcept = 0;

    /**
     *
     * @return
     */
    virtual bool modifiable() const = 0;

    /**
     *
     * @return
     */
    virtual bool nullCurrently() const = 0;

    // Specialize this if your data type can be visited as one of the normal methods
    // listed above.
    // If it isn't specialized, visit_special will be called.
    template <class T>
    struct VisitCaller;

    // Call this interface from types:

    template <class T>
    void visit(T& value) {
        using Type = ncraft::traits::RemoveCVRef<T>;
        VisitCaller<Type>::visit(*this, value);
    }

    template <class T>
    void operator()(T& value) {
        visit(value);
    }

    template <class T>
    void visit(const String& key, T& value) {
        using Type = ncraft::traits::RemoveCVRef<T>;
        this->key(key);
        this->visitValue(ncraft::AnyRef{value}, getType<Type>());
    }

    struct KeyValueVisitor {
        DataVisitor& visitor;
        String key;

        KeyValueVisitor(DataVisitor& visitor, String key)
            : visitor(visitor), key(::std::move(key)) {
        }

        template <class T>
        void operator()(T& value) {
            visitor.visit(key, value);
        }
    };

    KeyValueVisitor operator[](const String& key) {
        return KeyValueVisitor{*this, key};
    }

    KeyValueVisitor operator[](const char* key) {
        return KeyValueVisitor{*this, key};
    }

    template <class T>
    void visit(Int64 idx, T& value) {
        using Type = ncraft::traits::RemoveCVRef<T>;
        this->index(idx);
        this->visitValue(idx, AnyRef{value}, getType<Type>());
    }

    struct ElementVisitor {
        DataVisitor& visitor;
        Int64 index;

        ElementVisitor(DataVisitor& visitor, Int64 index)
            : visitor(visitor), index(index) {
        }

        template <class T>
        void operator()(T& value) {
            visitor.visit(index, value);
        }
    };

    ElementVisitor operator[](Int64 index) {
        return ElementVisitor{*this, index};
    }
};

template <>
struct DataVisitor::VisitCaller<Null> {
    static void visit(DataVisitor& visitor, Null) {
        visitor.visitNull();
    }
};

template <>
struct DataVisitor::VisitCaller<Bool> {
    static void visit(DataVisitor& visitor, Bool& value) {
        visitor.visitBool(value);
    }
};

template <>
struct DataVisitor::VisitCaller<Int8> {
    static void visit(DataVisitor& visitor, Int8& value) {
        visitor.visitInt8(value);
    }
};

template <>
struct DataVisitor::VisitCaller<Int16> {
    static void visit(DataVisitor& visitor, Int16& value) {
        visitor.visitInt16(value);
    }
};

template <>
struct DataVisitor::VisitCaller<Int32> {
    static void visit(DataVisitor& visitor, Int32& value) {
        visitor.visitInt32(value);
    }
};

template <>
struct DataVisitor::VisitCaller<Int64> {
    static void visit(DataVisitor& visitor, Int64& value) {
        visitor.visitInt64(value);
    }
};

template <>
struct DataVisitor::VisitCaller<UInt8> {
    static void visit(DataVisitor& visitor, UInt8& value) {
        visitor.visitUInt8(value);
    }
};

template <>
struct DataVisitor::VisitCaller<UInt16> {
    static void visit(DataVisitor& visitor, UInt16& value) {
        visitor.visitUInt16(value);
    }
};

template <>
struct DataVisitor::VisitCaller<UInt32> {
    static void visit(DataVisitor& visitor, UInt32& value) {
        visitor.visitUInt32(value);
    }
};

template <>
struct DataVisitor::VisitCaller<UInt64> {
    static void visit(DataVisitor& visitor, UInt64& value) {
        visitor.visitUInt64(value);
    }
};

template <>
struct DataVisitor::VisitCaller<Float> {
    static void visit(DataVisitor& visitor, Float& value) {
        visitor.visitFloat(value);
    }
};

template <>
struct DataVisitor::VisitCaller<Double> {
    static void visit(DataVisitor& visitor, Double& value) {
        visitor.visitDouble(value);
    }
};

template <>
struct DataVisitor::VisitCaller<String> {
    static void visit(DataVisitor& visitor, String& value) {
        visitor.visitString(value);
    }
};

template <class T>
struct DataVisitor::VisitCaller {
    using Type = ncraft::traits::RemoveCVRef<T>;

    static void visit(DataVisitor& visitor, T& value) {
        visitor.visit(value, getType<T>());
    }
};
}
}

#endif  // MOJO_LANG_DATA_VISITOR_HPP
