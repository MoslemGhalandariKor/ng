// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package checkbox

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func CheckboxListGroup() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"flex justify-center\"><div class=\" w-1/2 text border-4 p-5 border-white-600 place-content-center\"><h3 class=\"mb-4 font-semibold text-gray-900 dark:text-white\">Technology</h3><ul class=\"w-48 text-sm font-medium text-gray-900 bg-white border border-gray-200 rounded-lg dark:bg-gray-700 dark:border-gray-600 dark:text-white\"><li class=\"w-full border-b border-gray-200 rounded-t-lg dark:border-gray-600\"><div class=\"flex items-center ps-3\"><input id=\"vue-checkbox\" type=\"checkbox\" value=\"\" class=\"w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-700 dark:focus:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500\"> <label for=\"vue-checkbox\" class=\"w-full py-3 ms-2 text-sm font-medium text-gray-900 dark:text-gray-300\">Vue JS</label></div></li><li class=\"w-full border-b border-gray-200 rounded-t-lg dark:border-gray-600\"><div class=\"flex items-center ps-3\"><input id=\"react-checkbox\" type=\"checkbox\" value=\"\" class=\"w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-700 dark:focus:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500\"> <label for=\"react-checkbox\" class=\"w-full py-3 ms-2 text-sm font-medium text-gray-900 dark:text-gray-300\">React</label></div></li><li class=\"w-full border-b border-gray-200 rounded-t-lg dark:border-gray-600\"><div class=\"flex items-center ps-3\"><input id=\"angular-checkbox\" type=\"checkbox\" value=\"\" class=\"w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-700 dark:focus:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500\"> <label for=\"angular-checkbox\" class=\"w-full py-3 ms-2 text-sm font-medium text-gray-900 dark:text-gray-300\">Angular</label></div></li><li class=\"w-full border-b border-gray-200 rounded-t-lg dark:border-gray-600\"><div class=\"flex items-center ps-3\"><input id=\"laravel-checkbox\" type=\"checkbox\" value=\"\" class=\"w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-700 dark:focus:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500\"> <label for=\"laravel-checkbox\" class=\"w-full py-3 ms-2 text-sm font-medium text-gray-900 dark:text-gray-300\">Laravel</label></div></li></ul></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
