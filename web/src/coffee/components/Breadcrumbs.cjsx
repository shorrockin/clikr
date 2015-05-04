React        = require "react"
episodeStore = require "../stores/episodeStore"
sceneStore   = require "../stores/sceneStore"

module.exports = React.createClass
  getInitialState: () ->
    { episode: null, time: 0 }

  componentDidMount: () ->
    sceneStore.addChangeListener(@onSceneChanged)

  componentDidUnmount: () ->
    sceneStore.removeChangeListener(@onSceneChanged)

  onSceneChanged: () ->
    @setState { episode: episodeStore.episode, time: sceneStore.time() }

  render: () ->
    if @state.episode?
      <section className="breadcrumbs">
        <span><a href="#">{@state.episode.name}</a></span>
        <span><a href="#">Season {@state.episode.season}</a></span>
        <span><a href="#">Episode {@state.episode.episode}</a></span>
        <span><a href="#">Scene ({@state.time})</a></span>
      </section>
    else
      <section className="breadcrumbs" />