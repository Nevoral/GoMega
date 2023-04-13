package webJs

import (
	"fmt"
)

var Doctype = fmt.Sprintf("<!DOCTYPE html>")

/*
Comment function return string representing HTML tag witch represent comment. The value contains everything what you
want inside the tag (value like "string" or some tags).
*/
func Comment(value string) string {
	return fmt.Sprintf("<!--%s-->", value)
}

/*
A function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
  - download: Specifies that the target will be downloaded when a user clicks on the hyperlink.
  - href: Specifies the URL of the page the link goes to.
  - hreflang: Specifies the language of the linked document.
  - media: Specifies what media/device the linked document is optimized for
  - ping : Specifies a space-separated list of URLs to which, when the link is followed, post requests with the body ping will be sent by the browser (in the background). Typically used for tracking.
  - referrerpolicy: Specifies which referrer information to send with the link.
  - rel: Specifies the relationship between the current document and the linked document.
  - target: Specifies where to open the linked document.
  - type: Specifies the media type of the linked document.
*/
func A(atr string, value ...string) string {
	return CreateTag("a", atr, Val(value...))
}

/*
Abbr function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Abbr(atr string, value ...string) string {
	return CreateTag("abbr", atr, Val(value...))
}

/*
Address function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Address(atr string, value ...string) string {
	return CreateTag("address", atr, Val(value...))
}

/*
Area function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Area(atr string) string {
	return fmt.Sprintf("<area%s>", atr)
}

/*
Article function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Article(atr string, value ...string) string {
	return CreateTag("article", atr, Val(value...))
}

/*
Aside function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Aside(atr string, value ...string) string {
	return CreateTag("aside", atr, Val(value...))
}

/*
Audio function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Audio(atr string, value ...string) string {
	return CreateTag("audio", atr, Val(value...))
}

/*
B function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func B(atr string, value ...string) string {
	return CreateTag("b", atr, Val(value...))
}

/*
Base function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Base(atr string) string {
	return fmt.Sprintf("<base%s>", atr)
}

/*
Bdi function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Bdi(atr string, value ...string) string {
	return CreateTag("bdi", atr, Val(value...))
}

/*
Bdo function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Bdo(atr string, value ...string) string {
	return CreateTag("bdo", atr, Val(value...))
}

/*
Blockquote function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Blockquote(atr string, value ...string) string {
	return CreateTag("blockquote", atr, Val(value...))
}

/*
Body function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Body(atr string, value ...string) string {
	return CreateTag("body", atr, Val(value...))
}

/*
Br function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Br(atr string) string {
	return fmt.Sprintf("<br%s>", atr)
}

/*
Button function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Button(atr string, value ...string) string {
	return CreateTag("button", atr, Val(value...))
}

/*
Canvas function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Canvas(atr string, value ...string) string {
	return CreateTag("canvas", atr, Val(value...))
}

/*
Caption function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Caption(atr string, value ...string) string {
	return CreateTag("caption", atr, Val(value...))
}

/*
Cite function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Cite(atr string, value ...string) string {
	return CreateTag("cite", atr, Val(value...))
}

/*
Code function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Code(atr string, value ...string) string {
	return CreateTag("code", atr, Val(value...))
}

/*
Col function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Col(atr string) string {
	return fmt.Sprintf("<bol%s>", atr)
}

/*
Colgroup function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Colgroup(atr string, value ...string) string {
	return CreateTag("colgroup", atr, Val(value...))
}

/*
Data function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Data(atr string, value ...string) string {
	return CreateTag("data", atr, Val(value...))
}

/*
Datalist function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Datalist(atr string, value ...string) string {
	return CreateTag("datalist", atr, Val(value...))
}

/*
Dd function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Dd(atr string, value ...string) string {
	return CreateTag("dd", atr, Val(value...))
}

/*
Del function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Del(atr string, value ...string) string {
	return CreateTag("del", atr, Val(value...))
}

/*
Details function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Details(atr string, value ...string) string {
	return CreateTag("details", atr, Val(value...))
}

/*
Dfn function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Dfn(atr string, value ...string) string {
	return CreateTag("dfn", atr, Val(value...))
}

/*
Dialog function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Dialog(atr string, value ...string) string {
	return CreateTag("dialog", atr, Val(value...))
}

/*
Div function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Div(atr string, value ...string) string {
	return CreateTag("div", atr, Val(value...))
}

/*
Dl function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Dl(atr string, value ...string) string {
	return CreateTag("dl", atr, Val(value...))
}

/*
Dt function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Dt(atr string, value ...string) string {
	return CreateTag("dt", atr, Val(value...))
}

/*
Em function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Em(atr string, value ...string) string {
	return CreateTag("em", atr, Val(value...))
}

/*
Embed function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Embed(atr string) string {
	return fmt.Sprintf("<embed%s>", atr)
}

/*
Fieldset function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Fieldset(atr string, value ...string) string {
	return CreateTag("fieldset", atr, Val(value...))
}

/*
Figcaption function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Figcaption(atr string, value ...string) string {
	return CreateTag("figcaption", atr, Val(value...))
}

/*
Figure function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Figure(atr string, value ...string) string {
	return CreateTag("figure", atr, Val(value...))
}

/*
Footer function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Footer(atr string, value ...string) string {
	return CreateTag("footer", atr, Val(value...))
}

/*
Form function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Form(atr string, value ...string) string {
	return CreateTag("form", atr, Val(value...))
}

/*
H1 function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func H1(atr string, value ...string) string {
	return CreateTag("h1", atr, Val(value...))
}

/*
H2 function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func H2(atr string, value ...string) string {
	return CreateTag("h2", atr, Val(value...))
}

/*
H3 function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func H3(atr string, value ...string) string {
	return CreateTag("h3", atr, Val(value...))
}

/*
H4 function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func H4(atr string, value ...string) string {
	return CreateTag("h4", atr, Val(value...))
}

/*
H5 function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func H5(atr string, value ...string) string {
	return CreateTag("h5", atr, Val(value...))
}

/*
H6 function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func H6(atr string, value ...string) string {
	return CreateTag("h6", atr, Val(value...))
}

/*
Head function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Head(atr string, value ...string) string {
	return CreateTag("head", atr, Val(value...))
}

/*
Header function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Header(atr string, value ...string) string {
	return CreateTag("header", atr, Val(value...))
}

/*
Hgroup function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Hgroup(atr string, value ...string) string {
	return CreateTag("hgroup", atr, Val(value...))
}

/*
Hr function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Hr(atr string) string {
	return fmt.Sprintf("<hr%s>", atr)
}

/*
Html function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Html(atr string, value ...string) string {
	return CreateTag("html", atr, Val(value...))
}

/*
I function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func I(atr string, value ...string) string {
	return CreateTag("i", atr, Val(value...))
}

/*
Iframe function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Iframe(atr string, value ...string) string {
	return CreateTag("iframe", atr, Val(value...))
}

/*
Img function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Img(atr string) string {
	return fmt.Sprintf("<img%s>", atr)
}

/*
Input function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Input(atr string) string {
	return fmt.Sprintf("<input%s>", atr)
}

/*
Ins function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Ins(atr string, value ...string) string {
	return CreateTag("ins", atr, Val(value...))
}

/*
Kbd function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Kbd(atr string, value ...string) string {
	return CreateTag("kbd", atr, Val(value...))
}

/*
Label function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Label(atr string, value ...string) string {
	return CreateTag("label", atr, Val(value...))
}

/*
Legend function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Legend(atr string, value ...string) string {
	return CreateTag("legend", atr, Val(value...))
}

/*
Li function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Li(atr string, value ...string) string {
	return CreateTag("li", atr, Val(value...))
}

/*
Link function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Link(atr string) string {
	return fmt.Sprintf("<link%s>", atr)
}

/*
Main function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Main(atr string, value ...string) string {
	return CreateTag("main", atr, Val(value...))
}

/*
Map function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Map(atr string, value ...string) string {
	return CreateTag("map", atr, Val(value...))
}

/*
Mark function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Mark(atr string, value ...string) string {
	return CreateTag("mark", atr, Val(value...))
}

/*
Menu function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Menu(atr string, value ...string) string {
	return CreateTag("menu", atr, Val(value...))
}

/*
Meta function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Meta(atr string) string {
	return fmt.Sprintf("<meta%s>", atr)
}

/*
Meter function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Meter(atr string, value ...string) string {
	return CreateTag("meter", atr, Val(value...))
}

/*
Nav function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Nav(atr string, value ...string) string {
	return CreateTag("nav", atr, Val(value...))
}

/*
Noscript function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Noscript(atr string, value ...string) string {
	return CreateTag("noscript", atr, Val(value...))
}

/*
Object function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Object(atr string, value ...string) string {
	return CreateTag("object", atr, Val(value...))
}

/*
Ol function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Ol(atr string, value ...string) string {
	return CreateTag("ol", atr, Val(value...))
}

/*
Optgroup function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Optgroup(atr string, value ...string) string {
	return CreateTag("optgroup", atr, Val(value...))
}

/*
Option function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Option(atr string, value ...string) string {
	return CreateTag("option", atr, Val(value...))
}

/*
Output function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Output(atr string, value ...string) string {
	return CreateTag("output", atr, Val(value...))
}

/*
P function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func P(atr string, value ...string) string {
	return CreateTag("p", atr, Val(value...))
}

/*
Picture function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Picture(atr string, value ...string) string {
	return CreateTag("picture", atr, Val(value...))
}

/*
Pre function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Pre(atr string, value ...string) string {
	return CreateTag("pre", atr, Val(value...))
}

/*
Progress function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Progress(atr string, value ...string) string {
	return CreateTag("progress", atr, Val(value...))
}

/*
Q function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Q(atr string, value ...string) string {
	return CreateTag("q", atr, Val(value...))
}

/*
Rp function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Rp(atr string, value ...string) string {
	return CreateTag("rp", atr, Val(value...))
}

/*
Rt function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Rt(atr string, value ...string) string {
	return CreateTag("rt", atr, Val(value...))
}

/*
Ruby function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Ruby(atr string, value ...string) string {
	return CreateTag("ruby", atr, Val(value...))
}

/*
S function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func S(atr string, value ...string) string {
	return CreateTag("s", atr, Val(value...))
}

/*
Samp function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Samp(atr string, value ...string) string {
	return CreateTag("samp", atr, Val(value...))
}

/*
Script function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Script(atr string, value ...string) string {
	return CreateTag("script", atr, Val(value...))
}

/*
Section function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Section(atr string, value ...string) string {
	return CreateTag("section", atr, Val(value...))
}

/*
Select function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Select(atr string, value ...string) string {
	return CreateTag("select", atr, Val(value...))
}

/*
Slot function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Slot(atr string, value ...string) string {
	return CreateTag("slot", atr, Val(value...))
}

/*
Small function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Small(atr string, value ...string) string {
	return CreateTag("small", atr, Val(value...))
}

/*
Source function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Source(atr string) string {
	return fmt.Sprintf("<source%s>", atr)
}

/*
Span function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Span(atr string, value ...string) string {
	return CreateTag("span", atr, Val(value...))
}

/*
Strong function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Strong(atr string, value ...string) string {
	return CreateTag("strong", atr, Val(value...))
}

/*
Style function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Style(atr string, value ...string) string {
	return CreateTag("style", atr, Val(value...))
}

/*
Sub function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Sub(atr string, value ...string) string {
	return CreateTag("sub", atr, Val(value...))
}

/*
Summary function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Summary(atr string, value ...string) string {
	return CreateTag("summary", atr, Val(value...))
}

/*
Sup function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Sup(atr string, value ...string) string {
	return CreateTag("sup", atr, Val(value...))
}

/*
Svg function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Svg(atr string, value ...string) string {
	return CreateTag("svg", atr, Val(value...))
}

/*
Table function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Table(atr string, value ...string) string {
	return CreateTag("table", atr, Val(value...))
}

/*
Tbody function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Tbody(atr string, value ...string) string {
	return CreateTag("tbody", atr, Val(value...))
}

/*
Td function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Td(atr string, value ...string) string {
	return CreateTag("td", atr, Val(value...))
}

/*
Template function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Template(atr string, value ...string) string {
	return CreateTag("template", atr, Val(value...))
}

/*
Textarea function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Textarea(atr string, value ...string) string {
	return CreateTag("textarea", atr, Val(value...))
}

/*
Tfoot function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Tfoot(atr string, value ...string) string {
	return CreateTag("tfoot", atr, Val(value...))
}

/*
Th function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Th(atr string, value ...string) string {
	return CreateTag("th", atr, Val(value...))
}

/*
Thead function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Thead(atr string, value ...string) string {
	return CreateTag("thead", atr, Val(value...))
}

/*
Time function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Time(atr string, value ...string) string {
	return CreateTag("time", atr, Val(value...))
}

/*
Title function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Title(atr string, value ...string) string {
	return CreateTag("title", atr, Val(value...))
}

/*
Tr function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Tr(atr string, value ...string) string {
	return CreateTag("tr", atr, Val(value...))
}

/*
Track function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Track(atr string) string {
	return fmt.Sprintf("<track%s>", atr)
}

/*
U function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func U(atr string, value ...string) string {
	return CreateTag("u", atr, Val(value...))
}

/*
Ul function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Ul(atr string, value ...string) string {
	return CreateTag("ul", atr, Val(value...))
}

/*
Var function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Var(atr string, value ...string) string {
	return CreateTag("var", atr, Val(value...))
}

/*
Video function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Video(atr string, value ...string) string {
	return CreateTag("video", atr, Val(value...))
}

/*
Wbr function return string representing HTML tag with given arguments and value.
The "atr" contains all Global and Event attributes.
The value contains everything what you want inside the tag (value like "string" or some tags).
*/
func Wbr(atr string, value ...string) string {
	return CreateTag("wbr", atr, Val(value...))
}

func Animate(atr string, value ...string) string {
	return CreateTag("animate", atr, Val(value...))
}

func AnimateMotion(atr string, value ...string) string {
	return CreateTag("animateMotion", atr, Val(value...))
}

func AnimateTransform(atr string, value ...string) string {
	return CreateTag("animateTransform", atr, Val(value...))
}

func Circle(atr string, value ...string) string {
	return CreateTag("circle", atr, Val(value...))
}

func ClipPath(atr string, value ...string) string {
	return CreateTag("clipPath", atr, Val(value...))
}

func Defs(atr string, value ...string) string {
	return CreateTag("defs", atr, Val(value...))
}

func Desc(atr string, value ...string) string {
	return CreateTag("desc", atr, Val(value...))
}

func Discard(atr string, value ...string) string {
	return CreateTag("discard", atr, Val(value...))
}

func Ellipse(atr string, value ...string) string {
	return CreateTag("ellipse", atr, Val(value...))
}

func FeBlend(atr string, value ...string) string {
	return CreateTag("feBlend", atr, Val(value...))
}

func FeColorMatrix(atr string, value ...string) string {
	return CreateTag("feColorMatrix", atr, Val(value...))
}

func FeComponentTransfer(atr string, value ...string) string {
	return CreateTag("feComponentTransfer", atr, Val(value...))
}

func FeComposite(atr string, value ...string) string {
	return CreateTag("feComposite", atr, Val(value...))
}

func FeConvolveMatrix(atr string, value ...string) string {
	return CreateTag("feConvolveMatrix", atr, Val(value...))
}

func FeDiffuseLighting(atr string, value ...string) string {
	return CreateTag("feDiffuseLighting", atr, Val(value...))
}

func FeDisplacementMap(atr string, value ...string) string {
	return CreateTag("feDisplacementMap", atr, Val(value...))
}

func FeDistantLight(atr string, value ...string) string {
	return CreateTag("feDistantLight", atr, Val(value...))
}

func FeDropShadow(atr string, value ...string) string {
	return CreateTag("feDropShadow", atr, Val(value...))
}

func FeFlood(atr string, value ...string) string {
	return CreateTag("feFlood", atr, Val(value...))
}

func FeFuncA(atr string, value ...string) string {
	return CreateTag("feFuncA", atr, Val(value...))
}

func FeFuncB(atr string, value ...string) string {
	return CreateTag("feFuncB", atr, Val(value...))
}

func FeFuncG(atr string, value ...string) string {
	return CreateTag("feFuncG", atr, Val(value...))
}

func FeFuncR(atr string, value ...string) string {
	return CreateTag("feFuncR", atr, Val(value...))
}

func FeGaussianBlur(atr string, value ...string) string {
	return CreateTag("feGaussianBlur", atr, Val(value...))
}

func FeImage(atr string, value ...string) string {
	return CreateTag("feImage", atr, Val(value...))
}

func FeMerge(atr string, value ...string) string {
	return CreateTag("feMerge", atr, Val(value...))
}

func FeMergeNode(atr string, value ...string) string {
	return CreateTag("feMergeNode", atr, Val(value...))
}

func FeMorphology(atr string, value ...string) string {
	return CreateTag("feMorphology", atr, Val(value...))
}

func FeOffset(atr string, value ...string) string {
	return CreateTag("feOffset", atr, Val(value...))
}

func FePointLight(atr string, value ...string) string {
	return CreateTag("fePointLight", atr, Val(value...))
}

func FeSpecularLighting(atr string, value ...string) string {
	return CreateTag("feSpecularLighting", atr, Val(value...))
}

func FeSpotLight(atr string, value ...string) string {
	return CreateTag("feSpotLight", atr, Val(value...))
}

func FeTile(atr string, value ...string) string {
	return CreateTag("feTile", atr, Val(value...))
}

func FeTurbulence(atr string, value ...string) string {
	return CreateTag("feTurbulence", atr, Val(value...))
}

func Filter(atr string, value ...string) string {
	return CreateTag("filter", atr, Val(value...))
}

func ForeignObject(atr string, value ...string) string {
	return CreateTag("foreignObject", atr, Val(value...))
}

func G(atr string, value ...string) string {
	return CreateTag("g", atr, Val(value...))
}

func Image(atr string, value ...string) string {
	return CreateTag("image", atr, Val(value...))
}

func Line(atr string, value ...string) string {
	return CreateTag("line", atr, Val(value...))
}

func LinearGradient(atr string, value ...string) string {
	return CreateTag("linearGradient", atr, Val(value...))
}

func Marker(atr string, value ...string) string {
	return CreateTag("marker", atr, Val(value...))
}

func Mask(atr string, value ...string) string {
	return CreateTag("mask", atr, Val(value...))
}

func Metadata(atr string, value ...string) string {
	return CreateTag("metadata", atr, Val(value...))
}

func Mpath(atr string, value ...string) string {
	return CreateTag("mpath", atr, Val(value...))
}

func Path(atr string, value ...string) string {
	return CreateTag("path", atr, Val(value...))
}

func Pattern(atr string, value ...string) string {
	return CreateTag("pattern", atr, Val(value...))
}

func Polygon(atr string, value ...string) string {
	return CreateTag("polygon", atr, Val(value...))
}

func Polyline(atr string, value ...string) string {
	return CreateTag("polyline", atr, Val(value...))
}

func RadialGradient(atr string, value ...string) string {
	return CreateTag("radialGradient", atr, Val(value...))
}

func Rect(atr string, value ...string) string {
	return CreateTag("rect", atr, Val(value...))
}

func Set(atr string, value ...string) string {
	return CreateTag("set", atr, Val(value...))
}

func Stop(atr string, value ...string) string {
	return CreateTag("stop", atr, Val(value...))
}

func Switch(atr string, value ...string) string {
	return CreateTag("switch", atr, Val(value...))
}

func Symbol(atr string, value ...string) string {
	return CreateTag("symbol", atr, Val(value...))
}

func Text(atr string, value ...string) string {
	return CreateTag("text", atr, Val(value...))
}

func TextPath(atr string, value ...string) string {
	return CreateTag("textPath", atr, Val(value...))
}

func Tspan(atr string, value ...string) string {
	return CreateTag("tspan", atr, Val(value...))
}

func Use(atr string, value ...string) string {
	return CreateTag("use", atr, Val(value...))
}

func View(atr string, value ...string) string {
	return CreateTag("view", atr, Val(value...))
}
