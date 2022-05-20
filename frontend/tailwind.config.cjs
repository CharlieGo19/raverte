const colors = require("tailwindcss/colors");

const config = {
  content: ["./src/**/*.{html,js,svelte,ts}"],

  theme: {
		extend: {
			colors: {
				fuchsia: colors.fuchsia,
				canvas: "rgb(11, 10, 13)",
				lbtndark: "rgb(2, 15, 9)",
				rgreen: "rgb(38, 194, 33)",
				rred: "rgb(194, 33, 33)"
			},
			boxShadow: {
				lgreen: "35px 0px 30px 0px rgba(38, 194, 33, 0.7)",
				lred: "-35px 0px 30px 0px rgba(194, 33, 33, 0.7)",
				lbtn: "-5px 0px 15px 0px rgba(194, 33, 33, 0.7), 5px 0px 15px 0px rgba(38, 194, 33, 0.7);",
			},
			fontFamily: {
				poppins: "'Poppins'",
			},
		},
	},
  plugins: [],
};

module.exports = config;
