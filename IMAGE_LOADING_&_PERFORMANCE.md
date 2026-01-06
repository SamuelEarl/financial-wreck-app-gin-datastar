# IMAGE LOADING & PERFORMANCE NOTES

**NOTE: I have Chrome bookmarks (in my "Image Editing" bookmark folder) for videos that explain how to optimize images for the web using GIMP and WebP images.**

## 4K Images

A 4K image is an image with a resolution of approximately 4,000 pixels horizontally, most commonly specified as 3840 x 2160 pixels. This resolution contains about 8.3 million pixels, which is four times the number of pixels in a Full HD image, resulting in sharper, more detailed visuals. The "4K" term refers to the horizontal resolution of around 4,000 pixels, making it the standard for high-definition displays like TVs, monitors, and cameras. 

### Key characteristics of 4K images

* High pixel count: A 4K image has a resolution of 3840 x 2160 pixels, which equates to a total of 8,294,400 pixels.
* Increased detail: The higher number of pixels creates a crisper, more detailed, and more vivid picture compared to lower-resolution images.
* Four times Full HD: The total pixel count of a 4K image is four times that of a Full HD (1920 x 1080) image.
* Consumer vs. professional: While the common consumer standard is 3840 x 2160, the professional cinema standard for "4K" is slightly different at 4096 x 2160 pixels. 

## Images for Retina Displays

### Should I use 4K images for retina displays?

Yes, using 4K images for high-resolution displays like Apple's Retina is a good practice for sharpness, but it's crucial to consider file size and bandwidth, as 4K images are much larger. For web and email, consider using responsive images that deliver the right size to each device, optimizing file size and load times. In cases where image sharpness is critical, such as in professional design or on large screens, 4K is beneficial, but on smaller displays, a smaller, scaled-up image (like a "2x" image) can be sufficient. 

### When to use 4K images

* For maximum sharpness and clarity: 4K provides the highest detail, making it ideal for professional work or when you need future-proof content.
* On larger displays: A 4K display on a 27-inch or larger screen can provide "Retina" quality if the image is scaled correctly, often requiring a "2x" or even higher resolution image depending on the pixel density.
* For professional content: Use 4K for projects where high-quality images are essential, like graphic design, video production, or web development where color accuracy is important.
* For post-production flexibility: Shooting in 4K gives you more flexibility to crop, zoom, and reframe images in post-production without significant quality loss. 

### When to use smaller images

* For emails and small web banners: The benefits of a large 4K image are often lost on small screens, and the increased file size can lead to slow load times, frustrating users.
* When bandwidth is a concern: If your users have limited data or your website is on a low-bandwidth connection, a smaller, optimized image is better for performance.
* For less critical images: For complex images like lifestyle or landscape shots, the negative impact on load time may not be worth the increased sharpness. 

### Best practices

* Use responsive images: Deliver different image sizes based on the device's screen size and screen resolution to optimize performance for everyone. This is especially important for mobile devices.
* Use modern image formats: Consider using modern formats like WebP or AVIF instead of JPG for better compression and quality.
* Optimize with compression: Use image compression to reduce file size without a significant loss of quality.
* Consider your audience: If your audience is primarily using high-density displays, you may need to serve larger images, but always have smaller fallbacks for those with lower-spec devices.
* WebP now has broad support in the latest versions of major web browsers. If you don't need to support older browsers, then you can use WebP images throughout your web app instead of providing JPG or PNG fallbacks. (See https://developer.mozilla.org/en-US/docs/Web/Media/Guides/Formats/Image_types#webp_image.)
* Use the appropriate browser cache strategy for your app so that images load quickly. If your content does not change too often (i.e. if you are not a news site, for example), then make sure that your app will cache images in the user's browser. This way if the same image is needed again, the browser doesn't have to go over the network to get it, which is a massive performance boost for repeat viewings. Web browsers cache images by default, but you can change how the browser should cache images. Learn more at https://www.hostwinds.com/blog/what-is-browser-cache.
* Implement lazy loading for images, which is how you can avoid loading all the images on a page at once. Lazy loading means waiting to download an image until it is in or near the viewport. So, for example, an image way far down the page won't load if the user never scrolls there.

### How to upscale image resolution

The images for retina displays need to be 2x resolution, which means that the image has twice the width and twice the height of a standard "1x" image, resulting in a total of four times the number of pixels. You can use a free AI image upscaler online to convert standard-resolution images to high-resolution images. For example, Pixelcut.ai has a free image upscaler: https://www.pixelcut.ai/image-upscaler. Pixelcut.ai allows you to have 3 free downloads per day, but if something is wrong with that feature and you are unable to download more than once, then you can follow the instructions below.

* Make sure to use an image that has a 16:9 aspect ratio. If you are using Midjourney, then set the aspect ratio of the generated images to 16:9.
* Go to Pixelcut.ai and login with your Google account.
* After you login you will be redirected to the "Create" dashboard. (If you are not already there, then click the "Create" button in the left panel.)
* Click the "Upscale" feature.
* Upload an image.
* Select the "2x" upscale ratio (the "4x" ratio is a paid option). You will see something like this under the "Upscale" button: **1536x768** -> 3072x1536. The dark/bolded ratio is your image's current ratio and the light ratio is the ratio that your image will be upscaled too and you click "Upscale". 
* Click the "Upscale" button. The upscale process will take a little while to complete, especially if you are upscale from a large image to a really large image.
* After the image has completed the upscale process, click the "Download" button.
* Rename your downloaded image with the upscaled dimensions as part of the name.
* Repeat the upscale process until you get an image that is at least 4K (i.e. at least 3840 x 2160 pixels).

If you are unable to download your upscaled images more than once, then follow these instructions:

* After the image has completed the upscale process, press F12 to open DevTools.
* Under the "Elements" tab, expand the HTML elements until you find an <img> element. Note that there are 2 <img> elements in the DOM: The original that you uploaded and the scaled version. It looks like the first <img> listed in the DOM is the scaled version.
* Copy the src URL of the scaled image and paste it in another tab.
* Right click the image, select "Save image as...", and save your image to your desktop.
