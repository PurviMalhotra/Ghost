self.addEventListener("install", (event) => {
  console.log("Service Worker installed");
});

self.addEventListener("fetch", (event) => {

  const url = new URL(event.request.url);

  // Only Ghost-controlled assets
  if (!url.pathname.startsWith("/ghost-assets/")) {
    return;
  }

  event.respondWith(
    handleGhostRequest(url.pathname)
  );
});

async function handleGhostRequest(pathname) {

  const cache = await caches.open("ghost-cache");

  const cached = await cache.match(pathname);

  if (cached) {
    console.log("Cache hit");

    return cached;
  }

  console.log("Cache miss");

  const name = pathname.replace(
    "/ghost-assets/",
    ""
  );

  try {

    const helperResponse = await fetch(
      `http://localhost:8080/request?name=${name}`
    );

    if (helperResponse.ok) {

      console.log("Served from Ghost helper");

      cache.put(
        pathname,
        helperResponse.clone()
      );

      return helperResponse;
    }

  } catch (err) {

    console.log("Helper unavailable");
  }

  return new Response(
    "Ghost asset unavailable",
    { status: 404 }
  );
}