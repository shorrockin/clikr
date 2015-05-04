React        = require "react"
sceneStore   = require "../stores/sceneStore"
episodeStore = require "../stores/episodeStore"
actions      = require "../actions"
$            = require "jquery"

module.exports = React.createClass
  getInitialState: () ->
    { frame: 0, max: 0 }

  componentDidMount: () ->
    sceneStore.addChangeListener(@onSceneChanged)

  componentDidUnmount: () ->
    sceneStore.removeChangeListener(@onSceneChanged)

  onSceneChanged: () ->
    @setState { frame: sceneStore.frame, max: episodeStore.frames() }

  scrubberWidth: () ->
    640 # could probably figure this out by measuring if this was responsive

  controlPosition: () ->
    return 0 if @state.max == 0
    return (@scrubberWidth() * (@state.frame / @state.max)) - 10 # 10 is width of control

  onScrubberClick: (event) ->
    offset = $(event.target).offset()
    x      = Math.max(event.clientX - offset.left, 0)
    perc   = (x / @scrubberWidth())
    frame  = (@state.max * perc)

    actions.visitFrame(frame)
    return

  onMouseMove: (event) ->
    return unless @mouseDown == true
    @onScrubberClick(event)

  onMouseDown: (event) ->
    @mouseDown = true
    return

  onMouseUp: (event) ->
    @mouseDown = false
    return

  render: () ->
    <div className="scrubber" onClick={@onScrubberClick} onMouseMove={@onMouseMove} onMouseDown={@onMouseDown} onMouseUp={@onMouseUp} onMouseLeave={@onMouseUp}>
      <div className="control" style={{left: "#{@controlPosition()}px"}}></div>
    </div>
