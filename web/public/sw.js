self.addEventListener("install", (event) => {
  console.log("Service Worker installed");
});

self.addEventListener("fetch", (event) => {
  const url = new URL(event.request.url);

  // Ignore Vite internals
  if (
    url.pathname.startsWith("/@vite") ||
    url.pathname.startsWith("/src") ||
    url.pathname.startsWith("/node_modules")
  ) {
    return;
  }

  // Ignore websocket/HMR requests
  if (event.request.headers.get("upgrade") === "websocket") {
    return;
  }

  event.respondWith(handleRequest(event.request));
});

async function handleRequest(request) {
  const cache = await caches.open("ghost-cache");

  const cached = await cache.match(request);

  if (cached) {
    console.log("Cache hit");
    return cached;
  }

  console.log("Cache miss");

  try {
    const helperResponse = await fetch("http://localhost:8080/asset");

    if (helperResponse.ok) {
      console.log("Served from Ghost helper");

      cache.put(request, helperResponse.clone());

      return helperResponse;
    }
  } catch (err) {
    console.log("Helper unavailable");
  }

  console.log("Falling back to internet");

  const networkResponse = await fetch(request);

  cache.put(request, networkResponse.clone());

  return networkResponse;
}