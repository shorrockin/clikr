request    = require "superagent"
dispatcher = require "./dispatcher"

module.exports = {
  loadEpisode: (slug, season, episode) ->
    console.log "[actions] loadEpisode(#{slug}, #{season}, #{episode})"
    request.get("/api/series/#{slug}/season/#{season}/episode/#{episode}").end (error, result) ->
      console.log "[actions] loadEpisode ->", result.body
}