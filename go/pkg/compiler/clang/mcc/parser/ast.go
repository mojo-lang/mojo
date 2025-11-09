package parser

type Location struct {
	File                string `json:"file,omitempty"`
	Offset              int    `json:"offset,omitempty"`
	Line                int    `json:"line,omitempty"`
	PresumedLine        int    `json:"presumedLine,omitempty"`
	Col                 int    `json:"col,omitempty"`
	TokLen              int    `json:"tokLen,omitempty"`
	IsMacroArgExpansion bool   `json:"isMacroArgExpansion,omitempty"`

	IncludedFrom *Location `json:"includedFrom,omitempty"`
	SpellingLoc  *Location `json:"spellingLoc,omitempty"`
	ExpansionLoc *Location `json:"expansionLoc,omitempty"`
}

type LocationRange struct {
	Begin *Location `json:"begin,omitempty"`
	End   *Location `json:"end,omitempty"`
}

type Type struct {
	DesugaredQualType string `json:"desugaredQualType,omitempty"`
	QualType          string `json:"qualType,omitempty"`
	TypeAliasDeclId   string `json:"typeAliasDeclId,omitempty"`
}

type Term struct {
	Id                   string         `json:"id,omitempty"`
	Name                 string         `json:"name,omitempty"`
	MangledName          string         `json:"mangledName,omitempty"`
	CloseName            string         `json:"closeName,omitempty"`
	Kind                 string         `json:"kind,omitempty"`
	RenderKind           string         `json:"renderKind,omitempty"`
	Args                 []string       `json:"args,omitempty"`
	StorageClass         string         `json:"storageClass,omitempty"`
	ParentDeclContextId  string         `json:"parentDeclContextId,omitempty"`
	TargetLabelDeclId    string         `json:"targetLabelDeclId,omitempty"`
	DeclId               string         `json:"declId,omitempty"`
	Type                 *Type          `json:"type,omitempty"`
	ArgType              *Type          `json:"argType,omitempty"`
	ComputeLHSType       *Type          `json:"computeLHSType,omitempty"`
	ComputeResultType    *Type          `json:"computeResultType,omitempty"`
	Loc                  *Location      `json:"loc,omitempty"`
	Range                *LocationRange `json:"range,omitempty"`
	IsUsed               bool           `json:"isUsed,omitempty"`
	IsArrow              bool           `json:"isArrow,omitempty"`
	IsPostfix            bool           `json:"isPostfix,omitempty"`
	IsPartOfExplicitCast bool           `json:"isPartOfExplicitCast,omitempty"`
	IsImplicit           bool           `json:"isImplicit,omitempty"`
	IsReferenced         bool           `json:"isReferenced,omitempty"`
	IsBitfield           bool           `json:"isBitfield,omitempty"`
	Inline               bool           `json:"inline,omitempty"`
	Explicit             bool           `json:"explicit,omitempty"`
	Implicit             bool           `json:"implicit,omitempty"`
	Inherited            bool           `json:"inherited,omitempty"`
	CompleteDefinition   bool           `json:"completeDefinition,omitempty"`
	Variadic             bool           `json:"variadic,omitempty"`
	HasElse              bool           `json:"hasElse,omitempty"`
	CanOverflow          bool           `json:"canOverflow,omitempty"`
	Size                 int            `json:"size,omitempty"`
	Qualifiers           string         `json:"qualifiers,omitempty"`
	Text                 string         `json:"text,omitempty"`
	Cc                   string         `json:"cc,omitempty"`
	Init                 string         `json:"init,omitempty"`
	Opcode               string         `json:"opcode,omitempty"`
	Value                interface{}    `json:"value,omitempty"`
	ValueCategory        string         `json:"valueCategory,omitempty"`
	CastKind             string         `json:"castKind,omitempty"`
	PreviousDecl         string         `json:"previousDecl,omitempty"`
	ReferencedMemberDecl string         `json:"referencedMemberDecl,omitempty"`
	NonOdrUseReason      string         `json:"nonOdrUseReason,omitempty"`
	Decl                 *Term          `json:"decl,omitempty"`
	ReferencedDecl       *Term          `json:"referencedDecl,omitempty"`
	OwnedTagDecl         *Term          `json:"ownedTagDecl,omitempty"`
	TagUsed              string         `json:"tagUsed,omitempty"`
	Direction            string         `json:"direction,omitempty"`
	Param                string         `json:"param,omitempty"`
	ParamIdx             int            `json:"paramIdx,omitempty"`
	Inner                []*Term        `json:"inner,omitempty"`
	ArrayFiller          []*Term        `json:"array_filler,omitempty"`
}
