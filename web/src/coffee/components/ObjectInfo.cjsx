React        = require "react"
objectStore  = require "../stores/objectStore"
actions      = require "../actions"

ObjectInfo = React.createClass
  getInitialState: () ->
    @createState(null)

  componentDidMount: () ->
    objectStore.addChangeListener(@onObjectChanged)

  componentDidUnmount: () ->
    objectStore.removeChangeListener(@onObjectChanged)

  onObjectChanged: () ->
    @setState @createState(objectStore.object)

  createState: (object) ->
    if object
      image: object.image,
      name: object.name,
      description: object.description
      recommendations: object.recommendations
    else
      image: null,
      name: null,
      description: null,
      recommendations: null

  close: () ->
    actions.clearObject()

  render: () ->
    if @state.name?
      <section className="object-info">
        <a href="#" className="close" onClick={@close}>X</a>
        <img src={@state.image} />
        <h1>{@state.name}</h1>
        <p>{@state.description}</p>
        <Recommendations data={@state.recommendations} />
      </section>

    else
      return null


Recommendations = React.createClass
  render: () ->
    if @props.data? and @props.data.length > 0
      images = @props.data.map (r) ->
        <img key={r.id} src={r.image} />

      <div className="recommendations">
        <h2>Users who liked this, also enjoyed:</h2>
        {images}
      </div>
    else
      return null


module.exports = ObjectInfo