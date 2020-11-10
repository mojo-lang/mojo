Mojo主要定位于处理结构化数据处理的应用场景。


类似于PowerShell在shell方面的处理

```
msh> ps({}) | name == 'process_name' | {name:name, usages: count(usage)} | print
```

string -> pbf ast -> engine


类似于SQL在数据库操作方面的处理

```
roads |  grid_clip(3..15) |{
	grid: $
	point_count: $ | point_size | sum()
	geom: $.geom
} | attch_style({
	fill.color: scale(point_count, (100, #ffffff),
					              (300: #ffffdd),
						          (1000: #ffffaa),
						          (2000: #ffff00),
							      (_ #000000))

	symbol.text.name: point_count
}
```

mojo库代码文件命名及存储结构风格

1. module，一个mojo源代码文件即是一个模块，任何模块都会在一个包里
2. package 包，是指多个模块的集合，一般为目录；一个包允许存在多个目录
