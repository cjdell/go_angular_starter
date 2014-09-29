var angular = require('angular'),
  _ = require('underscore'),
  common = require('./common');

var UsersController = ['$injector', '$scope', '$state', '$stateParams', 'UserApi',
  function($injector, $scope, $state, $stateParams, UserApi) {
    var ctrl = this;

    // Inherit shared functionality from the ListController
    $injector.invoke(common.ListController, ctrl, {
      $scope: $scope,
      Api: UserApi
    });
  }
];

var UserController = ['$injector', '$scope', '$state', '$stateParams', 'UserApi',
  function($injector, $scope, $state, $stateParams, UserApi) {
    var ctrl = this;

    // Inherit shared functionality from the ItemController
    $injector.invoke(common.ItemController, ctrl, {
      $scope: $scope,
      Api: UserApi
    });
  }
];

module.exports = {
  UsersController: UsersController,
  UserController: UserController
};