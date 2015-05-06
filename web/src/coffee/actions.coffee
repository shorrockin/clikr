request    = require "superagent"
dispatcher = require "./dispatcher"

userId = new Date().getTime()

module.exports = {
  loadEpisode: (slug, season, episode) ->
    request.get("/api/series/#{slug}/season/#{season}/episode/#{episode}").end (error, result) ->
      console.log "[actions] loadEpisode(#{slug}, #{season}, #{episode}) -> ", result.body
      throw error if error? # tech demo error handling!

      dispatcher.dispatch
        actionType: "episode-loaded"
        episode: result.body

  visitFrame: (frame) ->
    dispatcher.dispatch
      actionType: "visit-frame"
      frame: frame

  clickInteraction: (interaction) ->
    request.put("/api/objects/#{interaction.object.id}/click")
           .set("X-Clikr-User", userId)
           .end (error, result) ->

      console.log "[actions] clickInteraction(#{interaction.object.name}:#{interaction.object.id}) -> ", result.body

      dispatcher.dispatch
        actionType: "click-interaction"
        interaction: interaction
        recommendations: result.body

  clearObject: () ->
    dispatcher.dispatch
      actionType: "clear-object"
}