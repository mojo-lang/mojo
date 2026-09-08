// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/// Block element
@discriminator('@type')
@label_format('{}')
type Block = Plain          @1
           | Paragraph      @2
           | LineBlock      @3
           | CodeBlock      @4
           | QuoteBlock     @5
           | OrderedList    @6
           | BulletList     @7
           | DefinitionList @8
           | Header         @9
           | Table          @10
           | Division       @11
           //| Null           @12


type Line: [Inline]

/// Plain text, not a paragraph
type Plain {
    inlines: [Inline] @2
}

/// Paragraph
type Paragraph {
    inlines: [Inline] @2
}

/// Multiple non-breaking lines
type LineBlock {
    lines : [Line] @2
}

type CodeBlock {
    attribute: Attribute @1
    language: String @2

    lines: [Line] @3
}

type QuoteBlock {
    blocks: [Block] @2
}

/// List attributes.
type ListAttribute {
    /// Style of list numbers.
    @case_style("kebab")
    enum NumberStyle {
        unspecified @0
        example     @1
        decimal     @2
        lower_roman @3
        upper_roman @4
        lower_alpha @5
        upper_alpha @6
    }

    @case_style("kebab")
    enum NumberDelimiter {
        unspecified @0
        period      @1
        one_parent  @2
        two_parents @3
    }

    begin_number: Int @1
    number_style: NumberStyle @2
    number_delimiter: NumberDelimiter @3
}

/// 
type OrderedList {
    attribute: ListAttribute @1
    items: [[Block]] @2
}

/// Bullet list (list of items, each a list of blocks)
type BulletList {
    items: [[Block]] @2
}

/// Definition list Each list item is a pair consisting of a term (a list of inlines)
/// and one or more definitions (each a list of blocks)
type DefinitionList {
    ///
    type Item {
        term: [Inline] @1
        definitions: [[Block]] @2
    }

    items: [Item] @2
}

type Header {
    attribute: Attribute @1

    level : Int @2
    text : [Inline] @3
}

///
type Table {
    /// Alignment of a table column.
    enum Alignment {
        unspecified @0
        left   @1
        right  @2
        center @3
    }

    /// Table cells are list of Blocks
    type Cell: [Block]

    /// row, list of blocks
    type Row: [Cell]

    type Header: [Cell]

    caption: [Inline] @1 //< caption for table

    alignment: Alignment @2 //< column alignments (required)

    width: Double @3 = 0 //< relative column widths

    header: Header @4 //< table header, each is column header (each a list of blocks)

    rows: [Row] @5 //< rows, a list of row
}

/// Generic block container with attributes
type Division {
    attribute: Attribute @1
    
    content: [Block] @2
}

type Blocks = [Block]
