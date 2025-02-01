// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package profile

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "nextgen/templates/dashboard/dashboardcomponents"

type ProfilePageInfoProps struct {
	FirstName   string
	LastName    string
	Username    string
	Email       string
	PhoneNumber string
	About       string
	UserPhoto   string
	CoverPhoto  string
	City        string
	Province    string
	PostalCode  string
	Addres      string
	CompanyLogo string
	CompanyName string
	IsLoggedIn  bool
}

type ProfilePageHeaderProp struct {
	Url   string
	Label string
	Class string
}

type ProfilePageProps struct {
	ProfilePageHeaderProp []ProfilePageHeaderProp
	LayoutProp            dashboardcomponents.LayoutProp
	ProfilePageInfoProps  ProfilePageInfoProps
}

func PersonalProfile(profilePageProps ProfilePageProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = dashboardcomponents.Layout(PersonalProfilePageContent(profilePageProps.ProfilePageInfoProps), profilePageProps.LayoutProp).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ProfilePageHeader(profilePageProps.ProfilePageHeaderProp).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func PersonalProfilePageContent(user ProfilePageInfoProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"mt-2\"><div class=\" flex  border-2 border-gray-200 border-dashed rounded-lg dark:border-gray-700 justify-center \"><form class=\"p-12 m-4 w-full bg-gray-200 rounded-2xl\"><div class=\"space-y-12\"><div class=\"border-b border-gray-900/10 pb-12\"><h2 class=\"text-base font-semibold leading-7 text-gray-900\">PersonalProfile</h2><p class=\"mt-1 text-sm leading-6 text-gray-600\">This information will be displayed publicly so be careful what you share.</p><div class=\"mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6\"><div class=\"sm:col-span-2\"><label for=\"username\" class=\"block text-sm font-medium leading-6 text-gray-900\">Username</label><div class=\"mt-2\"><input type=\"text\" name=\"username\" id=\"username\" autocomplete=\"user-name\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(user.Username)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/profile/ProfilePage.templ`, Line: 52, Col: 110}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></div></div><div class=\"col-span-full\"><label for=\"about\" class=\"block text-sm font-medium leading-6 text-gray-900\">About</label><div class=\"mt-2\"><textarea id=\"about\" name=\"about\" rows=\"3\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(user.About)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/profile/ProfilePage.templ`, Line: 58, Col: 76}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></textarea></div><p class=\"mt-3 text-sm leading-6 text-gray-600\">Write a few sentences about yourself.</p></div><div class=\"col-span-full\"><label for=\"photo\" class=\"block text-sm font-medium leading-6 text-gray-900\">Photo</label><div class=\"mt-2 flex items-center gap-x-3\"><svg src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(user.UserPhoto)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/profile/ProfilePage.templ`, Line: 65, Col: 34}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"h-12 w-12 text-gray-300\" viewBox=\"0 0 24 24\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M18.685 19.097A9.723 9.723 0 0021.75 12c0-5.385-4.365-9.75-9.75-9.75S2.25 6.615 2.25 12a9.723 9.723 0 003.065 7.097A9.716 9.716 0 0012 21.75a9.716 9.716 0 006.685-2.653zm-12.54-1.285A7.486 7.486 0 0112 15a7.486 7.486 0 015.855 2.812A8.224 8.224 0 0112 20.25a8.224 8.224 0 01-5.855-2.438zM15.75 9a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z\" clip-rule=\"evenodd\"></path></svg> <button type=\"button\" class=\"rounded-md bg-white px-2.5 py-1.5 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50\">Change</button></div></div><div class=\"col-span-full\"><label for=\"cover-photo\" class=\"block text-sm font-medium leading-6 text-gray-900\">Cover photo</label><div class=\"mt-2 flex justify-center rounded-lg border border-dashed border-gray-900/25 px-6 py-10\"><div class=\"text-center\"><svg class=\"mx-auto h-12 w-12 text-gray-300\" viewBox=\"0 0 24 24\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.5 6a2.25 2.25 0 012.25-2.25h16.5A2.25 2.25 0 0122.5 6v12a2.25 2.25 0 01-2.25 2.25H3.75A2.25 2.25 0 011.5 18V6zM3 16.06V18c0 .414.336.75.75.75h16.5A.75.75 0 0021 18v-1.94l-2.69-2.689a1.5 1.5 0 00-2.12 0l-.88.879.97.97a.75.75 0 11-1.06 1.06l-5.16-5.159a1.5 1.5 0 00-2.12 0L3 16.061zm10.125-7.81a1.125 1.125 0 112.25 0 1.125 1.125 0 01-2.25 0z\" clip-rule=\"evenodd\"></path></svg><div class=\"mt-4 flex text-sm leading-6 text-gray-600\"><label for=\"file-upload\" class=\"relative cursor-pointer rounded-md bg-white font-semibold text-indigo-600 focus-within:outline-none focus-within:ring-2 focus-within:ring-indigo-600 focus-within:ring-offset-2 hover:text-indigo-500\"><span>Upload a file</span> <input id=\"file-upload\" name=\"file-upload\" type=\"file\" class=\"sr-only\"></label><p class=\"pl-1\">or drag and drop</p></div><p class=\"text-xs leading-5 text-gray-600\">PNG, JPG, GIF up to 10MB</p></div></div></div></div></div><div class=\"border-b border-gray-900/10 pb-12\"><h2 class=\"text-base font-semibold leading-7 text-gray-900\">Personal Information</h2><p class=\"mt-1 text-sm leading-6 text-gray-600\">Use a permanent address where you can receive mail.</p><div class=\"mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6\"><div class=\"sm:col-span-3\"><label for=\"first-name\" class=\"block text-sm font-medium leading-6 text-gray-900\">First name</label><div class=\"mt-2\"><input type=\"text\" name=\"first-name\" id=\"first-name\" autocomplete=\"given-name\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(user.FirstName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/profile/ProfilePage.templ`, Line: 98, Col: 116}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></div></div><div class=\"sm:col-span-3\"><label for=\"last-name\" class=\"block text-sm font-medium leading-6 text-gray-900\">Last name</label><div class=\"mt-2\"><input type=\"text\" name=\"last-name\" id=\"last-name\" autocomplete=\"family-name\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(user.LastName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/profile/ProfilePage.templ`, Line: 104, Col: 114}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></div></div><div class=\"sm:col-span-4\"><label for=\"email\" class=\"block text-sm font-medium leading-6 text-gray-900\">Email address</label><div class=\"mt-2\"><input id=\"email\" name=\"email\" type=\"email\" autocomplete=\"email\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(user.Email)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/profile/ProfilePage.templ`, Line: 110, Col: 98}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></div></div><div class=\"sm:col-span-2 sm:col-start-1\"><label for=\"city\" class=\"block text-sm font-medium leading-6 text-gray-900\">City</label><div class=\"mt-2\"><input type=\"text\" name=\"city\" id=\"city\" autocomplete=\"address-level2\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(user.City)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/profile/ProfilePage.templ`, Line: 116, Col: 103}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></div></div><div class=\"sm:col-span-2\"><label for=\"region\" class=\"block text-sm font-medium leading-6 text-gray-900\">State / Province</label><div class=\"mt-2\"><input type=\"text\" name=\"region\" id=\"region\" autocomplete=\"address-level1\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(user.Province)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/profile/ProfilePage.templ`, Line: 122, Col: 111}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></div></div><div class=\"sm:col-span-2\"><label for=\"postal-code\" class=\"block text-sm font-medium leading-6 text-gray-900\">ZIP / Postal code</label><div class=\"mt-2\"><input type=\"text\" name=\"postal-code\" id=\"postal-code\" autocomplete=\"postal-code\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(user.PostalCode)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/profile/ProfilePage.templ`, Line: 128, Col: 120}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></div></div><div class=\"col-span-full\"><label for=\"street-address\" class=\"block text-sm font-medium leading-6 text-gray-900\">Street address</label><div class=\"mt-2\"><input type=\"text\" name=\"street-address\" id=\"street-address\" autocomplete=\"street-address\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(user.Addres)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/profile/ProfilePage.templ`, Line: 134, Col: 125}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></div></div></div></div><div class=\"border-b border-gray-900/10 pb-12\"><h2 class=\"text-base font-semibold leading-7 text-gray-900\">Notifications</h2><p class=\"mt-1 text-sm leading-6 text-gray-600\">We'll always let you know about important changes, but you pick what else you want to hear about.</p><div class=\"mt-10 space-y-10\"><fieldset><legend class=\"text-sm font-semibold leading-6 text-gray-900\">By Email</legend><div class=\"mt-6 space-y-6\"><div class=\"relative flex gap-x-3\"><div class=\"flex h-6 items-center\"><input id=\"comments\" name=\"comments\" type=\"checkbox\" class=\"h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600\"></div><div class=\"text-sm leading-6\"><label for=\"comments\" class=\"font-medium text-gray-900\">Comments</label><p class=\"text-gray-500\">Get notified when someones posts a comment on a posting.</p></div></div><div class=\"relative flex gap-x-3\"><div class=\"flex h-6 items-center\"><input id=\"candidates\" name=\"candidates\" type=\"checkbox\" class=\"h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600\"></div><div class=\"text-sm leading-6\"><label for=\"candidates\" class=\"font-medium text-gray-900\">Candidates</label><p class=\"text-gray-500\">Get notified when a candidate applies for a job.</p></div></div><div class=\"relative flex gap-x-3\"><div class=\"flex h-6 items-center\"><input id=\"offers\" name=\"offers\" type=\"checkbox\" class=\"h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600\"></div><div class=\"text-sm leading-6\"><label for=\"offers\" class=\"font-medium text-gray-900\">Offers</label><p class=\"text-gray-500\">Get notified when a candidate accepts or rejects an offer.</p></div></div></div></fieldset><fieldset><legend class=\"text-sm font-semibold leading-6 text-gray-900\">Push Notifications</legend><p class=\"mt-1 text-sm leading-6 text-gray-600\">These are delivered via SMS to your mobile phone.</p><div class=\"mt-6 space-y-6\"><div class=\"flex items-center gap-x-3\"><input id=\"push-everything\" name=\"push-notifications\" type=\"radio\" class=\"h-4 w-4 border-gray-300 text-indigo-600 focus:ring-indigo-600\"> <label for=\"push-everything\" class=\"block text-sm font-medium leading-6 text-gray-900\">Everything</label></div><div class=\"flex items-center gap-x-3\"><input id=\"push-email\" name=\"push-notifications\" type=\"radio\" class=\"h-4 w-4 border-gray-300 text-indigo-600 focus:ring-indigo-600\"> <label for=\"push-email\" class=\"block text-sm font-medium leading-6 text-gray-900\">Same as email</label></div><div class=\"flex items-center gap-x-3\"><input id=\"push-nothing\" name=\"push-notifications\" type=\"radio\" class=\"h-4 w-4 border-gray-300 text-indigo-600 focus:ring-indigo-600\"> <label for=\"push-nothing\" class=\"block text-sm font-medium leading-6 text-gray-900\">No push notifications</label></div></div></fieldset></div></div></div><div class=\"mt-6 flex items-center justify-end gap-x-6\"><button type=\"button\" class=\"text-sm font-semibold leading-6 text-gray-900\">Cancel</button> <button type=\"submit\" class=\"rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600\">Save</button></div></form></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func CompanyProfile(profilePageProps ProfilePageProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var13 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var13 == nil {
			templ_7745c5c3_Var13 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = dashboardcomponents.Layout(CompanyProfilePageContent(profilePageProps.ProfilePageInfoProps), profilePageProps.LayoutProp).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ProfilePageHeader(profilePageProps.ProfilePageHeaderProp).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func CompanyProfilePageContent(user ProfilePageInfoProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var14 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var14 == nil {
			templ_7745c5c3_Var14 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"mt-2\"><div class=\" flex  border-2 border-gray-200 border-dashed rounded-lg dark:border-gray-700 justify-center \"><form class=\"p-12 m-4 w-full bg-gray-200 rounded-2xl\"><div class=\"space-y-12\"><div class=\"border-b border-gray-900/10 pb-12\"><h2 class=\"text-base font-semibold leading-7 text-gray-900\">CompanyProfile</h2><p class=\"mt-1 text-sm leading-6 text-gray-600\">This information will be displayed publicly so be careful what you share.</p><div class=\" mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6\"><!-- name and type --><div class=\"sm:col-span-4 flex justify-between \"><div><label for=\"CompanyName\" class=\"block text-sm font-medium leading-6 text-gray-900\">Name</label><div class=\"mt-2\"><input type=\"text\" name=\"CompanyName\" id=\"CompanyName\" autocomplete=\"Company-Name\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></div></div><div><label for=\"username\" class=\"block text-sm font-medium leading-6 text-gray-900\">CompanyType</label><div class=\"mt-2\"><div class=\"flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 focus-within:ring-2 focus-within:ring-inset focus-within:ring-indigo-600 sm:max-w-md\"><select type=\"text\" name=\"CompanyType\" id=\"CompanyType\" autocomplete=\"CompanyType\" class=\"w-64 block flex-1 border-0 rounded-lg py-1.5 pl-1 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6\" placeholder=\"CompanyName\"><option selected>Choose a Type</option> <option value=\"M\">Market</option> <option value=\"LM\">LocaalMarket</option> <option value=\"HM\">HyperMarket</option></select></div></div></div></div><!-- about --><div class=\"col-span-full\"><label for=\"about\" class=\"block text-sm font-medium leading-6 text-gray-900\">About</label><div class=\"mt-2\"><textarea id=\"about\" name=\"about\" rows=\"3\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></textarea></div><p class=\"mt-3 text-sm leading-6 text-gray-600\">Write a few sentences about your job.</p></div><!-- Image --><div class=\"col-span-full\"><label for=\"photo\" class=\"block text-sm font-medium leading-6 text-gray-900\">Logo</label><div class=\"mt-2 flex items-center gap-x-3\"><svg class=\"h-12 w-12 text-gray-300\" viewBox=\"0 0 24 24\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M18.685 19.097A9.723 9.723 0 0021.75 12c0-5.385-4.365-9.75-9.75-9.75S2.25 6.615 2.25 12a9.723 9.723 0 003.065 7.097A9.716 9.716 0 0012 21.75a9.716 9.716 0 006.685-2.653zm-12.54-1.285A7.486 7.486 0 0112 15a7.486 7.486 0 015.855 2.812A8.224 8.224 0 0112 20.25a8.224 8.224 0 01-5.855-2.438zM15.75 9a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z\" clip-rule=\"evenodd\"></path></svg> <button type=\"button\" class=\"rounded-md bg-white px-2.5 py-1.5 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50\">Change</button></div></div><div class=\"col-span-full\"><label for=\"cover-photo\" class=\"block text-sm font-medium leading-6 text-gray-900\">Cover photo</label><div class=\"mt-2 flex justify-center rounded-lg border border-dashed border-gray-900/25 px-6 py-10\"><div class=\"text-center\"><svg class=\"mx-auto h-12 w-12 text-gray-300\" viewBox=\"0 0 24 24\" fill=\"currentColor\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.5 6a2.25 2.25 0 012.25-2.25h16.5A2.25 2.25 0 0122.5 6v12a2.25 2.25 0 01-2.25 2.25H3.75A2.25 2.25 0 011.5 18V6zM3 16.06V18c0 .414.336.75.75.75h16.5A.75.75 0 0021 18v-1.94l-2.69-2.689a1.5 1.5 0 00-2.12 0l-.88.879.97.97a.75.75 0 11-1.06 1.06l-5.16-5.159a1.5 1.5 0 00-2.12 0L3 16.061zm10.125-7.81a1.125 1.125 0 112.25 0 1.125 1.125 0 01-2.25 0z\" clip-rule=\"evenodd\"></path></svg><div class=\"mt-4 flex text-sm leading-6 text-gray-600\"><label for=\"file-upload\" class=\"relative cursor-pointer rounded-md bg-white font-semibold text-indigo-600 focus-within:outline-none focus-within:ring-2 focus-within:ring-indigo-600 focus-within:ring-offset-2 hover:text-indigo-500\"><span>Upload a file</span> <input id=\"file-upload\" name=\"file-upload\" type=\"file\" class=\"sr-only\"></label><p class=\"pl-1\">or drag and drop</p></div><p class=\"text-xs leading-5 text-gray-600\">PNG, JPG, GIF up to 10MB</p></div></div></div></div></div><!-- cantact information --><div class=\"border-b border-gray-900/10 pb-12\"><h2 class=\"text-base font-semibold leading-7 text-gray-900\">Company Information</h2><p class=\"mt-1 text-sm leading-6 text-gray-600\"></p><div class=\"mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6\"><div class=\"sm:col-span-3\"><label for=\"PhoneNumber\" class=\"block text-sm font-medium leading-6 text-gray-900\">Phone Number</label><div class=\"mt-2\"><input type=\"text\" name=\"PhoneNumber\" id=\"PhoneNumber\" autocomplete=\"PhoneNumber\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></div></div><div class=\"sm:col-span-3\"><label for=\"email\" class=\"block text-sm font-medium leading-6 text-gray-900\">Email address</label><div class=\"mt-2\"><input id=\"email\" name=\"email\" type=\"email\" autocomplete=\"email\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></div></div><div class=\"sm:col-span-2 sm:col-start-1\"><label for=\"city\" class=\"block text-sm font-medium leading-6 text-gray-900\">City</label><div class=\"mt-2\"><input type=\"text\" name=\"city\" id=\"city\" autocomplete=\"address-level2\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></div></div><div class=\"sm:col-span-2\"><label for=\"region\" class=\"block text-sm font-medium leading-6 text-gray-900\">State / Province</label><div class=\"mt-2\"><input type=\"text\" name=\"region\" id=\"region\" autocomplete=\"address-level1\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></div></div><div class=\"sm:col-span-2\"><label for=\"postal-code\" class=\"block text-sm font-medium leading-6 text-gray-900\">ZIP / Postal code</label><div class=\"mt-2\"><input type=\"text\" name=\"postal-code\" id=\"postal-code\" autocomplete=\"postal-code\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></div></div><div class=\"col-span-full\"><label for=\"street-address\" class=\"block text-sm font-medium leading-6 text-gray-900\">Street address</label><div class=\"mt-2\"><input type=\"text\" name=\"street-address\" id=\"street-address\" autocomplete=\"street-address\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></div></div></div></div><div class=\"border-b border-gray-900/10 pb-12\"><h2 class=\"text-base font-semibold leading-7 text-gray-900\">Notifications</h2><p class=\"mt-1 mb-4 text-sm leading-6 text-gray-600\">We'll always let you know about important changes, but you choose who get notifications .</p><div class=\"relative overflow-x-auto shadow-md sm:rounded-lg\"><table class=\"w-full text-sm text-left rtl:text-right text-gray-500\"><thead class=\"text-xs text-gray-700 uppercase bg-gray-300 \"><tr><th scope=\"col\" class=\"px-6 py-3\">Name</th><th scope=\"col\" class=\"px-6 py-3\">Role</th><th scope=\"col\" class=\"px-6 py-3\">emailaddres</th><th scope=\"col\" class=\"px-6 py-3\">PhoneNumber</th><th scope=\"col\" class=\"px-6 py-3\">Action</th></tr></thead> <tbody><tr class=\"odd:bg-gray-100 even:bg-gray-200 border-b dark:border-gray-700\"><th scope=\"row\" class=\"px-6 py-4 font-medium text-gray-900 whitespace-nowrap \"></th><td class=\"px-6 py-4\"></td><td class=\"px-6 py-4\"></td><td class=\"px-6 py-4\"></td><td class=\"px-6 py-4\"><a href=\"#\" class=\"font-medium text-blue-600 dark:text-blue-500 hover:underline\">Edit</a></td></tr><tr class=\"odd:bg-gray-100 even:bg-gray-200 border-b dark:border-gray-700\"><th scope=\"row\" class=\"px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white\"></th><td class=\"px-6 py-4\"></td><td class=\"px-6 py-4\"></td><td class=\"px-6 py-4\"></td><td class=\"px-6 py-4\"><a href=\"#\" class=\"font-medium text-blue-600 dark:text-blue-500 hover:underline\">Edit</a></td></tr></tbody></table></div></div></div><div class=\"mt-6 flex items-center justify-end gap-x-6\"><button type=\"button\" class=\"text-sm font-semibold leading-6 text-gray-900\">Cancel</button> <button type=\"submit\" class=\"rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600\">Save</button></div></form></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ProfilePageHeader(props []ProfilePageHeaderProp) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var15 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var15 == nil {
			templ_7745c5c3_Var15 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"text-sm font-medium text-center text-gray-500 border-b border-gray-200 dark:text-gray-400 dark:border-gray-700\"><ul class=\"flex flex-wrap -mb-px\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, prop := range props {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"me-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var16 = []any{prop.Class}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var16...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var17 templ.SafeURL = templ.URL(prop.Url)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var17)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var18 string
			templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var16).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/profile/ProfilePage.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var19 string
			templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs(prop.Label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/profile/ProfilePage.templ`, Line: 385, Col: 70}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
