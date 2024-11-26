package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("first_name", "last_name", "email")

	if form.Valid() {
		t.Error("it shouldn't be valid")
	}

	postedData := url.Values{}
	postedData.Add("email", "email")
	postedData.Add("first_name", "first_name")
	postedData.Add("last_name", "last_name")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("first_name", "last_name", "email")

	if !form.Valid() {
		t.Error("shows does not have required field when it does")
	}
}

func TestForm_New(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("email", "email")

	r := httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form := New(r.PostForm)

	if !form.Valid() {
		t.Error("got invalid when should have been valid")
	}

	form.Errors.Add("error1", "test errors")

	if form.Valid() {
		t.Error("got invalid when should have been valid")
	}

	if form.Values.Get("email") != "email" {
		t.Error("there is empty email")
	}

}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	if form.Has("last_name") {
		t.Error("There is no last_name yet")
	}

	postedData := url.Values{}
	postedData.Add("last_name", "last_name")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)

	if !form.Has("last_name") {
		t.Error("There is no last_name yet")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("email", "email")

	r := httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form := New(r.PostForm)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("It's not an email")
	}

	postedData.Del("email")
	postedData.Add("email", "sinyakov@email.com")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("It's not an email")
	}

}

func TestFrom_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("a", "a")

	r.PostForm = postedData
	form := New(r.PostForm)

	if form.MinLength("a", 3) {
		t.Error("Not enough chars")
	}

	r, _ = http.NewRequest("POST", "/whatever", nil)

	postedData = url.Values{}
	postedData.Add("b", "bbb")

	r.PostForm = postedData
	form = New(r.PostForm)

	if !form.MinLength("b", 3) {
		t.Error("There should be enough chars")
	}
}
