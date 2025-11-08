type Schema {
    type: Type @1
    format: String @2

    value: Schema @5 // array, map
    fields: {String: Schema}  @6 // struct
    enumerators: [String] @7 // enum

    required: Bool @10

    multiple_of: Double @14 // 浮点类型的精度位，比如 0.01，表示保留两位精度

    maximum: Double @15
    exclusive_maximum: Double @16

    minimum: Double @17
    exclusive_minimum: Double @18

    max_length: UInt @19

    default: Value @40
    example: Value @41
}
