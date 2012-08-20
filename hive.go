package main

import "github.com/hoisie/web"
import "github.com/hoisie/mustache"
import "io/ioutil"

//TODO: refactor all structs out to own files and create render and compile methods for each

//Render methods are for getting a page with just that items content
type User struct {
	Pic        string
	ProperName string
	UserName   string
	Email      string
	Bio        string
	Skills     string
	Github     string
	Html       string
}

type StreamItem struct {
	Pic      string
	UserName string
	JIT      string
	Id       string
	Html     string
}

type Stream struct {
	Items []*StreamItem
	Html  string
}

type Home struct {
	StreamRender *Stream
	CardRender   *User
	FollowRender *Follow
	Html         string
}

type Follow struct {
	FollowedBy []*User
	Following  []*User
	Html       string
}

func initDummys() []*User {
	var testNodes []*User
	const numdum = 8
	pics := [numdum]string{"picTEST", "http://4.bp.blogspot.com/-Q2hjS1dS1R8/T4YXpOfNjOI/AAAAAAAAAxQ/c-V_1FkMYmo/s1600/Bug.jpg",
		"https://encrypted-tbn1.google.com/images?q=tbn:ANd9GcRs5AS0g3hHRdJsO7gBgwu9v1Hr4grtuc_G1dh59MbxEVW3VH-GNw",
		"https://encrypted-tbn3.google.com/images?q=tbn:ANd9GcSJVzRTk5jiGvRIcKQZs-pm4__kMQOWae0WGGl3H32xZCTvci9U",
		"https://encrypted-tbn3.google.com/images?q=tbn:ANd9GcQ6VCAy3UhBqNohPBG1Dr5nVd2WfwTLnINK_pmh0Wo7RUPh7vwpjw",
		"https://encrypted-tbn1.google.com/images?q=tbn:ANd9GcQ677iObh3n9DhnfwvpFUH-ksX9mXv3kyS_h7npytmLACpe9EZX",
		"https://encrypted-tbn3.google.com/images?q=tbn:ANd9GcR94C_rLFc1arqiV_Dmi6LHIQzEVWvOFJg7TxdpdR-PxtVxVLAr",
		"https://encrypted-tbn3.google.com/images?q=tbn:ANd9GcRPaCON4nIzFMqrfCVuWAn8HoD0zH-ir-KovxFxwgy6ocUlYxHJ"}
	bios := [numdum]string{
		"bioTEST",
		"I don't exist",
		"I am from HERE",
		"I am from THERE",
		"I am from A",
		"I am from B",
		"I am from C",
		"I am from D"}
	names := [numdum]string{"nameTEST", "Wedunno", "Joe", "Bill", "Jane", "Sue", "Sally", "Tom"}
	users := [numdum]string{"userTEST", "Whothatis", "jmk", "bill-o-rama", "sparkles", "user", "user", "uzaaah"}
	github := [numdum]string{"githubTEST",
		"http://github.com/IAMAUSER",
		"http://github.com/IAMAUSER",
		"http://github.com/IAMAUSER",
		"http://github.com/IAMAUSER",
		"http://github.com/IAMAUSER",
		"http://github.com/IAMAUSER",
		"http://github.com/IAMAUSER"}

	for k, _ := range pics {

		newNode := &User{Pic: pics[k], ProperName: names[k], UserName: users[k], Email: "email@somewhere.com", Bio: bios[k], Skills: "skills", Github: github[k]}
		testNodes = append(testNodes, newNode)
	}

	return testNodes
}

func dummyStreamItem() *StreamItem {
	item := &StreamItem{
		Pic:      "http://0.gravatar.com/avatar/b7ddec29f78231d1a59752134b1f57fb",
		UserName: "ericfode",
		JIT:      "This might actaully end up being kinda cool/n bla /n    kfjaskldfjask  lfsdjflkasfjlk   lfjaskldfjasldkfjklasdfjklasdj alsdkfjjlskjaslfkslkja"}
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

func compileProfile(user *User) *User {
	user.Html = mustache.RenderFile("Pages/DisplayProfile.mustache", user)
	return user
}

func renderFollow() string {
	follow := new(Follow)
	follow.FollowedBy = initDummys()[0:4]
	follow.Following = initDummys()[4:7]
	return compileFollow(follow).Html
}

func compileFollow(f *Follow) *Follow {
	f.Html = mustache.RenderFile("Pages/Follow.mustache", f)
	return f
}

func renderStream() string {
	items := []*StreamItem{
		dummyStreamItem(), dummyStreamItem(), dummyStreamItem()}
	str := new(Stream)
	str.Items = make([]*StreamItem, len(items))
	for i, v := range items {
		str.Items[i] = compileStreamItem(v)
	}

	html := mustache.RenderFileInLayout("Pages/Stream.mustache", "Pages/layout.mustache", str)
	return html
}

func compileStream(s *Stream) *Stream {
	for _, v := range s.Items {
		compileStreamItem(v)
	}
	s.Html = mustache.RenderFile("Pages/Stream.mustache", s)
	return s
}

func renderStreamItem() string {
	item := dummyStreamItem()
	html := mustache.RenderFileInLayout("Pages/StreamItem.mustache", "Pages/layout.mustache", item)
	return html
}

func compileStreamItem(si *StreamItem) *StreamItem {
	si.Html = mustache.RenderFile("Pages/StreamItem.mustache", si)
	return si
}

func compileHome(home *Home) *Home {
	compileProfile(home.CardRender)
	compileStream(home.StreamRender)
	compileFollow(home.FollowRender)
	return home
}

func renderPage() string {
	home := new(Home)
	home.CardRender = dummyUser()
	home.StreamRender = &Stream{Items: []*StreamItem{
		dummyStreamItem(), dummyStreamItem(), dummyStreamItem()}}
	home.FollowRender = new(Follow)
	home.FollowRender.FollowedBy = initDummys()[0:4]
	home.FollowRender.Following = initDummys()[4:7]

	compileHome(home)
	home.Html = mustache.RenderFileInLayout("Pages/Home.mustache", "Pages/layout.mustache", home)
	return home.Html
}

func renderCSS(val string) string {

	if css, err := ioutil.ReadFile("Pages/CSS/" + val); err != nil {
		return err.Error()
	} else {
		return string(css)
	}
	return ""

}

func renderIMG(val string) []byte {

	if img, err := ioutil.ReadFile("Pages/IMG/" + val); err != nil {
		return []byte(err.Error())
	} else {
		return img
	}
	return []byte("")
}

func renderJS(val string) string {

	if js, err := ioutil.ReadFile("Pages/JS/" + val); err != nil {
		return err.Error()
	} else {
		return string(js)
	}
	return ""

}

func main() {
	web.Get("/", renderPage)
	web.Get("/ProfileCard", renderProfile)
	web.Get("/CSS/(.*)", renderCSS)
	web.Get("/JS/(.*)", renderJS)
	web.Get("/Stream", renderStream)
	web.Get("/StreamItem", renderStreamItem)
	web.Get("/Follows", renderFollow)
	web.Run("0.0.0.0:9998")
}
