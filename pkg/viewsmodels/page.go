package viewsmodels

import (
	"strconv"
	"time"

	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/pill.go/antiCSRF"
)

type Page struct {
	*Global
	Title           string
	Description     string
	Keywords        string
	Image           string
	PageName        string
	PageURL         string
	ActivePage      string
	SubDomain       string
	URLQueryParams  map[string]string
	IsAuthenticated bool
	CSRFToken       *CSRFToken
	Link            string
}

func (this *Page) SetTitle(value string) {
	if value == "" {
		this.Title = this.TextualContent.OfMeta("title")
	} else {
		this.Title = value + " | " + this.TextualContent.OfMeta("title")
	}
}

func (this *Page) SetDescription(value string) {
	if value == "" {
		this.Description = this.TextualContent.OfMeta("description")
	} else {
		this.Description = value
	}
}

func (this *Page) SetURL(URLPath string) {
	this.PageURL = webHost + URLPath
}

func (this *Page) SetImage(value string) {
	if value == "" {
		this.Image = webHost + defaultPageImage
	} else {
		this.Image = webHost + value
	}
}
func (this *Page) SetKeywords(value string) {
	if value == "" {
		this.Keywords = this.TextualContent.OfMeta("keywords")
	} else {
		this.Keywords = value
	}
}

func (this *Page) SetMetas(metas ...string) {
	var title, description, keywords, image string
	if metas != nil {
		if len(metas) > 0 {
			title = metas[0]
		}
		if len(metas) > 1 {
			description = metas[1]
		}
		if len(metas) > 2 {
			keywords = metas[2]
		}
		if len(metas) > 3 {
			image = metas[3]
		}
	}
	this.SetTitle(title)
	this.SetDescription(description)
	this.SetKeywords(keywords)
	this.SetImage(image)

}

func (this *Page) SetCSRFToken(csrfToken *antiCSRF.CSRFToken) {
	this.CSRFToken = &CSRFToken{}
	if csrfToken != nil {
		this.CSRFToken.HiddenField = csrfToken.HTMLInput()
		this.CSRFToken.Value = csrfToken.MaskedToken
	}
}

func NewPage(pageName string, context *appCtx.Context, active ...interface{}) *Page {
	page := &Page{}
	page.PageName = pageName
	page.Global = global
	if page.Environment == "development" {
		page.AppVersion = strconv.Itoa(int(time.Now().Unix()))
	}
	page.SetMetas()
	page.SetURL(context.URLPath)
	page.SetCSRFToken(context.CSRFToken)
	page.Link = "http://192.168.0.108:9000"
	urlQueryParams := map[string]string{}
	//set query params
	page.URLQueryParams = urlQueryParams
	if active != nil && len(active) > 0 {
		if isActive, ok := active[0].(bool); ok {
			if isActive {
				page.ActivePage = pageName
			}
		} else {
			page.ActivePage, _ = active[0].(string)
		}
	}
	return page
}
