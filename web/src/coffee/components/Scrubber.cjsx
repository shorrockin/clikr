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
    $("body").on "mouseup", @onMouseUp

  componentDidUnmount: () ->
    sceneStore.removeChangeListener(@onSceneChanged)
    $("body").off "mouseup", @onMouseUp

  onSceneChanged: () ->
    @setState { frame: sceneStore.frame, max: episodeStore.frames() }

  scrubberWidth: () ->
    640 # could probably figure this out by measuring if this was responsive

  controlPosition: () ->
    return 0 if @state.max == 0
    return (@scrubberWidth() * (@state.frame / @state.max)) - 10 # 10 is width of control

  onScrubberClick: (event) ->
    offset = $(".scrubber").offset().left
    x      = Math.max(event.clientX - offset, 0)
    x      = Math.min(x, @scrubberWidth())
    perc   = (x / @scrubberWidth())
    frame  = (@state.max * perc)

    actions.visitFrame(frame)
    return

  onMouseMove: (event) ->
    return unless @mouseDown == true
    @onScrubberClick(event)

  onMouseDown: (event) ->
    @mouseDown = true
    $("body").on "mousemove", @onMouseMove
    $("body").css("user-select", "none")
    return

  onMouseUp: (event) ->
    @mouseDown = false
    $("body").off "mousemove", @onMouseMove
    $("body").css("user-select", "")
    return

  render: () ->
    <div className="scrubber" onClick={@onScrubberClick} onMouseDown={@onMouseDown}>
      <div className="control" style={{left: "#{@controlPosition()}px"}}></div>
    </div>
