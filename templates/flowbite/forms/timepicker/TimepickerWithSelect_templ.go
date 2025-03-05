// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package timepicker

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func TimepickerWithSelect() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex justify-center\"><div class=\" w-1/2 text border-4 p-5 border-white-600 place-content-center\"><form><label for=\"time\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Select time:</label><div class=\"flex\"><input type=\"time\" id=\"time\" class=\"flex-shrink-0 rounded-none rounded-s-lg bg-gray-50 border text-gray-900 leading-none focus:ring-blue-500 focus:border-blue-500 block text-sm border-gray-300 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\" min=\"09:00\" max=\"18:00\" value=\"00:00\" required> <select id=\"timezones\" name=\"timezone\" class=\"bg-gray-50 border border-s-0 border-gray-300 text-gray-900 text-sm rounded-e-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\" required><option value=\"America/New_York\" selected>EST - GMT-5 (New York)</option> <option value=\"America/Los_Angeles\">PST - GMT-8 (Los Angeles)</option> <option value=\"Europe/London\">GMT - GMT+0 (London)</option> <option value=\"Europe/Paris\">CET - GMT+1 (Paris)</option> <option value=\"Asia/Tokyo\">JST - GMT+9 (Tokyo)</option> <option value=\"Australia/Sydney\">AEDT - GMT+11 (Sydney)</option> <option value=\"Canada/Mountain\">MST - GMT-7 (Canada)</option> <option value=\"Canada/Central\">CST - GMT-6 (Canada)</option> <option value=\"Canada/Eastern\">EST - GMT-5 (Canada)</option> <option value=\"Europe/Berlin\">CET - GMT+1 (Berlin)</option> <option value=\"Asia/Dubai\">GST - GMT+4 (Dubai)</option> <option value=\"Asia/Singapore\">SGT - GMT+8 (Singapore)</option></select></div></form></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
