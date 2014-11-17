module.exports = ['$rootScope', '$state', 'Authenticator',
  function($rootScope, $state, Authenticator) {
    $rootScope.getUser = function() {
      return Authenticator.getUser();
    };

    $rootScope.$on("$stateChangeStart", function(event, toState, toParams, fromState, fromParams) {
      var userType = 'Guest';

      if (Authenticator.getUser()) userType = Authenticator.getUser().Type;

      var allowed = !(toState.allow instanceof Array) || toState.allow.indexOf(userType) !== -1;

      // Check user is authenticated...
      if (!allowed) {
        // Log the state we were attepting to access
        toParams.attemptedStateParams = encodeURI(btoa(JSON.stringify(toParams)));
        toParams.attemptedStateName = toState.name;

        // Go to the signin state (login view)
        $state.go("auth.sign-in", toParams);

        // Don't do the thing we were going to do
        event.preventDefault();
      }
    });

    // App-wide accessible method for signing out
    $rootScope.signOut = function() {
      Authenticator.signOut(function() {
        $state.go('dashboard');
      });
    };
  }
];