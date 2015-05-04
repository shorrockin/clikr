ReactStore = require "./ReactStore"
dispatcher = require "../dispatcher"

class EpisodeStore extends ReactStore
  constructor: () ->
    @dispatchToken = dispatcher.register (payload) =>
      switch payload.actionType
        when "episode-loaded"
          @episode = payload.episode
          @changed()

  # returns the number of frames in the episode
  frames: () ->
    return 0 unless @episode? and @episode.scenes.length > 0
    scenes = @episode.scenes
    return scenes[scenes.length - 1].frame

module.exports = new EpisodeStore()