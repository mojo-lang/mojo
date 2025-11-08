
/// for compressed pured point tile.
type PointTile {
    /// the version of the point tile specification
    version: Int32  @15 = 1

    id: String @1 //< tile quad-key

    /// the point data type
    type: String  @2 //< the user defined point data type

    /// the properties for this tile
    properties: {String: VectorTile.Value} @3

    /// Dictionary encoding for keys
    keys: [String] @4

    /// Dictionary encoding for values
    values: [VectorTile.Value] @5

    /// attribute index for which value is raw integer
    raw_values: [Int32] @6 @protobuf.packed

    /// for string ids, using keys or md5(string id)
    ids: [UInt64] @7 @protobuf.packed

    /// Tags of this feature are encoded as repeated pairs of
    /// integers.
    tags: [UInt32] @10 @protobuf.packed

    /// the coordinates for the point, using E7 of the latitude & longitude
    xs: [Int32] @11 @protobuf.packed
    ys: [Int32] @12 @protobuf.packed
}

type TileSet {

}

type PointTileSet: TileSet {


}
