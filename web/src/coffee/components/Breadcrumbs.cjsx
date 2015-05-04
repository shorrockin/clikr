React = require "react"

module.exports = React.createClass
  render: () ->
    <section className="breadcrumbs">
      <span><a href="#">New Girl</a></span>
      <span><a href="#">Season 4</a></span>
      <span><a href="#">Episode 21</a></span>
      <span><a href="#">Scene (4:13)</a></span>
    </section>