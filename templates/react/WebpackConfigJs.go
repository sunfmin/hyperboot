package react

import (
	"bytes"
)

func WebpackConfigJs(data ReactData) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\nvar webpack = require(\"webpack\");\nvar ExtractTextPlugin = require(\"extract-text-webpack-plugin\");\n\nmodule.exports = {\n\tcache: true,\n\tentry: {\n\t\tweb: './app.js'\n\t},\n\t// devtool: 'inline-source-map',\n\toutput: {\n\t\tpath: \"../www/assets/\",\n\t\tfilename: \"main.js\"\n\t},\n\tmodule: {\n\t\tloaders: [\n\t\t\t{ test: /\\.jsx?$/, loader: 'jsx-loader?harmony&stripTypes' },\n\t\t\t{ test: /\\.less$/, loader: ExtractTextPlugin.extract(\"css-loader!less-loader\") },\n\t\t\t{ test: /\\.woff\\d?(\\?v=[0-9]\\.[0-9]\\.[0-9])?$/, loader: \"url-loader?limit=10000&minetype=application/font-woff\" },\n\t\t\t{ test: /\\.(ttf|eot|svg)(\\?v=[0-9]\\.[0-9]\\.[0-9])?$/, loader: \"file-loader\" }\n\t\t]\n\t},\n\tplugins: [\n\t\tnew webpack.IgnorePlugin(/vertx/),\n\t\tnew ExtractTextPlugin(\"main.css\", { allChunks: true })\n\t]\n};")

	return _buffer.String()
}
