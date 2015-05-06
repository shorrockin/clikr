React        = require "react"
episodeStore = require "../stores/episodeStore"
sceneStore   = require "../stores/sceneStore"
actions      = require "../actions"

InteractionElement = React.createClass
  onClick: () ->
    actions.clickInteraction @props.data

  render: () ->
    size = switch @props.data.size
      when 1 then "small"
      when 2 then "medium"
      else "large"
    <div onClick={@onClick} className={size} style={{left: "#{@props.data.x}px", top: "#{@props.data.y}px" }}></div>


module.exports = React.createClass
  getInitialState: () ->
    { interactions: null }

  componentDidMount: () ->
    sceneStore.addChangeListener(@onSceneChanged)

  componentDidUnmount: () ->
    sceneStore.removeChangeListener(@onSceneChanged)

  onSceneChanged: () ->
    @setState { interactions: sceneStore.scene?.interactions }

  render: () ->
    if @state.interactions?
      interactions = @state.interactions.map (interaction) ->
        <InteractionElement key={interaction.id} data={interaction} />

      <div className="interactions">
        {interactions}
      </div>
    else
      return null
