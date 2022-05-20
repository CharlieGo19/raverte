import preprocess from "svelte-preprocess";
import adapter from "@sveltejs/adapter-static";

/** @type {import('@sveltejs/kit').Config} */
const config = {
  kit: {
    adapter: adapter({
      // Dist to stick with WAILS naming convention (embed  - main.go).
      pages: 'dist',
      assets: 'dist'
    }),

    prerender: {
      // SPA mode, so index.html is rendered for WAILS to grab.
      default: true,
    },
  },

  preprocess: [
    preprocess({
      postcss: true,
    }),
  ],
};

export default config;
