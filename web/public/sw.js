self.addEventListener("install", (event) => {
  console.log("Service Worker installed");
});

self.addEventListener("fetch", (event) => {
  event.respondWith(
    caches.open("ghost-cache").then(async (cache) => {
      const cached = await cache.match(event.request);

      if (cached) {
        console.log("Cache hit");
        return cached;
      }

      console.log("Cache miss");

      const response = await fetch(event.request);

      cache.put(event.request, response.clone());

      return response;
    })
  );
});