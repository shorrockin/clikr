React = require "react"

module.exports = React.createClass
  render: () ->
    <section className="scene">
      <img src="/images/new_girl/s4e21/scene05569.jpeg" />

      <div className="scrubber">
        <div className="control"></div>
      </div>

      <div className="interactions">
        <div className="small" style={{left: "300px", top: "320px"}}></div>
        <div className="medium" style={{left: "50px", top: "75px"}}></div>
        <div className="large" style={{left: "530px", top: "300px"}}></div>
      </div>
    </section>