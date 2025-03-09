// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package warehouse

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"nextgen/templates/components"
	"nextgen/templates/dashboard/dashboardcomponents"
)

type InventoryPageProps struct {
	LayoutProp               dashboardcomponents.LayoutProp
	WarehouseHeaderPageProps []WarehouseHeaderPageProps
	FormLayoutSimpleProp     components.FormLayoutSimpleProp
}

func InventoryPage(inventoryPageProps InventoryPageProps) templ.Component {
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
		templ_7745c5c3_Err = dashboardcomponents.Layout(InventoryPageContent(inventoryPageProps), inventoryPageProps.LayoutProp).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func InventoryPageContent(inventoryPageProps InventoryPageProps) templ.Component {
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
		templ_7745c5c3_Err = WarehouseHeaderPage(inventoryPageProps.WarehouseHeaderPageProps).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"mt-2\"><div class=\" flex  border-2 border-gray-200 border-dashed rounded-lg justify-center \"><div class=\" m-2 w-full bg-gray-200 rounded-2xl\"><div class=\"space-y-12 p-4\"><div class=\"max-w-6xl mx-auto bg-white p-6 rounded-lg shadow-lg\"><h2 class=\"text-2xl font-bold mb-4 text-center\">Clothing Inventory</h2><!-- Add New Product Button --><button id=\"addNewButton\" class=\"bg-gray-500 hover:bg-gray-600 text-white px-4 py-2 rounded mb-4\">+ Add New Product</button><!-- Inventory Table --><div class=\"overflow-x-auto\"><table class=\"w-full border-collapse border border-gray-300\" id=\"inventoryTable\"><thead class=\"bg-gray-200\"><tr><th class=\"border p-2\">#</th><th class=\"border p-2\">Name</th><th class=\"border p-2\">Stock</th><th class=\"border p-2\">Color</th><th class=\"border p-2\">Size</th><th class=\"border p-2\">Material</th><th class=\"border p-2\">User</th><th class=\"border p-2\">Price</th><th class=\"border p-2\">Warehouse</th><th class=\"border p-2\">Actions</th></tr></thead> <tbody id=\"inventoryBody\"><tr class=\"bg-white\"><td class=\"border p-2 text-center\">1</td><td class=\"border p-2 text-center\">Denim Jacket</td><td class=\"border p-2 text-center\">15</td><td class=\"border p-2 text-center\">Blue</td><td class=\"border p-2 text-center\">M</td><td class=\"border p-2 text-center\">Cotton Blend</td><td class=\"border p-2 text-center\">Admin</td><td class=\"border p-2 text-center\">$49.99</td><td class=\"border p-2 text-center\">Warehouse A</td><td class=\"border p-2 text-center\"><button class=\"edit-btn bg-blue-700 hover:bg-blue-800 text-white px-3 py-1 rounded\">Edit</button> <button class=\"delete-btn bg-red-700 hover:bg-red-800 text-white px-3 py-1 rounded\">Delete</button></td></tr></tbody></table></div><!-- Add/Edit Form --><div id=\"productForm\" class=\"hidden mt-6\"><h3 class=\"text-xl font-bold mb-2\" id=\"formTitle\">Add New Product</h3><form id=\"productFormElement\" class=\"grid grid-cols-2 gap-4\"><input type=\"hidden\" id=\"editIndex\"><div><label class=\"block text-sm font-semibold\">Name</label> <input type=\"text\" id=\"productName\" class=\"w-full p-2 border rounded-lg\" required></div><div><label class=\"block text-sm font-semibold\">Stock</label> <input type=\"number\" id=\"productStock\" class=\"w-full p-2 border rounded-lg\" required></div><div><label class=\"block text-sm font-semibold\">Color</label> <input type=\"text\" id=\"productColor\" class=\"w-full p-2 border rounded-lg\" required></div><div><label class=\"block text-sm font-semibold\">Size</label> <input type=\"text\" id=\"productSize\" class=\"w-full p-2 border rounded-lg\" required></div><div><label class=\"block text-sm font-semibold\">Material</label> <input type=\"text\" id=\"productMaterial\" class=\"w-full p-2 border rounded-lg\" required></div><div><label class=\"block text-sm font-semibold\">User</label> <input type=\"text\" id=\"productUser\" class=\"w-full p-2 border rounded-lg\" required></div><div><label class=\"block text-sm font-semibold\">Price ($)</label> <input type=\"number\" id=\"productPrice\" class=\"w-full p-2 border rounded-lg\" required></div><div><label class=\"block text-sm font-semibold\">Warehouse</label> <input type=\"text\" id=\"productWarehouse\" class=\"w-full p-2 border rounded-lg\" required></div><div class=\"col-span-2\"><button type=\"submit\" class=\"w-full bg-blue-500 text-white p-2 rounded-lg\">Save Product</button></div></form></div></div></div></div></div></div><script>\r\n        document.addEventListener(\"DOMContentLoaded\", function () {\r\n            const inventoryTable = document.getElementById(\"inventoryBody\");\r\n            const productForm = document.getElementById(\"productForm\");\r\n            const productFormElement = document.getElementById(\"productFormElement\");\r\n            const addNewButton = document.getElementById(\"addNewButton\");\r\n            const formTitle = document.getElementById(\"formTitle\");\r\n            let editRow = null;\r\n\r\n            // Show Add New Product Form\r\n            addNewButton.addEventListener(\"click\", function () {\r\n                editRow = null;\r\n                formTitle.textContent = \"Add New Product\";\r\n                productFormElement.reset();\r\n                productForm.classList.remove(\"hidden\");\r\n            });\r\n\r\n            // Handle Edit\r\n            inventoryTable.addEventListener(\"click\", function (event) {\r\n                if (event.target.classList.contains(\"edit-btn\")) {\r\n                    let row = event.target.closest(\"tr\");\r\n                    let cells = row.querySelectorAll(\"td\");\r\n\r\n                    document.getElementById(\"productName\").value = cells[1].innerText;\r\n                    document.getElementById(\"productStock\").value = cells[2].innerText;\r\n                    document.getElementById(\"productColor\").value = cells[3].innerText;\r\n                    document.getElementById(\"productSize\").value = cells[4].innerText;\r\n                    document.getElementById(\"productMaterial\").value = cells[5].innerText;\r\n                    document.getElementById(\"productUser\").value = cells[6].innerText;\r\n                    document.getElementById(\"productPrice\").value = parseFloat(cells[7].innerText.replace(\"$\", \"\"));\r\n                    document.getElementById(\"productWarehouse\").value = cells[8].innerText;\r\n\r\n                    editRow = row;\r\n                    formTitle.textContent = \"Edit Product\";\r\n                    productForm.classList.remove(\"hidden\");\r\n                }\r\n            });\r\n\r\n            // Handle Delete\r\n            inventoryTable.addEventListener(\"click\", function (event) {\r\n                if (event.target.classList.contains(\"delete-btn\")) {\r\n                    let row = event.target.closest(\"tr\");\r\n                    row.remove();\r\n                }\r\n            });\r\n\r\n            // Handle Form Submission\r\n            productFormElement.addEventListener(\"submit\", function (event) {\r\n                event.preventDefault();\r\n\r\n                let productData = {\r\n                    name: document.getElementById(\"productName\").value,\r\n                    stock: document.getElementById(\"productStock\").value,\r\n                    color: document.getElementById(\"productColor\").value,\r\n                    size: document.getElementById(\"productSize\").value,\r\n                    material: document.getElementById(\"productMaterial\").value,\r\n                    user: document.getElementById(\"productUser\").value,\r\n                    price: parseFloat(document.getElementById(\"productPrice\").value).toFixed(2),\r\n                    warehouse: document.getElementById(\"productWarehouse\").value,\r\n                };\r\n\r\n                if (editRow) {\r\n                    let cells = editRow.querySelectorAll(\"td\");\r\n                    cells[1].innerText = productData.name;\r\n                    cells[2].innerText = productData.stock;\r\n                    cells[3].innerText = productData.color;\r\n                    cells[4].innerText = productData.size;\r\n                    cells[5].innerText = productData.material;\r\n                    cells[6].innerText = productData.user;\r\n                    cells[7].innerText = `$${productData.price}`;\r\n                    cells[8].innerText = productData.warehouse;\r\n                } else {\r\n                    inventoryTable.insertAdjacentHTML(\"beforeend\", `<tr>\r\n                        <td class=\"border p-2 text-center\">${inventoryTable.rows.length + 1}</td>\r\n                        <td class=\"border p-2\">${productData.name}</td>\r\n                        <td class=\"border p-2\">${productData.stock}</td>\r\n                        <td class=\"border p-2\">${productData.color}</td>\r\n                        <td class=\"border p-2\">${productData.size}</td>\r\n                        <td class=\"border p-2\">${productData.material}</td>\r\n                        <td class=\"border p-2\">${productData.user}</td>\r\n                        <td class=\"border p-2\">$${productData.price}</td>\r\n                        <td class=\"border p-2\">${productData.warehouse}</td>\r\n                        <td class=\"border p-2\"><button class=\"edit-btn bg-yellow-500 text-white px-3 py-1 rounded\">Edit</button> <button class=\"delete-btn bg-red-500 text-white px-3 py-1 rounded\">Delete</button></td>\r\n                    </tr>`);\r\n                }\r\n\r\n                productForm.classList.add(\"hidden\");\r\n                productFormElement.reset();\r\n            });\r\n        });\r\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
