package main

import "github.com/hoisie/web"
import "github.com/hoisie/mustache"
import "io/ioutil"

type User struct {
	Pic        string
	ProperName string
	UserName   string
	Email      string
	Bio        string
	Skills     string
	Github     string
}

func dummyUser() *User {
	user := &User{
		Pic:        "http://0.gravatar.com/avatar/b7ddec29f78231d1a59752134b1f57fb",
		ProperName: "Eric Fode",
		UserName:   "ericfode",
		Email:      "ericfode@gmail.com",
		Bio:        "My Awsome Bio",
		Skills:     "Everthing, Being Awesome, Spelling badly, Stack Smashing, Makeing explosives",
		Github:     "github.com/ericfode"}
	return user
}

func renderProfile() string {
	html := mustache.RenderFileInLayout("Pages/DisplayProfile.mustache", "Pages/layout.mustache", dummyUser())
	return html
}

func renderCSS(val string) string {

	if css, err := ioutil.ReadFile("Pages/CSS/" + val); err != nil {
		return err.Error()
	} else {
		return string(css)
	}
	return ""

}

func main() {
	web.Get("/", renderProfile)
	web.Get("/CSS/(.*)", renderCSS)
	web.Run("0.0.0.0:9998")
}
