React        = require "react"
episodeStore = require "../stores/episodeStore"
sceneStore   = require "../stores/sceneStore"

module.exports = React.createClass
  getInitialState: () ->
    { image: null }

  componentDidMount: () ->
    sceneStore.addChangeListener(@onSceneChanged)

  componentDidUnmount: () ->
    sceneStore.removeChangeListener(@onSceneChanged)

  onSceneChanged: () ->
    @setState { image: sceneStore.scene?.image }

  render: () ->
    if @state.image?
      <img src={@state.image} />
    else
      <div className="img"/>