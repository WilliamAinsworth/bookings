package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {
	request := httptest.NewRequest("POST", "/test", nil)
	form := New(request.PostForm)

	has := form.Has("test")
	if has {
		t.Error("forms shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("test", "a")

	form = New(postedData)

	has = form.Has("test")

	if !has {
		t.Error("shows form does not have field when it should")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("test")
	if form.Valid() {
		t.Error("form shows valid when field does not exist")
	}

	postedData = url.Values{}
	postedData.Add("test", "test@test.com")
	form = New(postedData)
	form.IsEmail("test")
	if !form.Valid() {
		t.Error("form shows invalid when correct email is given")
	}
}

func TestForm_MinLength(t *testing.T) {
	request := httptest.NewRequest("POST", "/test", nil)
	form := New(request.PostForm)

	form.MinLength("test", 3)
	if form.Valid() {
		t.Error("form shows valid when field does not exist")
	}

	isError := form.Errors.Get("test")
	if isError == "" {
		t.Error("should have an error but did not get one")
	}

	postedData := url.Values{}
	postedData.Add("test1", "aa")
	form = New(postedData)
	form.MinLength("test1", 3)
	if form.Valid() {
		t.Error("form shows valid when field does not meet minimum length")
	}

	postedData = url.Values{}
	postedData.Add("test2", "aaaa")
	form = New(postedData)
	form.MinLength("test2", 3)

	if !form.Valid() {
		t.Error("form is invalid when field meets minimum length")
	}
	isError = form.Errors.Get("test2")
	if isError != "" {
		t.Error("should not have an error but got one")
	}

}

func TestForm_Required(t *testing.T) {
	request := httptest.NewRequest("POST", "/test", nil)
	form := New(request.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	request, _ = http.NewRequest("POST", "/test", nil)

	request.PostForm = postedData
	form = New(request.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have valid field when it does")
	}
}

func TestForm_Valid(t *testing.T) {

	request := httptest.NewRequest("POST", "/test", nil)
	form := New(request.PostForm)

	if !form.Valid() {
		t.Error("got invalid when should have been valid")
	}
}

func TestNew(t *testing.T) {

}
