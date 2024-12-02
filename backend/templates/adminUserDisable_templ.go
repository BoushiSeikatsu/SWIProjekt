// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "swi-warehouse/models"

func AdminDisableUser(users []models.User) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Disable User</title><link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css\" rel=\"stylesheet\"><link href=\"https://cdnjs.cloudflare.com/ajax/libs/bootstrap-icons/1.8.1/font/bootstrap-icons.min.css\" rel=\"stylesheet\"><style>\r\n        body, html {\r\n            height: 100%;\r\n            margin: 0;\r\n            overflow: hidden;\r\n        }\r\n        .navbar {\r\n            position: fixed;\r\n            top: 0;\r\n            width: 100%;\r\n            z-index: 1000;\r\n            background-color: #be274b;\r\n            padding: 10px 0;\r\n        }\r\n        .navbar .nav-link,\r\n        .navbar .navbar-icon {\r\n            color: #e2e6ea;\r\n            padding: 10px 15px;\r\n            font-weight: bold;\r\n            text-align: center;\r\n            display: flex;\r\n            flex-direction: column;\r\n            line-height: 1.2;\r\n            position: relative;\r\n        }\r\n        .navbar-icon {\r\n            font-size: 1.5rem;\r\n            cursor: default;\r\n        }\r\n        .nav-link[onclick] {\r\n            cursor: pointer;\r\n            user-select: none;\r\n        }\r\n        .d-flex {\r\n            align-items: center;\r\n        }\r\n        .navbar .nav-link:hover {\r\n            background-color: #da3060;\r\n            border-radius: 5px;\r\n            transition: background-color 0.3s ease;\r\n        }\r\n        .content-container {\r\n            display: flex;\r\n            justify-content: center;\r\n            align-items: center;\r\n            min-height: 100vh;\r\n            padding-top: 78px;\r\n        }\r\n        .justify-between {\r\n            display: flex;\r\n            align-items: center;\r\n            justify-content: space-between;\r\n            width: 100%;\r\n        }\r\n        .submenu {\r\n            display: none;\r\n            position: absolute;\r\n            background-color: #f8f9fa;\r\n            border-radius: 5px;\r\n            padding: 10px;\r\n            top: 100%;\r\n            left: 50%;\r\n            transform: translateX(-50%);\r\n            margin-top: 5px;\r\n            z-index: 1001;\r\n        }\r\n        .submenu a {\r\n            color: #333;\r\n            text-decoration: none;\r\n            display: block;\r\n            padding: 5px 15px;\r\n        }\r\n        .submenu a:hover {\r\n            background-color: #ddd;\r\n            border-radius: 5px;\r\n        }\r\n        .setup-container {\r\n            height: calc(100% + 78px);\r\n            display: flex;\r\n            justify-content: center;\r\n            align-items: center;\r\n        }\r\n        .setup-form {\r\n            width: 100%;\r\n            max-width: 400px;\r\n            padding: 15px;\r\n            border: 1px solid #ccc;\r\n            border-radius: 10px;\r\n            box-shadow: 0px 0px 10px rgba(0,0,0,0.1);\r\n        }\r\n        .error-text {\r\n            color: red;\r\n            font-size: 0.875em;\r\n            margin-top: 5px;\r\n        }\r\n    </style></head><body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = AdminNav().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container setup-container\"><div class=\"setup-form\"><h2 class=\"text-center mb-4\">Disable User</h2><form method=\"post\"><div class=\"mb-3\"><label for=\"userSelect\" class=\"form-label\">Select User</label> <select class=\"form-select\" id=\"userSelect\" name=\"userSelect\" required><option value=\"\" disabled selected>Select a user</option> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, user := range users {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(user.Username)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/adminUserDisable.templ`, Line: 123, Col: 56}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(user.Username)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/adminUserDisable.templ`, Line: 123, Col: 72}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select><div class=\"error-text\">Optional user error</div></div><button type=\"submit\" class=\"btn btn-primary w-100\"><i class=\"bi bi-shield-lock me-2\"></i>Disable User</button></form></div></div><script>\r\n        function toggleSubmenu(id) {\r\n            const submenu = document.getElementById(id);\r\n            const isVisible = submenu.style.display === 'block';\r\n            document.querySelectorAll('.submenu').forEach(menu => menu.style.display = 'none');\r\n            submenu.style.display = isVisible ? 'none' : 'block';\r\n        }\r\n        document.addEventListener('click', function(e) {\r\n            const isClickInside = e.target.closest('.nav-link') || e.target.closest('.submenu');\r\n            if (!isClickInside) {\r\n                document.querySelectorAll('.submenu').forEach(menu => menu.style.display = 'none');\r\n            }\r\n        });\r\n    </script><script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
