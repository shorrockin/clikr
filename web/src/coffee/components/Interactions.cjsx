React = require "react"

module.exports = React.createClass
  render: () ->
    <div className="interactions">
      <div className="small" style={{left: "300px", top: "320px"}}></div>
      <div className="medium" style={{left: "50px", top: "75px"}}></div>
      <div className="large" style={{left: "530px", top: "300px"}}></div>
    </div>
