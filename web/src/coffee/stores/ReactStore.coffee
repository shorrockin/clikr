EventEmitter = require('events').EventEmitter

class ReactStore extends EventEmitter
  changed: () ->
    @emit "changed"

  addChangeListener: (cb) ->
    @on "changed", cb

  removeChangeListener: (cb) ->
    @removeListener "changed", cb

module.exports = ReactStore