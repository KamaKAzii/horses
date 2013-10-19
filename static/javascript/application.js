var api = {
  getPlayers: function() {
    var request = $.ajax({
      url: "/api/getPlayers",
      dataType: "json"
    });
    request.done(function(data) {
      utils.renderPlayers(data.Players);
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
});
