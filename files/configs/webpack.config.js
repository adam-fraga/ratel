const path = require("path"); // Node Js module path management

module.exports = {
  devtool: "eval-source-map", //Provide debug in inspector in source file and not tidious readable bundle.js file (add source map in tsconfig file)
  entry: "./scripts/index.ts", //Entrypoint for webpack compiler
  module: {
    rules: [
      {
        test: /\.ts$/,
        use: "ts-loader",
        include: [path.resolve(__dirname, "scripts")],
        exclude: /node_modules/,
      },
    ],
  },
  resolve: {
    extensions: ["js", "ts"],
  },
  output: {
    publicPath: "scripts", // Specify the folder for webpack dev server to reload properly
    filename: "bundle.js", // Name of the outputed file after compile
    path: path.resolve(__dirname, "static/js"), // Abs path for the outputed file
  },
};
