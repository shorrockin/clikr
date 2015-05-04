console.log "booting clikr clikr"

React = require "react"
Clikr = require "./components/clikr"

React.renderComponent Clikr(null), document.getElementById("body")