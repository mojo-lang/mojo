

csv('')
csv_file('')

json('')
json_file('')

loads<json>('')
load<json>('')

var positions = {}
read_line('xxx') | json | !$.calctype | {tid, } | positions.append

read_line('xxx') | json | $.calctype | {
	date
	devid
	tid
	lat: $.position.lat
	lon: $.position.lon
	source: postions[$.tid] | $.status == 1 | source
} | write_csv("")