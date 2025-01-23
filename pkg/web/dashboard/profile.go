package dashboard

import (

	"nextgen/templates/dashboard/pages"
	"sort"
	"github.com/gin-gonic/gin"
	"nextgen/templates/dashboard/dashboardcomponents"
)

var headerElementsClass = "inline-block p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300"
var currentHeaderElementClass = "inline-block p-4 text-blue-600 border-b-2 border-blue-600 rounded-t-lg active dark:text-gray-200 dark:border-gray-200"
var userInfoProps dashboardcomponents.UserInfoProps

var ProfilePageHeaderProps = map[string]*pages.ProfilePageHeaderProp{
	"PersonalProfile": &pages.ProfilePageHeaderProp{Label: "Personal Profile",
		Url:   "/dashboard/personal-profile",
		Class: ""},
	"CompanyProfile": &pages.ProfilePageHeaderProp{Label: "Company Profile",
		Url:   "/dashboard/company-profile",
		Class: ""},
}

func PersonalProfilePage(c *gin.Context) {
	ProfilePageHeaderProp := make([]pages.ProfilePageHeaderProp, 0)

	for key, value := range ProfilePageHeaderProps {
		if key == "PersonalProfile" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		ProfilePageHeaderProp = append(ProfilePageHeaderProp, *value)
	}
	sort.Slice(ProfilePageHeaderProp, func(i, j int) bool {
		return ProfilePageHeaderProp[i].Url > ProfilePageHeaderProp[j].Url
	})
	props := pages.ProfilePageProps{}
	layoutProp, exists := c.Get("LayoutProp")
	
	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	props.ProfileHeaderProps = ProfilePageHeaderProp
	pages.PersonalProfile(layoutProp.(dashboardcomponents.LayoutProp), props).Render(c.Request.Context(), c.Writer)

}

func CompanyProfilePage(c *gin.Context) {
	ProfilePageHeaderProp := make([]pages.ProfilePageHeaderProp, 0)

	for key, value := range ProfilePageHeaderProps {
		if key == "CompanyProfile" {
			value.Class = currentHeaderElementClass
		} else {
			value.Class = headerElementsClass
		}
		ProfilePageHeaderProp = append(ProfilePageHeaderProp, *value)
	}
	sort.Slice(ProfilePageHeaderProp, func(i, j int) bool {
		return ProfilePageHeaderProp[i].Url > ProfilePageHeaderProp[j].Url
	})
	
	props := pages.ProfilePageProps{}
	layoutProp, exists := c.Get("LayoutProp")
	
	if !exists {
		layoutProp = dashboardcomponents.LayoutProp{}
	}

	props.ProfileHeaderProps = ProfilePageHeaderProp
	pages.CompanyProfile(layoutProp.(dashboardcomponents.LayoutProp), props).Render(c.Request.Context(), c.Writer)

}
