package core

import (
	"errors"
	"strings"
)

func (x *WildcardString) Format() string {
	if x != nil {
		sb := strings.Builder{}
		for _, wildcard := range x.Wildcards {
			switch wildcard.Type {
			case WildcardString_Wildcard_TYPE_SINGLE_CHARACTER:
				sb.WriteByte('.')
			case WildcardString_Wildcard_TYPE_ZERO_OR_MORE_CHARACTERS:
				sb.WriteByte('*')
			case WildcardString_Wildcard_TYPE_IN_CHARACTERS:
				sb.WriteByte('[')
				sb.WriteString(wildcard.Content)
				sb.WriteByte(']')
			case WildcardString_Wildcard_TYPE_NOT_IN_CHARACTERS:
				sb.WriteByte('[')
				sb.WriteByte('^')
				sb.WriteString(wildcard.Content)
				sb.WriteByte(']')
			default:
				sb.WriteString(wildcard.Content)
			}
		}
		return sb.String()
	}
	return ""
}

func (x *WildcardString) ToString() string {
	return x.Format()
}

func ParseWildcardString(value string) (*WildcardString, error) {
	ts := &WildcardString{}
	if err := ts.Parse(value); err != nil {
		return nil, err
	}
	return ts, nil
}

func (x *WildcardString) Parse(value string) error {
	if x != nil {
		wildcard := &WildcardString_Wildcard{}

		beginWildcard := func() {
			x.Wildcards = append(x.Wildcards, wildcard)
			wildcard = &WildcardString_Wildcard{}
		}

		inRange := false
		for i, ch := range value {
			if inRange {
				if ch == ']' && value[i-1] != '\\' {
					inRange = false
					beginWildcard()
				} else if ch == '^' && value[i-1] == '[' {
					// noop
				} else {
					wildcard.Content = wildcard.Content + string(ch)
				}
				continue
			}

			if ch == '.' && (i == 0 || value[i-1] != '\\') {
				if len(wildcard.Content) > 0 {
					beginWildcard()
				}

				wildcard.Type = WildcardString_Wildcard_TYPE_SINGLE_CHARACTER
				beginWildcard()
			} else if ch == '*' && (i == 0 || value[i-1] != '\\') {
				if len(wildcard.Content) > 0 {
					beginWildcard()
				}

				wildcard.Type = WildcardString_Wildcard_TYPE_ZERO_OR_MORE_CHARACTERS
				beginWildcard()
			} else if ch == '[' && (i == 0 || value[i-1] != '\\') {
				if len(wildcard.Content) > 0 {
					beginWildcard()
				}

				if i == len(value)-1 {
					return errors.New("syntax error for the WildcardString: " + value)
				}

				if value[i+1] == '^' {
					wildcard.Type = WildcardString_Wildcard_TYPE_NOT_IN_CHARACTERS
				} else {
					wildcard.Type = WildcardString_Wildcard_TYPE_IN_CHARACTERS
				}
				inRange = true
			} else if ch == '\\' {
				// noop
			} else {
				wildcard.Content = wildcard.Content + string(ch)
			}
		}

		if wildcard.Type > 0 || len(wildcard.Content) > 0 {
			x.Wildcards = append(x.Wildcards, wildcard)
		}
	}
	return nil
}
