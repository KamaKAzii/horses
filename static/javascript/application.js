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
  placeBet: function($form) {
    var request = $.ajax({
      url: "/api/placeBet",
      type: "post",
      data: $form.serialize()
    });
    request.done(function(data) {
      api.getWorld();
    });
    request.fail(function(data) {
      $(".error span").html(data.responseText);
      setTimeout(function() {
        $(".error span").slideUp("slow", function() {
          $(this).empty().show();
        });
      }, 2000);
    });
  },
  runRace: function() {
    var request = $.ajax({
      url: "/api/runRace",
      type: "post",
      dataType: "json"
    });
    request.done(function(data) {
      utils.animateRaceResults(data.RaceOrder);
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

    // Render betting forms
    var $bets = $(".bets").empty();
    var $horsesSelect = $("<select>")
      .attr("name", "horse");

    $.each(horses, function(index, horse) {
      var $opt = $("<option>")
        .html(horse.Name)
        .attr("value", horse.Id);
      $horsesSelect.append($opt);
    });
    $.each(players, function(index, player) {
      var $form = $("<form>");
      var $inputAmount = $("<input>")
        .attr({
          "name": "amount",
          "placeholder": "Amount"
        });
      var $playerIdHidden = $("<input>")
        .attr({
          "type": "hidden",
          "name": "player",
          "value": player.Id,
        });
      var $submit = $("<button>")
        .html("Place Bet for " + player.Name)
        .addClass("placeBet");

      $form.append($horsesSelect.clone());
      $form.append($playerIdHidden);
      $form.append($inputAmount);
      $form.append($submit);

      $bets.append($form);

      // Hook up form events
      $submit.on("click", function(e) {
        api.placeBet($form);
        e.preventDefault();
      });
    });

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
    console.log(resultSet);

    $.each(resultSet, function(index, value) {

      var position = index + 1;
      var baseTime = 7000;
      var targetId = parseInt(value);

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
  // Load the world data
  api.getWorld();

  // Race button
  $("button.race").on("click", function(e) {
    api.runRace();
    e.preventDefault();
  });

});
