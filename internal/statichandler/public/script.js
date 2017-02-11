(function() {
	'use strict'

	// root element:
	var httpfsx = document.querySelector('.httpfsx')

	// file-system item nodes:
	var items = httpfsx.querySelectorAll('.httpfsx .list .item')

	// paths, that are listed on current page:
	var existingPaths = []

	for( var i = 0; i < items.length; i++ ) {

		// item wrapper node:
		var item = items[i]

		var star = item.querySelector('.star') // starring element
		var link = item.querySelector('.link') // link element (for item's path detection)

		// item's path:
		var path = link.getAttribute('href')

		// what key is used for storing current item's starred state:
		var starKey = 'httpfsx:star:' + path

		// remember "starring" key to simplify "toggle starring" click handler:
		star.setAttribute('data-httpfsx-star-key', starKey)

		// change star's view, if item is starred:
		if( localStorage.getItem(starKey) ) {
			star.classList.add('active')
		}

		// remember this item's path to remove deleted items from localStorage down the code:
		existingPaths.push(path.replace(/\/{2,}|\/$/g, ''))

	}

	// current request (location) path:
	var currentPath = location.pathname.replace(/\/{2,}|\/$/g, '')

	// traversing localStorage items to clean up deleted ones:
	for( var key in localStorage ) {

		// ignoring any foreign keys:
		if( key.indexOf('httpfsx:') == -1 ) {
			continue
		}

		// extracting stored item path from key:
		var storedPath = key.replace(/httpfsx:[^:]+?:/, '')

		// got item from other path, cannot touch it:
		if( storedPath.indexOf(currentPath) != 0 ) {
			continue
		}

		// does current localStorage item exists (not deleted)?
		var exists = false

		// checking, if current localStorage item is present on current location (page):
		for( var i = 0; i < existingPaths.length; i++ ) {
			var existingPath = existingPaths[i]
			if( storedPath.indexOf(existingPath) == 0 ) {
				exists = true
				break
			}
		}

		// removing deleted file-system items from localStorage:
		if( !exists ) {
			localStorage.removeItem(key)
		}

	}

	// capturing "star" and "clear-storage" clicks:
	httpfsx.addEventListener('click', function(event) {

		if( event.target.classList.contains('star') ) {

			var star = event.target

			var starKey = star.getAttribute('data-httpfsx-star-key')

			// toggle starring status:
			if( localStorage.getItem(starKey) ) {
				localStorage.removeItem(starKey)
				star.classList.remove('active')
			} else {
				localStorage.setItem(starKey, 'T')
				star.classList.add('active')
			}

		} else if( event.target.classList.contains('clear-storage') ) {

			// confirm and clear localStorage:
			if( confirm('Clear storage?') ) {

				localStorage.clear()

				// apply loosing stars to UI:
				var stars = httpfsx.querySelectorAll('.star')
				for( var i = 0; i < stars.length; i++ ) {
					var star = stars[i]
					star.classList.remove('active')
				}

				alert('Storage cleared')
			}

		}

	})

})()
