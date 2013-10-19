var api = {
  getWorld: function() {
    var request = $.ajax({
      url: "/api/getWorld",
      dataType: "json"
    });
    request.done(function(data) {
      utils.renderWorld(data.Players, data.Horses);
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
    });
  },
  runRace: function() {
    var request = $.ajax({
      url: "/api/runRace",
      type: "post"
    });
    request.always(function(data) {
      // utils.animateRaceResults(data.Results);
      utils.animateRaceResults([1, 4, 3, 2]);
    });
  }
};

var utils = {
  renderWorld: function(players, horses) {
    // Render players
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

    // Render horses
    var $racecourse = $(".racecourse").empty();
    var $racecourseUl = $("<ul>");
    $.each(horses, function(index, value) {
      var $horseLi = $("<li>");
      var $nameSpan = $("<span>")
        .html(value.Name)
        .addClass("name");
      var $horseSpan = $("<span>")
        .addClass("horse")
        .data("id",  value.Id);
      $horseLi.append($horseSpan).append($nameSpan);
      $racecourseUl.append($horseLi);
    });
    $racecourse.append($racecourseUl);

  },
  animateRaceResults: function(resultSet) {
    var $horses = $(".horse");

    $.each(resultSet, function(index, value) {

      var position = index + 1;
      var baseTime = 7000;
      var targetId = value;

      $horses.each(function() {
        var $currentHorse = $(this);
        if ($currentHorse.data("id") == targetId) {
          $currentHorse.animate(
            { left: "100%" },
            {
              duration: baseTime + (position * 500),
              complete: function() {
                if (position == resultSet.length) {
                  setTimeout(function() { api.getWorld(); }, 1000);
                }
              }
            }
          );
        }
      });

    });

  }
};

$(function () {
  // Place fake bet button
  $(".placeFakeBet").on("click", function(e) {
    api.placeFakeBet();
    e.preventDefault();
  });
  
  // Load the world data
  api.getWorld();

  // Race button
  $("button.race").on("click", function(e) {
    api.runRace();
    e.preventDefault();
  });

});
