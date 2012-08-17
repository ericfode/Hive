package main

import "github.com/hoisie/web"
import "github.com/hoisie/mustache"

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
		Skills:     "Everthing",
		Github:     "github.com/ericfode"}
	return user
}

func renderProfile(val string) string {
	html := mustache.RenderFileInLayout("DisplayProfile.mustache", "layout.mustache", dummyUser())
	return html
}

func main() {
	web.Get("/(.*)", renderProfile)
	web.Run("0.0.0.0:9998")
}
