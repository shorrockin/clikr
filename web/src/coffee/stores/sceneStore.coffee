ReactStore   = require "./ReactStore"
dispatcher   = require "../dispatcher"
episodeStore = require "./episodeStore"

class SceneStore extends ReactStore
  constructor: () ->
    @dispatchToken = dispatcher.register (payload) =>
      switch payload.actionType
        when "episode-loaded"
          dispatcher.waitFor([episodeStore.dispatchToken])
          @frame = 0
          @changed()
        when "visit-frame"
          @frame = payload.frame
          @changed()

  image: () ->
    out = null
    for scene in episodeStore.episode.scenes
      out = scene.image if (@frame > scene.frame or out == null)
    out

  time: () ->
    seconds = (@frame / episodeStore.episode.fps)
    minutes = parseInt(seconds / 60)
    seconds = seconds - (minutes * 60)

    out = ""
    out += "#{parseInt(minutes)}m " if minutes > 0
    out += "#{parseInt(seconds)}s" if (seconds != 0 or minutes == 0)
    out.trim()

module.exports = new SceneStore()