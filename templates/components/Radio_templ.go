// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

type RadioListWithDescProps struct {
	Title       string
	Discription string
}

func RadioListWithDesc(props RadioListWithDescProps) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex justify-center bg-white px-4 py-12\"><fieldset><legend class=\"sr-only\">Plan</legend><div class=\"space-y-5\"><div class=\"relative flex items-start\"><div class=\"flex h-6 items-center\"><input id=\"small\" aria-describedby=\"small-description\" name=\"plan\" type=\"radio\" checked=\"\" class=\"h-4 w-4 border-gray-300 text-indigo-600 focus:ring-indigo-600\"></div><div class=\"ml-3 text-sm leading-6\"><label for=\"small\" class=\"font-medium text-gray-900\">Small</label><p id=\"small-description\" class=\"text-gray-500\">4 GB RAM / 2 CPUS / 80 GB SSD Storage</p></div></div><div class=\"relative flex items-start\"><div class=\"flex h-6 items-center\"><input id=\"medium\" aria-describedby=\"medium-description\" name=\"plan\" type=\"radio\" class=\"h-4 w-4 border-gray-300 text-indigo-600 focus:ring-indigo-600\"></div><div class=\"ml-3 text-sm leading-6\"><label for=\"medium\" class=\"font-medium text-gray-900\">Medium</label><p id=\"medium-description\" class=\"text-gray-500\">8 GB RAM / 4 CPUS / 160 GB SSD Storage</p></div></div><div class=\"relative flex items-start\"><div class=\"flex h-6 items-center\"><input id=\"large\" aria-describedby=\"large-description\" name=\"plan\" type=\"radio\" class=\"h-4 w-4 border-gray-300 text-indigo-600 focus:ring-indigo-600\"></div><div class=\"ml-3 text-sm leading-6\"><label for=\"large\" class=\"font-medium text-gray-900\">Large</label><p id=\"large-description\" class=\"text-gray-500\">16 GB RAM / 8 CPUS / 320 GB SSD Storage</p></div></div></div></fieldset></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
