request    = require "superagent"
dispatcher = require "./dispatcher"

module.exports = {
  loadEpisode: (slug, season, episode) ->
    console.log "[actions] loadEpisode(#{slug}, #{season}, #{episode})"
    request.get("/api/series/#{slug}/season/#{season}/episode/#{episode}").end (error, result) ->
      console.log "[actions] loadEpisode returned", result.body
      throw error if error? # tech demo error handling!

      dispatcher.dispatch
        actionType: "episode-loaded"
        episode: result.body

  visitFrame: (frame) ->
    dispatcher.dispatch
      actionType: "visit-frame"
      frame: frame

  clickInteraction: (interaction) ->
    dispatcher.dispatch
      actionType: "click-interaction"
      interaction: interaction
}