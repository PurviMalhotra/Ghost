# Ghost
The "Ghost" Personal CDN is essentially a "stealth" speed booster for your local Wi-Fi. It’s like having a shared folder that everyone on your team uses, but without anyone actually having to manually drag and drop files.
It helps devices on the same Wi-Fi network share cached assets with each other instead of repeatedly downloading them from the internet.

## Using:

- mDNS for automatic peer discovery
- Service Workers for intercepting and redirecting requests

When the browser says "I need to download image.png from example.com," the Service Worker pauses the request. It checks if a teammate nearby has it first. If yes, it pulls it from them; if no, it lets the request go to the real internet.

