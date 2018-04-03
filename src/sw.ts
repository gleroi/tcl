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
        fetch(event.request)
        .then(resp => {
            return caches.open("v1").then(cache => {
                cache.put(event.request, resp.clone());
                return resp;
            });
        }).catch(err => {
            return caches.match(event.request);
        }));
});
