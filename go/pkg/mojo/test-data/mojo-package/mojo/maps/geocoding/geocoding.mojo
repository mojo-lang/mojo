/// Geocoding API
///
@http.base_path('/geocoding/v1{/user}')
interface Geocoding {
    ///
    ///
    @http.get('/address/{address}{.format}')
    geocode(address: String @1) -> AddressSet

    ///
    ///
    @http.get('/locations/{location}')
    @error_codes([{
        value: 405
        brief: 'Invalid ID supplied'
        category: 'geocoding'
    }, {
        value: 404
        brief: 'Not found'
        category: 'geocoding'
    }])
    reverseGeocode (location:LngLat@5,
        nearby:Bool @10,
        nearby_radius:Int @11
    ) -> Address @2

    ///
    /// @http.query('', '')
    /// @http.header('', '') // header name, header value format
    @http.get('/addresses/{addresses}...[.{format}]')
    batchGeocode(addresses:String@1) -> AddressSet

    ///
    ///
    @http.get('/addresses/{lng,lat}[.{format}]')
    batch_reverse_geocode (
        locations: [LngLat] @1,
        include_poi: Bool @2
    ) -> [Address]
}