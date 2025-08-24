import { copy } from "esbuild-plugin-copy";
import * as esbuild from "esbuild";

let ctx = await esbuild.context({
  entryPoints: [
    "src/common.js",
    "src/components/**/*.ts",
    "src/views/**/*.ts",
    "src/components/**/*.js",
    "src/views/**/*.js"
  ],
  loader: {
    ".png": "file",
  },
  bundle: true,
  minify: true,
  sourcemap: 'external',
  format: "esm",
  target: ["chrome58", "firefox57", "safari11", "edge18"],
  outdir: "../assets/",
  publicPath: "/assets/",
  plugins: [
    copy({
      assets: {
        from: ["./node_modules/shareon/dist/*.{js,css}"],
        to: ["../assets/vendor/shareon"],
      },
    }),
    copy({
      assets: {
        from: ["./node_modules/htmx.org/dist/htmx.min.js"],
        to: ["../assets/vendor/htmx"],
      },
    }),
  ],
});

if (process.argv.indexOf(`--watch`) !== -1) {
  await ctx.watch();
  console.log("watching...");
} else {
  await ctx.rebuild();
  await ctx.dispose();
}
