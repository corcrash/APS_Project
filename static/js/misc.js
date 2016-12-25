$(document).ready(function(){

/*navbar active class manipulation*/
  $(".nav a").on("click", function() {
    $(".nav").find(".active").removeClass("active");
    $(this).parent().addClass("active");
  });

/*Drawing type active class manipulation*/
  // $("#shapesToolbar button").on("click", function() {
  //   $(this).addClass("active").siblings().removeClass("active");
  // });


  /*Line type active class manipulation*/
  $("#stylesToolbar button").on("click", function () {
      $(this).addClass("active").siblings().removeClass("active");
  });

  // /*Line Width slider initialization*/
  // var slider = $("#lineWidthSlider").slider({
  //      formatter: function(value) {
  //        return "Current value: " + value;
  //      }
  // });

});
