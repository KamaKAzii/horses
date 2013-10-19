var api = {
  getPlayers: function() {
    var request = $.ajax({
      url: "/api/getPlayers",
      dataType: "json"
    });
    request.done(function(data) {
      utils.renderPlayers(data.Players);
    });
  },
  placeFakeBet: function() {
    var request = $.ajax({
      url: "/api/placeBet",
      type: "post",
      data: {
        "player": 1,
        "horse": 1,
        "amount": 20
      }
    });
    request.done(function(data) {
      console.log("placeBet done");
    });
  }
};

var utils = {
  renderPlayers: function(players) {
    var $players = $(".players").empty();
    var $playerUl = $("<ul>");
    $.each(players, function(playerNumber, player) {
      var playerLiMarkup = 
        "<li>" +
        "<dl class='player'>" + 
        "<dt>Name</dt>" +
        "<dd>" + player.Name + "</dd>" +
        "<dt>Money</dt>" +
        "<dd>$" + player.Money + "</dd>" +
        "</dl>" +
        "</li>";
      var $playerLi = $(playerLiMarkup);
      $playerUl.append($playerLi);
    });
    $players.append($playerUl);
  }
};

$(function () {
  // Load the player data.
  api.getPlayers();

  // Place fake bet button
  $(".placeFakeBet").on("click", function(e) {
    api.placeFakeBet();
    e.preventDefault();
  });
});
