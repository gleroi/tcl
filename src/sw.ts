self.addEventListener('install', function (event: any) {
    event.waitUntil(
        caches.open('v1').then(cache => {
            return cache.addAll([
                "/",
                "/index.html",
                "/main.js",
            ]);
        })
    );
});

self.addEventListener('fetch', function (event: any) {
    event.respondWith(
        caches.match(event.request)
    );
});
