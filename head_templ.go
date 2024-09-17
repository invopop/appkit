// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package appkit

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func DefaultCSS() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Inter font --><link rel=\"preconnect\" href=\"https://fonts.googleapis.com\"><link rel=\"preconnect\" href=\"https://fonts.gstatic.com\" crossorigin><link href=\"https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600&amp;display=swap\" rel=\"stylesheet\"><style>\n    [x-cloak] { display: none !important; }\n    .ui-select {\n        appearance: none;\n        -webkit-appearance: none;\n        -moz-appearance: none;\n        background-image: url('data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAiIGhlaWdodD0iMjAiIHZpZXdCb3g9IjAgMCAyMCAyMCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHJlY3QgeD0iMiIgeT0iMiIgd2lkdGg9IjE2IiBoZWlnaHQ9IjE2IiByeD0iNCIgZmlsbD0iI0YzRjRGNiIvPgo8cGF0aCBkPSJNNi41IDguMjUwMDRMMTAgMTEuNzVMMTMuNSA4LjI1IiBzdHJva2U9IiMwMzA3MTIiIHN0cm9rZS13aWR0aD0iMS4xIi8+Cjwvc3ZnPg==');\n        background-repeat: no-repeat;\n        background-position: center right 6px;\n    }\n\n    .ui-button:hover  {\n        box-shadow: inset 0 0 0 1000px rgba(0, 0, 0, 0.16);\n    }\n\n    .ui-button:active {\n        box-shadow: inset 0 0 0 1000px rgba(0, 0, 0, 0.32);\n    }\n\n    div[contenteditable=\"true\"] .tag {\n        display: inline-flex;\n        padding: 2px 4px;\n        justify-content: center;\n        align-items: center;\n        gap: 4px;\n        border-radius: 4px;\n        border: 1px solid #E5E7EB;\n        background: #F9FAFB;\n        pointer-events: none;\n        color: #030712;\n        font-family: Inter;\n        font-size: 12px;\n        font-style: normal;\n        font-weight: 500;\n        line-height: 16px;\n    }\n    [popover] {\n        inset: unset;\n        width: auto;\n        height: auto;\n        padding: 0;\n        background: transparent;\n        border: none;\n        position: fixed;\n        top: 0px;\n        right: 0px;\n        pointer-events: none;\n\n        overlay: none;\n    }\n\n    [popover]:popover-open + .popover-menu {\n        display: block;\n    }\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func DefaultScripts() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n        const CONSOLE_SDK_URL = 'https://cdn.jsdelivr.net/npm/@invopop/console-ui-sdk@0.0.4/index.js'\n    </script><script src=\"https://cdn.jsdelivr.net/npm/alpinejs/dist/cdn.min.js\" defer></script><script type=\"module\">\n        import { install } from 'https://esm.run/@twind/core';\n        import presetAutoprefix from 'https://esm.run/@twind/preset-autoprefix';\n        import presetTailwind from 'https://esm.run/@twind/preset-tailwind';\n        import presetTypography from 'https://esm.run/@twind/preset-typography';\n        import theme from 'https://cdn.jsdelivr.net/npm/@invopop/ui-kit/dist/tw.theme.js'\n        install({\n            hash: false,\n            presets: [presetAutoprefix(), presetTailwind(), presetTypography()],\n            theme\n        })\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
