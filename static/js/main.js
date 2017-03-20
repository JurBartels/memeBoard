$(document).ready(function(){
  $("#playbutton").click(function(){
    $(".audioTest").trigger('load');
    $(".audioTest").trigger('play');
  });
});
