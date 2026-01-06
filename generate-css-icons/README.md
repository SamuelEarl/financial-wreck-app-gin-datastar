# Generate CSS Icons

IMPORTANT: You need to have Node.js installed on your computer in order for this module to work.

You have to update the `styles/_icons.scss` file with any new icons that you want to use.

This is how you update the `styles/_icons.scss` file:

1. Add the icon set and icon name to the `icons` object in the `generate-css-icons.js` file (in the root directory).
2. Run `node generate-css-icons.js` or `make generate-css-icons` from the root directory. This will overwrite the `styles/_icons.scss` file with the latest CSS code.
