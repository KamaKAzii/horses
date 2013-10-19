var utils = {
  testCall: function() {
    var request = $.ajax({
      url: "/static/index.html"
    });
    request.done(function(data) {
      console.log(data);
    });
  }
};

$(function () {
  $("button.race").on("click", function(e) {
    utils.testCall();
    e.preventDefault();
  });
});
