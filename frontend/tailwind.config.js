/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				accent: '#4FD8D8',
				secondary: '#F1F5F9',
				backgroundFrom: '#B374B8',
				backgroundTo: '#C25858'
			}
		}
	},
	plugins: []
};
