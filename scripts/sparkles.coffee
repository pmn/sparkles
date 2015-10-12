# Description
#  Manage :sparkle: points since
#
# Commands:
#  hubot sparkle <username> - give <username> a :sparkle: point
#  hubot sparkles - give the leaderboard for :sparkle: points
#  hubot sparkle me <username> - get your scores, optionally <username>'s scores
#  hubot sparkle party - randomly sparkle some people in the chat room
#  hubot who sparkled <username> - get <username>'s scores (use 'me' for yours)
#
_ = require("underscore")

module.exports = (robot) ->
  formatDate = (input) ->
    date = new Date(input)
    output = (date.getMonth() + 1) + "/" + date.getDate() + "/" + date.getFullYear()
    return output

  server = "http://sparklies.herokuapp.com/"

  robot.respond /sparkle party$/i, (msg) ->
    url = server + "sparkles"

    sender = msg.message.user
    room = robot.rooms[sender.room]
    unless room?
      msg.send "I have no memory of this place..."
      msg.finish()
      return

    ppl = room.users
    unless ppl?
      msg.send "Nobody in this room :("
      msg.finish()
      return

    randomized_ppl = _.shuffle(ppl)
    winner_count = Math.floor((Math.random()*10)+1)

    winners = randomized_ppl[0..winner_count]
    msg.send "OMG SPARKLE PARTY!!1!"
    for winner in winners
      sparkle = { "sparkler": msg.message.user.name.toLowerCase(), "sparklee": winner.name.toLowerCase(), "reason": "Sparkle Party!", "room": msg.message.room}
      msg.http(url).post(JSON.stringify(sparkle)) (err, res, body) ->
        response = JSON.parse(body)
        msg.send ":boom::#{response.name}: gets a :sparkle:!"
    msg.finish()

  robot.respond /who sparkled (.*)?/i, (msg) ->
    # Get someone's sparkle scores
    target = msg.match[1]
    if not target? or target is "me"
      target = msg.message.user.name
    target = target.toLowerCase()

    url = server + "sparkles/#{target}"
    msg.finish()
    msg.http(url).get() (err, res, body) ->
      if err or body is "null"
        msg.send ":crying_cat_face: nobody loves #{target}"
      else
        scores = JSON.parse(body)
        if target is msg.message.user.name
          response = "#{msg.message.user.name}, here's how you got famous:"
        else
          response = "Here's who gave :#{target}: a :sparkle:"
        for score in scores
          response += "\n#{score.sparkler} on #{formatDate(score.time)}"
          if score.reason
            response += " (#{score.reason})"
        msg.send response


  robot.respond /sparkle(?:s)? (.*)?/i, (msg) ->
    # Give a sparkle to a thing. Or a person. Whatever.
    url = server + "sparkles"
    body = msg.match[1].split(" ")
    recipient = _.first(body)
    reason = _.rest(body).join(" ")
    if recipient.match(/^scottjg/i)
      msg.send "scottjg doesn't want your fucking sparkles. get back to work."
      msg.finish()
      return
    else if recipient.match(/^dreww/i)
      msg.send "http://f.cl.ly/items/0N3t2B1i091Z3y3e2g1E/drew-disappoint-loop.gif"
      msg.finish()
      return

    sparkle = { "sparkler": msg.message.user.name.toLowerCase(), "sparklee": recipient.toLowerCase(), "reason": reason, "room": msg.message.room}
    msg.finish()
    msg.http(url).post(JSON.stringify(sparkle)) (err, res, body) ->
      if res.statusCode == 200
        response = JSON.parse(body)
        if response.score is 1
          msg.send ":tada: :#{response.name}: has just gotten their first :sparkle:! :tada:"
        else
          msg.send "Awww yiss, :#{response.name}: now has #{response.score} :sparkle: points!"

        if recipient.toLowerCase() is msg.message.user.name.toLowerCase()
          msg.send "Nothing wrong with a pat on the back, eh #{recipient}?"
      else
        msg.send res.statusCode
        msg.send body
        msg.send "Something went wrong. No :sparkle: for anyone :cry:"

  robot.respond /sparkle/i, (msg) ->
    response = "Give someone a :sparkle: with /sparkle <username>!\n"
    response += "See someone elses sparkles with /who sparkled <username>\n"
    response += "See who sparkled you with /who sparkled me\n"
    response += "See the leaderboard with /sparkles\n"
    response += "Have a sparkle party with /sparkle party"
    msg.send response

  robot.respond /sparkles/i, (msg) ->
    # Return the leaderboard for all the sparkles. ALL THE SPARKLES.
    url = server + "top/receiver"
    msg.http(url).get() (err, res, body) ->
      leaders = JSON.parse(body)
      output = "Top sparkles over the last 60 days:\n"
      for leader in leaders
        if leader.name != "me" and leader.name != "testing"
          output += "#{leader.name}: #{leader.score} points\n"
      msg.send output
