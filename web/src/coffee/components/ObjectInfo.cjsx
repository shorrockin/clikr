React        = require "react"
objectStore  = require "../stores/objectStore"
actions      = require "../actions"

module.exports = React.createClass
  getInitialState: () ->
    { image: null, name: null, description: null }

  componentDidMount: () ->
    objectStore.addChangeListener(@onObjectChanged)

  componentDidUnmount: () ->
    objectStore.removeChangeListener(@onObjectChanged)

  onObjectChanged: () ->
    @setState {
      image: objectStore.object.image,
      name: objectStore.object.name,
      description: objectStore.object.description
    }

  render: () ->
    if @state.name?
      <section className="object-info">
        <img src={@state.image} />
        <h1>{@state.name}</h1>
        <p>{@state.description}</p>
      </section>
    else
      <section className="empty-object-info" />
