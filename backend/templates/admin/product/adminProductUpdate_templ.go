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

func AdminUpdateProduct(product models.Product, manufacturers []models.Manufacturer) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Update Product</title><link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css\" rel=\"stylesheet\"><link href=\"https://cdnjs.cloudflare.com/ajax/libs/bootstrap-icons/1.8.1/font/bootstrap-icons.min.css\" rel=\"stylesheet\"><style>\n        body, html {\n            height: 100%;\n            margin: 0;\n            overflow: hidden;\n        }\n        .navbar {\n            position: fixed;\n            top: 0;\n            width: 100%;\n            z-index: 1000;\n            background-color: #be274b;\n            padding: 10px 0;\n        }\n        .navbar .nav-link,\n        .navbar .navbar-icon {\n            color: #e2e6ea;\n            padding: 10px 15px;\n            font-weight: bold;\n            text-align: center;\n            display: flex;\n            flex-direction: column;\n            line-height: 1.2;\n            position: relative;\n        }\n        .navbar-icon {\n            font-size: 1.5rem;\n            cursor: default;\n        }\n        .nav-link[onclick] {\n            cursor: pointer;\n            user-select: none;\n        }\n        .d-flex {\n            align-items: center;\n        }\n        .navbar .nav-link:hover {\n            background-color: #da3060;\n            border-radius: 5px;\n            transition: background-color 0.3s ease;\n        }\n        .content-container {\n            display: flex;\n            justify-content: center;\n            align-items: center;\n            min-height: 100vh;\n            padding-top: 78px;\n        }\n        .justify-between {\n            display: flex;\n            align-items: center;\n            justify-content: space-between;\n            width: 100%;\n        }\n        .submenu {\n            display: none;\n            position: absolute;\n            background-color: #f8f9fa;\n            border-radius: 5px;\n            padding: 10px;\n            top: 100%;\n            left: 50%;\n            transform: translateX(-50%);\n            margin-top: 5px;\n            z-index: 1001;\n        }\n        .submenu a {\n            color: #333;\n            text-decoration: none;\n            display: block;\n            padding: 5px 15px;\n        }\n        .submenu a:hover {\n            background-color: #ddd;\n            border-radius: 5px;\n        }\n        .setup-container {\n            height: calc(100% + 78px);\n            display: flex;\n            justify-content: center;\n            align-items: center;\n        }\n        .setup-form {\n            width: 100%;\n            max-width: 800px;\n            padding: 15px;\n            border: 1px solid #ccc;\n            border-radius: 10px;\n            box-shadow: 0px 0px 10px rgba(0,0,0,0.1);\n        }\n        .back-arrow {\n            font-size: 1.5rem;\n            color: gray;\n            text-decoration: none;\n        }\n        .error-text {\n            color: red;\n            font-size: 0.875em;\n            margin-top: 5px;\n        }\n    </style></head><body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templates.AdminNav().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container setup-container\"><div class=\"setup-form\"><div class=\"d-flex align-items-center mb-4\"><a href=\"SearchProducts.html\" class=\"back-arrow position-absolute\"><i class=\"bi bi-arrow-left-circle\"></i></a><h2 class=\"text-center w-100\">Update Product</h2></div><form method=\"post\"><input type=\"hidden\" id=\"actionType\" name=\"actionType\" value=\"update\"><div class=\"row\"><div class=\"col-md-6\"><div class=\"mb-3\"><label for=\"description\" class=\"form-label\">Description</label> <input type=\"text\" class=\"form-control\" id=\"description\" name=\"description\" placeholder=\"Description\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(product.Description)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductUpdate.templ`, Line: 136, Col: 156}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" required><div class=\"error-text\">Optional description error</div></div><div class=\"mb-3\"><label for=\"manufacturerSelect\" class=\"form-label\">Select Manufacturer</label> <select class=\"form-select\" id=\"manufacturerSelect\" name=\"manufacturerSelect\" required>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, manufacturer := range manufacturers {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatUint(uint64(manufacturer.ID), 10))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductUpdate.templ`, Line: 143, Col: 94}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if manufacturer.ID == product.Manufacturer.ID {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" selected")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(manufacturer.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductUpdate.templ`, Line: 147, Col: 47}
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select><div class=\"error-text\">Optional manufacturer error</div></div><div class=\"mb-3\"><label for=\"batch\" class=\"form-label\">Batch</label> <input type=\"text\" class=\"form-control\" id=\"batch\" name=\"batch\" placeholder=\"Batch\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(product.Batch)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductUpdate.templ`, Line: 154, Col: 132}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" required><div class=\"error-text\">Optional batch error</div></div></div><div class=\"col-md-6\"><div class=\"mb-3\"><div class=\"d-flex\"><div class=\"me-2\"><label for=\"dimX\" class=\"form-label\">Width</label> <input type=\"number\" class=\"form-control\" id=\"dimX\" name=\"dimX\" placeholder=\"0\" min=\"0\" step=\"1\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatUint(uint64(product.XDimension), 10))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductUpdate.templ`, Line: 163, Col: 190}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" required oninput=\"this.value = this.value.replace(/[^0-9]/g, &#39;&#39;)\"></div><div class=\"me-2\"><label for=\"dimY\" class=\"form-label\">Height</label> <input type=\"number\" class=\"form-control\" id=\"dimY\" name=\"dimY\" placeholder=\"0\" min=\"0\" step=\"1\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatUint(uint64(product.YDimension), 10))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductUpdate.templ`, Line: 167, Col: 190}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" required oninput=\"this.value = this.value.replace(/[^0-9]/g, &#39;&#39;)\"></div><div><label for=\"dimZ\" class=\"form-label\">Depth</label> <input type=\"number\" class=\"form-control\" id=\"dimZ\" name=\"dimZ\" placeholder=\"0\" min=\"0\" step=\"1\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatUint(uint64(product.ZDimension), 10))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductUpdate.templ`, Line: 171, Col: 190}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" required oninput=\"this.value = this.value.replace(/[^0-9]/g, &#39;&#39;)\"></div></div><div class=\"error-text\">Optional dimension error</div></div><div class=\"mb-3\"><label for=\"unit\" class=\"form-label\">Unit</label> <input type=\"text\" class=\"form-control\" id=\"unit\" name=\"unit\" placeholder=\"Unit\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(product.Unit)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/admin/product/adminProductUpdate.templ`, Line: 178, Col: 128}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" required><div class=\"error-text\">Optional unit error</div></div></div></div><!-- TODO??? --><!-- If product is not stored yet --><div class=\"d-flex justify-content-between mt-3\"><button type=\"submit\" class=\"btn btn-primary w-100 me-2\"><i class=\"bi bi-shield-lock me-2\"></i>Update</button> <button type=\"button\" class=\"btn btn-danger w-100\" onclick=\"handleDelete()\"><i class=\"bi bi-shield-lock me-2\"></i>Delete</button></div><!-- If product is already stored (no delete) --><!-- <div class=\"d-flex justify-content-center\">\n                    <div style=\"max-width: 500px; width: 100%;\">\n                        <button type=\"submit\" class=\"btn btn-primary w-100 mt-3\">\n                            <i class=\"bi bi-shield-lock me-2\"></i>Update\n                        </button>\n                    </div>\n                </div> --></form></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
