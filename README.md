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
  href="https://cdn.jsdelivr.net/npm/@invopop/alpine-kit@0.0.11/style.css"
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
  <img :src="config.logo_url || 'https://placehold.co/40x40?text=Logo'" />
  <span x-text="loading ? 'Uploading...' : 'Select logo...'"></span>
  <button @click="selectImage" :disabled="loading">
    <svg
      width="20"
      height="20"
      viewBox="0 0 20 20"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        id="Icon"
        d="M11.1667 5.33333L12.4393 4.06066C13.0251 3.47487 13.9749 3.47487 14.5607 4.06066L15.9393 5.43934C16.5251 6.02513 16.5251 6.97487 15.9393 7.56066L14.6667 8.83333M11.1667 5.33333L3.29289 13.2071C3.10536 13.3946 3 13.649 3 13.9142V16.5C3 16.7761 3.22386 17 3.5 17H6.08579C6.351 17 6.60536 16.8946 6.79289 16.7071L14.6667 8.83333M11.1667 5.33333L14.6667 8.83333"
        stroke="#0A0A0A"
        stroke-width="1.1"
        stroke-linecap="round"
      />
    </svg>
  </button>
  <button @click="deleteImage" :disabled="loading">
    <svg
      width="20"
      height="20"
      viewBox="0 0 20 20"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        id="Icon"
        d="M4.83455 15.9828L5.38341 15.9474L4.83455 15.9828ZM15.1655 15.9828L14.6166 15.9474L15.1655 15.9828ZM2.5 4.38243H1.95V5.48243H2.5V4.38243ZM17.5 5.48243H18.05V4.38243H17.5V5.48243ZM8.72568 8.98649V8.43649H7.62568V8.98649H8.72568ZM7.62568 13.4459V13.9959H8.72568V13.4459H7.62568ZM12.3743 8.98649V8.43649H11.2743V8.98649H12.3743ZM11.2743 13.4459V13.9959H12.3743V13.4459H11.2743ZM12.6084 5.06952C12.6842 5.36369 12.984 5.54079 13.2782 5.46507C13.5723 5.38936 13.7494 5.08951 13.6737 4.79534L12.6084 5.06952ZM6.32631 4.79534C6.25059 5.08951 6.42769 5.38936 6.72186 5.46507C7.01603 5.54079 7.31588 5.36369 7.39159 5.06952L6.32631 4.79534ZM3.57276 4.96784L4.28569 16.0182L5.38341 15.9474L4.67048 4.89702L3.57276 4.96784ZM6.4528 18.05H13.5472V16.95H6.4528V18.05ZM15.7143 16.0182L16.4272 4.96784L15.3295 4.89702L14.6166 15.9474L15.7143 16.0182ZM15.8784 4.38243H4.12162V5.48243H15.8784V4.38243ZM2.5 5.48243H4.12162V4.38243H2.5V5.48243ZM15.8784 5.48243H17.5V4.38243H15.8784V5.48243ZM13.5472 18.05C14.6923 18.05 15.6406 17.1609 15.7143 16.0182L14.6166 15.9474C14.5802 16.5113 14.1123 16.95 13.5472 16.95V18.05ZM4.28569 16.0182C4.35941 17.1609 5.30773 18.05 6.4528 18.05V16.95C5.88775 16.95 5.41979 16.5113 5.38341 15.9474L4.28569 16.0182ZM7.62568 8.98649V13.4459H8.72568V8.98649H7.62568ZM11.2743 8.98649V13.4459H12.3743V8.98649H11.2743ZM10 3.05C11.2542 3.05 12.3094 3.90779 12.6084 5.06952L13.6737 4.79534C13.2527 3.15941 11.7683 1.95 10 1.95V3.05ZM7.39159 5.06952C7.6906 3.90779 8.74583 3.05 10 3.05V1.95C8.23176 1.95 6.74737 3.15941 6.32631 4.79534L7.39159 5.06952Z"
        fill="#EC4E46"
      />
    </svg>
  </button>
</div>

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
```

![image](https://github.com/invopop/alpine-kit/assets/12644599/26e26461-644c-4b89-b3a0-018d6447beec)

### Common Styles

The global stylesheet provides styling for common elements like inputs and buttons and complex component described above. Simply include it in your project, and use the styles accordingly.

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

  <!-- ui-checkbox class will apply styles for checkboxes -->
  <div class="ui-checkbox">
    <input
      type="checkbox"
      x-model="config.legal_terms_accepted"
      id="legalTermsAccepted"
    />
    <label for="legalTermsAccepted"> Accept Legal Terms and Conditions </label>
  </div>

  <!-- ui-tags will apply styles to a list of tags -->
  <ul class="ui-tags">
    <template x-for="(tag, idx) in config.tags" :key="idx">
      <li>
        <span x-text="tag"></span>
        <button data-tooltip="remove tag" @click="remTag(idx)">x</button>
      </li>
    </template>
  </ul>

  <!-- ui-save-button will apply styles to the submit button for the form -->
  <button class="ui-save-button" @click="saveConfig">Save</button>
</div>
```

![image](https://github.com/invopop/alpine-kit/assets/12644599/0c36dd98-894f-4352-8304-f259de5c91f9)
