// Allows for nicer validation so that the user isn't pestered with validation hints until they have been given the chance to complete the form input
module.exports = [

  function() {
    return {
      restrict: 'E',
      require: '?ngModel',
      link: function(scope, elm, attr, ctrl) {
        if (!ctrl) {
          return;
        }

        elm.on('focus', function() {
          elm.addClass('has-focus');
          ctrl.$hasFocus = true;
        });

        elm.on('blur', function() {
          elm.removeClass('has-focus');
          elm.addClass('has-visited');
          ctrl.$hasFocus = false;
          ctrl.$hasVisited = true;
        });

        // elm.closest('form').on('submit', function() {
        //   elm.addClass('has-visited');

        //   scope.$apply(function() {
        //     ctrl.hasFocus = false;
        //     ctrl.hasVisited = true;
        //   });
        // });
      }
    };
  }
];