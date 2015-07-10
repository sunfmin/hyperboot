package react

import (
	"bytes"
)

func HelloJs(data ReactData) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n/* @flow */\n\nvar React = require(\"react\");\n\nvar Hello = React.createClass({\n\tgetInitialState: function() {\n\t\treturn {\n\t\t\tname: \"\",\n\t\t}\n\t},\n\n\tcomponentDidMount: function() {\n\t},\n\n\trender: function() : React.ReactElement {\n\t\treturn (\n\t\t\t\t<div className=\"alert alert-success\">\n\t\t\t\t\tHello {this.state.name}, <input ref=\"name\" onChange={this.changeName}/>\n\t\t\t\t</div>\n\t\t);\n\t},\n\n\tchangeName: function() {\n\t\tvar userserv = this.props.service.GetUserService(\"\");\n\t\tvar self = this;\n    \tuserserv.PutUser(88, this.refs.name.getDOMNode().value, function(user, err){\n    \t\tself.setState({name: user.Name});\n    \t})\n  \t}\n\n});\n\nmodule.exports = Hello;")

	return _buffer.String()
}
