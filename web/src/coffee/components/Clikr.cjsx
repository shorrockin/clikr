React       = require "react"
Breadcrumbs = require "./Breadcrumbs"
Scene       = require "./Scene"

module.exports = React.createClass
  render: () ->
    <div>
      <Breadcrumbs />
      <Scene />
    </div>
