service StakeMarkService {
    get_stake_mark(location: LngLat) -> StakeMark {
        edges | min_dist(location, 100) | ( project_point ) -> {
                stake_mark_edge | id == project_point.id | stake_marks[project_point.index] | {
                    road_number: road_number
                    road_name: road_name
                    location: project_point.point
                    value: {
                        km: value.km
                        meter: value.meter + project_point.distance
                    }
                }}
    }
}
