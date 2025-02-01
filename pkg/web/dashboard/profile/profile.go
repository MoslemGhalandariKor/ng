package profile

import (
	"fmt"
	"net/http"
	"nextgen/internals/gintemplrenderer"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages/profile"
	"sort"

	"github.com/gin-gonic/gin"
)

var headerElementsClass = "inline-block p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
var currentHeaderElementClass = "inline-block p-4 text-blue-600 border-b-2 border-blue-600 rounded-t-lg active dark:text-gray-200 dark:border-gray-200"

var ProfilePageHeaderProps = map[string]*profile.ProfilePageHeaderProp{
	"PersonalProfile": &profile.ProfilePageHeaderProp{Label: "Personal Profile",
		Url:   "/dashboard/personal-profile",
		Class: ""},
	"CompanyProfile": &profile.ProfilePageHeaderProp{Label: "Company Profile",
		Url:   "/dashboard/company-profile",
		Class: ""},
}

func PersonalProfilePage(c *gin.Context) {

	profilePageProps := profile.ProfilePageProps{}

	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}
	profilePageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	profilePageHeaderProp := make([]profile.ProfilePageHeaderProp, 0)

	for key, value := range ProfilePageHeaderProps {
		if key == "PersonalProfile" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		profilePageHeaderProp = append(profilePageHeaderProp, *value)
	}
	sort.Slice(profilePageHeaderProp, func(i, j int) bool {
		return profilePageHeaderProp[i].Url > profilePageHeaderProp[j].Url
	})

	profilePageProps.ProfilePageHeaderProp = profilePageHeaderProp

	users, err := GetUser()
	if err != nil {
		fmt.Println(err)
	}

	userInfo := []profile.ProfilePageInfoProps{}

	for _, user := range *users {
		userInfo = append(userInfo, profile.ProfilePageInfoProps{
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Username:    user.Username,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			About:       user.About,
			UserPhoto:   user.UserPhoto,
			CoverPhoto:  user.CoverPhoto,
			City:        user.City,
			Province:    user.Province,
			PostalCode:  user.PostalCode,
			Addres:      user.Addres,
			CompanyLogo: user.CompanyLogo,
			CompanyName: user.CompanyName,
		})
	}

	profilePageProps.ProfilePageInfoProps = userInfo[0]

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, profile.PersonalProfile(profilePageProps))
	c.Render(http.StatusOK, r)
}

func CompanyProfilePage(c *gin.Context) {
	
	profilePageProps := profile.ProfilePageProps{}

	layoutProp, exists := c.Get("LayoutProp")

	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}
	profilePageProps.LayoutProp = layoutProp.(dashboardcomponents.LayoutProp)

	profilePageHeaderProp := make([]profile.ProfilePageHeaderProp, 0)

	for key, value := range ProfilePageHeaderProps {
		if key == "CompanyProfile" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		profilePageHeaderProp = append(profilePageHeaderProp, *value)
	}
	sort.Slice(profilePageHeaderProp, func(i, j int) bool {
		return profilePageHeaderProp[i].Url > profilePageHeaderProp[j].Url
	})

	profilePageProps.ProfilePageHeaderProp = profilePageHeaderProp
	

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, profile.CompanyProfile(profilePageProps))
	c.Render(http.StatusOK, r)

}
