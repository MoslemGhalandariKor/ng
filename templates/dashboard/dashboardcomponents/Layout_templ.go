// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package dashboardcomponents

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "nextgen/templates/components"

type UserInfoProps struct {
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
	CompanyLogo string
	CompanyName string
	IsLoggedIn  bool
}

type LayoutProp struct {
	UserInfoProps UserInfoProps
	Alerts        []components.AlertProps
}

func Layout(contents templ.Component, layoutProp LayoutProp) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" class=\"h-full bg-white-50\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title></title><!-- Tailwind CSS --><script src=\"https://cdn.tailwindcss.com\"> </script><link rel=\"stylesheet\" href=\"../../../static/css/output.css\"><!-- Flowbite --><link href=\"https://cdn.jsdelivr.net/npm/flowbite@2.5.2/dist/flowbite.min.css\" rel=\"stylesheet\"><!-- HTMX --></head><body class=\"h-full dark:bg-gray-800\"><div><div class=\"relative z-40\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = SidebarNew(contents, layoutProp.UserInfoProps).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"absolute z-50 top-0 left-72  justify-between \">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Alert(layoutProp.Alerts).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div><script src=\"../../../static/js/flowbite/flowbite.min.js\"></script><script src=\"https://cdn.jsdelivr.net/npm/flowbite@2.5.2/dist/flowbite.min.js\"></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/flowbite/2.3.0/datepicker.min.js\"></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/flowbite/1.6.3/flowbite.min.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
