// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package templatesAdminManufacturer

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"swi-warehouse/templates"
)

func AdminAddManufacturer() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Add Manufacturer</title><link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css\" rel=\"stylesheet\"><link href=\"https://cdnjs.cloudflare.com/ajax/libs/bootstrap-icons/1.8.1/font/bootstrap-icons.min.css\" rel=\"stylesheet\"><style>\n        body, html {\n            height: 100%;\n            margin: 0;\n            overflow: hidden;\n        }\n        .navbar {\n            position: fixed;\n            top: 0;\n            width: 100%;\n            z-index: 1000;\n            background-color: #be274b;\n            padding: 10px 0;\n        }\n        .navbar .nav-link,\n        .navbar .navbar-icon {\n            color: #e2e6ea;\n            padding: 10px 15px;\n            font-weight: bold;\n            text-align: center;\n            display: flex;\n            flex-direction: column;\n            line-height: 1.2;\n            position: relative;\n        }\n        .navbar-icon {\n            font-size: 1.5rem;\n            cursor: default;\n        }\n        .nav-link[onclick] {\n            cursor: pointer;\n            user-select: none;\n        }\n        .d-flex {\n            align-items: center;\n        }\n        .navbar .nav-link:hover {\n            background-color: #da3060;\n            border-radius: 5px;\n            transition: background-color 0.3s ease;\n        }\n        .content-container {\n            display: flex;\n            justify-content: center;\n            align-items: center;\n            min-height: 100vh;\n            padding-top: 78px;\n        }\n        .justify-between {\n            display: flex;\n            align-items: center;\n            justify-content: space-between;\n            width: 100%;\n        }\n        .submenu {\n            display: none;\n            position: absolute;\n            background-color: #f8f9fa;\n            border-radius: 5px;\n            padding: 10px;\n            top: 100%;\n            left: 50%;\n            transform: translateX(-50%);\n            margin-top: 5px;\n            z-index: 1001;\n        }\n        .submenu a {\n            color: #333;\n            text-decoration: none;\n            display: block;\n            padding: 5px 15px;\n        }\n        .submenu a:hover {\n            background-color: #ddd;\n            border-radius: 5px;\n        }\n        .add-manufacturer-container {\n            height: calc(100% + 78px);\n            display: flex;\n            justify-content: center;\n            align-items: center;\n        }\n        .add-manufacturer-form {\n            width: 100%;\n            max-width: 400px;\n            padding: 15px;\n            border: 1px solid #ccc;\n            border-radius: 10px;\n            box-shadow: 0px 0px 10px rgba(0,0,0,0.1);\n        }\n        .error-text {\n            color: red;\n            font-size: 0.875em;\n            margin-top: 5px;\n        }\n    </style></head><body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templates.AdminNav().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container add-manufacturer-container\"><div class=\"add-manufacturer-form\"><h2 class=\"text-center mb-4\">Add Manufacturer</h2><form method=\"post\"><div class=\"mb-3\"><label for=\"name\" class=\"form-label\">Manufacturer Name</label> <input type=\"text\" class=\"form-control\" id=\"name\" name=\"name\" placeholder=\"Manufacturer Name\" required><div class=\"error-text\">Optional manufacturer name error</div></div><button type=\"submit\" class=\"btn btn-primary w-100\"><i class=\"bi bi-shield-lock me-2\"></i>Add</button></form></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templates.AdminScript().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
