$( document ).ready(function() {
	console.log( 'ready!' );


	$.get( "api/jsoncats/json?name=arnold&type=fluffy", function( data ) {
  		console.log(  data );
  		$("#name").text(data.name);
  		$("#type").text(data.type);
	});
});