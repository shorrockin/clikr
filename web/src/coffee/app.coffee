console.log "booting clikr"

React      = require "react"
Clikr      = require "./components/Clikr"

React.render React.createElement(Clikr, {}), document.getElementById("container")
