package main

import "github.com/hoisie/web"
import "github.com/hoisie/mustache"
import "io/ioutil"

import "github.com/HackerSchool12/SpiderDB"
import "github.com/HackerSchool12/SpiderDB/socialGraph"

//TODO: refactor all structs out to own files and create render and compile methods for each
//Render methods are for getting a page with just that items content

var gm *spiderDB.GraphManager

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
	Items  []*StreamItem
	UserID string
	Html   string
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

func initDummys() {
	gm = new(spiderDB.GraphManager)
	gm.Initialize()

	const numdum = 8
	pics := [numdum]string{"http://content8.flixster.com/question/28/64/25/2864258_std.jpg",
		"http://4.bp.blogspot.com/-Q2hjS1dS1R8/T4YXpOfNjOI/AAAAAAAAAxQ/c-V_1FkMYmo/s1600/Bug.jpg",
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
		newNode := socialGraph.NewSocialNode(pics[k],
			names[k], users[k], "email@somewhere.com",
			bios[k], "skills", github[k], gm)
		gm.AddNode(newNode)
	}

	edgF := socialGraph.NewSocialEdge(1, "follows", gm)
	edgFB := socialGraph.NewSocialEdge(1, "follower", gm)

	gm.AddEdge(edgF)
	gm.AddEdge(edgFB)

	var err error

	node0, err := gm.GetNode("1", socialGraph.SocialNodeConst)
	if err != nil {
		print(err.Error())
	}
	node1, err := gm.GetNode("2", socialGraph.SocialNodeConst)
	if err != nil {
		print(err.Error())
	}
	node2, err := gm.GetNode("3", socialGraph.SocialNodeConst)
	if err != nil {
		print(err.Error())
	}

	gm.Attach(node0, node1, edgF)
	gm.Attach(node0, node2, edgFB)

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

func AddJitter(userID string, jit string) {
	currUser, _ := gm.GetNode(userID, socialGraph.SocialNodeConst)
	jitter := socialGraph.NewMessageNode(jit)
	edge := socialGraph.NewSocialEdge(8182012, "jittered", gm)
	gm.Attach(currUser, jitter, edge)

}

//node -> social node -> user???? There's got to be a better way.
// FIX ME
func FetchUserInfo(userID string) *User {
	n, err := gm.GetNode(userID, socialGraph.SocialNodeConst)

	if err != nil {
		return nil
	}

	sn, ok := n.(*socialGraph.SocialNode)
	if !ok {
		return nil
	}

	return SocialNodeToUser(sn)
}

func renderProfile() string {
	html := mustache.RenderFileInLayout("Pages/DisplayProfile.mustache", "Pages/layout.mustache", dummyUser())
	return html
}

func compileProfile(user *User) *User {
	user.Html = mustache.RenderFile("Pages/DisplayProfile.mustache", user)
	return user
}

func GetFollow(userID string) ([]*User, []*User, error) {
	following := make([]*User, 0)
	followedBy := make([]*User, 0)

	node, err1 := gm.GetNode(userID, socialGraph.SocialNodeConst)
	nbr, err2 := gm.GetNeighbors(node, socialGraph.SocialEdgeConst, socialGraph.SocialNodeConst)

	if err1 != nil {
		return nil, nil, err1
	}

	if err2 != nil {
		return nil, nil, err2
	}

	for _, v := range nbr {

		sn, ok := v.NodeB.(*socialGraph.SocialNode)
		if !ok {
			return nil, nil, &hiveError{"Cannot cast to SocialNode"}
		}
		usr := SocialNodeToUser(sn)

		if v.Edg.GetType() == "follows" {
			following = append(following, usr)
		}
		if v.Edg.GetType() == "follower" {
			followedBy = append(followedBy, usr)
		}
	}

	return following, followedBy, nil
}

func SocialNodeToUser(sn *socialGraph.SocialNode) *User {
	usr := &User{}

	usr.Pic = sn.GetPic()
	usr.ProperName = sn.GetProperName()
	usr.UserName = sn.GetUserName()
	usr.Email = sn.GetEmail()
	usr.Bio = sn.GetBio()
	usr.Skills = sn.GetSkills()
	usr.Github = sn.GetGit()

	return usr
}

func renderFollow(userID string) string {
	follow := new(Follow)
	followedBy, following, err := GetFollow(userID)

	if err != nil {
		return "Failed to get Follow Lists"
	}

	follow.FollowedBy = followedBy
	follow.Following = following

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

func addStreamItem(ctx *web.Context) string {
	JIT := ctx.Params["JIT"]
	userID := ctx.Params["userID"]
	userNode, err := gm.GetNode(userID, socialGraph.SocialNodeConst)
	if err != nil {
		panic(err.Error())
	}
	messageNode := &socialGraph.MessageNode{}
	messageNode.SetText(JIT)
	jitEdge := socialGraph.NewSocialEdge(1, "jitted", gm)
	gm.AddNode(messageNode)
	gm.AddEdge(jitEdge)
	gm.Attach(userNode, messageNode, jitEdge)
	return renderPage(userID)
}

func compileHome(home *Home) *Home {
	compileProfile(home.CardRender)
	compileStream(home.StreamRender)
	compileFollow(home.FollowRender)
	return home
}

func compileSplash() {

}

func renderSplash() string {
	html := mustache.RenderFileInLayout("Pages/Splash.mustache", "Pages/layout.mustache")
	return html
}

func renderPage(user string) string {
	home := new(Home)

	followedBy, following, err := GetFollow(user)

	if err != nil {
		print(err.Error())
	}
	home.CardRender = FetchUserInfo(user)
	home.StreamRender = &Stream{UserID: user, Items: []*StreamItem{
		dummyStreamItem(), dummyStreamItem(), dummyStreamItem()}}
	home.FollowRender = new(Follow)
	home.FollowRender.FollowedBy = followedBy
	home.FollowRender.Following = following

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
	initDummys()
	web.Get("/Home&userID=(.*)", renderPage)
	web.Get("/Splash", renderSplash)
	web.Get("/ProfileCard", renderProfile)
	web.Get("/CSS/(.*)", renderCSS)
	web.Get("/JS/(.*)", renderJS)
	web.Get("/Stream", renderStream)
	web.Get("/StreamItem", renderStreamItem)
	web.Post("/StreamItem", addStreamItem)
	web.Get("/Follows", renderFollow)
	web.Run("0.0.0.0:9998")
}
