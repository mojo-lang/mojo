/// [Vector Tile Specification v2.1](https://github.com/mapbox/vector-tile-spec/tree/master/2.1)
type VectorTile {
    /// GeomType is described in section 4.3.4 of the specification
    enum GeomType {
        unknown @0 @alias("UNKNOWN")
        point @1 @alias("POINT")
        linestring @2 @alias("LINESTRING")
        polygon @3 @alias("POLYGON")
    }

    // Variant type encoding
    // The use of values is described in section 4.1 of the specification
    type Value {
        // Exactly one of these values must be present in a valid message
        string_value: String @1
        float_value: Float @2
        double_value: Double @3
        int_value: Int64 @4
        uint_value: UInt64 @5
        sint_value: Int64 @6 @protobuf.signed
        bool_value: Bool @7
    }

    // Features are described in section 4.2 of the specification
    type Feature {
        id: UInt64 @1 = 0

        // Tags of this feature are encoded as repeated pairs of
        // integers.
        // A detailed description of tags is located in sections
        // 4.2 and 4.4 of the specification
        tags: [UInt32] @2 @protobuf.packed

        // The type of geometry stored in this feature.
        type: GeomType @3 = GeomType.unknown

        // Contains a stream of commands and parameters (vertices).
        // A detailed description on geometry encoding is located in
        // section 4.3 of the specification.
        geometry: [UInt32] @4 @protobuf.packed
    }

    // Layers are described in section 4.1 of the specification
    type Layer {
        // Any compliant implementation must first read the version
        // number encoded in this message and choose the correct
        // implementation for this version number before proceeding to
        // decode other parts of this message.
        version: UInt32 @15 @required = 1

        name: String @1 @required

        // The actual features in this tile.
        features: [Feature] @2

        // Dictionary encoding for keys
        keys: [String] @3

        // Dictionary encoding for values
        values: [Value] @4

        // Although this is an "optional" field it is required by the specification.
        // See https://github.com/mapbox/vector-tile-spec/issues/47
        extent: UInt32 @5 = 4096
    }

    layers: [Layer] @3
}