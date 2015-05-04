React        = require "react"
Interactions = require "./Interactions"
Scrubber     = require "./Scrubber"

module.exports = React.createClass
  render: () ->
    <section className="scene">
      <img src="/images/new_girl/s4e21/scene05569.jpeg" />

      <Scrubber />
      <Interactions />
    </section>