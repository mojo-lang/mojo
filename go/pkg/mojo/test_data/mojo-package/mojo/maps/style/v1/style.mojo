///
interface Style {
    @http.post('')
    make() -> Map

    @http.put('')
    update() -> Map

    @http.delete('')
    remove() -> Null
}