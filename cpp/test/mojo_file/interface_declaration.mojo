@template("{longitude},{latitude}")
type LngLat {
	longitude: Double @1
	latitude  : Double @2
}

type Highway {
	number: String @1 //< 道路国家标准编号
	start : String @2
	end   : String @3
}

type District {
	id: Int64 @1
	name: String @2
	full_name: String @3
}

type StakeMark {
	
	/// 桩号值对象。用以表示K23+345的桩号，分成`kilometer`部分和`meter`部分。
	@template('T{kilometer}[+.]{meter}')
	type Value {
	    kilometer: Int32 @1  //< kilometer part of the stake mark value. e.g. 23
	    meter    : Int32 @2  //< meter part of the stake mark value. e.g. 345
	}

	road_number: String @1 //< 道路国家标准编号
	road_name  : String @2 //< 道路标准名称

	district   : District@3 |
	             Int64   @4 @field(id) |
	             String  @5 @field(full_name)

	/// 道路桩号值，可以使字符串表示，例如`K23+345`或者Value对象表示`{kilometer: 23, meter: 345}`
	value      : Value@6 | String@7

	/// 桩号所在的经纬度
	location   : Union<LngLat@8, [Double]@9 @packed, String@10>

	accuracy   : Int32@11 //< 桩号的精度范围
}

/// 与Edge每个Shape Point绑定后的路桩号。
/// 当Edge为单向Edge时，其Shape Point的次序必须和道路方向一致；
/// 当Edge为双向Edge时，其Shape Point的次序必须和路庄数据一致，即`direction`为`forward`
type BoundStakeMark : StakeMark {
	enum Direction {
		forward   //< 道路方向与路桩的次序一致，如G2京沪高速，从北京至上海方向的edge为forward
		backward
	}

	type Id {
      	edge_id : Int64 //< 道路Edge的ID
      	index   : Int32 //< 当前Shape Point在该道路Edge中Index，从0开始
	}

	id    : Id        //< 唯一ID

	/// 当Edge为单向道路时，如果方向与桩号次序一致，则为`forward`，反之为`backward`
	/// 当Edge为双向道路时，其Shape Point的次序必须调整成和桩号次序保持一致，即其永远为`forward`
	direction  : Direction = forward
}

func to(id:BoundStakeMark.Id) -> Int64 {
	return (id.edge_id <<< 10) ||| id.index
}

type BoundStakeMarkEdge {
	start: Int64
	end: Int64
	geometry : Line 
}

/// interface for stake mark
interface StakeMarkService {
	
	/// comments
	@http_get('/stake_marks')
	func get_stake_mark(location: LngLat) -> StakeMark {
        edges | min_dist(location, 100) | project_point -> {
                stake_mark_edge | id == project_point.id | stake_marks[project_point.index] | {
                    road_number
                    road_name
                    location: project_point.point
                    value: {
                        km: value.km
                        meter: value.meter + project_point.distance
                    }
                }}
    }

	/// comments
	@http_get('/stakemarks/{road}/{stake_mark}')
	func locate_stake_mark(road:String, stake_mark:String, direction:RoadDirection=forward) -> StakeMark {
		var sm_value = StakeMark.Value(stake_mark)
		bound_stake_marks |
			$.road_number == road &&
				$.value.kilometer == sm_value.kilometer &&
				$.direction == direction |
			range |
			sm_value.meter in [$0.value.meter, $1.value.meter] |
			interpolation
	}
}
