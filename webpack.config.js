module.exports = {
    entry: "./assets/js/src/index.js",
    output: {
        path: __dirname,
        filename: "assets/js/dist/bundle.js"
    },
    module: {
        loaders: [
            { test: /\.js?/, loader: 'babel-loader', exclude: /node_modules/ }
        ]
    }
};