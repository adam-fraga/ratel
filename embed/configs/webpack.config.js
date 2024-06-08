import { resolve as _resolve } from "path"; // Node Js module path management

export const devtool = "eval-source-map";
export const entry = "./src/script/index.ts";
export const module = {
  rules: [
    {
      test: /\.ts$/,
      use: "ts-loader",
      include: [_resolve(__dirname, "src/script")],
      exclude: /node_modules/,
    },
  ],
};
export const resolve = {
  extensions: ["js", "ts"],
};
export const output = {
  publicPath: "src/script", // Specify the folder for webpack dev server to reload properly
  filename: "bundle.js", // Name of the outputed file after compile
  path: _resolve(__dirname, "static/js"), // Abs path for the outputed file
};
