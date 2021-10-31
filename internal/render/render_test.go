package render

import (
	"github.com/WilliamAinsworth/bookings/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var templateData models.TemplateData

	request, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(request.Context(), "flash", "123")

	result := AddDefaultData(&templateData, request)

	if result.Flash != "123" {
		t.Error("flash value of 123 not found")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = tc

	request, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplate(&ww, request, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser")
	}

	err = RenderTemplate(&ww, request, "non-existent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("rendered template that doesn't exist")
	}
}

func getSession() (*http.Request, error) {
	request, err := http.NewRequest("GET", "some-url", nil)
	if err != nil {
		return nil, err
	}
	ctx := request.Context()
	ctx, _ = session.Load(ctx, request.Header.Get("X-Session"))
	request = request.WithContext(ctx)

	return request, nil
}

func TestNewTemplate(t *testing.T) {
	NewTemplate(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
