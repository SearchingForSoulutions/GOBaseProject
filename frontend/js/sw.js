const cacheName = 'gobaseproject';


// per PWA: Registering our Service worker
if('serviceWorker' in navigator) {
  navigator.serviceWorker.register('js/sw.js', { scope: './js/' })
}

// Cache all the files to make a PWA
self.addEventListener('install', e => {
  e.waitUntil(
    caches.open(cacheName).then(cache => {
      // Our application only has two files here index.html and manifest.json
      // but you can add more such as style.css as your app grows
      return cache.addAll([
        './',
        './index.html',
        './manifest.json',
        './js/*',
        './css/style.css',  
        './assets/logo.png'
      ]);
    })
  );
});

// Our service worker will intercept all fetch requests
// and check if we have cached the file
// if so it will serve the cached file
self.addEventListener('fetch', event => {
  event.respondWith(
    caches.open(cacheName)
      .then(cache => cache.match(event.request, { ignoreSearch: true }))
      .then(response => {
        return response || fetch(event.request);
      })
  );
});


// per live reload
/**
 * Apro una connessione WebSocket con il server, quando viene chiusa ricarico la pagina.
 * Questa è una implementazione semplicistica del live reload.
 * @file site/client-websocket.js
 */
(() => {
	const socketUrl = 'wss://' + window.location.host + "/ws"
    //console.log("socketURL: " + socketUrl)

	let socket = new WebSocket(socketUrl);
	socket.addEventListener('close', () => {
		// Then the server has been turned off,
		// either due to file-change-triggered reboot,
		// or to truly being turned off.

		// Attempt to re-establish a connection until it works,
		// failing after a few seconds (at that point things are likely
		// turned off/permanantly broken instead of rebooting)
		const interAttemptTimeoutMilliseconds = 100;
		const maxDisconnectedTimeMilliseconds = 3000;
		const maxAttempts = Math.round(
			maxDisconnectedTimeMilliseconds / interAttemptTimeoutMilliseconds,
		);
		let attempts = 0;
		const reloadIfCanConnect = () => {
			attempts++;
			if (attempts > maxAttempts) {
				console.error('Could not reconnect to dev server.');
				return;
			}
			socket = new WebSocket(socketUrl);
			socket.addEventListener('error', () => {
				setTimeout(reloadIfCanConnect, interAttemptTimeoutMilliseconds);
			});
			socket.addEventListener('open', () => {
				location.reload();
			});
		};
		reloadIfCanConnect();
	});
})();
