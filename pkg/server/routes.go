package server

import (
	"nextgen/pkg/accounts"
	"nextgen/pkg/auth"
	product_management "nextgen/pkg/product_management"
	"nextgen/pkg/web/aboutus"
	"nextgen/pkg/web/app"
	appmiddleware "nextgen/pkg/web/app/middleware"
	"nextgen/pkg/web/blog"
	"nextgen/pkg/web/dashboard"
	dashboardmiddleware "nextgen/pkg/web/dashboard/middleware"
	"nextgen/pkg/web/dashboard/product"
	"nextgen/pkg/web/dashboard/profile"
	"nextgen/pkg/web/dashboard/sell"
	"nextgen/pkg/web/dashboard/team"
	"nextgen/pkg/web/dashboard/warehouse"
)

type Routes struct {
	server *Server
}

func (r *Routes) webRouter(server *Server) {
	r.server = server
	webRoutes := server.Router.Group("")
	webRoutes.Use(appmiddleware.LayoutPropMiddelware)
	webRoutes.GET("/home", app.HomePage)
	webRoutes.GET("/login", app.LoginPage)
	webRoutes.GET("/singup", app.RegisterPage)

}

func (r *Routes) blogRouter(server *Server) {
	r.server = server
	blogRoutes := server.Router.Group("blog")
	blogRoutes.Use(appmiddleware.LayoutPropMiddelware)
	blogRoutes.GET("/", blog.BlogPage)
	blogRoutes.GET("/post/:id", blog.PostPage)
}

func (r *Routes) aboutusRouter(server *Server) {
	r.server = server
	aboutusRouter := server.Router.Group("aboutus")
	aboutusRouter.Use(appmiddleware.LayoutPropMiddelware)
	aboutusRouter.GET("/ourteam", aboutus.TeamPage)
}

func (r *Routes) accountsRouter(server *Server) {
	r.server = server
	aboutusRouter := server.Router.Group("accounts")
	aboutusRouter.POST("/membership", accounts.AddMembershipHandler)
	aboutusRouter.POST("/addemployee", team.AddEmployeeHandler)
	aboutusRouter.POST("/addtask", team.AddTaskHandler)

}

func (r *Routes) dashboardRouter(server *Server) {
	r.server = server
	dashboardRoutes := server.Router.Group("dashboard")
	dashboardRoutes.Use(auth.AuthMiddleware)
	dashboardRoutes.Use(dashboardmiddleware.LayoutPropMiddelware)
	dashboardRoutes.GET("/", dashboard.DashboardPage)
	// Profile Routes
	dashboardRoutes.GET("/personal-profile", profile.PersonalProfilePage)
	dashboardRoutes.GET("/company-profile", profile.CompanyProfilePage)
	// Employees Routes
	dashboardRoutes.GET("/employee/employees", team.EmployeesPage)
	dashboardRoutes.GET("/employee/add-employee", team.AddEmployeePage)
	dashboardRoutes.POST("/employee/delete-employee/:id", team.DeleteEmployeeHandler)
	dashboardRoutes.GET("/employee/tasks", team.TasksPage)
	dashboardRoutes.GET("/employee/add-task", team.AddTaskPage)
	dashboardRoutes.POST("/employee/delete-task/:id", team.DeleteTaskHandler)
	// Products Routes
	dashboardRoutes.GET("/products", product.ProductsPage)
	dashboardRoutes.GET("/add-product", product.AddProductPage)
	dashboardRoutes.POST("/add-product", product_management.AddProductHandler)
	dashboardRoutes.POST("/delete-product/:id", product_management.DeleteProductHandler)
	dashboardRoutes.GET("/categories", product.CategoryPage)
	dashboardRoutes.GET("/add-category", product.AddCategoryPage)
	dashboardRoutes.POST("/add-category", product_management.AddCategoryHandler)
	dashboardRoutes.POST("/delete-category/:id", product_management.DeleteCategoryHandler)
	dashboardRoutes.GET("brands", product.BrandPage)
	dashboardRoutes.GET("/add-brand", product.AddBrandPage)
	// dashboardRoutes.POST("/add-brands", product_management.AddBrandHandler)
	// dashboardRoutes.POST("/delete-brands/:id", product_management.DeleteBrandHandler)

	// Sell Routes
	dashboardRoutes.GET("/sellingPage", sell.SellingPage)
	// Warehouse Routes
	dashboardRoutes.GET("/warehouses", warehouse.WarehousesPage)
	dashboardRoutes.GET("/add-warehouse", warehouse.AddWarehousePage)
	dashboardRoutes.GET("/inventory", warehouse.InventoryPage)
	// Extra Routes
	dashboardRoutes.GET("/calendar", dashboard.CalendarPage)
	dashboardRoutes.GET("/loyalty", dashboard.LoyaltyPage)
	dashboardRoutes.GET("/campaign", dashboard.CampaignPage)
	dashboardRoutes.GET("/settings", dashboard.SettingsPage)

}

func (r *Routes) authRouter(server *Server) {
	r.server = server
	authRoutes := server.Router.Group("/auth")
	authRoutes.POST("/register", auth.RegisterHandler)
	authRoutes.POST("/login", auth.LoginHandler)
	authRoutes.GET("/logout", auth.LogoutHandler)
}
