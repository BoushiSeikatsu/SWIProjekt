// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package templatesAdminProduct

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"strconv"
	"swi-warehouse/models"
	"swi-warehouse/templates"
)

func AdminSearchProduct(products []models.Product) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Manage Products</title><link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css\" rel=\"stylesheet\"><link href=\"https://cdnjs.cloudflare.com/ajax/libs/bootstrap-icons/1.8.1/font/bootstrap-icons.min.css\" rel=\"stylesheet\"><style>\n        body, html {\n            height: 100%;\n            margin: 0;\n            overflow: hidden;\n        }\n        .navbar {\n            position: fixed;\n            top: 0;\n            width: 100%;\n            z-index: 1000;\n            background-color: #be274b;\n            padding: 10px 0;\n        }\n        .navbar .nav-link,\n        .navbar .navbar-icon {\n            color: #e2e6ea;\n            padding: 10px 15px;\n            font-weight: bold;\n            text-align: center;\n            display: flex;\n            flex-direction: column;\n            line-height: 1.2;\n            position: relative;\n        }\n        .navbar-icon {\n            font-size: 1.5rem;\n            cursor: default;\n        }\n        .nav-link[onclick] {\n            cursor: pointer;\n            user-select: none;\n        }\n        .d-flex {\n            align-items: center;\n        }\n        .navbar .nav-link:hover {\n            background-color: #da3060;\n            border-radius: 5px;\n            transition: background-color 0.3s ease;\n        }\n        .content-container {\n            display: flex;\n            justify-content: center;\n            align-items: center;\n            min-height: 100vh;\n            padding-top: 78px;\n        }\n        .justify-between {\n            display: flex;\n            align-items: center;\n            justify-content: space-between;\n            width: 100%;\n        }\n        .submenu {\n            display: none;\n            position: absolute;\n            background-color: #f8f9fa;\n            border-radius: 5px;\n            padding: 10px;\n            top: 100%;\n            left: 50%;\n            transform: translateX(-50%);\n            margin-top: 5px;\n            z-index: 1001;\n        }\n        .submenu a {\n            color: #333;\n            text-decoration: none;\n            display: block;\n            padding: 5px 15px;\n        }\n        .submenu a:hover {\n            background-color: #ddd;\n            border-radius: 5px;\n        }\n        .setup-container {\n            height: calc(100% + 78px);\n            display: flex;\n            justify-content: center;\n            align-items: center;\n        }\n        .inventory-container {\n            margin-top: -75px;\n        }\n        .table-responsive {\n            max-height: 400px;\n            overflow-y: auto;\n        }\n        .table th, .table td {\n            text-align: center;\n            vertical-align: middle;\n        }\n        .pagination {\n            justify-content: center;\n        }\n    </style></head><body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templates.AdminNav().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container setup-container\"><div class=\"container inventory-container\"><h2 class=\"mb-4 text-center\">Manage Products</h2><div class=\"d-flex justify-content-between align-items-center mb-3\"><a href=\"/admin/addProduct\"><button class=\"btn btn-success\"><i class=\"bi bi-shield-lock me-2\"></i>Create New Item</button></a><form class=\"d-flex align-items-center\" action=\"#\" method=\"GET\"><input type=\"text\" name=\"search\" class=\"form-control me-2\" placeholder=\"Search...\" style=\"width: 200px;\"> <select name=\"manufacturer\" class=\"form-select me-2\" style=\"width: 200px;\"><option selected value=\"all\">All Manufacturers</option> <option value=\"1\">Manufacturer 1</option> <option value=\"2\">Manufacturer 2</option> <option value=\"3\">Manufacturer 3</option></select> <button type=\"submit\" class=\"btn btn-primary\"><i class=\"bi bi-search\"></i> Search</button></form></div><div class=\"table-responsive\"><table class=\"table table-bordered table-striped\"><thead class=\"table-dark\"><tr><th>Description</th><th>Manufacturer</th><th>Width</th><th>Height</th><th>Depth</th><th>Batch</th><th>Unit</th><th>Identification</th><th>Actions</th></tr></thead> <tbody><!-- Placeholder Item -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, product := range products {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(product.Description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductSearch.templ`, Line: 162, Col: 52}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(product.Manufacturer.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductSearch.templ`, Line: 163, Col: 58}
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
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatUint(uint64(product.XDimension), 10))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductSearch.templ`, Line: 164, Col: 83}
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
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatUint(uint64(product.YDimension), 10))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductSearch.templ`, Line: 165, Col: 83}
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
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatUint(uint64(product.ZDimension), 10))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductSearch.templ`, Line: 166, Col: 83}
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
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(product.Batch)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductSearch.templ`, Line: 167, Col: 46}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(product.Unit)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductSearch.templ`, Line: 168, Col: 45}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatUint(uint64(product.ID), 10))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductSearch.templ`, Line: 169, Col: 75}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 templ.SafeURL = templ.URL(fmt.Sprintf("/admin/locateProduct/%d", product.ID))
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var10)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><button class=\"btn btn-secondary btn-sm me-2\">Locate</button></a> <a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 templ.SafeURL = templ.URL(fmt.Sprintf("/admin/updateProduct/%d", product.ID))
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var11)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><button class=\"btn btn-primary btn-sm me-2\"><i class=\"bi bi-shield-lock me-2\"></i>Update</button></a></td></tr>")
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