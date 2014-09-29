module.exports = ['$rootScope', '$state',
  function($rootScope, $state) {
    $rootScope.getStateClass = function() {
      return $state.current.name.replace(/\./g, '--'); // Replace dots with hyphens
    };
  }
];