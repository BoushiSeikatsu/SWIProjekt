// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package templatesAdminProduct

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"strconv"
	"swi-warehouse/models"
	"swi-warehouse/templates"
)

func AdminProductHistory(adds []models.Added) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Supply and Withdraw History</title><link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css\" rel=\"stylesheet\"><link href=\"https://cdnjs.cloudflare.com/ajax/libs/bootstrap-icons/1.8.1/font/bootstrap-icons.min.css\" rel=\"stylesheet\"><style>\n        body, html {\n            height: 100%;\n            margin: 0;\n            overflow: hidden;\n        }\n        .navbar {\n            position: fixed;\n            top: 0;\n            width: 100%;\n            z-index: 1000;\n            background-color: #be274b;\n            padding: 10px 0;\n        }\n        .navbar .nav-link,\n        .navbar .navbar-icon {\n            color: #e2e6ea;\n            padding: 10px 15px;\n            font-weight: bold;\n            text-align: center;\n            display: flex;\n            flex-direction: column;\n            line-height: 1.2;\n            position: relative;\n        }\n        .navbar-icon {\n            font-size: 1.5rem;\n            cursor: default;\n        }\n        .nav-link[onclick] {\n            cursor: pointer;\n            user-select: none;\n        }\n        .d-flex {\n            align-items: center;\n        }\n        .navbar .nav-link:hover {\n            background-color: #da3060;\n            border-radius: 5px;\n            transition: background-color 0.3s ease;\n        }\n        .content-container {\n            display: flex;\n            justify-content: center;\n            align-items: center;\n            min-height: 100vh;\n            padding-top: 78px;\n        }\n        .justify-between {\n            display: flex;\n            align-items: center;\n            justify-content: space-between;\n            width: 100%;\n        }\n        .submenu {\n            display: none;\n            position: absolute;\n            background-color: #f8f9fa;\n            border-radius: 5px;\n            padding: 10px;\n            top: 100%;\n            left: 50%;\n            transform: translateX(-50%);\n            margin-top: 5px;\n            z-index: 1001;\n        }\n        .submenu a {\n            color: #333;\n            text-decoration: none;\n            display: block;\n            padding: 5px 15px;\n        }\n        .submenu a:hover {\n            background-color: #ddd;\n            border-radius: 5px;\n        }\n        .setup-container {\n            height: calc(100% + 78px);\n            display: flex;\n            justify-content: center;\n            align-items: center;\n        }\n        .inventory-container {\n            margin-top: -75px;\n        }\n        .table-responsive {\n            max-height: 400px;\n            overflow-y: auto;\n        }\n        .table th, .table td {\n            text-align: center;\n            vertical-align: middle;\n        }\n        .pagination {\n            justify-content: center;\n        }\n    </style></head><body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templates.AdminNav().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container setup-container\"><div class=\"container inventory-container\"><h2 class=\"mb-4 text-center\">Supply and Withdraw History</h2><div class=\"d-flex justify-content-end align-items-center mb-3\"><form class=\"d-flex align-items-center\" action=\"#\" method=\"GET\"><input type=\"text\" name=\"search\" class=\"form-control me-2\" placeholder=\"Search...\" style=\"width: 200px;\"> <select name=\"user\" class=\"form-select me-2\" style=\"width: 150px;\"><option selected value=\"all\">All Users</option> <option value=\"1\">Users 1</option> <option value=\"2\">Users 2</option> <option value=\"3\">Users 3</option></select> <select name=\"moveAction\" class=\"form-select me-2\" style=\"width: 150px;\"><option selected value=\"all\">All Actions</option> <option value=\"1\">Supply</option> <option value=\"2\">Withdraw</option></select> <select name=\"manufacturer\" class=\"form-select me-2\" style=\"width: 200px;\"><option selected value=\"all\">All Manufacturers</option> <option value=\"1\">Manufacturer 1</option> <option value=\"2\">Manufacturer 2</option> <option value=\"3\">Manufacturer 3</option></select> <select name=\"storage\" class=\"form-select me-2\" style=\"width: 150px;\"><option selected value=\"all\">All Storages</option> <option value=\"1\">Storage 1</option> <option value=\"2\">Storage 2</option> <option value=\"3\">Storage 3</option></select> <button type=\"submit\" class=\"btn btn-primary\"><i class=\"bi bi-search\"></i> Search</button></form></div><div class=\"table-responsive\"><table class=\"table table-bordered table-striped\"><thead class=\"table-dark\"><tr><th>Date</th><th>User</th><th>Action</th><th>Manufacturer</th><th>Product</th><th>Amount</th><th>Unit</th><th>Storage</th></tr></thead> <tbody><!-- Placeholder Item -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, add := range adds {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr><td>Někdy</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(add.User.Username)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductHistory.templ`, Line: 174, Col: 50}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if add.Amount > 0 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<td>Supply</td>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<td>Withdraw</td>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(add.Product.Manufacturer.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductHistory.templ`, Line: 180, Col: 62}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(add.Product.Description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductHistory.templ`, Line: 181, Col: 56}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(add.Amount))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductHistory.templ`, Line: 182, Col: 57}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(add.Product.Unit)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductHistory.templ`, Line: 183, Col: 49}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(add.Storage.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductHistory.templ`, Line: 184, Col: 49}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td></tr>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tbody></table></div><nav aria-label=\"Inventory pagination\"><ul class=\"pagination\"><li class=\"page-item disabled\"><a class=\"page-link\" href=\"#\" tabindex=\"-1\">Previous</a></li><li class=\"page-item\"><a class=\"page-link\" href=\"#\">1</a></li><li class=\"page-item\"><a class=\"page-link\" href=\"#\">2</a></li><li class=\"page-item\"><a class=\"page-link\" href=\"#\">3</a></li><li class=\"page-item\"><a class=\"page-link\" href=\"#\">Next</a></li></ul></nav></div></div><script>\n        function toggleSubmenu(id) {\n            const submenu = document.getElementById(id);\n            const isVisible = submenu.style.display === 'block';\n            document.querySelectorAll('.submenu').forEach(menu => menu.style.display = 'none');\n            submenu.style.display = isVisible ? 'none' : 'block';\n        }\n        document.addEventListener('click', function(e) {\n            const isClickInside = e.target.closest('.nav-link') || e.target.closest('.submenu');\n            if (!isClickInside) {\n                document.querySelectorAll('.submenu').forEach(menu => menu.style.display = 'none');\n            }\n        });\n    </script><script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
