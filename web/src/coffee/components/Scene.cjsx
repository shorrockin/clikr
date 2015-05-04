React        = require "react"
Interactions = require "./Interactions"
Scrubber     = require "./Scrubber"
Image        = require "./Image"

module.exports = React.createClass
  render: () ->
    <section className="scene">
      <Image />
      <Scrubber />
      <Interactions />
    </section>