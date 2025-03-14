// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package product

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"nextgen/pkg/product_management"
	"nextgen/templates/components"
	"nextgen/templates/dashboard/dashboardcomponents"
	"nextgen/templates/dashboard/pages/product/productcategorycomponents"
)

type AddProductPageProps struct {
	LayoutProp                  dashboardcomponents.LayoutProp
	AddProductPageContentsProps AddProductPageContentsProps
}

type AddProductPageContentsProps struct {
	ProductPageHeaderProps []ProductPageHeaderProp
	FormLayoutSimpleProp   components.FormLayoutSimpleProp
	CategoryInfo           []product_management.CategoryView
}

func AddProductPage(addProductPageProps AddProductPageProps) templ.Component {
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
		templ_7745c5c3_Err = dashboardcomponents.Layout(AddProductPageContents(addProductPageProps.AddProductPageContentsProps), addProductPageProps.LayoutProp).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func AddProductPageContents(addProductPageContentsProp AddProductPageContentsProps) templ.Component {
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
		templ_7745c5c3_Err = ProductPageHeader(addProductPageContentsProp.ProductPageHeaderProps).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.FormLayoutSimple(AddProductPageForm(addProductPageContentsProp), addProductPageContentsProp.FormLayoutSimpleProp).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func AddProductPageForm(categoryInfo AddProductPageContentsProps) templ.Component {
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = productcategorycomponents.FashionAndClothingPage(categoryInfo.CategoryInfo).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

//	templ AddProductPageForm() {
//		<div class="mt-2">
//			<div class=" flex  border-2 border-gray-200 border-dashed rounded-lg justify-center ">
//				<div class=" m-2 w-full bg-gray-200 rounded-2xl">
//					<div class="space-y-12">
//						<div class="relative p-4 w-full  h-full md:h-auto">
//							<!-- Modal content -->
//							<div class="relative p-4 bg-white rounded-lg shadow  sm:p-5">
//								<!-- Modal header -->
//								<div class="flex justify-between items-center pb-4 mb-4 rounded-t border-b sm:mb-5 ">
//									<h3 class="text-lg font-semibold text-gray-900">Add Product</h3>
//								</div>
//								<!-- Modal body -->
//								<form action="#">
//									<div class="grid gap-4 mb-4 sm:grid-cols-2">
//										<div>
//											<label for="name" class="block mb-2 text-sm font-medium text-gray-900 ">Product Name</label>
//											<input type="text" name="name" id="name" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5" placeholder="Type product name" required=""/>
//										</div>
//										<div><label for="category" class="block mb-2 text-sm font-medium text-gray-900 ">Category</label><select id="category" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 "><option selected="">Select category</option><option value="TV">TV/Monitors</option><option value="PC">PC</option><option value="GA">Gaming/Console</option><option value="PH">Phones</option></select></div>
//										<div>
//											<label for="brand" class="block mb-2 text-sm font-medium text-gray-900 ">Brand</label>
//											<input type="text" name="brand" id="brand" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 " placeholder="Product brand" required=""/>
//										</div>
//										<div>
//											<label for="price" class="block mb-2 text-sm font-medium text-gray-900 ">Price</label>
//											<input type="number" name="price" id="price" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 " placeholder="$2999" required=""/>
//										</div>
//										<div class="grid gap-4 sm:col-span-2 md:gap-6 sm:grid-cols-4">
//											<div>
//												<label for="weight" class="block mb-2 text-sm font-medium text-gray-900 ">Item weight (kg)</label>
//												<input type="number" name="weight" id="itemweight" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 " placeholder="12" required=""/>
//											</div>
//											<div>
//												<label for="length" class="block mb-2 text-sm font-medium text-gray-900 ">Lenght (cm)</label>
//												<input type="number" name="length" id="itemlength" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 " placeholder="105" required=""/>
//											</div>
//											<div>
//												<label for="breadth" class="block mb-2 text-sm font-medium text-gray-900 ">Breadth (cm)</label>
//												<input type="number" name="breadth" id="itembreadth" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 " placeholder="15" required=""/>
//											</div>
//											<div>
//												<label for="width" class="block mb-2 text-sm font-medium text-gray-900 ">Width (cm)</label>
//												<input type="number" name="width" id="itemwidth" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 " placeholder="23" required=""/>
//											</div>
//										</div>
//										<div class="sm:col-span-2"><label for="description" class="block mb-2 text-sm font-medium text-gray-900 ">Description</label><textarea id="description" rows="4" class="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-primary-500 focus:border-primary-500 " placeholder="Write product description here"></textarea></div>
//									</div>
//									<div class="mb-4 space-y-4 sm:flex sm:space-y-0">
//										<div class="flex items-center mr-4">
//											<input id="inline-checkbox" type="checkbox" value="" name="sellingType" class="w-4 h-4 bg-gray-100 rounded border-gray-300 text-primary-600 focus:ring-primary-500  focus:ring-2"/>
//											<label for="inline-checkbox" class="ml-2 text-sm font-medium text-gray-900 ">In-store only</label>
//										</div>
//										<div class="flex items-center mr-4">
//											<input id="inline-2-checkbox" type="checkbox" value="" name="sellingType" class="w-4 h-4 bg-gray-100 rounded border-gray-300 text-primary-600 focus:ring-primary-500 focus:ring-2"/>
//											<label for="inline-2-checkbox" class="ml-2 text-sm font-medium text-gray-900 ">Online selling only</label>
//										</div>
//										<div class="flex items-center mr-4">
//											<input checked="" id="inline-checked-checkbox" type="checkbox" value="" name="sellingType" class="w-4 h-4 bg-gray-100 rounded border-gray-300 text-primary-600 focus:ring-primary-500 focus:ring-2 "/>
//											<label for="inline-checked-checkbox" class="ml-2 text-sm font-medium text-gray-900 ">Both in-store and online</label>
//										</div>
//									</div>
//									<div class="mb-4">
//										<span class="block mb-2 text-sm font-medium text-gray-900 ">Product Images</span>
//										<div class="flex justify-center items-center w-full">
//											<label for="dropzone-file" class="flex flex-col justify-center items-center w-full h-64 bg-gray-50 rounded-lg border-2 border-gray-300 border-dashed cursor-pointer  hover:bg-gray-100 ">
//												<div class="flex flex-col justify-center items-center pt-5 pb-6">
//													<svg aria-hidden="true" class="mb-3 w-10 h-10 text-gray-400" fill="none" stroke="currentColor" viewbox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
//														<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
//													</svg>
//													<p class="mb-2 text-sm text-gray-500 ">
//														<span class="font-semibold">Click to upload</span>
//														or drag and drop
//													</p>
//													<p class="text-xs text-gray-500 ">SVG, PNG, JPG or GIF (MAX. 800x400px)</p>
//												</div>
//												<input id="dropzone-file" type="file" class="hidden"/>
//											</label>
//										</div>
//									</div>
//									<div class="items-center space-y-4 sm:flex sm:space-y-0 sm:space-x-4">
//										<button type="submit" class="w-full justify-center sm:w-auto text-gray-500 inline-flex items-center bg-white hover:bg-gray-100 focus:ring-4 focus:outline-none focus:ring-primary-300 rounded-lg border border-gray-200 text-sm font-medium px-5 py-2.5 hover:text-gray-900 focus:z-10 ">Add product</button>
//									</div>
//								</form>
//							</div>
//						</div>
//					</div>
//				</div>
//			</div>
//		</div>
//	}
var _ = templruntime.GeneratedTemplate
