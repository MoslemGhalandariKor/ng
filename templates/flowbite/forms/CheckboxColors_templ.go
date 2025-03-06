// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package forms

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func CheckboxColors() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"flex justify-center\"><div class=\" w-1/2 text border-4 p-5 border-white-600 place-content-center\"><div class=\"flex items-center me-4\"><input checked id=\"red-checkbox\" type=\"checkbox\" value=\"\" class=\"w-4 h-4 text-red-600 bg-gray-100 border-gray-300 rounded focus:ring-red-500 dark:focus:ring-red-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600\"> <label for=\"red-checkbox\" class=\"ms-2 text-sm font-medium text-gray-900 dark:text-gray-300\">Red</label></div><div class=\"flex items-center me-4\"><input checked id=\"green-checkbox\" type=\"checkbox\" value=\"\" class=\"w-4 h-4 text-green-600 bg-gray-100 border-gray-300 rounded focus:ring-green-500 dark:focus:ring-green-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600\"> <label for=\"green-checkbox\" class=\"ms-2 text-sm font-medium text-gray-900 dark:text-gray-300\">Green</label></div><div class=\"flex items-center me-4\"><input checked id=\"purple-checkbox\" type=\"checkbox\" value=\"\" class=\"w-4 h-4 text-purple-600 bg-gray-100 border-gray-300 rounded focus:ring-purple-500 dark:focus:ring-purple-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600\"> <label for=\"purple-checkbox\" class=\"ms-2 text-sm font-medium text-gray-900 dark:text-gray-300\">Purple</label></div><div class=\"flex items-center me-4\"><input checked id=\"teal-checkbox\" type=\"checkbox\" value=\"\" class=\"w-4 h-4 text-teal-600 bg-gray-100 border-gray-300 rounded focus:ring-teal-500 dark:focus:ring-teal-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600\"> <label for=\"teal-checkbox\" class=\"ms-2 text-sm font-medium text-gray-900 dark:text-gray-300\">Teal</label></div><div class=\"flex items-center me-4\"><input checked id=\"yellow-checkbox\" type=\"checkbox\" value=\"\" class=\"w-4 h-4 text-yellow-400 bg-gray-100 border-gray-300 rounded focus:ring-yellow-500 dark:focus:ring-yellow-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600\"> <label for=\"yellow-checkbox\" class=\"ms-2 text-sm font-medium text-gray-900 dark:text-gray-300\">Yellow</label></div><div class=\"flex items-center me-4\"><input checked id=\"orange-checkbox\" type=\"checkbox\" value=\"\" class=\"w-4 h-4 text-orange-500 bg-gray-100 border-gray-300 rounded focus:ring-orange-500 dark:focus:ring-orange-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600\"> <label for=\"orange-checkbox\" class=\"ms-2 text-sm font-medium text-gray-900 dark:text-gray-300\">Orange</label></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
