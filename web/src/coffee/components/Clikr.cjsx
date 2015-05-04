React       = require "react"
Breadcrumbs = require "./Breadcrumbs"
Scene       = require "./Scene"
dispatcher  = require "../dispatcher"
actions     = require "../actions"

module.exports = React.createClass
  getInitialState: () ->
    actions.loadEpisode("new_girl", 4, 21)

  render: () ->
    <div>
      <Breadcrumbs />
      <Scene />
    </div>
