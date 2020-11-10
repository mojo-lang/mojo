type Country {
    id: PlaceId @1
    customized_id: String @2
    geometry: geom.Geometry @3
    type: String @4
    name: String @5
    code: String @6
    aliases: [String] @7
    brief: String @8

    tags: [Tag] @14
    images: [Image] @19

    name_zh: String @20
    name_en: String @21
    name_py: String @22

    create_time: DateTime @1000
    update_time: DateTime @1001
}