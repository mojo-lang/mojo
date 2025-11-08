package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const MarkdownText = "## 1. 像素点（Point）\n\n### 构造函数\n\n```javascript\nnew aimap.Point(x: number, y: number)\n```\n\n#### 参数说明\n| 字段名           | 类型          | 是否必须 | 默认值         | 说明                              |\n| ---------------- | ------------- | -------- | -------------- | --------------------------------- |\n| x               | number        | 是       |  | 像素横坐标值      |\n| y              | number           | 是       |                |   像素纵坐标值                                |\n\n\n#### 示例代码\n\n```javascript\nconst p = new aimap.Point(116, 39);\n```\n\n### 成员函数\n#### `equals(point: Point)`\n##### 参数\n`point (Point)` \n##### 返回值\n`boolean`\n##### 示例\n```javascript\nconst p = new aimap.Point(116, 39);\nconst p1 = new aimap.Point(116, 39);\nconst b = p.equals(p1);\n```\n#### `toArray()`\n##### 返回值\n`Array<number>`\n##### 示例\n```javascript\nconst p = new aimap.Point(12, 100);\np.toArray(); // = [12, 100]\n```\n#### `toString()`\n##### 返回值\n`string`\n##### 示例\n```javascript\nconst p = new aimap.Point(12, 100);\np.toString(); // = \"(12, 100)\"\n```"

func TestParser_ParseBytes(t *testing.T) {
	parser := NewParser()
	source := []byte(MarkdownText)
	parser.ParseBytes(source)

	astParser := &AstParser{}
	astParser.Parse(source, parser.Node)

	// md := New()
	// file, _ := ioutil.TempFile(".", "*.pdf")
	// md.ConvertToPdf(source, file)
	//
	// file, _ = ioutil.TempFile(".", "*.html")
	// md.ConvertToHtml(source, file)
}

func TestMarkdown_Render(t *testing.T) {
	md := New()
	d, err := md.Parse(MarkdownText)
	assert.NoError(t, err)

	str, _ := md.RenderToString(d)
	assert.True(t, len(str) > 0)
}

const Unordered = "* Item *1*\n  descriptions\n* Item 2\n  * Item 2a\n  * Item 2b"

func TestAstParser_ParseUnordered(t *testing.T) {
	parser := NewParser()
	source := []byte(Unordered)
	parser.ParseBytes(source)

	astParser := &AstParser{}
	astParser.Parse(source, parser.Node)
}

const OrderedList = "1. 在页面引入`aimap`的*sdk*;\n\n   ```html\n   <link rel='stylesheet' href='https://location-dev.newayz.com/aimap-gl-js/v1.3.10/aimap-gl.css' />\n   <script src='https://location-dev.newayz.com/aimap-gl-js/v1.3.10/aimap-gl.js'></script>\n   ```\n\n2. 添加div标签坐标地图容器，同时为该div指定id属性；\n\n   ```html\n   <div id='map'></div>\n   ```\n\n3. 指定div容器的样式，为地图容器指定高度、宽度；\n\n   ```css\n   body {\n   \tmargin: 0;\n   \tpadding: 0;\n   }\n   \n   html,\n   body,\n   #map {\n   \theight: 100%;\n   }\n   ```"

func TestAstParser_ParseOrderedList(t *testing.T) {
	parser := NewParser()
	source := []byte(OrderedList)
	parser.ParseBytes(source)

	astParser := &AstParser{}
	astParser.Parse(source, parser.Node)
}
