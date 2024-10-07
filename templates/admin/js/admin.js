(function (window, document, undefined) {
	'use strict';

	/*==============================
	Header
	==============================*/
	if (document.querySelector('.header')) {
		const headerBtn = document.querySelector('.header__btn');
		const headerNav = document.querySelector('.sidebar');

		function toggleHeaderMenu() {
			headerBtn.classList.toggle('header__btn--active');
			headerNav.classList.toggle('sidebar--active');
		}

		headerBtn.addEventListener('click', toggleHeaderMenu);
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

	if (document.querySelector('.dashbox__table-wrap--3')) {
		Scrollbar.init(document.querySelector('.dashbox__table-wrap--3'), {
			damping: 0.1,
			renderByPixels: true,
			alwaysShowTracks: true,
			continuousScrolling: true
		});
	}

	if (document.querySelector('.dashbox__table-wrap--4')) {
		Scrollbar.init(document.querySelector('.dashbox__table-wrap--4'), {
			damping: 0.1,
			renderByPixels: true,
			alwaysShowTracks: true,
			continuousScrolling: true
		});
	}

	if (document.querySelector('.catalog--1')) {
		Scrollbar.init(document.querySelector('.catalog--1'), {
			damping: 0.1,
			renderByPixels: true,
			alwaysShowTracks: true,
			continuousScrolling: true
		});
	}

	if (document.querySelector('.catalog--2')) {
		Scrollbar.init(document.querySelector('.catalog--2'), {
			damping: 0.1,
			renderByPixels: true,
			alwaysShowTracks: true,
			continuousScrolling: true
		});
	}

	/*==============================
	Filter
	==============================*/
	if (document.querySelector('#filter__sort')) {
		new SlimSelect({
			select: '#filter__sort',
			settings: {
				showSearch: false,
			}
		});
	}

	/* add page */
	if (document.querySelector('#sign__quality')) {
		new SlimSelect({
			select: '#sign__quality',
			settings: {
				showSearch: false,
			}
		});
	}

	if (document.querySelector('#sign__genre')) {
		new SlimSelect({
			select: '#sign__genre',
			settings: {
				placeholderText: 'Choose genre',
			}
		});
	}

	if (document.querySelector('#sign__country')) {
		new SlimSelect({
			select: '#sign__country',
			settings: {
				placeholderText: 'Choose country',
			}
		});
	}

	if (document.querySelector('#sign__director')) {
		new SlimSelect({
			select: '#sign__director',
			settings: {
				placeholderText: 'Choose director',
			}
		});
	}

	if (document.querySelector('#sign__actors')) {
		new SlimSelect({
			select: '#sign__actors',
			settings: {
				placeholderText: 'Choose actors',
			}
		});
	}

	/*==============================
	Upload
	==============================*/
	if (document.getElementById('sign__gallery-upload')) {
		var galleryUpload = document.getElementById('sign__gallery-upload');

		galleryUpload.addEventListener('change', function(event) {
			var length = event.target.files.length;
			var galleryLabel = galleryUpload.getAttribute('data-name');
			var label = document.querySelector(galleryLabel);

			if (length > 1) {
				label.textContent = length + " files selected";
			} else {
				label.textContent = event.target.files[0].name;
			}
		});
	}

	if (document.querySelector('.sign__video-upload')) {
		document.querySelectorAll('.sign__video-upload').forEach(function(element) {
			element.addEventListener('change', function() {
				var videoLabel = element.getAttribute('data-name');
				var vlabel = document.querySelector(videoLabel);

				if (element.value !== '') {
					vlabel.textContent = element.files[0].name;
				} else {
					vlabel.textContent = "Upload video";
				}
			});
		});
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

})(window, document);