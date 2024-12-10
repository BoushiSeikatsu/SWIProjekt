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

func AdminSelectSupplyProduct(manufacturer models.Manufacturer, products []models.Product, storages []models.Storage) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Supply Product</title><link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css\" rel=\"stylesheet\"><link href=\"https://cdnjs.cloudflare.com/ajax/libs/bootstrap-icons/1.8.1/font/bootstrap-icons.min.css\" rel=\"stylesheet\"><style>\n        body, html {\n            height: 100%;\n            margin: 0;\n            overflow: hidden;\n        }\n        .navbar {\n            position: fixed;\n            top: 0;\n            width: 100%;\n            z-index: 1000;\n            background-color: #be274b;\n            padding: 10px 0;\n        }\n        .navbar .nav-link,\n        .navbar .navbar-icon {\n            color: #e2e6ea;\n            padding: 10px 15px;\n            font-weight: bold;\n            text-align: center;\n            display: flex;\n            flex-direction: column;\n            line-height: 1.2;\n            position: relative;\n        }\n        .navbar-icon {\n            font-size: 1.5rem;\n            cursor: default;\n        }\n        .nav-link[onclick] {\n            cursor: pointer;\n            user-select: none;\n        }\n        .d-flex {\n            align-items: center;\n        }\n        .navbar .nav-link:hover {\n            background-color: #da3060;\n            border-radius: 5px;\n            transition: background-color 0.3s ease;\n        }\n        .content-container {\n            display: flex;\n            justify-content: center;\n            align-items: center;\n            min-height: 100vh;\n            padding-top: 78px;\n        }\n        .justify-between {\n            display: flex;\n            align-items: center;\n            justify-content: space-between;\n            width: 100%;\n        }\n        .submenu {\n            display: none;\n            position: absolute;\n            background-color: #f8f9fa;\n            border-radius: 5px;\n            padding: 10px;\n            top: 100%;\n            left: 50%;\n            transform: translateX(-50%);\n            margin-top: 5px;\n            z-index: 1001;\n        }\n        .submenu a {\n            color: #333;\n            text-decoration: none;\n            display: block;\n            padding: 5px 15px;\n        }\n        .submenu a:hover {\n            background-color: #ddd;\n            border-radius: 5px;\n        }\n        .setup-container {\n            height: calc(100% + 78px);\n            display: flex;\n            justify-content: center;\n            align-items: center;\n        }\n        .setup-form {\n            width: 100%;\n            max-width: 400px;\n            padding: 15px;\n            border: 1px solid #ccc;\n            border-radius: 10px;\n            box-shadow: 0px 0px 10px rgba(0,0,0,0.1);\n        }\n        .back-arrow {\n            font-size: 1.5rem;\n            color: gray;\n            text-decoration: none;\n        }\n        .error-text {\n            color: red;\n            font-size: 0.875em;\n            margin-top: 5px;\n        }\n    </style></head><body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templates.AdminNav().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container setup-container\"><div class=\"setup-form\"><div class=\"d-flex align-items-center mb-4\"><a href=\"SelectSupplyManufacturer.html\" class=\"back-arrow position-absolute\"><i class=\"bi bi-arrow-left-circle\"></i></a><h2 class=\"text-center w-100\">Supply Product</h2></div><form method=\"post\"><div class=\"mb-3\"><label for=\"manufacturerSelect\" class=\"form-label\">Manufacturer</label> <select class=\"form-select\" id=\"manufacturerSelect\" disabled><option selected>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(manufacturer.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductSupply.templ`, Line: 133, Col: 59}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option></select></div><div class=\"mb-3\"><label for=\"productSelect\" class=\"form-label\">Select Product</label> <select class=\"form-select\" id=\"productSelect\" name=\"productSelect\" required><option value=\"\" disabled selected>Select Product</option> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, product := range products {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatUint(uint64(product.ID), 10))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductSupply.templ`, Line: 141, Col: 85}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(product.Description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductSupply.templ`, Line: 141, Col: 107}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select><div class=\"error-text\">Optional product error</div></div><div class=\"mb-3\"><label for=\"storageSelect\" class=\"form-label\">Select Storage</label> <select class=\"form-select\" id=\"storageSelect\" name=\"storageSelect\" required><option value=\"\" disabled selected>Select Storage</option> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, storage := range storages {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatUint(uint64(storage.ID), 10))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductSupply.templ`, Line: 151, Col: 85}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(storage.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductSupply.templ`, Line: 151, Col: 100}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select><div class=\"error-text\">Optional storage error</div></div><div class=\"mb-3\"><label for=\"amount\" class=\"form-label\">Amount</label><div class=\"input-group\"><input type=\"number\" class=\"form-control\" id=\"amount\" name=\"amount\" min=\"0\" step=\"1\" value=\"0\" required oninput=\"this.value = this.value.replace(/[^0-9]/g, &#39;&#39;)\"></div><div class=\"error-text\">Optional amount error</div></div><button type=\"submit\" class=\"btn btn-primary w-100\">Supply</button></form></div></div><script>\n        function toggleSubmenu(id) {\n            const submenu = document.getElementById(id);\n            const isVisible = submenu.style.display === 'block';\n            document.querySelectorAll('.submenu').forEach(menu => menu.style.display = 'none');\n            submenu.style.display = isVisible ? 'none' : 'block';\n        }\n        document.addEventListener('click', function(e) {\n            const isClickInside = e.target.closest('.nav-link') || e.target.closest('.submenu');\n            if (!isClickInside) {\n                document.querySelectorAll('.submenu').forEach(menu => menu.style.display = 'none');\n            }\n        });\n    </script><script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
