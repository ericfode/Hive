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

type StreamItem struct {
	Pic      string
	UserName string
	JIT      string
}

func dummyStreamItem() *StreamItem {
	item := &StreamItem{
		Pic:      "http://0.gravatar.com/avatar/b7ddec29f78231d1a59752134b1f57fb",
		UserName: "ericfode",
		JIT:      "This might actaully end up being kinda cool"}
	return item
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

func renderSideBar() string {

}

func renderFollowers() string {

}

func renderFollowing() string {

}

func renderStream() string {
	items := []*streamItem{
		dummyStreamItem(), dummyStreamItem(), dummyStreamItem()}
	html := mustache.RenderFileInLayout("Pages/Stream.mustache", "Pages/layout.mustache", items)
	return html
}

func renderStreamItem() string {
	item := dummyStreamItem()
	html := mustache.RenderInLayout("Pages/StreamItem.mustache", "Pages/layout.mustache", item)
}

func renderCSS(val string) string {

	if css, err := ioutil.ReadFile("Pages/CSS/" + val); err != nil {
		return err.Error()
	} else {
		return string(css)
	}
	return ""

}

func renderIMG(val string) string{
	
}

func renderJS(val string) string {

	if css, err := ioutil.ReadFile("Pages/JS/" + val); err != nil {
		return err.Error()
	} else {
		return string(css)
	}
	return ""

}

func main() {
	web.Get("/", renderProfile)
	web.Get("/CSS/(.*)", renderCSS)
	web.Get("/JS/(.*)", renderJS)
	web.Get("/Stream", renderStream)
	web.Run("0.0.0.0:9998")
}
