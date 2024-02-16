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
  href="https://cdn.jsdelivr.net/npm/@invopop/alpine-kit/style.css"
/>
```

#### 3. Include Alpine component plugin (Optional)

If you are going to make use of a reusable component (See below) you need to add the Alpine component plugin.

```html
<script src="https://unpkg.com/alpinejs-component@latest/dist/component.min.js"></script>
```

## Usage

### Components

Browse through the components directory to find reusable Alpine.js components. Each component is self-contained, including logic and styles, and can be easily integrated into your module. Just copy and paste the html file relative to your index.html file and use it like this

```html
<x-component
  url="/path/to-the-file/image-picker.html"
  x-data="{ param: 'Whatever param your component may need' }"
  @delete="console.log('listening to component emitted events')"
></x-component>
```

### Common Styles

The global stylesheet provides styling for common elements like inputs and buttons. Simply include it in your project, and use the styles accordingly.

```html
<!-- ui-config-container class will apply some margins and paddings to the container -->
<div x-data="component" class="ui-config-container">
  <!-- ui-input class will apply styles for inputs and selects -->
  <div class="ui-input">
    <label>
      <span>Logo height:</span>
      <input
        type="number"
        x-model.number="config.logo_height"
        min="10"
        max="50"
      />
    </label>
  </div>

  <!-- ui-save-button will apply styles to the submit button for the form -->
  <button class="ui-save-button" @click="saveConfig">Save</button>
</div>
```
