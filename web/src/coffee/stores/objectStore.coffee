ReactStore = require "./ReactStore"
dispatcher = require "../dispatcher"

class ObjectStore extends ReactStore
  constructor: () ->
    @dispatchToken = dispatcher.register (payload) =>
      switch payload.actionType
        when "click-interaction"
          @object = payload.interaction.object
          @object.recommendations = payload.recommendations
          @changed()
        when "clear-object"
          @object = null
          @changed()

module.exports = new ObjectStore()