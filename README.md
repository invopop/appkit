# Invopop Appkit

This repository contains a collection of reusable Templ components and a centralized stylesheet for common elements such as inputs and buttons.

## Development

### Generate Templates

Appkit uses [templ](https://templ.guide/) to define a set of components in Go. To generate the templates, run:

```bash
templ generate
```

During development, it can help massive to have hot reload to be able to make changes and see them quickly. We're currently using Air for this:

#### Air

Air is a great tool to auto reload potentially any project, but works great with Go. Install with:

```bash
go install github.com/air-verse/air@latest
```

The `.toml` is already configured and ready in this repository, so simply run:

```bash
air
```

Air is a bit more reliable at detecting file changes, especially for stylesheets. You'll always need to wait a few seconds before page reloads to give the system chance to recompile. A proxy is available with Air, but we didn't find it to be very reliable and was breaking with query parameters, it obvously also wouldn't work for PDF reloads.

#### Code examples of common HTML elements

- [Text Input](https://codepen.io/javiermartinez/pen/ZEdYXLr)
- [Numeric Input](https://codepen.io/javiermartinez/pen/bGPNoEe)
- [Select](https://codepen.io/javiermartinez/pen/mdZyByN)
- [Select with button](https://codepen.io/javiermartinez/pen/XWwvaxy)
- [Checkbox](https://codepen.io/javiermartinez/pen/NWZPapd)
- [Default Button](https://codepen.io/javiermartinez/pen/bGPNorg)
- [Save Button](https://codepen.io/javiermartinez/pen/rNEaGGK)
- [Small Button](https://codepen.io/javiermartinez/pen/MWMYEOQ)
- [Tags](https://codepen.io/javiermartinez/pen/mdZyBwy)
- [Horizontal Separator](https://codepen.io/javiermartinez/pen/LYKEzeo)
- [Section Title](https://codepen.io/javiermartinez/pen/jOjEGzG)
- [Section Title with button](https://codepen.io/javiermartinez/pen/vYqEerO)
- [Label for optional fields](https://codepen.io/javiermartinez/pen/BagyvyG)
- [Helper label](https://codepen.io/javiermartinez/pen/RwzNLyW)
- [Info Paragraph](https://codepen.io/javiermartinez/pen/MWMYZWW)
- [Paragraph with highlighted words](https://codepen.io/javiermartinez/pen/dyBPVjb)
- [Table](https://codepen.io/javiermartinez/pen/MWMYEBG)
- [Code Snippet](https://codepen.io/javiermartinez/pen/YzoPrJX)
