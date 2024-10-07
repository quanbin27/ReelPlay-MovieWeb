(function (window, document, undefined) {
	'use strict';

	/*==============================
	Header
	==============================*/
	if (document.querySelector('.header')) {
		const headerBtn = document.querySelector('.header__btn');
		const headerNav = document.querySelector('.header__nav');
		const headerSrc = document.querySelector('.header__search');
		const headerSrcBtn = document.querySelector('.header__search-btn');
		const headerSrcCls = document.querySelector('.header__search-close');

		function toggleHeaderMenu() {
			headerBtn.classList.toggle('header__btn--active');
			headerNav.classList.toggle('header__nav--active');

			/* z-index fix */
			if (document.querySelector('.filter--fixed')) {
				var filterFixed = document.querySelector('.filter--fixed');
				filterFixed.classList.toggle('filter--hidden');
			}
		}
		headerBtn.addEventListener('click', toggleHeaderMenu);

		function toggleSearch() {
			headerSrc.classList.toggle('header__search--active');
		}
		headerSrcBtn.addEventListener('click', toggleSearch);
		headerSrcCls.addEventListener('click', toggleSearch);
	}

	/*==============================
	Filter
	==============================*/
	if (document.querySelector('.mfilter')) {
		var mfilterBtn = document.querySelector('.filter__menu');
		var mfilterClose = document.querySelector('.mfilter__close');
		var mfilter = document.querySelector('.mfilter');

		function toggleMfilter() {
			mfilter.classList.toggle('mfilter--active');
		}

		mfilterBtn.addEventListener('click', toggleMfilter);
		mfilterClose.addEventListener('click', toggleMfilter);
	}

	/* z-index fix */
	if (document.querySelector('.filter--fixed')  && window.innerWidth >= 1200 ) {
		const filterFixed = document.querySelector('.filter--fixed');

		filterFixed.classList.add('filter--hidden');

		window.addEventListener('scroll', function() {
			

			if (filterFixed && window.innerWidth >= 1200) {
				var distanceToTop = filterFixed.getBoundingClientRect().top;

				if (distanceToTop <= 80) {
					filterFixed.classList.remove('filter--hidden');
				} else {
					filterFixed.classList.add('filter--hidden');
				}
			}
		});
	}

	/*==============================
	Carousel
	==============================*/
	if (document.querySelector('.home__carousel')) {
		new Splide('.home__carousel', {
			type: 'loop',
			perPage: 5,
			drag: true,
			pagination: false,
			autoWidth: false,
			autoHeight: false,
			speed: 800,
			gap: 30,
			arrows: false,
			focus: 0,
			breakpoints: {
				575: {
					gap: 24,
					arrows: false,
					perPage: 2,
				},
				767: {
					gap: 24,
					arrows: false,
					perPage: 3,
				},
				991: {
					arrows: false,
					perPage: 3,
					gap: 24,
				},
				1199: {
					arrows: false,
					perPage: 4,
					gap: 24,
				},
				1399: {
					arrows: false,
					perPage: 5,
					gap: 24,
				},
			}
		}).mount();
	}

	if (document.querySelector('.hero')) {
		new Splide('.hero', {
			type: 'loop',
			perPage: 1,
			drag: true,
			pagination: true,
			speed: 1200,
			gap: 24,
			arrows: false,
			focus: 0,
		}).mount();
	}

	if (document.querySelector('.section__carousel')) {
		var elms = document.getElementsByClassName('section__carousel');

		for ( var i = 0; i < elms.length; i++ ) {
			new Splide(elms[ i ], {
				type: 'loop',
				perPage: 6,
				drag: true,
				pagination: false,
				autoWidth: false,
				autoHeight: false,
				speed: 800,
				gap: 24,
				arrows: false,
				focus: 0,
				breakpoints: {
					575: {
						arrows: false,
						perPage: 2,
					},
					767: {
						arrows: false,
						perPage: 3,
					},
					991: {
						arrows: false,
						perPage: 3,
					},
					1199: {
						arrows: false,
						perPage: 4,
					},
				}
			}).mount();
		}
	}

	if (document.querySelector('.section__roadmap')) {
		var elms = document.getElementsByClassName('section__roadmap');

		for ( var i = 0; i < elms.length; i++ ) {
			new Splide(elms[ i ], {
				type: 'loop',
				perPage: 3,
				drag: true,
				pagination: false,
				autoWidth: false,
				autoHeight: true,
				speed: 800,
				gap: 30,
				arrows: false,
				focus: 0,
				breakpoints: {
					767: {
						gap: 20,
						arrows: false,
						perPage: 1,
					},
					991: {
						arrows: false,
						perPage: 2,
					},
					1199: {
						arrows: false,
						perPage: 3,
					},
				}
			}).mount();
		}
	}

	/*==============================
	Section bg
	==============================*/
	if (document.querySelector('.section--bg')) {
		var mainBg = document.querySelector('.section--bg');

		if (mainBg.getAttribute('data-bg')) {
			mainBg.style.background = `url(${mainBg.getAttribute('data-bg')})`;
			mainBg.style.backgroundPosition = 'center center';
			mainBg.style.backgroundRepeat = 'no-repeat';
			mainBg.style.backgroundSize = 'cover';
		}
	}

	if (document.querySelector('.hero__slide')) {
		document.querySelectorAll('.hero__slide').forEach(function(element) {
			if (element.getAttribute("data-bg")) {
				element.style.background = 'url(' + element.getAttribute('data-bg') + ')';
				element.style.backgroundPosition = 'center center';
				element.style.backgroundRepeat = 'no-repeat';
				element.style.backgroundSize = 'cover';
			}
		});
	}

	if (document.querySelector('.section__details-bg')) {
		var mainBg = document.querySelector('.section__details-bg');

		if (mainBg.getAttribute('data-bg')) {
			mainBg.style.background = `url(${mainBg.getAttribute('data-bg')})`;
			mainBg.style.backgroundPosition = 'center center';
			mainBg.style.backgroundRepeat = 'no-repeat';
			mainBg.style.backgroundSize = 'cover';
		}
	}

	/*==============================
	Select
	==============================*/
	if (document.querySelector('#filter__genre')) {
		new SlimSelect({
			select: '#filter__genre'
		});
	}

	if (document.querySelector('#filter__quality')) {
		new SlimSelect({
			select: '#filter__quality',
			settings: {
				showSearch: false,
			}
		});
	}

	if (document.querySelector('#filter__rate')) {
		new SlimSelect({
			select: '#filter__rate',
			settings: {
				showSearch: false,
			}
		});
	}

	if (document.querySelector('#filter__sort')) {
		new SlimSelect({
			select: '#filter__sort',
			settings: {
				showSearch: false,
			}
		});
	}

	/* mobile filter */
	if (document.querySelector('#mfilter__genre')) {
		new SlimSelect({
			select: '#mfilter__genre'
		});
	}

	if (document.querySelector('#mfilter__quality')) {
		new SlimSelect({
			select: '#mfilter__quality',
			settings: {
				showSearch: false,
			}
		});
	}

	if (document.querySelector('#mfilter__rate')) {
		new SlimSelect({
			select: '#mfilter__rate',
			settings: {
				showSearch: false,
			}
		});
	}

	if (document.querySelector('#mfilter__sort')) {
		new SlimSelect({
			select: '#mfilter__sort',
			settings: {
				showSearch: false,
			}
		});
	}

	/* tv series */
	if (document.querySelector('#filter__season')) {
		new SlimSelect({
			select: '#filter__season',
			settings: {
				showSearch: false,
			}
		});
	}

	if (document.querySelector('#filter__series')) {
		new SlimSelect({
			select: '#filter__series',
			settings: {
				showSearch: false,
			}
		});
	}

	if (document.querySelector('#filter__sync')) {
		new SlimSelect({
			select: '#filter__sync',
			settings: {
				showSearch: false,
			}
		});
	}

	/*==============================
	Favorite
	==============================*/
	document.querySelectorAll('.item__favorite').forEach(function (element) {
		element.addEventListener('click', function () {
			element.classList.toggle('item__favorite--active');
		});
	});

	/*==============================
	Player
	==============================*/
	if (document.querySelector('#player')) {
		const player = new Plyr(document.querySelector('#player'));
	}

	/*==============================
	Scrollbar
	==============================*/
	var Scrollbar = window.Scrollbar;

	if (document.querySelector('.dashbox__table-wrap--1')) {
		Scrollbar.init(document.querySelector('.dashbox__table-wrap--1'), {
			damping: 0.1,
			renderByPixels: true,
			alwaysShowTracks: true,
			continuousScrolling: true
		});
	}

	if (document.querySelector('.dashbox__table-wrap--2')) {
		Scrollbar.init(document.querySelector('.dashbox__table-wrap--2'), {
			damping: 0.1,
			renderByPixels: true,
			alwaysShowTracks: true,
			continuousScrolling: true
		});
	}

	if (document.querySelector('.item__description')) {
		Scrollbar.init(document.querySelector('.item__description'), {
			damping: 0.1,
			renderByPixels: true,
			alwaysShowTracks: true,
			continuousScrolling: true
		});
	}

	/*==============================
	Gallery
	==============================*/
	if (document.querySelector('.gallery')) {
		var initPhotoSwipeFromDOM = function(gallerySelector) {

			// parse slide data (url, title, size ...) from DOM elements 
			// (children of gallerySelector)
			var parseThumbnailElements = function(el) {
				var thumbElements = el.childNodes,
					numNodes = thumbElements.length,
					items = [],
					figureEl,
					linkEl,
					size,
					item;

				for(var i = 0; i < numNodes; i++) {

					figureEl = thumbElements[i]; // <figure> element

					// include only element nodes 
					if(figureEl.nodeType !== 1) {
						continue;
					}

					linkEl = figureEl.children[0]; // <a> element

					size = linkEl.getAttribute('data-size').split('x');

					// create slide object
					item = {
						src: linkEl.getAttribute('href'),
						w: parseInt(size[0], 10),
						h: parseInt(size[1], 10)
					};



					if(figureEl.children.length > 1) {
						// <figcaption> content
						item.title = figureEl.children[1].innerHTML; 
					}

					if(linkEl.children.length > 0) {
						// <img> thumbnail element, retrieving thumbnail url
						item.msrc = linkEl.children[0].getAttribute('src');
					} 

					item.el = figureEl; // save link to element for getThumbBoundsFn
					items.push(item);
				}

				return items;
			};

			// find nearest parent element
			var closest = function closest(el, fn) {
				return el && ( fn(el) ? el : closest(el.parentNode, fn) );
			};

			// triggers when user clicks on thumbnail
			var onThumbnailsClick = function(e) {
				e = e || window.event;
				e.preventDefault ? e.preventDefault() : e.returnValue = false;

				var eTarget = e.target || e.srcElement;

				// find root element of slide
				var clickedListItem = closest(eTarget, function(el) {
					return (el.tagName && el.tagName.toUpperCase() === 'FIGURE');
				});

				if(!clickedListItem) {
					return;
				}

				// find index of clicked item by looping through all child nodes
				// alternatively, you may define index via data- attribute
				var clickedGallery = clickedListItem.parentNode,
					childNodes = clickedListItem.parentNode.childNodes,
					numChildNodes = childNodes.length,
					nodeIndex = 0,
					index;

				for (var i = 0; i < numChildNodes; i++) {
					if(childNodes[i].nodeType !== 1) { 
						continue; 
					}

					if(childNodes[i] === clickedListItem) {
						index = nodeIndex;
						break;
					}
					nodeIndex++;
				}



				if(index >= 0) {
					// open PhotoSwipe if valid index found
					openPhotoSwipe( index, clickedGallery );
				}
				return false;
			};

			// parse picture index and gallery index from URL (#&pid=1&gid=2)
			var photoswipeParseHash = function() {
				var hash = window.location.hash.substring(1),
				params = {};

				if(hash.length < 5) {
					return params;
				}

				var vars = hash.split('&');
				for (var i = 0; i < vars.length; i++) {
					if(!vars[i]) {
						continue;
					}
					var pair = vars[i].split('=');  
					if(pair.length < 2) {
						continue;
					}           
					params[pair[0]] = pair[1];
				}

				if(params.gid) {
					params.gid = parseInt(params.gid, 10);
				}

				return params;
			};

			var openPhotoSwipe = function(index, galleryElement, disableAnimation, fromURL) {
				var pswpElement = document.querySelectorAll('.pswp')[0],
					gallery,
					options,
					items;

				items = parseThumbnailElements(galleryElement);

				// define options (if needed)
				options = {

					// define gallery index (for URL)
					galleryUID: galleryElement.getAttribute('data-pswp-uid'),

					getThumbBoundsFn: function(index) {
						// See Options -> getThumbBoundsFn section of documentation for more info
						var thumbnail = items[index].el.getElementsByTagName('img')[0], // find thumbnail
							pageYScroll = window.pageYOffset || document.documentElement.scrollTop,
							rect = thumbnail.getBoundingClientRect(); 

						return {x:rect.left, y:rect.top + pageYScroll, w:rect.width};
					}

				};

				// PhotoSwipe opened from URL
				if(fromURL) {
					if(options.galleryPIDs) {
						// parse real index when custom PIDs are used 
						// http://photoswipe.com/documentation/faq.html#custom-pid-in-url
						for(var j = 0; j < items.length; j++) {
							if(items[j].pid == index) {
								options.index = j;
								break;
							}
						}
					} else {
						// in URL indexes start from 1
						options.index = parseInt(index, 10) - 1;
					}
				} else {
					options.index = parseInt(index, 10);
				}

				// exit if index not found
				if( isNaN(options.index) ) {
					return;
				}

				if(disableAnimation) {
					options.showAnimationDuration = 0;
				}

				// Pass data to PhotoSwipe and initialize it
				gallery = new PhotoSwipe( pswpElement, PhotoSwipeUI_Default, items, options);
				gallery.init();
			};

			// loop through all gallery elements and bind events
			var galleryElements = document.querySelectorAll( gallerySelector );

			for(var i = 0, l = galleryElements.length; i < l; i++) {
				galleryElements[i].setAttribute('data-pswp-uid', i+1);
				galleryElements[i].onclick = onThumbnailsClick;
			}

			// Parse URL and open gallery if it contains #&pid=3&gid=1
			var hashData = photoswipeParseHash();
			if(hashData.pid && hashData.gid) {
				openPhotoSwipe( hashData.pid ,  galleryElements[ hashData.gid - 1 ], true, true );
			}
		};

		initPhotoSwipeFromDOM('.gallery');
	}

	/*==============================
	Back to top
	==============================*/
	if (document.querySelector('.footer__back')) {
		document.querySelector('.footer__back').addEventListener('click', function() {
			window.scrollTo({
				top: 0,
				behavior: 'smooth'
			});
		});
	}

})(window, document);