// IMPORTANT: You need to have Node.js installed on your computer in order for this module to work.

// The code for this module and the instructions for using CSS icons are found here:
// https://iconify.design/docs/usage/css/utils/

// To update the CSS icons, run `node generate-css-icons/index.js` from the root directory or `make generate-css-icons`. This will update the `../styles/_icons.scss` file.

// To use those icons in HTML, pass the `icon-set:icon-name` string (from Iconify) to the custom <Icon /> component in the `src/components` folder.
// You can also use individual <span> elements with two class names: class name for the icon set, class name for the icon.
// <span class="icon--mdi-light icon--mdi-light--alert-circle"></span>

import { readFile, writeFile } from "node:fs/promises";
import { getIconsCSS } from "@iconify/utils";
import { locate } from "@iconify/json";

(async function () {
  console.log("Generating CSS icons...");

  /**
   * List of icons. Key is icon set prefix, value is array of icons
   *
   * @type {Record<string, string[]>}
   */
  const icons = {
    "material-symbols": [
      "search",
      "person-outline",
      "shopping-bag-outline-sharp",
      "arrow-cool-down",
      "menu",
      "close",
      "chevron-left",
      "chevron-right",
      "save-outline-sharp",
    ],
    "gg": [
      "spinner-two-alt",
    ],
  };

  // Parse each icon set
  let code = "";
  for (const prefix in icons) {
    // Find location of .json file
    const filename = locate(prefix);

    // Load file and parse it
    /** @type {import("@iconify/types").IconifyJSON} */
    const iconSet = JSON.parse(await readFile(filename, "utf8"));

    // Get CSS
    const css = getIconsCSS(iconSet, icons[prefix]);

    // Add it to code
    code += css;
  }

  // Save CSS file
  await writeFile("../styles/_icons.scss", code, "utf8");
  console.log(`styles/_icons.scss file has been updated! ${code.length} bytes saved.`);
})();
