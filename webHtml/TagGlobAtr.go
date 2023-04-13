package webJs

import "fmt"

func AccessKey(atr string) string {
	return fmt.Sprintf(" accesskey=\"%s\"", atr)
}

func Aria(name string, atr string) string {
	if name == "" {
		return fmt.Sprintf(" aria=\"%s\"", atr)
	}
	return fmt.Sprintf(" aria-%s=\"%s\"", name, atr)
}

func Autocapitalize(atr string) string {
	return fmt.Sprintf(" autocapitalize=\"%s\"", atr)
}

func Autofocus(atr string) string {
	return fmt.Sprintf(" autofocus=\"%s\"", atr)
}

func Class(atr string) string {
	return fmt.Sprintf(" class=\"%s\"", atr)
}

func Contenteditable(atr string) string {
	return fmt.Sprintf(" contenteditable=\"%s\"", atr)
}

func DataAtr(name string, atr string) string {
	if name == "" {
		return fmt.Sprintf(" data=\"%s\"", atr)
	}
	return fmt.Sprintf(" data-%s=\"%s\"", name, atr)
}

func Dir(atr string) string {
	return fmt.Sprintf(" dir=\"%s\"", atr)
}

func Draggable(atr string) string {
	return fmt.Sprintf(" draggable=\"%s\"", atr)
}

func EnterKeyHint(atr string) string {
	return fmt.Sprintf(" enterkeyhint=\"%s\"", atr)
}

func ExportParts(atr string) string {
	return fmt.Sprintf(" exportparts=\"%s\"", atr)
}

func Hidden(atr string) string {
	return fmt.Sprintf(" hidden=\"%s\"", atr)
}

func Id(atr string) string {
	return fmt.Sprintf(" id=\"%s\"", atr)
}

func Inert(atr string) string {
	return fmt.Sprintf(" inert=\"%s\"", atr)
}

func InputMode(atr string) string {
	return fmt.Sprintf(" inputmode=\"%s\"", atr)
}

func Is(atr string) string {
	return fmt.Sprintf(" is=\"%s\"", atr)
}

func ItemId(atr string) string {
	return fmt.Sprintf(" itemid=\"%s\"", atr)
}

func ItemProp(atr string) string {
	return fmt.Sprintf(" itemprop=\"%s\"", atr)
}

func ItemRef(atr string) string {
	return fmt.Sprintf(" itemref=\"%s\"", atr)
}

func ItemScope(atr string) string {
	return fmt.Sprintf(" itemscope=\"%s\"", atr)
}

func ItemType(atr string) string {
	return fmt.Sprintf(" itemtype=\"%s\"", atr)
}

func Lang(atr string) string {
	return fmt.Sprintf(" lang=\"%s\"", atr)
}

func Nonce(atr string) string {
	return fmt.Sprintf(" nonce=\"%s\"", atr)
}

func Part(atr string) string {
	return fmt.Sprintf(" part=\"%s\"", atr)
}

func Role(atr string) string {
	return fmt.Sprintf(" role=\"%s\"", atr)
}

func SlotAtr(atr string) string {
	return fmt.Sprintf(" slot=\"%s\"", atr)
}

func Spellcheck(atr string) string {
	return fmt.Sprintf(" spellcheck=\"%s\"", atr)
}

func StyleAtr(atr string) string {
	return fmt.Sprintf(" style=\"%s\"", atr)
}

func Tabindex(atr string) string {
	return fmt.Sprintf(" tabindex=\"%s\"", atr)
}

func TitleAtr(atr string) string {
	return fmt.Sprintf(" title=\"%s\"", atr)
}

func Translate(atr string) string {
	return fmt.Sprintf(" translate=\"%s\"", atr)
}

func VirtualKeyBoardPolicy(atr string) string {
	return fmt.Sprintf(" virtualkeyboardpolicy=\"%s\"", atr)
}
