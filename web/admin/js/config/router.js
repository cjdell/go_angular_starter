module.exports = ['$stateProvider', '$urlRouterProvider',
  function($stateProvider, $urlRouterProvider) {
    $urlRouterProvider.otherwise('/records/products/new');

    var routes = require('../../config/routes.json');

    // Amend static routes here if functions needed etc.
    // routes["records.categories.view"].resolve = {
    //   abc: function() {
    //     return 123;
    //   }
    // };

    for (var stateName in routes) {
      var stateInfo = routes[stateName];

      $stateProvider.state(stateName, stateInfo);
    }
  }
];