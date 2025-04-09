// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package productcategorycomponents

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "nextgen/pkg/product_management"

func FashionAndClothingPage(categoryInfo []product_management.CategoryView) templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"mt-2\"><div class=\" flex  border-2 border-gray-200 border-dashed rounded-lg justify-center \"><div class=\" m-2 w-full bg-gray-200 rounded-2xl\"><div class=\"space-y-12 p-4\"><div class=\"p-4 w-full h-full md:h-auto bg-gray-200 rounded-2xl p-4\"><!-- Modal content --><div class=\"p-4 shadow bg-white relative sm:rounded-lg overflow-hidden ring-1 ring-black ring-opacity-5 rounded-2xl\"><!-- Modal header --><!-- Modal body --><div class=\" gap-4 mb-4 sm:grid-cols-2 space-y-4\"><div class=\"grid gap-4 sm:col-span-2 md:gap-6 sm:grid-cols-4\"><label for=\"category_id\" class=\"block mb-2 text-sm font-medium text-gray-900 \">Product Category :</label><div><select id=\"category_id\" name=\"category_id\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 \"><option selected=\"\">Select category</option> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, category := range categoryInfo {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(category.RowID)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/product/productcategorycomponents/FashionAndClothingPage.templ`, Line: 32, Col: 42}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(category.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/product/productcategorycomponents/FashionAndClothingPage.templ`, Line: 32, Col: 60}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</select></div></div><div class=\"grid gap-4 sm:col-span-2 md:gap-6 sm:grid-cols-4\"><div><label for=\"name\" class=\"block mb-2 text-sm font-medium text-gray-900 \">Product Name</label> <input type=\"text\" name=\"name\" id=\"name\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5\" placeholder=\"Type product name here\" required=\"\"></div><div><label for=\"barcode\" class=\"block mb-2 text-sm font-medium text-gray-900 \">Product Barcode</label> <input type=\"text\" name=\"barcode\" id=\"barcode\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5\" placeholder=\"Type product barcode here\" required=\"\"></div><div><label for=\"brand_id\" class=\"block mb-2 text-sm font-medium text-gray-900 \">Brand</label> <input type=\"text\" name=\"brand_id\" id=\"brand_id\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 \" placeholder=\"Nike\" required=\"\"></div></div><div class=\"grid gap-4 sm:col-span-2 md:gap-6 sm:grid-cols-4\"><div><label for=\"price\" class=\"block mb-2 text-sm font-medium text-gray-900 \">Price</label> <input type=\"number\" name=\"price\" id=\"price\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 \" placeholder=\"$2999\" required=\"\"></div><div><label for=\"amount\" class=\"block mb-2 text-sm font-medium text-gray-900 \">Amount</label> <input type=\"number\" name=\"amount\" id=\"amount\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 \" placeholder=\"34\" required=\"\"></div></div><div class=\"grid gap-4 sm:col-span-2 md:gap-6 sm:grid-cols-4\"><div><label for=\"prod_size\" class=\"block mb-2 text-sm font-medium text-gray-900 \">Item Size</label> <input type=\"text\" name=\"prod_size\" id=\"prod_size\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 \" placeholder=\"xxl\" required=\"\"></div><div><label for=\"prod_length\" class=\"block mb-2 text-sm font-medium text-gray-900 \">Item Lenght (cm)</label> <input type=\"number\" name=\"prod_length\" id=\"prod_length\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 \" placeholder=\"105\" required=\"\"></div><div><label for=\"prod_material\" class=\"block mb-2 text-sm font-medium text-gray-900 \">Item Material</label> <input type=\"text\" name=\"prod_material\" id=\"prod_material\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 \" placeholder=\"Cotton\" required=\"\"></div><div><label for=\"prod_color\" class=\"block mb-2 text-sm font-medium text-gray-900 \">Item Color</label> <input type=\"text\" name=\"prod_color\" id=\"prod_color\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 \" placeholder=\"Red\" required=\"\"></div></div><div class=\"sm:col-span-2\"><label for=\"description\" class=\"block mb-2 text-sm font-medium text-gray-900 \">Description</label> <input type=\"text\" name=\"description\" id=\"description\" class=\"block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-primary-500 focus:border-primary-500 \" placeholder=\"Write product description here\" required=\"\"></div></div><div class=\"mb-4 space-y-4 sm:flex sm:space-y-0\"><div class=\"flex items-center mr-4\"><input id=\"inline-checkbox\" type=\"checkbox\" value=\"\" name=\"sellingType\" class=\"w-4 h-4 bg-gray-100 rounded border-gray-300 text-primary-600 focus:ring-primary-500  focus:ring-2\"> <label for=\"inline-checkbox\" class=\"ml-2 text-sm font-medium text-gray-900 \">In-store only</label></div><div class=\"flex items-center mr-4\"><input id=\"inline-2-checkbox\" type=\"checkbox\" value=\"\" name=\"sellingType\" class=\"w-4 h-4 bg-gray-100 rounded border-gray-300 text-primary-600 focus:ring-primary-500 focus:ring-2\"> <label for=\"inline-2-checkbox\" class=\"ml-2 text-sm font-medium text-gray-900 \">Online selling only</label></div><div class=\"flex items-center mr-4\"><input checked=\"\" id=\"inline-checked-checkbox\" type=\"checkbox\" value=\"\" name=\"sellingType\" class=\"w-4 h-4 bg-gray-100 rounded border-gray-300 text-primary-600 focus:ring-primary-500 focus:ring-2 \"> <label for=\"inline-checked-checkbox\" class=\"ml-2 text-sm font-medium text-gray-900 \">Both in-store and online</label></div></div><div class=\"mb-4\"><span class=\"block mb-2 text-sm font-medium text-gray-900 \">Product Images</span><div class=\"flex justify-center items-center w-full\"><label for=\"image_file\" class=\"flex flex-col justify-center items-center w-full h-64 bg-gray-50 rounded-lg border-2 border-gray-300 border-dashed cursor-pointer  hover:bg-gray-100 \"><div class=\"flex flex-col justify-center items-center pt-5 pb-6\"><svg aria-hidden=\"true\" class=\"mb-3 w-10 h-10 text-gray-400\" fill=\"none\" stroke=\"currentColor\" viewbox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12\"></path></svg><p class=\"mb-2 text-sm text-gray-500 \"><span class=\"font-semibold\">Click to upload</span> or drag and drop</p><p class=\"text-xs text-gray-500 \">SVG, PNG, JPG or GIF (MAX. 800x400px)</p></div><input id=\"image_file\" name=\"image_file\" type=\"file\" class=\"hidden\"></label></div></div></div></div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
