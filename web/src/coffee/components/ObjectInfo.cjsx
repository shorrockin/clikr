React        = require "react"
objectStore  = require "../stores/objectStore"
actions      = require "../actions"

module.exports = React.createClass
  getInitialState: () ->
    # { image: null, name: null, description: null }
    { image: "http://placehold.it/140&text=Tights", name: "Tights", description: "Lorem ipsum dolor sit amet, placerat laoreet commodo nec id. Ut ultrices accumsan odio magna risus, duis dolorem turpis, faucibus convallis fringilla orci sit erat, lacus vel venenatis soluta duis convallis eleifend, nullam repellendus natoque phasellus interdum. Dui vel dolor vitae nunc posuere mauris, ut lacus dolor, malesuada orci bibendum incididunt mi, aenean cras id, arcu facilisis tellus quis facilisis. Vel aliquet adipiscing, a massa suscipit aliquam id felis gravida. Pretium montes rutrum sociis magna ante, magna praesent integer morbi sem diam rhoncus, egestas inceptos ornare sapien libero fusce dui. Tristique qui et, non pede hendrerit massa, sed magna justo turpis quis nunc sit, erat vestibulum etiam sed aliquet." }

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
