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
        caches.match(event.request).then(resp => {
            if (resp) {
                console.log("from cache:", event.request.url);
                return resp;
            }
            return fetch(event.request).then(resp => {
                if (resp) {
                    caches.open("v1").then(cache => {
                        cache.put(event.request, resp.clone());
                    }).then(() => {
                        console.log("put in cache:", event.request.url);
                    }).catch(err=> {
                        console.error("put in cache:", event.request.url, err);
                    });
                }
                return resp;
            });
        })
    );
});
