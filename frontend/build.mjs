import * as esbuild from "esbuild";

let ctx = await esbuild.context({
  entryPoints: [
    "src/common.js",
    "src/map.js",
    "src/route_segments.js",
    "src/components/**/*.js",
    "src/views/**/*.js",
  ],
  loader: {
    ".png": "file",
  },
  bundle: true,
  minify: true,
  sourcemap: true,
  format: "esm",
  target: ["chrome58", "firefox57", "safari11", "edge18"],
  outdir: "../assets/",
});

if (process.argv.indexOf(`--watch`) !== -1) {
  await ctx.watch();
  console.log("watching...");
} else {
  await ctx.rebuild();
  await ctx.dispose();
}
