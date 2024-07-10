# Invopop Alpine Kit

This repository contains a collection of reusable components for Alpine.js and a centralized stylesheet for common elements such as inputs and buttons.

## Getting Started

### Prerequisites

- [Alpine.js v3](https://github.com/alpinejs/alpine)

#### 1. Include Alpine.js

Make sure to include Alpine.js in your module. You can include it from a CDN.

```html
<script src="https://cdn.jsdelivr.net/npm/alpinejs@v3"></script>
```

#### 2. Include common CSS stylesheet

Include the stylesheet for common elements, such as inputs and buttons.

```html
<link
  rel="stylesheet"
  href="https://cdn.jsdelivr.net/npm/@invopop/alpine-kit@0.0.28/style.css"
/>
```

#### 3. Include Inter font from Google Fonts

```html
<link rel="preconnect" href="https://fonts.googleapis.com" />
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
<link
  href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600&display=swap"
  rel="stylesheet"
/>
```

## Usage

### Components

Browse through the components directory to find reusable Alpine.js components. Each component is self-contained, including logic and styles, and can be easily integrated into your module. Just copy and paste the html content. Make sure to replace any reference to placeholders to actual implementation.

```html
<!-- image-picker-html -->
<div class="ui-image-picker">
  <img x-show="config.logo_url" :src="config.logo_url" />
  <div x-show="!config.logo_url" class="ui-logo-placeholder"></div>
  <span x-text="loading ? 'Uploading...' : 'Select logo'"></span>
  <button
    class="ui-save-button ui-button-small"
    @click="selectImage"
    :disabled="loading"
  >
    Upload
  </button>
  <button
    class="ui-secondary-button ui-button-small"
    @click="deleteImage"
    :disabled="loading || !config.logo_url"
  >
    Remove
  </button>
</div>
```

![image](https://github.com/invopop/alpine-kit/assets/12644599/475d4bc2-0d44-418d-9b76-e159bc034f91)

### Common Styles

The global stylesheet provides styling for common elements like inputs and buttons and complex component described above. Simply include it in your project, and use the styles accordingly.

![image](https://github.com/invopop/alpine-kit/assets/12644599/fca008f2-538c-4817-ad7a-e99dea459dc7)

#### Code examples of common HTML elements

[Text Input](https://codepen.io/javiermartinez/pen/ZEdYXLr)
[Numeric Input](https://codepen.io/javiermartinez/pen/bGPNoEe)
[Select](https://codepen.io/javiermartinez/pen/mdZyByN)
[Select with button](https://codepen.io/javiermartinez/pen/XWwvaxy)
[Checkbox](https://codepen.io/javiermartinez/pen/NWZPapd)
[Default Button](https://codepen.io/javiermartinez/pen/bGPNorg)
[Save Button](https://codepen.io/javiermartinez/pen/rNEaGGK)
[Small Button](https://codepen.io/javiermartinez/pen/MWMYEOQ)
[Tags](https://codepen.io/javiermartinez/pen/mdZyBwy)
[Horizontal Separator](https://codepen.io/javiermartinez/pen/LYKEzeo)
[Section Title](https://codepen.io/javiermartinez/pen/jOjEGzG)
[Section Title with button](https://codepen.io/javiermartinez/pen/vYqEerO)
[Helper label](https://codepen.io/javiermartinez/pen/RwzNLyW)
[Paragraph with highlighted words](https://codepen.io/javiermartinez/pen/dyBPVjb)
[Table](https://codepen.io/javiermartinez/pen/MWMYEBG)
[Code Snippet](https://codepen.io/javiermartinez/pen/YzoPrJX)
