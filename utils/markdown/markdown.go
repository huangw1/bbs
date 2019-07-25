/**
 * @Author: huangw1
 * @Date: 2019/7/26 10:59
 */

package markdown

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/huangw1/bbs/utils/extension"
	"github.com/iris-contrib/blackfriday"
	"github.com/microcosm-cc/bluemonday"
	"regexp"
	"strings"
)

type MDResult struct {
	ContentHtml string
	SummaryText string
	TocHtml     string
	ThumbUrl    string
}

type MDOption struct {
	toc               bool
	thumb             bool
	summaryTextLength int
}

type MDOptionFunc func(*MDOption)

func MdWithTOC() MDOptionFunc {
	return func(mdOption *MDOption) {
		mdOption.toc = true
	}
}

func MdWithThumb() MDOptionFunc {
	return func(mdOption *MDOption) {
		mdOption.thumb = true
	}
}

func MdWithSummaryLength(summaryLength int) MDOptionFunc {
	return func(mdOption *MDOption) {
		mdOption.summaryTextLength = summaryLength
	}
}

func NewMD(options ...MDOptionFunc) *MDOption {
	mdOption := &MDOption{
		toc:               false,
		thumb:             false,
		summaryTextLength: 256,
	}
	if len(options) > 0 {
		for _, optionFunc := range options {
			optionFunc(mdOption)
		}
	}
	return mdOption
}

func (md *MDOption) Run(mdText string) *MDResult {
	mdText = strings.Replace(mdText, "\r\n", "\n", -1)
	var unsafe []byte
	if md.toc {
		option := blackfriday.WithRenderer(blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
			Flags: blackfriday.CommonHTMLFlags | blackfriday.TOC,
		}))
		unsafe = blackfriday.Run([]byte(mdText), option)
	} else {
		unsafe = blackfriday.Run([]byte(mdText))
	}

	content := string(unsafe)

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
	doc.Find("*").Contents().FilterFunction(func(i int, selection *goquery.Selection) bool {
		// todo
		return true
	})
	tocHtml := md.buildTOC(doc)
	doc.Find("nav").Remove()
	contentHTML, _ := doc.Find("body").Html()
	return &MDResult{
		TocHtml:     tocHtml,
		SummaryText: md.buildSummary(doc),
		ContentHtml: md.sanitize(contentHTML),
		ThumbUrl:    md.buildThumbURL(doc),
	}
}

func (md *MDOption) buildTOC(doc *goquery.Document) string {
	if !md.toc {
		return ""
	}
	top := doc.Find("nav > ul > li")
	if top.Size() == 0 {
		return ""
	}
	contentTOC, _ := doc.Find("nav").First().Html()
	return contentTOC
}

func (md *MDOption) buildSummary(doc *goquery.Document) string {
	if md.summaryTextLength <= 0 {
		return ""
	}
	summary := doc.Text()
	summary = strings.TrimSpace(summary)
	return extension.GetSummary(summary, md.summaryTextLength)
}

func (md *MDOption) buildThumbURL(doc *goquery.Document) string {
	if !md.thumb {
		return ""
	}
	selection := doc.Find("img").First()
	thumbnailURL, _ := selection.Attr("src")
	if thumbnailURL == "" {
		thumbnailURL, _ = selection.Attr("data-src")
	}
	return thumbnailURL
}

func (md *MDOption) sanitize(html string) string {
	return bluemonday.UGCPolicy().AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code").
		AllowAttrs("data-src").OnElements("img").
		AllowAttrs("class", "target", "id", "style").Globally().
		AllowAttrs("src", "width", "height", "border", "marginwidth", "marginheight").OnElements("iframe").
		AllowAttrs("controls", "src").OnElements("audio").
		AllowAttrs("color").OnElements("font").
		AllowAttrs("controls", "src", "width", "height").OnElements("video").
		AllowAttrs("src", "media", "type").OnElements("source").
		AllowAttrs("width", "height", "data", "type").OnElements("object").
		AllowAttrs("name", "value").OnElements("param").
		AllowAttrs("src", "type", "width", "height", "wmode", "allowNetworking").OnElements("embed").
		Sanitize(html)
}
