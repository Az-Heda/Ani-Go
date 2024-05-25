class Colors {
	static #_darkThemeClass = 'dark-theme';
	static #_currentTheme = 'light';
	static #_colors = {
		'light': { text: '#000000' },
		'dark': { text: '#FFFFFF' },
	}
	static get textColorLightTheme() { return '#000000' };
	static get textColorDarkTheme()  { return '#FFFFFF' };

	static get isDarkTheme() { return document.body.classList.contains(this.#_darkThemeClass) }

	/**
	 * @no_docs
	 * @returns {string} Color to be applied on the text based of the current theme (light or dark)
	 */
	static get textColor() {
		this.#_currentTheme = (document.body.classList.contains(this.#_darkThemeClass)) ? 'dark' : 'light';
		return this.#_colors[this.#_currentTheme].text;
	}

	/**
	 * @param {string} color1 First hex color to blend (the one you get with percentage=0)
	 * @param {string} color2 Second hex color to blend (the one you get with percentage=1)
	 * @param {float} percentage Percentage of mixing between the 2 colors
	 * @returns {string} Hex color
	 * 
	 * __Example__
	 * ```js
	 * Color.blend('#FF0000', '#00FF00', 0.2); // #CC3300
	 * Color.blend('#FFAAEE', '#FFF999', 0.5); // #FFD2C4
	 * Color.blend('#726AFF', '#009112', 0.7); // #228559
	 * ```
	 */
	static blend(color1, color2, percentage) {
		const int_to_hex = function(num) {
			let hex = Math.round(num).toString(16);
			return (hex.length == 1) ? '0'+hex : hex;
		}

		if (color1.length != 4 && color1.length != 7)
			throw new Error('colors must be provided as hexes');
		if (color2.length != 4 && color2.length != 7)
			throw new Error('colors must be provided as hexes');
		if (percentage > 1 || percentage < 0)
			throw new Error('percentage must be between 0 and 1');

		if (color1.length == 4)
			color1 = color1[1] + color1[1] + color1[2] + color1[2] + color1[3] + color1[3];
		else
			color1 = color1.substring(1);
		if (color2.length == 4)
			color2 = color2[1] + color2[1] + color2[2] + color2[2] + color2[3] + color2[3];
		else
			color2 = color2.substring(1);

		color1 = [parseInt(color1[0] + color1[1], 16), parseInt(color1[2] + color1[3], 16), parseInt(color1[4] + color1[5], 16)];
		color2 = [parseInt(color2[0] + color2[1], 16), parseInt(color2[2] + color2[3], 16), parseInt(color2[4] + color2[5], 16)];
		let color3 = [
			(1 - percentage) * color1[0] + percentage * color2[0],
			(1 - percentage) * color1[1] + percentage * color2[1],
			(1 - percentage) * color1[2] + percentage * color2[2]
		];
		color3 = '#' + color3.map(int_to_hex).join('')
		return color3;
	}

	/**
	 * @description Given an HEX Color as input, this function will invert the color; #012345 --> #FEDCBA
	 * @param {string} hexColor Hex color to invert
	 * @returns {string} Hex color
	 * 
	 * __Example__
	 * ```js
	 * Color.invert('#CC3300'); // #33CCFF
	 * Color.invert('#FFD2C4'); // #002D3B
	 * Color.invert('#228559'); // #DD7AA6
	 * ```
	 */
	static invert(hexColor) {
		if (hexColor.length != 4 && hexColor.length != 7)
			throw new Error('Color must be provided as a hex');
		if (hexColor.length == 4)
			hexColor = '#' + hexColor[1] + hexColor[1] + hexColor[2] + hexColor[2] + hexColor[3] + hexColor[3];
		let colors = hexColor.split('');
		let newColor = '#';
		colors.shift();
		for (let p = 0; p < colors.length; p += 2) {
			let c = colors[p] + colors[p+1];
			let nc = (Math.abs(parseInt(c, 16) - 255)).toString(16)
			newColor += (nc.length == 1) ? '0'+nc : nc
		}
		return newColor.toUpperCase();
	}

	/**
	 * @no_docs
	 * @param {string} hexColor Hex color to invert
	 * @returns {string} Hex color
	 */
	static invertIfDarkTheme(hexColor) {
		return (!this.isDarkTheme) ? hexColor : this.invert(hexColor);
	}

	/**
	 * @description This function will convert all of the 3 digits HEX Colors into a 6 digits
	 * @param {string} hex 3 or 6 digit HEX Color
	 * @returns {string} Returns a 6 digit HEX Color
	 * 
	 * __Example__
	 * ```js
	 * Color.fullhex('#FA0'); // #FFAA00
	 * Color.fullhex('#123456'); // #123456
	 * ```
	 */
	static fullhex(hex) {
		let c = hex.replace('#', '');
		return `#${(c.length == 6) ? c : c.split('').map((i) => i+i).join('')}`;
	}

	/**
	 * @description Transform the given HEX Color into the same color as [r, g, b]
	 * @param {string} hexColor Hex color to transform to RGB
	 * @returns {Array<r,g,b>} RGB Value of the given hex
	 * 
	 * __Example__
	 * ```js
	 * Color.hex2rgb('#33CCFF'); // [51, 204, 255]
	 * Color.hex2rgb('#002D3B'); // [0, 45, 59]
	 * Color.hex2rgb('#DD7AA6'); // [221, 122, 166]
	 * ```
	 */
	static hex2rgb(hex) {
		if (hex.length != 4 && hex.length != 7)
			throw new Error('colors must be provided as hexes');
		let vals = hex.split('');
		let newVals = [];
		vals.shift();
		for (let x = 0; x < vals.length; x += (vals.length == 6) ? 2 : 1) {
			newVals.push(parseInt(vals[x] + ((vals.length == 6) ? vals[x+1] : ''), 16));
		}
		return newVals;
	}

	/**
	 * @description Transform the given RGB Color into the same color with the HEX format
	 * @param {int} r R value for the color
	 * @param {int} g G value for the color
	 * @param {int} b B value for the color
	 * @returns {string} Hex color from the given RGB values
	 * 
	 * __Example__
	 * ```js
	 * Color.rgb2hex(51, 204, 255); // #33CCFF
	 * Color.rgb2hex(0, 45, 59); // #002D3B
	 * Color.rgb2hex(...[221, 122, 166]); // #DD7AA6
	 * ```
	 */
	static rgb2hex(r, g, b) {
		if (!(r >= 0 && r <= 255)) { throw new Error(`Red value must be between 0 and 255, received ${r}`); }
		if (!(g >= 0 && g <= 255)) { throw new Error(`Green value must be between 0 and 255, received ${g}`); }
		if (!(b >= 0 && b <= 255)) { throw new Error(`Blue value must be between 0 and 255, received ${b}`); }
		let hex = [r, g, b].map((item) => { return item.toString(16) }).map((item) => { return (item.length == 1) ? '0'+item : item })
		return ('#' + hex.join('')).toUpperCase();
	}

	/**
	 * @description Apply a Lighter/Darker shade at the given color;
	 * @param {string} color Hex of the color 
	 * @param {float} percentage Percentage of Brightness/Shade to apply (positive for brightness, negative for shade)
	 * @returns {string} Hex color with the given filter
	 * 
	 * 1. if percentage is 0: return the given color
	 * 2. between 0 and 1: color goes to the given one to white
	 * 3. between 0 and -1: color goes to the given one to black
	 * 
	 * 
	 * __Example__
	 * ```js
	 * console.log(Color.shade('#33CCFF', -0.7)); // #0f3d4d
	 * console.log(Color.shade('#002D3B', 0.2));  // #335762
	 * console.log(Color.shade('#DD7AA6', 0.5));  // #eebdd3 
	 * ```
	 */
	static shade(color, percentage) {
		if (!(percentage >= -1 && percentage <= 1)) { throw new Error('Percentage must be between -1 and 1')}
		if (percentage > 0) {
			return this.blend(color, '#FFFFFF', percentage);
		}
		if (percentage < 0) {
			return this.blend(color, '#000000', Math.abs(percentage));
		}
		return color;
	}

	/**
	 * @description Show in the console a box with the color passed, with the text in both white and black colors
	 * @param {string | int} hex Hex of the color to display or __RED__ color of RGB Value
	 * @param  {...int} ...gb __GREEN__ and __BLUE__ values of the color into separate arguments (or with spread operator "...[g, b]")
	 * 
	 * __Example__
	 * ```js
	 * Color.logColor('#FFAA00');
	 * Color.logColor(200, 50, 100);
	 * ```
	 */
	static logColor(hex, ...gb) {
		const consoleColor = (c) => {
			const common = 'padding: 15px 100px; font-size: 30px; text-align: center;';
			const s1 = `${common} background-color: ${c}; color: #FFF; border-right: 1px solid #FFF;`
			const s2 = `${common} background-color: ${c}; color: #000; border-left: 1px solid #000;`
			console.log(`%c${c}%c${c}`, s1, s2);
		}
		if (gb.length == 0) {
			consoleColor(this.fullhex(hex));
		}
		else if (gb.length == 2) {
			hex = this.rgb2hex(hex, ...gb)
			consoleColor(this.fullhex(hex));
		}
	}
}