React        = require "react"
episodeStore = require "../stores/episodeStore"
sceneStore   = require "../stores/sceneStore"

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
        size = switch interaction.size
          when 1 then "small"
          when 2 then "medium"
          else "large"

        <div key={interaction.id} className={size} style={{left: "#{interaction.x}px", top: "#{interaction.y}px" }}></div>

      <div className="interactions">
        {interactions}
      </div>
    else
      <div className="interactions"/>