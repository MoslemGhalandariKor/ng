// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package team

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "nextgen/templates/dashboard/dashboardcomponents"

type EmployeeInfoProps struct {
	FirstName   string
	LastName    string
	Address     string
	PhoneNumber string
	Email       string
	NationalId  string
	Salary      string
	Username    string
	Role        string
	Birthday    string
	Photo       string
	DeleteUrl   string
}

type EmployeePageProp struct {
	LayoutProp          dashboardcomponents.LayoutProp
	TeamPageHeaderProps []TeamPageHeaderProp
	EmployeesInfoProps  []EmployeeInfoProps
	// AddEmployeePageProp []AddEmployeePageProp
}

func EmployeesPage(employeePageProp EmployeePageProp) templ.Component {
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
		templ_7745c5c3_Err = dashboardcomponents.Layout(EmployeesPageContent(employeePageProp), employeePageProp.LayoutProp).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func EmployeesPageContent(employeePageProp EmployeePageProp) templ.Component {
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
		templ_7745c5c3_Err = TeamPageHeader(employeePageProp.TeamPageHeaderProps).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"mt-2 w-full\"><div class=\" border-2 border-gray-200 border-dashed rounded-lg dark:border-gray-700 justify-center \"><div class=\"p-12 m-4 bg-gray-200 dark:bg-gray-200 rounded-2xl\"><div class=\"space-y-7  m-2 pl-4  bg-gray-200 rounded-lg justify-center\"><div class=\"sm:flex sm:items-center\"><div class=\"sm:flex-auto\"><p class=\"mt-4 text-sm text-gray-700 \">A list of all the users in your account including their name, title, email and role.</p></div></div><div class=\"mt-8 flow-root\"><div class=\"-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8\"><div class=\"inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8\"><div class=\"overflow-hidden shadow ring-1 ring-black ring-opacity-5 sm:rounded-lg\"><table class=\"min-w-full divide-y divide-gray-300\"><thead class=\"bg-gray-50\"><tr><th scope=\"col\" class=\"py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-6\">Name</th><th scope=\"col\" class=\"px-3 py-3.5 text-left text-sm font-semibold text-gray-900\">Username</th><th scope=\"col\" class=\"px-3 py-3.5 text-left text-sm font-semibold text-gray-900\">PhoneNumber</th><th scope=\"col\" class=\"px-3 py-3.5 text-left text-sm font-semibold text-gray-900\">Role</th><th scope=\"col\" class=\"relative py-3.5 pl-3 pr-4 sm:pr-6\"><span class=\"sr-only\">Edit</span></th></tr></thead> <tbody class=\"divide-y divide-gray-200 bg-white\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, employee := range employeePageProp.EmployeesInfoProps {
			templ_7745c5c3_Err = EmployeesInfo(employee).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</tbody></table></div></div></div></div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func EmployeesInfo(employee EmployeeInfoProps) templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<tr><td class=\"whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-6\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(employee.FirstName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 119, Col: 109}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, " ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(employee.LastName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 119, Col: 131}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</td><td class=\"whitespace-nowrap px-3 py-4 text-sm text-gray-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(employee.Username)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 120, Col: 83}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "</td><td class=\"whitespace-nowrap px-3 py-4 text-sm text-gray-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(employee.PhoneNumber)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 121, Col: 86}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</td><td class=\"whitespace-nowrap px-3 py-4 text-sm text-gray-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(employee.Role)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 122, Col: 79}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "</td><td class=\"relative whitespace-nowrap py-4 pl-3 pr-4 text-right text-sm font-medium sm:pr-6\"><button data-modal-target=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(employee.Username)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 124, Col: 48}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\" data-modal-toggle=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(employee.Username)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 124, Col: 88}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "\" class=\"block text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800\" type=\"button\">More Info</button><!-- Main modal --><div id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(employee.Username)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 128, Col: 30}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "\" aria-hidden=\"true\" class=\" hidden overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-50 justify-center items-center w-full md:inset-0 h-[calc(100%-1rem)] max-h-full\"><div class=\"relative p-4 w-full max-w-2xl max-h-full\"><!-- Modal content --><div class=\"relative bg-white rounded-lg  dark:bg-gray-700\"><!-- Modal header --><div class=\"flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600\"><h3 class=\"text-xl font-semibold text-gray-900 dark:text-white\">Employee Information</h3><button type=\"button\" class=\"text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white\" data-modal-hide=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(employee.Username)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 137, Col: 259}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "\"><svg class=\"w-3 h-3\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 14 14\"><path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6\"></path></svg> <span class=\"sr-only\">Close modal</span></button></div><!-- Modal body --><div class=\" border-2 border-gray-200 border-dashed rounded-lg dark:border-gray-700 justify-center \"><div class=\"p-12 m-4 bg-gray-200 dark:bg-gray-200 rounded-2xl\"><div class=\"space-y-7  m-2 pl-4  bg-gray-200 rounded-lg justify-center\"><div class=\" grid gap-4 sm:grid-cols-2 sm:gap-6 \"><div class=\"space-y-8 border-b border-gray-900/10  pb-12 sm:space-y-0 sm:divide-y sm:divide-gray-900/10 sm:border-t sm:pb-0\"><div class=\"sm:grid sm:grid-cols-3 sm:items-start sm:gap-2 sm:py-4\"><label for=\"name\" class=\"block text-sm font-medium leading-6 text-gray-900  sm:pt-1.5\">Name</label><div class=\"mt-2 sm:col-span-2 sm:mt-0\"><input type=\"text\" name=\"name\" id=\"name\" autocomplete=\"given-name\" class=\" p-2  block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(employee.FirstName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 153, Col: 336}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "\"></div></div></div><div class=\"space-y-8 border-b border-gray-900/10  pb-12 sm:space-y-0 sm:divide-y sm:divide-gray-900/10 sm:border-t sm:pb-0\"><div class=\"sm:grid sm:grid-cols-3 sm:items-start sm:gap-2 sm:py-4\"><label for=\"last_name\" class=\"block text-sm font-medium leading-6 text-gray-900  sm:pt-1.5\">Last Name</label><div class=\"mt-2 sm:col-span-2 sm:mt-0\"><input type=\"text\" name=\"last_name\" id=\"last_name\" autocomplete=\"given-name\" class=\" p-2  block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(employee.LastName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 161, Col: 345}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "\"></div></div></div><div class=\"space-y-8 border-b border-gray-900/10  pb-12 sm:space-y-0 sm:divide-y sm:divide-gray-900/10 sm:border-t sm:pb-0\"><div class=\"sm:grid sm:grid-cols-3 sm:items-start sm:gap-2 sm:py-4\"><label for=\"address\" class=\"block text-sm font-medium leading-6 text-gray-900  sm:pt-1.5\">Address</label><div class=\"mt-2 sm:col-span-2 sm:mt-0\"><input type=\"text\" name=\"address\" id=\"address\" autocomplete=\"given-name\" class=\" p-2  block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(employee.Address)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 169, Col: 340}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "\"></div></div></div><div class=\"space-y-8 border-b border-gray-900/10  pb-12 sm:space-y-0 sm:divide-y sm:divide-gray-900/10 sm:border-t sm:pb-0\"><div class=\"sm:grid sm:grid-cols-3 sm:items-start sm:gap-2 sm:py-4\"><label for=\"phone_number\" class=\"block text-sm font-medium leading-6 text-gray-900  sm:pt-1.5\">Phone Number</label><div class=\"mt-2 sm:col-span-2 sm:mt-0\"><input type=\"text\" name=\"phone_number\" id=\"phone_number\" autocomplete=\"given-name\" class=\" p-2  block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var16 string
		templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(employee.PhoneNumber)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 177, Col: 354}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "\"></div></div></div><div class=\"space-y-8 border-b border-gray-900/10  pb-12 sm:space-y-0 sm:divide-y sm:divide-gray-900/10 sm:border-t sm:pb-0\"><div class=\"sm:grid sm:grid-cols-3 sm:items-start sm:gap-2 sm:py-4\"><label for=\"email\" class=\"block text-sm font-medium leading-6 text-gray-900  sm:pt-1.5\">Email</label><div class=\"mt-2 sm:col-span-2 sm:mt-0\"><input type=\"email\" name=\"email\" id=\"email\" autocomplete=\"given-name\" class=\" p-2  block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var17 string
		templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(employee.Email)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 185, Col: 335}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "\"></div></div></div><div class=\"space-y-8 border-b border-gray-900/10  pb-12 sm:space-y-0 sm:divide-y sm:divide-gray-900/10 sm:border-t sm:pb-0\"><div class=\"sm:grid sm:grid-cols-3 sm:items-start sm:gap-2 sm:py-4\"><label for=\"national_id\" class=\"block text-sm font-medium leading-6 text-gray-900  sm:pt-1.5\">National Id</label><div class=\"mt-2 sm:col-span-2 sm:mt-0\"><input type=\"text\" name=\"national_id\" id=\"national_id\" autocomplete=\"given-name\" class=\" p-2  block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var18 string
		templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(employee.NationalId)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 193, Col: 351}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "\"></div></div></div><div class=\"space-y-8 border-b border-gray-900/10  pb-12 sm:space-y-0 sm:divide-y sm:divide-gray-900/10 sm:border-t sm:pb-0\"><div class=\"sm:grid sm:grid-cols-3 sm:items-start sm:gap-2 sm:py-4\"><label for=\"salary\" class=\"block text-sm font-medium leading-6 text-gray-900  sm:pt-1.5\">Salary</label><div class=\"mt-2 sm:col-span-2 sm:mt-0\"><input type=\"number\" name=\"salary\" id=\"salary\" autocomplete=\"given-name\" class=\" p-2  block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var19 string
		templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs(employee.Salary)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 201, Col: 339}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "\"></div></div></div><div class=\"space-y-8 border-b border-gray-900/10  pb-12 sm:space-y-0 sm:divide-y sm:divide-gray-900/10 sm:border-t sm:pb-0\"><div class=\"sm:grid sm:grid-cols-3 sm:items-start sm:gap-2 sm:py-4\"><label for=\"username\" class=\"block text-sm font-medium leading-6 text-gray-900  sm:pt-1.5\">UserName</label><div class=\"mt-2 sm:col-span-2 sm:mt-0\"><input type=\"text\" name=\"username\" id=\"username\" autocomplete=\"given-name\" class=\" p-2  block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var20 string
		templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(employee.Username)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 209, Col: 343}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "\"></div></div></div><div class=\"space-y-8 border-b border-gray-900/10  pb-12 sm:space-y-0 sm:divide-y sm:divide-gray-900/10 sm:border-t sm:pb-0\"><div class=\"sm:grid sm:grid-cols-3 sm:items-start sm:gap-2 sm:py-4\"><label for=\"role\" class=\"block text-sm font-medium leading-6 text-gray-900  sm:pt-1.5\">Role</label><div class=\"mt-2 sm:col-span-2 sm:mt-0\"><input type=\"text\" name=\"role\" id=\"role\" autocomplete=\"given-role\" class=\" p-2  block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var21 string
		templ_7745c5c3_Var21, templ_7745c5c3_Err = templ.JoinStringErrs(employee.Role)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 217, Col: 331}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var21))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "\"></div></div></div><div class=\"relative space-y-8 border-b border-gray-900/10  pb-12 sm:space-y-0 sm:divide-y sm:divide-gray-900/10 sm:border-t sm:pb-0\"><div class=\" relative sm:grid sm:grid-cols-3 sm:items-start sm:gap-2 sm:py-4\"><label for=\"birthday\" class=\"block text-sm font-medium leading-6 text-gray-900  sm:pt-1.5\">Birthday</label><div class=\"mt-2 sm:col-span-2 sm:mt-0\"><input type=\"text\" name=\"birthday\" id=\"birthday\" autocomplete=\"given-birthday\" class=\" p-2  block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var22 string
		templ_7745c5c3_Var22, templ_7745c5c3_Err = templ.JoinStringErrs(employee.Birthday)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 225, Col: 347}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var22))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "\"></div></div></div><div class=\"col-span-full flex justify-between border-t border-b p-8 border-gray-900/10\"><label for=\"photo\" class=\"block text-sm font-medium leading-6 text-gray-900  sm:pt-1.5\">Photo</label><div class=\" flex items-center gap-x-3\"><img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var23 string
		templ_7745c5c3_Var23, templ_7745c5c3_Err = templ.JoinStringErrs(employee.Photo)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/dashboard/pages/team/EmployeesPage.templ`, Line: 232, Col: 37}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var23))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "\" alt=\"employee photo\" class=\"h-12 w-12 text-gray-300\" viewBox=\"0 0 24 24\" fill=\"currentColor\" aria-hidden=\"true\"> <path fill-rule=\"evenodd\" d=\"M18.685 19.097A9.723 9.723 0 0021.75 12c0-5.385-4.365-9.75-9.75-9.75S2.25 6.615 2.25 12a9.723 9.723 0 003.065 7.097A9.716 9.716 0 0012 21.75a9.716 9.716 0 006.685-2.653zm-12.54-1.285A7.486 7.486 0 0112 15a7.486 7.486 0 015.855 2.812A8.224 8.224 0 0112 20.25a8.224 8.224 0 01-5.855-2.438zM15.75 9a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z\" clip-rule=\"evenodd\"></path> <button type=\"button\" class=\"rounded-md bg-white px-2.5 py-1.5 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50\">Change</button></div></div></div></div></div></div><!-- Modal footer --><div class=\"flex justify-between p-4 md:p-5 border-t border-gray-200 rounded-b dark:border-gray-600\"><form action=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var24 templ.SafeURL = templ.URL(employee.DeleteUrl)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var24)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "\" method=\"post\"><button data-modal-hide=\"default-modal\" type=\"submit\" class=\"text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800\">Delete</button></form><form action=\"\" method=\"post\"><button data-modal-hide=\"default-modal\" type=\"submit\" class=\"text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800\">Edit</button></form></div></div></div></div></td></tr>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

// href={ templ.URL(post.Link) }
var _ = templruntime.GeneratedTemplate
