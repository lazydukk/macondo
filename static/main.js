/* global jQuery, _, console, JSON */
(function($, _) {
  "use strict";
  // On DOM initialize:
  $(function() {
    var $results;

    $('#input-method-name').change(function() {
      var method = $('#input-method-name').val();
      if (method === 'AnagramService.Anagram') {
        $('#input-args').val('{"mode": "build", "letters": "AEROLITH"}');
      } else if (method === 'GaddagService.Generate') {
        $('#input-args').val([
          '{"filename": "/Users/Cesar/coding/ujamaa/words/OWL2.txt",',
          ' "minimize": true}'
        ].join(''));
      } else if (method === 'GaddagService.LoadDawg') {
        $('#input-args').val(
          '{"filename": "/Users/Cesar/coding/ujamaa/words/OWL2.dawg"}');
      } else if (method === 'GaddagService.GenerateDawg') {
        $('#input-args').val([
          '{"filename": "/Users/Cesar/coding/ujamaa/words/OWL2.txt",',
          ' "minimize": true}'
        ].join(''));
      }
    });


    $results = $('#textarea-results');
    $.jsonRPC.setup({
      endPoint: '/rpc'
    });
    $('#submit-rpc').click(function() {
      var methodName, args;
      methodName = $('#input-method-name').val();
      args = $('#input-args').val();
      try {
        args = JSON.parse(args);
      } catch (e) {
        $results.val("Could not parse arguments: " + e);
        return;
      }
      if (!_.isObject(args)) {
        $results.val("You must enter an object as an argument.");
        return;
      }
      $.jsonRPC.request(methodName, {
        params: args,
        success: function(result) {
          $results.val("Success! Result was: " + JSON.stringify(result.result));
          console.log('Success! Result was:', result);
        },
        error: function(result) {
          $results.val("Failure :( Result was: " + result.error.message);
          console.log('Failure. Result:', result);
        }
      });
    });

    // Trigger change event.
    $('#input-method-name').change();

  });

}(jQuery, _));