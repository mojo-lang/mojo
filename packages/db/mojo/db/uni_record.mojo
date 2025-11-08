
type UniRecord {
    id: core.Id @1 
    
    data: core.Object @2 @db.json

    create_time: core.Timestamp @5
    update_time: core.Timestamp @6

    delete_time: DeleteTime @7

    table_name: String @10 @db.ignore
}