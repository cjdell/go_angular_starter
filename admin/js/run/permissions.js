module.exports = ['$rootScope', '$state', 'Authenticator',
  function($rootScope, $state, Authenticator) {
    $rootScope.canView = function(thing) {
      var user = Authenticator.getUser();

      if (user.Type === 'Admin') return true;

      return false;
    };
  }
];